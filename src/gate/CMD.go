package gate

import (
	. "proto/pb"
	"tcpServer"

	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/tcp"
)

func cmdInit() {
	for key, _ := range CMD_name {
		if key > int32(CMD_PlayerStart) && key < int32(CMD_PlayerEnd) {
			tcpServer.RegisterCMD(
				CMD(key),
				func(s tcp.ISession, msg *message.Message) bool {
					playerID := msg.ReadUInt32()
					if playerID == 0 {
						return true
					}
					p, ok := getInstance().players[playerID]
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
			playerID := msg.ReadUInt32()
			err := msg.ReadUInt16()
			if playerID == 0 {
				return true
			}
			p, ok := getInstance().players[playerID]
			if !ok {
				return true
			}
			if p.gameServerID != session.Data.ServerID {
				return true
			}

			if err == 0 {
				p.status = 2
			}
			p.SendMessage(msg)
			return true
		})
}
