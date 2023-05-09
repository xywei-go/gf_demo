package hi

import (
	"context"
	"hello_gf/internal/service"
)

type sHi struct{}

func init() {
	service.RegisterHi(New())
}

func New() service.IHi {
	return &sHi{}
}

func (s *sHi) SayHello(ctx context.Context) {
}
