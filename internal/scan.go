//Package internal .
package internal

import "encoding/json"

// Scan representation into data struct .
type Scan struct {
	ScanID      int       `json:"id"`
	Login       string    `json:"login"`
	NodeID      string    `json:"node_id"`
	URL         string    `json:"url"`
	Type      *Scantype   `json:"type"`
}

//Scantype definition .
type Scantype int

const (
	//Unknown Scantype iota .
	Unknown Scantype = iota
	//BasicScan Scantype iota .
	BasicScan
	//ServiceDetection Scantype iota .
	ServiceDetection
)


func (s Scantype) String() string {
	return toString[s]
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


// UnmarshalJSON convert type from json to scanType .
func (s *Scantype) UnmarshalJSON(b []byte) error {
	var j string
	err := json.Unmarshal(b, &j)
	if err != nil {
		return err
	}
	*s = toID[j]
	return nil
}

// ScanRepo definition of methods to access data .
type ScanRepo interface {
	GetScans() ([]Scan, error)
}
