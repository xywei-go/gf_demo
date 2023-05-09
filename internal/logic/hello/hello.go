package hello

import (
	"fmt"
	"hello_gf/internal/service"
)

type sHello struct{}

func init() {
	fmt.Println("-------------------------------------------------hello.go------------------------------------")
	service.RegisterHello(New())
}

func New() service.IHello {
	return &sHello{}
}

func (sHello *sHello) SayHello() string {
	return "hello! from logic hello"
}
