// ================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IHi interface {
		SayHello(ctx context.Context)
	}
)

var (
	localHi IHi
)

func Hi() IHi {
	if localHi == nil {
		panic("implement not found for interface IHi, forgot register?")
	}
	return localHi
}

func RegisterHi(i IHi) {
	localHi = i
}
