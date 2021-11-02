package network

import (
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

		//! 3.阻塞等待客户端连接，处理业务
		for {
			tcpConn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error: ", err)
				continue
			}

			go func() {
				for {
					buf := make([]byte, 512)
					size, err := tcpConn.Read(buf)
					if err != nil {
						fmt.Println("receive buf error: ", err)
						continue
					}

					if _, err := tcpConn.Write(buf[:size]); err != nil {
						fmt.Println("send error: ", err)
						continue
					}
				}
			}()
		}
	}()
}

func (s *Server) Stop() {

}

// Run 运行服务
func (s *Server) Run() {
	s.Start()

	// 阻塞
	select{}
}
