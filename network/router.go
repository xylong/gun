package network

import "github.com/xylong/gun/iface"

// BaseRouter 基础路由
type BaseRouter struct{}

func (br *BaseRouter) Before(request iface.IRequest) {}
func (br *BaseRouter) Handle(request iface.IRequest) {}
func (br *BaseRouter) After(request iface.IRequest)  {}
