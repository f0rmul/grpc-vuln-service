package models

import "github.com/f0rmul/vuln-service/pkg/netvuln_v1"

type Service struct {
	Name    string
	Version string
	Port    int32
	Vulns   []Vulnerability
}

func (s *Service) ToProto() *netvuln_v1.Service {
	vulns := make([]*netvuln_v1.Vulnerability, 0, len(s.Vulns))

	for _, vuln := range s.Vulns {
		vulns = append(vulns, vuln.ToProto())
	}

	return &netvuln_v1.Service{
		Name:    s.Name,
		Version: s.Version,
		TcpPort: s.Port,
		Vulns:   vulns,
	}
}
