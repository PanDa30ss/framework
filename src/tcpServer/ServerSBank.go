package tcpServer

import (
	"github.com/PanDa30ss/core/tcp"
)

type serverSBank struct {
	tcp.SessionSBank
}

func (this *serverSBank) CreateSession() tcp.ISessionS {
	return newServerS()
}

func makeServerSBank() *serverSBank {
	s := &serverSBank{}
	return s
}
