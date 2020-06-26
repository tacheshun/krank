//Package fetching ...
package fetching

import (
	"context"
	"fmt"
	"log"
	"math"
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
	RunBasicScan() (*nmap.Run, []string, error)
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
func (s *service) RunBasicScan() (result *nmap.Run, warnings []string, err error){

	ctx, cancel := context.WithTimeout(context.Background(), TIMES*time.Minute)
	defer cancel()

	// Equivalent to `/usr/local/bin/nmap -p 80,443,843 google.com facebook.com youtube.com`,
	// with a 5 minute timeout.
	scanner, err := nmap.NewScanner(
		nmap.WithTargets("0.0.0.0"),
		nmap.WithPorts("80,443,22,843"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, _, err = scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	// Return result, optional warnings but no error
	return result, warnings, nil
}

func numOfRoutines(numOfScans, scansPerRouting int) int {
	return int(math.Ceil(float64(numOfScans) / float64(scansPerRouting)))
}
