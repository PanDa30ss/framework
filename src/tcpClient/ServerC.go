package tcpClient

import (
	"config"
	"fmt"
	"proto/pb"

	"github.com/PanDa30ss/core/tcp"

	"github.com/PanDa30ss/core/message"
)

type ServerC struct {
	tcp.SessionC
	Data *config.Server
}

func (this *ServerC) OnClose() {
	fmt.Println("ServerC OnClose")
}

func (this *ServerC) OnOpen() {
	fmt.Println("ServerC OnOpen")
	pkg := &pb.Server{}
	config.GetInstance().Server.Dump(pkg)
	this.SendPBMessage(pb.CMD_REQ_RegisterRole, pkg)
}

func (this *ServerC) SendPBMessage(cmd pb.CMD, pkg interface{}) {
	msg := message.MakePBMessage(uint16(cmd), pkg)
	this.SendMessage(msg)
}

func newServerC(data *config.Server) *ServerC {
	ic := &ServerC{}
	ic.Data = data
	ic.SetAddress(data.Internal)
	ic.BindPingManager(tcp.DefaultPingTimeData)
	ic.SetFactory(ServerCID)
	return ic
}
