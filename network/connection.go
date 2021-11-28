package network

import (
	"fmt"
	"net"

	"github.com/xylong/gun/iface"
)

type Connection struct {
	// 当前连接套接字
	Conn *net.TCPConn
	// 连接🆔
	ConnID uint32
	// 连接是否关闭
	isClosed bool
	// 告知当前连接已退出
	ExitChan chan bool
	// 该连接处理的路由
	Router iface.IRouter
}

func NewConnection(conn *net.TCPConn, connID uint32, router iface.IRouter) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		ExitChan: make(chan bool, 1),
		Router: router,
	}
}

// Read 读取数据
func (c *Connection) Read() {
	fmt.Println("read goroutine is running")
	defer fmt.Printf("ConnID: %d read exit, remote address is %s \n", c.ConnID, c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		_, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive buf error:", err)
			continue
		}

		// *执行注册的路由方法
		go func(req iface.IRequest) {
			c.Router.Before(req)
			c.Router.Handle(req)
			c.Router.After(req)
		}(&Request{
			conn: c,
			data: buf,
		})
	}
}

func (c *Connection) Start() {
	fmt.Println("conn start... ConnID: ", c.ConnID)
	go c.Read()
}

// Stop 关闭连接
func (c *Connection) Stop() {
	if c.isClosed == true {
		return
	}
	c.isClosed = true
	//? 关闭socket连接
	c.Conn.Close()
	//? 回收资源
	close(c.ExitChan)
}

// GetTCPConnection 获取tcp连接
func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

// GetConnID 获取连接🆔
func (c *Connection) GetConnID() uint32 {
	return c.ConnID
}

// GetRemoteAddr 获取远程地址信息
func (c *Connection) GetRemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}

// Send 发送数据
func (c *Connection) Send(data []byte) error {
	return nil
}
