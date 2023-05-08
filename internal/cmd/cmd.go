package cmd

import (
	"context"
	"hello_gf/internal/controller/hello"
	"hello_gf/internal/controller/hi"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					hello.HelloController,
				)
				group.Group("/my", func(group *ghttp.RouterGroup) {
					group.Bind(
						hi.New(),
					)
				})

			})
			s.Run()
			return nil
		},
	}
)
