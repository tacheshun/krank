package storage

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	scanscli "github.com/tacheshun/krank/internal"
)

const (
	nmapEndpoint = "/users/tacheshun/orgs"
	rmmURL       = "https://api.github.com"
)

type scanRepo struct {
	url string
}

func NewScanRepository() (scanscli.ScanRepo){
	return &scanRepo{url: rmmURL}
}

func (s *scanRepo) GetScans() (scans []scanscli.Scan, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", s.url, nmapEndpoint))
	if err != nil {
		return nil, err
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(contents, &scans)
	if err != nil {
		return nil, err
	}
	return
}
