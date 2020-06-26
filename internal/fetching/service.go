//Package fetching ...
package fetching

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/tacheshun/krank/internal/storage"

	"github.com/Ullaakut/nmap"

	scanscli "github.com/tacheshun/krank/internal"
)

//WHAT THE HELL DOES THE GOLINT WANT FROM MY LIFE ...
const (
	TIMES = 5
	SECONDS = 5000
)

// Service provides scans fetching operations.
type Service interface {
	FetchScans() ([]scanscli.Scan, error)
	RunBasicScan(string) (map[string]string, []string, error)
}

type service struct {
	sR scanscli.ScanRepo
}

// NewService creates an adding service with the necessary dependencies.
func NewService(r scanscli.ScanRepo) *service {
	return &service{r}
}

func (s *service) FetchScans() ([]scanscli.Scan, error) {
	return s.sR.GetScans()
}

// RunBasicScan scans given target hosts for open ports.
func (s *service) RunBasicScan(jobID string) (resultMap map[string]string, warnings []string, err error) {

	ctx, cancel := context.WithTimeout(context.Background(), TIMES*time.Minute)
	defer cancel()

	scanner, err := nmap.NewScanner(
		nmap.WithTargets("0.0.0.0"),
		nmap.WithPorts("80,443,22,554,843,8554"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	var result *nmap.Run
	result, _, err = scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}
	resultMap = make(map[string]string)
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}
		resultMap["deviceId"] = storage.DeviceID
		resultMap["jobId"] = jobID

		for _, port := range host.Ports {
			resultMap["response"] += strconv.Itoa(int(port.ID)) + port.Protocol + " " + port.Service.Name + " " + port.State.String() + "<br />"
		}
	}

	return resultMap, warnings, nil
}
