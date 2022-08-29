package tcpServer

import (
	"config"
	"fmt"
	"proto/pb"

	"github.com/PanDa30ss/core/event"
	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/tcp"
)

type ServerS struct {
	tcp.SessionS
	Data *config.Server
}

func (this *ServerS) OnClose() {
	fmt.Println("ServerS OnClose")
	if this.Data == nil {
		return
	}
	event.DispatchEvent(event.MakeEvent(EID_ServerSClose, this))
	// fmt.Println("ServerS OnClose", this.GetConv())
}

func (this *ServerS) OnOpen() {
	// fmt.Println("ServerS OnOpen", this.GetConv())
	fmt.Println("ServerS OnOpen")
}

func newServerS() *ServerS {
	us := &ServerS{}
	us.SetFactory(ServerSID)
	return us
}

func (this *ServerS) GetServerID() uint32 {
	if this.Data == nil {
		return 0
	}
	return this.Data.ServerID
}

func (this *ServerS) SendPBMessage(cmd pb.CMD, pkg interface{}) {
	msg := message.MakePBMessage(uint16(cmd), pkg)
	this.SendMessage(msg)
}
