package game

import (
	. "entry/base/proto/pb"
	"entry/base/tcpClient"

	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/tcp"
)

func cmdInit() {
	registerCMD(
		CMD_REQ_Test,
		func(s tcp.ISession, msg *message.Message) bool {
			playerID := msg.ReadUInt32()
			pkg := &Test{}
			msg.Unmarshal(pkg)
			p := getInstance().players[playerID]
			p.sendPBMessage(CMD_RSP_Test, 0, pkg)
			return true
		})

	registerCMD(
		CMD_REQ_EnterGame,
		func(s tcp.ISession, msg *message.Message) bool {
			session := s.(*tcpClient.ServerC)
			pkg := &EnterGame{}
			msg.Unmarshal(pkg)

			p := getInstance().addPlayer(pkg.GetTestID())
			p.gate = session.Data.ServerID
			p.sendPBMessage(CMD_RSP_EnterGame, 0, pkg)
			return true
		})
}
