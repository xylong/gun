package iface

// IRouter 路由
type IRouter interface {
	// 处理conn业务之前的🪝方法
	Before(IRequest)
	// 处理业务🪝方法
	Handle(IRequest)
	// 处理conn业务后的🪝方法
	After(IRequest)
}