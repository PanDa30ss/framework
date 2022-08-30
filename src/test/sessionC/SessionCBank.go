package main

import (
	"entry/base/proto/pb"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/service"
	"github.com/PanDa30ss/core/tcp"
	"github.com/PanDa30ss/core/timeUtil"
)

var SessionCID = tcp.GetSessionUID()

type SessionC struct {
	tcp.SessionC
	index  uint32
	ticker *timeUtil.Ticker
}

func (this *SessionC) OnClose() {
	fmt.Println("SessionC OnClose", this.index)
	if this.ticker != nil {
		this.ticker.Stop()
	}
}

func (this *SessionC) OnOpen() {
	fmt.Println("SessionC OnOpen", this.index)
	pkg := &pb.EnterGame{}
	pkg.TestID = &this.index
	this.SendPBMessage(pb.CMD_REQ_EnterGame, pkg)
	this.ticker = timeUtil.MakeTicker(1*time.Second, runTest, this)
}
func (this *SessionC) SendPBMessage(cmd pb.CMD, pkg interface{}) {
	msg := message.MakePBMessage(uint16(cmd), pkg)
	this.SendMessage(msg)
}

var _ = tcp.RegisterCMD(SessionCID, uint16(pb.CMD_RSP_EnterGame), func(s tcp.ISession, msg *message.Message) bool {
	pkg := &pb.Test{}
	pkg.ServerID = &s.(*SessionC).index
	s.(*SessionC).SendPBMessage(pb.CMD_REQ_Test, pkg)

	return true
})
var _ = tcp.RegisterCMD(SessionCID, uint16(pb.CMD_RSP_Test), func(s tcp.ISession, msg *message.Message) bool {
	msg.ReadUInt32()
	msg.ReadUInt16()
	pkg := &pb.Test{}
	msg.Unmarshal(pkg)
	fmt.Println(pkg)

	return true
})

func runTest(params ...interface{}) {
	this := params[0].(*SessionC)
	pkg := &pb.Test{}
	pkg.ServerID = &this.index
	this.SendPBMessage(pb.CMD_REQ_Test, pkg)
}
func newSessionC(addr string) *SessionC {
	ic := &SessionC{}
	ic.SetAddress(addr)
	ic.BindPingManager(tcp.DefaultPingTimeData)
	ic.SetFactory(SessionCID)
	return ic
}

func main() {

	service.Init()
	if !service.Start() {
		return
	}
	bank := &tcp.SessionCBank{}
	bank.Init()
	bank.Start()
	var i uint32 = 0
	for ; i < 10240; i++ {
		// for ; i < 10240; i++ {
		session := newSessionC("127.0.0.1:9997")
		session.index = i + 100
		bank.AddConnectSession(session)
	}
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)
	<-sigs
}
