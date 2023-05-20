package usecase

import (
	"context"
	"time"

	"github.com/Ullaakut/nmap/v3"
	"github.com/f0rmul/vuln-service/internal/models"
	"github.com/f0rmul/vuln-service/pkg/logger"
	"github.com/pkg/errors"
)

type VulnUsecase struct {
	logger logger.Logger
}

func NewVulnUsecase(logger logger.Logger) *VulnUsecase {
	return &VulnUsecase{logger: logger}
}

func (v *VulnUsecase) CheckTargets(ctx context.Context, targets []string, ports []string) ([]models.Target, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(ctx,
		nmap.WithTargets(targets...),
		nmap.WithPorts(ports...),
		nmap.WithCustomArguments("-sV"),
		nmap.WithScripts("/usr/share/nmap/scripts/vulners"),
		nmap.WithScriptArguments(map[string]string{"mincvss": "5.0"}),
	)

	if err != nil {
		v.logger.Error(err)
		return nil, errors.Wrap(err, "nmap.NewScanner()")
	}

	scannedTargets, warnings, err := scanner.Run()

	if len(*warnings) > 0 {
		v.logger.Infof("Warnings occured: %s", *warnings)
	}
	if err != nil {
		v.logger.Errorf("unable torun scanner %v", v)
		return nil, errors.Wrap(err, "scanner.Run()")
	}

	v.logger.Infof("Retrieving data from  scanned hosts...")

	results := make([]models.Target, 0, len(targets))

	// So bad package.... :( 6 inner for-loops
	for _, host := range scannedTargets.Hosts {

		var target models.Target
		target.IP = host.Addresses[0].String()
		for _, port := range host.Ports {

			var service models.Service
			service.Port = int32(port.ID)
			service.Name = port.Service.Name
			service.Version = port.Service.Version

			for _, script := range port.Scripts {
				if script.ID != "vulners" {
					continue
				}
				for _, table := range script.Tables {
					for _, innerTable := range table.Tables {
						var vuln models.Vulnerability

						for _, el := range innerTable.Elements {
							switch el.Key {
							case "id":
								vuln.ID = el.Value
							case "cvss":
								vuln.CVSS = el.Value
							}
						}
						service.Vulns = append(service.Vulns, vuln)
					}
				}
			}
			target.Services = append(target.Services, service)
		}
		results = append(results, target)
	}
	return results, nil
}
