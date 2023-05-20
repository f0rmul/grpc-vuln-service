package server

import (
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/f0rmul/vuln-service/config"
	vulnGrpc "github.com/f0rmul/vuln-service/internal/delivery/grpc"
	"github.com/f0rmul/vuln-service/internal/usecase"
	"github.com/f0rmul/vuln-service/pkg/logger"
	"github.com/f0rmul/vuln-service/pkg/netvuln_v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

type Server struct {
	logger logger.Logger
	cfg    *config.Config
}

func NewVulnServer(logger logger.Logger, cfg *config.Config) *Server {
	return &Server{logger: logger, cfg: cfg}
}

func (s *Server) Run() error {
	vulnUseCase := usecase.NewVulnUsecase(s.logger)

	l, err := net.Listen("tcp", net.JoinHostPort(s.cfg.GrpcServer.Host, s.cfg.GrpcServer.Port))
	if err != nil {
		s.logger.Errorf("net.Listen(): %v", err)
		return err
	}
	defer l.Close()

	server := grpc.NewServer(grpc.KeepaliveParams(keepalive.ServerParameters{
		MaxConnectionIdle: s.cfg.GrpcServer.MaxConnectionIdle,
		MaxConnectionAge:  s.cfg.GrpcServer.MaxConnectionAge,
		Timeout:           s.cfg.GrpcServer.Timeout,
		Time:              s.cfg.GrpcServer.Time,
	}))

	vulnService := vulnGrpc.NewVulnService(vulnUseCase, s.cfg, s.logger)

	netvuln_v1.RegisterNetVulnServiceServer(server, vulnService)

	s.logger.Info("[+] Grpc server initialized")

	go func() {
		s.logger.Infof("[+] Statring Grpc server on port: %s", s.cfg.GrpcServer.Port)
		s.logger.Fatal(server.Serve(l))
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	v := <-quit
	s.logger.Errorf("signal.Notify: %v", v)

	server.GracefulStop()

	s.logger.Info("Server exited properly")

	return nil
}
