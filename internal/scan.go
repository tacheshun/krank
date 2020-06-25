package internal

// Scan representation into data struct
type Scan struct {
	DeviceID    int       `json:"device_id"`
	DeviceHost  int       `json:"device_host"`
	Name        string    `json:"name"`
	Type      *Scantype   `json:"type"`
}

type Scantype int

const (
	Unknown Scantype = iota
	BasicScan
	ServiceDetection
	SpoofWithDecoys
)