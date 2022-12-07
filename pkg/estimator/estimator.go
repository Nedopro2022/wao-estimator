package estimator

import (
	"context"
	"errors"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	ErrInvalidRequest    = errors.New("ErrInvalidRequest")
	ErrEstimatorNotFound = errors.New("ErrEstimatorNotFound")
	ErrEstimatorError    = errors.New("ErrEstimatorError")
)

type Estimator struct {
	k8sClient client.Client
}

func NewEstimator() *Estimator {
	return &Estimator{}
}

func (e *Estimator) EstimatePowerConsumption(ctx context.Context, namespace, estimator string, cpuMilli, numWorkload int) (*[]int, error) {
	var wattIncrease []int
	// TODO
	return &wattIncrease, nil
}
