//Package storage ...
package storage

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	scanscli "github.com/tacheshun/krank/internal"
	"github.com/tacheshun/krank/internal/errors"
)

const (
	NmapEndpointCheckRun = "/check-run/"
	NmapEndpointAck      = "http://localhost/dashboard/api/nmap/acknowledge/"
	RmmURL               = "http://localhost/dashboard/api/nmap"
	DeviceID             = "65898"
)

type scanRepo struct {
	url string
}

//NewScanRepository public method constructor .
func NewScanRepository() scanscli.ScanRepo {
	return &scanRepo{url: RmmURL}
}

func (s *scanRepo) GetScans() ([]scanscli.Scan, error) {
	requestBody, err := json.Marshal(map[string]string{
		"deviceId": DeviceID,
	})

	response, err := http.Post(fmt.Sprintf("%v%v", s.url, NmapEndpointCheckRun), "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response to %s", NmapEndpointCheckRun)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading the response from %s", NmapEndpointCheckRun)
	}

	var decoded []scanscli.Scan
	err = json.Unmarshal(contents, &decoded)
	if err != nil {
		return nil, nil
	}

	err = response.Body.Close()
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't close response body")
	}

	return decoded, nil
}
