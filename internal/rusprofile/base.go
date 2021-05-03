package rusprofile

import (
	"context"
	"github.com/timfame/rusprofile-service/internal/cache"
	"github.com/timfame/rusprofile-service/internal/config"
	"github.com/timfame/rusprofile-service/internal/models"
	"github.com/timfame/rusprofile-service/pkg/html_utils"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
)

type base struct {
	httpClient *http.Client
	cache      cache.Storage
	config     *config.Rusprofile
}

func NewBase(cfg *config.Rusprofile, c cache.Storage) *base {
	return &base{
		httpClient: &http.Client{
			Timeout: time.Second * 3,
		},
		cache:  c,
		config: cfg,
	}
}

func (b *base) GetCompanyByINN(ctx context.Context, inn string) (*models.Company, error) {
	company, err := b.cache.GetCompany(ctx, inn)
	if err == nil {
		return company, nil
	}

	mainDiv, err := b.getMainDivByURL(b.config.SearchURL + inn)
	if err != nil {
		return nil, err
	}

	// check if search returns ambiguous results (zero or more than one)
	if value, err := html_utils.GetAttributeValueByKey(mainDiv, html_utils.ClassAttrKey); err == nil &&
		value == b.config.SearchAmbiguousResultDivClass {
		mainDiv, err = b.processAmbiguousResult(mainDiv)
		if err != nil {
			return nil, err
		}
	}

	var name, kpp, director string

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		if result, err := b.findName(mainDiv); err != nil {
			return err
		} else {
			name = result
			return nil
		}
	})

	g.Go(func() error {
		if result, err := b.findKPP(mainDiv); err != nil {
			return err
		} else {
			kpp = result
			return nil
		}
	})

	// Info about director is compulsory, so it can be empty
	g.Go(func() error {
		director, _ = b.findDirector(mainDiv)
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, err
	}

	company = &models.Company{
		Name:     name,
		Director: director,
		INN:      inn,
		KPP:      kpp,
	}
	_ = b.cache.StoreCompany(ctx, company)

	return company, nil
}

func (b *base) findName(mainDiv *html.Node) (string, error) {
	if nameDiv, err := html_utils.FindDivByAttribute(mainDiv, html_utils.ClassAttrKey, b.config.CompanyNameDivClass); err != nil {
		return "", err
	} else {
		return html_utils.GetText(nameDiv), nil
	}
}

func (b *base) findKPP(mainDiv *html.Node) (string, error) {
	if kppSpan, err := html_utils.FindSpanByAttribute(mainDiv, html_utils.IDAttrKey, b.config.CompanyKPPSpanID); err != nil {
		return "", err
	} else {
		return html_utils.GetText(kppSpan), nil
	}
}

func (b *base) findDirector(mainDiv *html.Node) (string, error) {
	if director, err := html_utils.FindSpanByClassAndText(
		mainDiv,
		b.config.CompanyInfoTitleClass,
		b.config.CompanyInfoTitleDirectorText); err != nil {
		return "", err
	} else if text, err := html_utils.FindAmongNextSiblingsByAttribute(
		director,
		html_utils.ClassAttrKey,
		b.config.CompanyInfoTextClass); err != nil {
		return "", err
	} else {
		text = text.FirstChild
		if text.DataAtom == atom.A {
			text = text.FirstChild
		}
		return html_utils.GetText(text), nil
	}
}

func (b *base) processAmbiguousResult(mainDiv *html.Node) (*html.Node, error) {
	if _, err := html_utils.FindDivByAttribute(mainDiv, html_utils.ClassAttrKey, b.config.SearchEmptyResultDivClass); err != nil {
		return nil, ErrNotFound
	}
	if companyTitle, err := html_utils.FindDivByAttribute(
		mainDiv,
		html_utils.ClassAttrKey,
		b.config.SearchCompanyItemTitleDivClass); err != nil {
		return nil, err
	} else if a, err := html_utils.FindTagAmongChildren(companyTitle, atom.A); err != nil {
		return nil, err
	} else if href, err := html_utils.GetAttributeValueByKey(a, html_utils.HrefAttrKey); err != nil {
		return nil, err
	} else {
		mainDiv, err := b.getMainDivByURL(b.config.BaseURL + href)
		if err != nil {
			return nil, err
		}
		return mainDiv, nil
	}
}

func (b *base) getMainDivByURL(url string) (*html.Node, error) {
	resp, err := b.httpClient.Get(url)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode {
	case http.StatusInternalServerError:
		return nil, ErrInternalRusprofile
	case http.StatusTooManyRequests:
		return nil, ErrTooManyRequests
	case http.StatusOK:
	default:
		return nil, ErrRuspofileResponseStatus
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	if mainDiv, err := html_utils.FindDivByAttribute(doc, html_utils.IDAttrKey, b.config.MainDivID); err != nil {
		return nil, err
	} else {
		return mainDiv, nil
	}
}
