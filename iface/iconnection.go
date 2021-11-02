package iface

import "net"

// IConnection 连接
type IConnection interface {
	// 启动连接
	Start()
	// 停止连接
	Stop()
	// 获取当前连接
	GetTCPConnection() *net.TCPConn
	// 获取当前连接模块的🆔
	GetConnID() uint32
	// 获取远程客户端的tcp状态
	GetRemoteAddr() net.Addr
	// 发送数据 
	Send(data []byte) error
}

// HandleFunc 业务处理函数
type HandleFunc func(*net.TCPConn,[]byte,int) error 