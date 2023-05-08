package hello

import (
	"context"

	"github.com/gogf/gf/v2/frame/g"

	v1 "hello_gf/api/hello/v1"
)

type helloController struct{}

var HelloController = helloController{}

func (c *helloController) Hello(ctx context.Context, req *v1.Req) (res *v1.Res, err error) {
	g.RequestFromCtx(ctx).Response.Writeln("Hello World! First Demo!第一个项目项目")
	return
}
