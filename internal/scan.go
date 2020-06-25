package internal

import "encoding/json"

// Scan representation into data struct
type Scan struct {
	ScanID      int       `json:"id"`
	Login       string    `json:"login"`
	NodeID      string    `json:"node_id"`
	URL         string    `json:"url"`
	Type      *Scantype   `json:"type"`
}

type Scantype int

const (
	Unknown Scantype = iota
	BasicScan
	ServiceDetection
)

func (s Scantype) String() string {
	return toString[s]
}

// NewScanType initialize a type from enum beerTypes
func NewScanType(t string) *Scantype {
	scantype := toID[t]
	return &scantype
}

var toString = map[Scantype]string{
	BasicScan:         "BasicScan",
	ServiceDetection:  "ServiceDetection",
	Unknown:           "unknown",
}

var toID = map[string]Scantype{
	"BasicScan":          BasicScan,
	"ServiceDetection":   ServiceDetection,
	"unknown":            Unknown,
}


// UnmarshalJSON convert type from json to scanType
func (s *Scantype) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = toID[j]
	return nil
}

// ScanRepo definiton of methods to access data
type ScanRepo interface {
	GetScans() ([]Scan, error)
}

// NewScan initialize struct scan
func NewScan(id int, login string, nodeId string, url string, scanType *Scantype) (s Scan) {
	s = Scan{
		ScanID:    id,
		Login:     login,
		NodeID:    nodeId,
		URL:       url,
		Type:      scanType,
	}
	return
}
