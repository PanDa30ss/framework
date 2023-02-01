package gate

import (
	. "entry/base/proto"
	. "entry/base/proto/pb"
	"entry/base/tcpServer"

	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/tcp"
)

func cmdInitial() {
	for key, _ := range CMD_name {
		if key > int32(CMD_PlayerStart) && key < int32(CMD_PlayerEnd) {
			tcpServer.RegisterCMD(
				CMD(key),
				func(s tcp.ISession, msg *message.Message) bool {
					head := MakeGame2GateHead()
					head.Load(msg)
					if head.PlayerID == 0 {
						return true
					}
					p, ok := getInstance().players[head.PlayerID]
					if !ok {
						return true
					}
					p.SendMessage(msg)
					return true
				})
		}
	}
	tcpServer.RegisterCMD(
		CMD_RSP_EnterGame,
		func(s tcp.ISession, msg *message.Message) bool {
			session := s.(*tcpServer.ServerS)
			head := MakeGame2GateHead()
			head.Load(msg)
			if head.PlayerID == 0 {
				return true
			}
			p, ok := getInstance().players[head.PlayerID]
			if !ok {
				return true
			}
			if p.gameServerID != session.Data.ServerID {
				return true
			}

			if head.Err == 0 {
				p.status = 2
			}
			p.SendMessage(msg)
			return true
		})
}
