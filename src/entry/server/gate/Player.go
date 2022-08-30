package gate

import (
	"entry/base/proto/pb"
	"fmt"

	"github.com/PanDa30ss/core/tcp"

	"github.com/PanDa30ss/core/message"
)

var playerSessionID = tcp.GetSessionUID()

func init() {

}

type player struct {
	tcp.SessionS
	gameServerID uint32
	playerId     uint32
	status       int
}

func (this *player) OnClose() {
	fmt.Println("gate player OnClose")
	// fmt.Println("ServerS OnClose", this.GetConv())
}

func (this *player) OnOpen() {
	// fmt.Println("ServerS OnOpen", this.GetConv())
	// fmt.Println("player OnOpen")
}

func newPlayer() *player {
	this := &player{}
	this.gameServerID = 0
	this.playerId = 0
	this.SetFactory(playerSessionID)
	this.status = 0
	return this
}

func (this *player) GetServerID() uint32 {
	return this.gameServerID
}

func (this *player) SendPBMessage(cmd pb.CMD, pkg interface{}) {
	msg := message.MakePBMessage(uint16(cmd), pkg)
	this.SendMessage(msg)
}

func makeGameMessage(cmd int32, playerId uint32, src *message.Message) *message.Message {
	gMsg := message.MakeMessage()
	gMsg.SetID(uint16(cmd))
	gMsg.Write(playerId)
	gMsg.Write(src.GetBody())
	gMsg.Done()
	return gMsg
}
