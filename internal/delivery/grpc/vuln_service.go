package grpc

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/f0rmul/vuln-service/config"
	"github.com/f0rmul/vuln-service/internal/delivery/grpc/validate"
	"github.com/f0rmul/vuln-service/internal/models"
	"github.com/f0rmul/vuln-service/pkg/logger"
	netvuln "github.com/f0rmul/vuln-service/pkg/netvuln_v1"
	"github.com/f0rmul/vuln-service/pkg/utils"
)

type VulnUsecase interface {
	CheckTargets(ctx context.Context, targets []string, ports []string) ([]models.Target, error)
}

type VulnServive struct {
	netvuln.UnimplementedNetVulnServiceServer
	vulnUseCase VulnUsecase
	logger      logger.Logger
	cfg         *config.Config
}

func NewVulnService(uc VulnUsecase, cfg *config.Config, logger logger.Logger) *VulnServive {
	return &VulnServive{vulnUseCase: uc, cfg: cfg, logger: logger}
}

func (service *VulnServive) CheckVuln(ctx context.Context, request *netvuln.CheckVulnRequest) (*netvuln.CheckVulnResponse, error) {
	err := validate.ValidateProtoRequest(request)
	if err != nil {
		service.logger.Errorf("validate.ValidateProtoRequest(): %v", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	service.logger.Info("Starting processing targets")

	results, err := service.vulnUseCase.CheckTargets(
		ctx,
		request.GetTargets(),
		utils.IntToStringSlice(request.GetTcpPorts()),
	)

	if err != nil {
		service.logger.Errorf("service.vulnUseCase.CheckTargets(): %v", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	pbTargetResults := make([]*netvuln.TargetResult, 0, len(results))
	for _, target := range results {
		pbTargetResults = append(pbTargetResults, target.ToProto())
	}

	return &netvuln.CheckVulnResponse{
		Results: pbTargetResults,
	}, nil
}
