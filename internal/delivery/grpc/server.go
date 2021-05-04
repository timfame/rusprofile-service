package grpc

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	rs "github.com/timfame/rusprofile-service/gen/rusprofile_service"
	"github.com/timfame/rusprofile-service/internal/config"
	"github.com/timfame/rusprofile-service/internal/rusprofile"
	"github.com/timfame/rusprofile-service/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"net/http"
)

const (
	localHost = "0.0.0.0"
)

type server struct {
	*rs.UnimplementedRusprofileServiceServer
	grpcServer *grpc.Server
	gwServer   *http.Server
	service    rusprofile.Service
	logger     *logger.Logger
	config     *config.Grpc
}

func NewServer(cfg *config.Grpc, s rusprofile.Service, l *logger.Logger) *server {
	return &server{
		grpcServer: grpc.NewServer(),
		service:    s,
		logger:     l,
		config:     cfg,
	}
}

func (s *server) Run() error {
	lis, err := net.Listen("tcp", ":" + s.config.Port)
	if err != nil {
		return err
	}

	rs.RegisterRusprofileServiceServer(s.grpcServer, s)

	s.logger.Info("Grpc server started", zap.String("port", s.config.Port))
	if err := s.grpcServer.Serve(lis); err != nil {
		return err
	}
	return nil
}

func (s *server) RunGateway(ctx context.Context) error {
	conn, err := grpc.DialContext(ctx, localHost + ":" + s.config.Port, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return err
	}

	gwMux := runtime.NewServeMux()
	err = rs.RegisterRusprofileServiceHandler(ctx, gwMux, conn)
	if err != nil {
		return err
	}

	s.gwServer = &http.Server{
		Addr:    ":" + s.config.GatewayPort,
		Handler: gwMux,
	}

	s.logger.Info("Grpc gateway started", zap.String("port", s.config.GatewayPort))
	if err := s.gwServer.ListenAndServe(); err != nil {
		return err
	}
	return nil
}

func (s *server) GracefulStop(ctx context.Context) error {
	if s.grpcServer != nil {
		s.grpcServer.GracefulStop()
	}
	if s.gwServer != nil {
		if err := s.gwServer.Shutdown(ctx); err != nil {
			return err
		}
	}
	return nil
}
