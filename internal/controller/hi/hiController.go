package hi

import (
	"context"
	v1 "hello_gf/api/hello/v1"

	"github.com/gogf/gf/v2/frame/g"
)

type hiController struct{}

func New() *hiController {
	return &hiController{}
}

func (c *hiController) SayHello(ctx context.Context, req *v1.HiReq) (res *v1.HiRes, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("我是hi包")
	return
}
