package iface

// IRequest 客户端请求的连接和数据
type IRequest interface {
	// 获取当前连接
	GetConnection() IConnection

	// 获取请求数据
	GetData() []byte
} 