package fetching

import (
	"context"
	"fmt"
	"log"
	"time"

	scanscli "github.com/tacheshun/krank/internal"
	"github.com/pkg/errors"
	"github.com/Ullaakut/nmap"
)

// Service provides scans fetching operations.
type Service interface {
	// FetchScans fetch all scans from external repository.
	FetchScans() ([]scanscli.Scan, error)
	// FetchByID filter all scans and get only the scan that match with given id.
	FetchByID(id int) (scanscli.Scan, error)
}

type service struct {
	sR scanscli.ScanRepo
}

func (s *service) FetchScans() ([]scanscli.Scan, error) {
	return s.sR.GetScans()
}

func (s *service) FetchByID(id int) (scanscli.Scan, error) {
	scans, err := s.FetchScans()
	if err != nil {
		return scanscli.Scan{}, err
	}

	for _, scan := range scans {
		if scan.ScanID == id {
			return scan, nil
		}
	}

	return scanscli.Scan{}, errors.Errorf("Scan %d not found", id)
}

// NewService creates an adding service with the necessary dependencies.
func NewService(r scanscli.ScanRepo) Service {
	return &service{r}
}

// RunBasicScan scans given target hosts for open ports.
func RunBasicScan()  {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
	defer cancel()

	// Equivalent to `/usr/local/bin/nmap -p 80,443,843 google.com facebook.com youtube.com`,
	// with a 5 minute timeout.
	scanner, err := nmap.NewScanner(
		nmap.WithTargets("google.com", "facebook.com", "youtube.com"),
		nmap.WithPorts("80,443,843"),
		nmap.WithContext(ctx),
	)
	if err != nil {
		log.Fatalf("unable to create nmap scanner: %v", err)
	}

	result, _, err := scanner.Run()
	if err != nil {
		log.Fatalf("unable to run nmap scan: %v", err)
	}

	// Use the results to print an example output
	for _, host := range result.Hosts {
		if len(host.Ports) == 0 || len(host.Addresses) == 0 {
			continue
		}

		fmt.Printf("Host %q:\n", host.Addresses[0])

		for _, port := range host.Ports {
			fmt.Printf("\tPort %d/%s %s %s\n", port.ID, port.Protocol, port.State, port.Service.Name)
		}
	}

	fmt.Printf("Nmap done: %d hosts up scanned in %.2f seconds\n", len(result.Hosts), result.Stats.Finished.Elapsed)
}

