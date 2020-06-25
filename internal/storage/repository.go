//Package storage ...
package storage

import (
	"encoding/json"
	"fmt"
	"github.com/tacheshun/krank/internal/errors"
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

//NewScanRepository public method constructor .
func NewScanRepository() scanscli.ScanRepo {
	return &scanRepo{url: rmmURL}
}

func (s *scanRepo) GetScans() (scans []scanscli.Scan, err error) {
	response, err := http.Get(fmt.Sprintf("%v%v", s.url, nmapEndpoint))
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error getting response to %s", nmapEndpoint)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "error reading the response from %s", nmapEndpoint)
	}

	err = json.Unmarshal(contents, &scans)
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't parsing response into scans data")
	}

	err = response.Body.Close()
	if err != nil {
		return nil, errors.WrapDataUnreacheable(err, "can't close response body")
	}

	return
}
