package tcpClient

import (
	"entry/base/config"
	"entry/base/proto/pb"
	"fmt"

	"github.com/PanDa30ss/core/tcp"

	"github.com/PanDa30ss/core/message"
)

type MasterC struct {
	tcp.SessionC
}

func (this *MasterC) OnClose() {
	fmt.Println("MasterC OnClose")
}

func (this *MasterC) OnOpen() {
	fmt.Println("MasterC OnOpen")
	pkg := &pb.Server{}
	config.GetInstance().Server.Dump(pkg)
	this.SendPBMessage(pb.CMD_REQ_RegisterServer, pkg)
}

func (this *MasterC) SendPBMessage(cmd pb.CMD, pkg interface{}) {
	msg := message.MakePBMessage(uint16(cmd), pkg)
	this.SendMessage(msg)
}

func newMasterC(addr string) *MasterC {
	ic := &MasterC{}
	ic.SetAddress(addr)
	ic.BindPingManager(tcp.DefaultPingTimeData)
	ic.SetFactory(ServerCID)
	return ic
}
