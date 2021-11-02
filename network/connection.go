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
	// 业务处理函数
	handle   iface.HandleFunc
	ExitChan chan bool
}

func NewConnection(conn *net.TCPConn, connID uint32, callback iface.HandleFunc) *Connection {
	return &Connection{
		Conn:     conn,
		ConnID:   connID,
		isClosed: false,
		handle:   callback,
		ExitChan: make(chan bool, 1),
	}
}

// Read 读取数据
func (c *Connection) Read() {
	fmt.Println("read goroutine is running")
	defer fmt.Printf("ConnID: %d read exit, remote address is %s \n", c.ConnID, c.GetRemoteAddr().String())
	defer c.Stop()

	for {
		buf := make([]byte, 512)
		size, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("receive buf error:", err)
			continue
		}

		//* 业务处理
		if err := c.handle(c.Conn, buf, size); err != nil {
			fmt.Println("ConnID:", c.ConnID, "handle error:", err)
			break
		}
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
func (c *Connection) Send() {

}
