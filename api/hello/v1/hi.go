package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type HiReq struct {
	g.Meta `path:"/hi/say-hello" method:"get" tages:"hi包下的方法" summary:"api/hello/v1包的hi.go"`
}

type HiRes struct {
	g.Meta `mime:"text/html" example:"string"`
}

