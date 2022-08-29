package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/PanDa30ss/core/service"
	"github.com/PanDa30ss/core/tcp"
)

var SessionSID = tcp.GetSessionUID()

type SessionS struct {
	tcp.SessionS
}

func (this *SessionS) OnClose() {
	fmt.Println("SessionS OnClose")
}

func (this *SessionS) OnOpen() {
	fmt.Println("SessionS OnOpen")
}

func newSessionS() *SessionS {
	us := &SessionS{}
	us.SetFactory(SessionSID)
	return us
}

type SessionSBank struct {
	tcp.SessionSBank
}

func (this *SessionSBank) CreateSession() tcp.ISessionS {
	return newSessionS()
}

func MakeSessionSBank() *SessionSBank {
	s := &SessionSBank{}
	return s
}

func main() {

	service.Init()
	if !service.Start() {
		return
	}

	bank := &SessionSBank{}
	bank.BindAddr("127.0.0.1:9998")
	bank.Init(bank)
	bank.Start()

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)
	<-sigs
}
