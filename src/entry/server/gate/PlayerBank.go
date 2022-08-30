package gate

import (
	"github.com/PanDa30ss/core/tcp"
)

type playerBank struct {
	tcp.SessionSBank
}

func (this *playerBank) CreateSession() tcp.ISessionS {
	return newPlayer()
}

func makePlayerBank() *playerBank {
	s := &playerBank{}
	return s
}
