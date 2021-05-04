package test

import (
	"context"
	"encoding/json"
	"github.com/pkg/errors"
	"github.com/timfame/rusprofile-service/internal/config"
	"github.com/timfame/rusprofile-service/internal/delivery/grpc"
	"github.com/timfame/rusprofile-service/internal/models"
	"golang.org/x/sync/errgroup"
	"net/http"
	"testing"
	"time"
)

func TestGateway(t *testing.T) {
	service, l := getTestServiceAndLogger()
	grpcServ := grpc.NewServer(&config.Grpc{
		Port:        "8087",
		GatewayPort: "8097",
	}, service, l)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return errors.WithMessage(grpcServ.Run(), "Grpc server")
	})
	g.Go(func() error {
		return errors.WithMessage(grpcServ.RunGateway(ctx), "Grpc gateway server")
	})

	//time.Sleep(5 * time.Second)

	resp, err := http.Get("http://localhost:8097/v1/companies?inn=7707083893")
	if err != nil {
		t.Fatal(err)
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Response status code: %d", resp.StatusCode)
	}

	var cmp struct {
		Name     string `json:"name"`
		INN      string `json:"inn"`
		KPP      string `json:"kpp"`
		Director string `json:"director"`
	}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&cmp); err != nil {
		t.Fatalf("Cannot decode response: %v", err)
	}

	checkCompany(t, &models.Company{
		Name:     "ПУБЛИЧНОЕ АКЦИОНЕРНОЕ ОБЩЕСТВО \"СБЕРБАНК РОССИИ\"",
		Director: "Греф Герман Оскарович",
		INN:      "7707083893",
		KPP:      "773601001",
	}, &models.Company{
		Name:     cmp.Name,
		Director: cmp.Director,
		INN:      cmp.INN,
		KPP:      cmp.KPP,
	})

	cancel()

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer shutdownCancel()

	if err := grpcServ.GracefulStop(shutdownCtx); err != nil {
		t.Fatalf("Graceful stop failed: %v", err)
	}

	_ = g.Wait()

}
