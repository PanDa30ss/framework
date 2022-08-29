package game

import (
	"proto/pb"

	"github.com/PanDa30ss/core/message"
)

type player struct {
	playerID uint32
	gate     uint32
}

func makePlayer(playerID uint32) *player {
	ret := &player{}
	ret.playerID = playerID
	ret.gate = 0
	return ret
}

func (this *player) sendMessage(msg *message.Message) {
	if this.gate == 0 {
		return
	}
	gate, ok := getInstance().gateCs[this.gate]
	if !ok {
		return
	}
	gate.SendMessage(msg)
}

func (this *player) sendPBMessage(cmd pb.CMD, err uint16, pkg interface{}) {
	msg := message.MakeMessage()
	msg.SetID(uint16(cmd))
	msg.Write(this.playerID)
	msg.Write(err)
	if pkg != nil {
		msg.Marshal(pkg)
	}
	msg.Done()
	this.sendMessage(msg)
}
