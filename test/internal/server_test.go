package internal

import (
	"fmt"
	"net"
	"testing"
	"time"

	"github.com/xylong/gun/gnet"
)

func TestServer_Start(t *testing.T) {
	s := gnet.NewServer("[gun v1.0.1]")
	s.Run()
}

func TestClient(t *testing.T) {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("client dial error: ", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hello world"))
		if err != nil {
			fmt.Println("client send error: ", err)
			return
		}

		buf := make([]byte, 512)
		size, err := conn.Read(buf)
		if err != nil {
			fmt.Println("client read error: ", err)
			return
		}

		fmt.Printf("server msg:%s(size=%d)\n", buf, size)
		time.Sleep(time.Second * 1)
	}
}
