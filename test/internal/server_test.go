package internal

import (
	"testing"

	"github.com/xylong/gun/gnet"
)

func TestServer_Start(t *testing.T) {
	s:=gnet.NewServer("[gun v1.0.1]")
	s.Run()
}
