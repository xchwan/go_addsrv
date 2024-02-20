package calculate

import (
	"context"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

func covint(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return num
}

type addRequest struct {
	A string `json:"a"`
	B string `json:"b"`
}

type addResponse struct {
	Value string `json:"value"`
	Err   string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makeaddEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(addRequest)
		value, err := svc.AddValue(ctx, covint(req.A), covint(req.B))
		if err != nil {
			return addResponse{strconv.Itoa(0), err.Error()}, err
		}
		return addResponse{strconv.Itoa(value), ""}, err
	}
}

type subRequest struct {
	A string `json:"a"`
	B string `json:"b"`
}

type subResponse struct {
	Value string `json:"value"`
	Err   string `json:"err,omitempty"` // errors don't JSON-marshal, so we use a string
}

func makesubEndpoint(svc Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(subRequest)
		value, err := svc.SubValue(ctx, covint(req.A), covint(req.B))
		if err != nil {
			return subResponse{strconv.Itoa(0), err.Error()}, err
		}
		return subResponse{strconv.Itoa(value), ""}, err
	}
}
