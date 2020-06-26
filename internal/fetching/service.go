//Package fetching ...
package fetching

import (
	"context"
	"log"
	"math"
	"strconv"
	"sync"
	"time"

	"github.com/Ullaakut/nmap"

	scanscli "github.com/tacheshun/krank/internal"
)

//WHAT THE HELL DOES THE GOLINT WANT FROM MY LIFE ...
const (
	TIMES = 5
)

// Service provides scans fetching operations.
type Service interface {
	// FetchScans fetches all scans from external repository.
	FetchScans() ([]scanscli.Scan, error)
	// FetchByID filter all scans and get only the scan that match with given id.
	FetchByID(id int) (scanscli.Scan, error)
	RunBasicScan() (map[string]string, []string, error)
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

func (s *service) FetchByID(id int) (scanscli.Scan, error) {
	scans, err := s.FetchScans()
	if err != nil {
		return scanscli.Scan{}, err
	}

	scansPerRoutine := 10
	numRoutines := numOfRoutines(len(scans), scansPerRoutine)

	wg := &sync.WaitGroup{}
	wg.Add(numRoutines)

	var b scanscli.Scan

	for i := 0; i < numRoutines; i++ {
		go func(id, begin, end int, scans []scanscli.Scan, b *scanscli.Scan, wg *sync.WaitGroup) {
			for i := begin; i <= end; i++ {
				if scans[i].ScanID == id {
					*b = scans[i]
				}
			}
			wg.Done()
		}(id, i, i+scansPerRoutine, scans, &b, wg)
	}

	wg.Wait()

	return b, nil
}

// RunBasicScan scans given target hosts for open ports.
func (s *service) RunBasicScan() (resultMap map[string]string, warnings []string, err error) {

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

		for _, port := range host.Ports {
			resultMap[strconv.Itoa(int(port.ID))] = port.Protocol + "/" + port.Service.Name + "/" + port.State.String()
		}
	}
	return resultMap, warnings, nil
}

func numOfRoutines(numOfScans, scansPerRouting int) int {
	return int(math.Ceil(float64(numOfScans) / float64(scansPerRouting)))
}
