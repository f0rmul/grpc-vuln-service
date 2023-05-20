package models

import "github.com/f0rmul/vuln-service/pkg/netvuln_v1"

type Target struct {
	IP       string
	Services []Service
}

func (t *Target) ToProto() *netvuln_v1.TargetResult {

	services := make([]*netvuln_v1.Service, 0, len(t.Services))

	for _, service := range t.Services {
		services = append(services, service.ToProto())
	}

	return &netvuln_v1.TargetResult{
		Target:   t.IP,
		Services: services,
	}
}
