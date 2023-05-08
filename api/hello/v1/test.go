package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type TestReq struct {
	g.Meta `path:"/test/sayHello" method:"get" tages:"test包下的方法" summary:"api/hello/v1包的test.go"`
}

type TestRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
