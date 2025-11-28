package request

import "github.com/inexio/thola/internal/device"

// ReadSystemRequest
//
// ReadSystemRequest is the request struct for the read system request.
//
// swagger:model
type ReadSystemRequest struct {
	ReadRequest
}

// ReadSystemResponse
//
// ReadSystemResponse is the response struct for the read system response.
//
// swagger:model
type ReadSystemResponse struct {
	System device.SystemComponent `yaml:"system" json:"system" xml:"system"`
	ReadResponse
}
