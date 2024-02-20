package calculate

import (
	"context"
	"errors"
	"math"
)

type Service interface {
	AddValue(ctx context.Context, a, b int) (int, error)
	SubValue(ctx context.Context, a, b int) (int, error)
}

var (
	ErrIntOverflow = errors.New("integer overflow")
)

type service struct{}

func NewService() *service {
	return &service{}
}

func (s *service) AddValue(ctx context.Context, a, b int) (int, error) {
	if (b > 0 && a > (math.MaxInt-b)) || (b < 0 && a < (math.MinInt-b)) {
		return 0, ErrIntOverflow
	}
	return a + b, nil
}

func (s *service) SubValue(ctx context.Context, a, b int) (int, error) {
	if (a > math.MaxInt) || (b > math.MaxInt || a < math.MinInt || b < math.MinInt) {
		return 0, ErrIntOverflow
	}
	return a - b, nil
}
