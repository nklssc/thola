//go:build !client
// +build !client

package request

import (
	"context"

	"github.com/pkg/errors"
)

func (r *ReadSystemRequest) process(ctx context.Context) (Response, error) {
	com, err := GetCommunicator(ctx, r.BaseRequest)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get communicator")
	}

	result, err := com.GetSystemComponent(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "can't get system components")
	}

	return &ReadSystemResponse{
		System: result,
	}, nil
}
