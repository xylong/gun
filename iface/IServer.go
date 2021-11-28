package iface

type IServer interface {
	// Start 启动服务
	Start()
	// Stop 停止服务
	Stop()
	// Run 运行服务
	Run()
	// AddRouter 添加路由
	AddRouter(router IRouter)
}