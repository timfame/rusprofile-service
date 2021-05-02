package rusprofile

import (
	"context"
	"github.com/pkg/errors"
	"github.com/timfame/rusprofile-service/internal/cache"
	"github.com/timfame/rusprofile-service/internal/config"
	"github.com/timfame/rusprofile-service/internal/models"
	"github.com/timfame/rusprofile-service/pkg/html_utils"
	"golang.org/x/net/html"
	"golang.org/x/sync/errgroup"
	"net/http"
	"strings"
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
		return nil, errors.Wrap(err, "Get company by INN failed")
	}

	// check if search returns ambiguous results (zero or more than one)
	if value, ok := html_utils.GetAttributeValueByKey(mainDiv, html_utils.ClassAttrKey); ok &&
		value == b.config.SearchAmbiguousResultDivClass {
		mainDiv, err = b.processAmbiguousResult(mainDiv)
		if err != nil {
			return nil, errors.Wrap(err, "Get company by INN failed")
		}
	}

	var name, kpp, director string

	g, _ := errgroup.WithContext(ctx)

	g.Go(func() error {
		if result, ok := b.findName(mainDiv); ok {
			name = result
			return nil
		}
		return ErrNotFound
	})

	g.Go(func() error {
		if result, ok := b.findKPP(mainDiv); ok {
			kpp = result
			return nil
		}
		return ErrNotFound
	})

	// Info about director is compulsory, so it can be empty
	g.Go(func() error {
		director, _ = b.findDirector(mainDiv)
		return nil
	})

	if err := g.Wait(); err != nil {
		return nil, errors.Wrap(err, "Get company by INN failed")
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

func (b *base) findName(mainDiv *html.Node) (string, bool) {
	if nameDiv, ok := html_utils.FindDivByAttribute(mainDiv, html_utils.ClassAttrKey, b.config.CompanyNameDivClass); ok {
		return strings.TrimSpace(nameDiv.Data), true
	}
	return "", false
}

func (b *base) findKPP(mainDiv *html.Node) (string, bool) {
	if kppSpan, ok := html_utils.FindSpanByAttribute(mainDiv, html_utils.IDAttrKey, b.config.CompanyKPPSpanID); ok {
		return strings.TrimSpace(kppSpan.Data), true
	}
	return "", false
}

func (b *base) findDirector(mainDiv *html.Node) (string, bool) {
	if director, ok := html_utils.FindNthByClass(
		mainDiv,
		b.config.CompanyInfoTitleWithDirectorPosition,
		b.config.CompanyInfoTitleClass); ok {
		if text, ok := html_utils.FindSpanByAttribute(director, html_utils.ClassAttrKey, b.config.CompanyInfoTextClass); ok {
			text = text.FirstChild
			if html_utils.IsA(text) {
				text = text.FirstChild
			}
			return strings.TrimSpace(text.Data), true
		}
	}
	return "", false
}

func (b *base) processAmbiguousResult(mainDiv *html.Node) (*html.Node, error) {
	if _, ok := html_utils.FindDivByAttribute(mainDiv, html_utils.ClassAttrKey, b.config.SearchEmptyResultDivClass); ok {
		return nil, ErrNotFound
	}
	if companyTitle, ok := html_utils.FindDivByAttribute(
		mainDiv,
		html_utils.ClassAttrKey,
		b.config.SearchCompanyTitleDivClass); ok {
		if href, ok := html_utils.GetHref(companyTitle.FirstChild); ok {
			mainDiv, err := b.getMainDivByURL(b.config.BaseURL + href)
			if err != nil {
				return nil, err
			}
			return mainDiv, nil
		}
	}
	return nil, ErrNotFound
}

func (b *base) getMainDivByURL(url string) (*html.Node, error) {
	resp, err := b.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	mainDiv, ok := html_utils.FindDivByAttribute(doc, html_utils.IDAttrKey, b.config.MainDivID)
	if !ok {
		return nil, ErrNotFound
	}

	return mainDiv, nil
}
