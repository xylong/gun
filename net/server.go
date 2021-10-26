package net

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

func NewServer(name string) *Server {
	return &Server{
		Name: name,
		IPVersion: "tcp4",
		IP: "0.0.0.0",
		Port: 10000,
	}
}

func (s *Server) Start() {
	
}

func (s *Server) Stop() {
	
}

func (s *Server) Run() {
	
}