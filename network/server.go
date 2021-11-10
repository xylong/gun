package network

import (
	"errors"
	"fmt"
	"net"

	"github.com/xylong/gun/iface"
)

// Server 服务
type Server struct {
	// 服务名称
	Name string
	// ip版本
	IPVersion string
	// ip地址
	IP string
	// 端口
	Port int
}

func NewServer(name string) iface.IServer {
	return &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      10000,
	}
}

// ClientCallBack 客户端回调
// todo 定义当前客户端连接所绑定的handle，以后应该由用户自定义handle方法
func ClientCallBack(conn *net.TCPConn, bytes []byte, size int) error {
	fmt.Println("[connection handle] callback to client")
	if _, err := conn.Write(bytes); err != nil {
		return errors.New("callback to client error")
	}

	return nil
}

// Start 启动服务
func (s *Server) Start() {
	go func() {
		//* 1.获取tcp的地址
		fmt.Printf("[Start] server listener at %s:%d\n", s.IP, s.Port)
		tcpAddr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp address error: ", err)
			return
		}

		//* 2.监听服务地址
		listener, err := net.ListenTCP(s.IPVersion, tcpAddr)
		if err != nil {
			fmt.Println("listen error: ", err)
			return
		}

		// 客户端🆔
		var clientID uint32 = 0

		//! 3.阻塞等待客户端连接，处理业务
		for {
			tcpConn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error: ", err)
				continue
			}

			connection := NewConnection(tcpConn, clientID, ClientCallBack)
			clientID++

			go connection.Start()
		}
	}()
}

func (s *Server) Stop() {

}

// Run 运行服务
func (s *Server) Run() {
	s.Start()

	// 阻塞
	select {}
}
