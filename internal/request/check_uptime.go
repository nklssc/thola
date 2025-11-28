package request

import (
	"context"

	"github.com/inexio/go-monitoringplugin"
)

// CheckUptimeRequest
//
// CheckUptimeRequest is the request struct for the check uptime request.
//
// swagger:model
type CheckUptimeRequest struct {
	CheckDeviceRequest
	UptimeThreshold monitoringplugin.Thresholds `json:"uptimeThreshold" xml:"uptimeThreshold"`
}

func (r *CheckUptimeRequest) validate(ctx context.Context) error {
	if err := r.UptimeThreshold.Validate(); err != nil {
		return err
	}

	return r.CheckDeviceRequest.validate(ctx)
}
