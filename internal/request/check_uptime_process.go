//go:build !client
// +build !client

package request

import (
	"context"

	"github.com/inexio/go-monitoringplugin"
)

func (r *CheckUptimeRequest) process(ctx context.Context) (Response, error) {
	r.init()

	com, err := GetCommunicator(ctx, r.BaseRequest)
	if r.mon.UpdateStatusOnError(err, monitoringplugin.UNKNOWN, "error while getting communicator", true) {
		return &CheckResponse{r.mon.GetInfo()}, nil
	}

	system, err := com.GetSystemComponent(ctx)
	if r.mon.UpdateStatusOnError(err, monitoringplugin.UNKNOWN, "error while reading system", true) {
		return &CheckResponse{r.mon.GetInfo()}, nil
	}

	if system.Uptime != nil {
		err = r.mon.AddPerformanceDataPoint(monitoringplugin.NewPerformanceDataPoint("uptime", *system.Uptime).SetThresholds(r.UptimeThreshold))
		if r.mon.UpdateStatusOnError(err, monitoringplugin.UNKNOWN, "error while adding performance data point", true) {
			r.mon.PrintPerformanceData(false)
			return &CheckResponse{r.mon.GetInfo()}, nil
		}
	} else {
		r.mon.UpdateStatus(monitoringplugin.UNKNOWN, "uptime not available")
	}

	return &CheckResponse{r.mon.GetInfo()}, nil
}
