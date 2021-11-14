package network

import "github.com/xylong/gun/iface"

// Request 请求
type Request struct {
	conn iface.IConnection
	data []byte
}

// GetConnection 获取当前连接信息
func (r *Request) GetConnection() iface.IConnection {
	return r.conn
}

// GetData 获取请求数据 
func (r *Request) GetData() []byte {
	return r.data
}