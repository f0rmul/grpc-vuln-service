package validate

import (
	"errors"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/f0rmul/vuln-service/pkg/netvuln_v1"
)

var (
	ErrInvalidTarget = errors.New("invalid target")
	ErrInvalidPort   = errors.New("invalid port")
)

func ValidateProtoRequest(request *netvuln_v1.CheckVulnRequest) error {

	type Request struct {
		Hosts []string
		Ports []int32
	}

	req := Request{Hosts: request.GetTargets(), Ports: request.GetTcpPorts()}

	for _, host := range req.Hosts {
		if !govalidator.IsDNSName(host) && !govalidator.IsIPv4(host) {
			return ErrInvalidTarget
		}
	}

	for _, port := range req.Ports {
		if !govalidator.IsPort(strconv.Itoa(int(port))) {
			return ErrInvalidPort
		}
	}
	return nil
}
