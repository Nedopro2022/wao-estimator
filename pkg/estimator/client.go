package estimator

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Nedopro2022/wao-estimator/pkg/estimator/api"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type Client struct {
	c api.ClientWithResponsesInterface
	e client.ObjectKey
}

func NewClient(server string, estimatorNamespace, estimatorName string, opts ...api.ClientOption) (*Client, error) {
	c, err := api.NewClientWithResponses(server, opts...)
	if err != nil {
		return nil, err
	}
	ec := Client{c: c, e: client.ObjectKey{Namespace: estimatorNamespace, Name: estimatorName}}
	return &ec, nil
}

func (c *Client) EstimatePowerConsumption(ctx context.Context, cpuMilli, numWorkload int) (*api.PowerConsumption, error) {
	body := api.PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionJSONRequestBody{
		CpuMilli:     cpuMilli,
		NumWorkload:  numWorkload,
		WattIncrease: nil,
	}
	resp, err := c.c.PostNamespacesNamespaceEstimatorsEstimatorResourcesPowerconsumptionWithResponse(ctx, c.e.Namespace, c.e.Name, body)
	if err != nil {
		return nil, err
	}
	switch resp.StatusCode() {
	case http.StatusOK:
		return resp.JSON200, nil
	default:
		return nil, fmt.Errorf("%d %s", resp.StatusCode(), resp.Status())
	}
}
