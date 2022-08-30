package gate

import (
	. "entry/base/proto"
	. "entry/base/proto/pb"

	"github.com/PanDa30ss/core/message"

	"github.com/PanDa30ss/core/tcp"
)

func init() {
	for key, _ := range CMD_name {
		if key > int32(CMD_PlayerStart) && key < int32(CMD_PlayerEnd) {
			playerCmd(key)

		}
	}

}

func playerCmd(cmd int32) bool {
	return tcp.RegisterCMD(
		playerSessionID,
		uint16(cmd),
		func(s tcp.ISession, msg *message.Message) bool {
			if s.(*player).status != 2 {
				return true
			}
			if s.(*player).playerId == 0 {
				return true
			}
			if s.(*player).gameServerID == 0 {
				return true
			}
			var game, ok = getInstance().games[s.(*player).gameServerID]
			if !ok {
				return true
			}
			head := MakeGate2GameHead()
			head.PlayerID = s.(*player).playerId
			game.SendMessage(makeGameMessage(cmd, head, msg))
			return true
		})
}

var _ = tcp.RegisterCMD(
	playerSessionID,
	uint16(CMD_REQ_EnterGame),
	func(s tcp.ISession, msg *message.Message) bool {
		pkg := &EnterGame{}
		msg.Unmarshal(pkg)
		gameServerID := getInstance().assignGameServer()
		if gameServerID == 0 {
			return false
		}
		playerId := pkg.GetTestID()
		p := s.(*player)
		p.gameServerID = gameServerID
		p.playerId = playerId
		p.status = 1
		// curP, ok := getInstance().players[playerId]
		// if ok {
		// 	// fmt.Println(curP)
		// }
		getInstance().players[playerId] = p
		var game = getInstance().games[gameServerID]
		game.SendMessage(msg)
		return true
	})
