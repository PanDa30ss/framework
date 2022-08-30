package gate

import (
	"entry/base/config"
	"entry/base/tcpServer"
	"sync"

	"github.com/PanDa30ss/core/service"
)

type mGate struct {
	service.Module
	playerBank *playerBank
	games      map[uint32]*tcpServer.ServerS
	players    map[uint32]*player
}

var instance *mGate
var once sync.Once

func getInstance() *mGate {
	once.Do(func() {
		instance = makeInstance()
	})
	return instance
}

func makeInstance() *mGate {
	ret := &mGate{}
	ret.games = make(map[uint32]*tcpServer.ServerS)
	ret.playerBank = makePlayerBank()
	ret.players = make(map[uint32]*player)
	return ret
}

func (this *mGate) Init() {
	config.Register(config.Gate)
	this.playerBank.Init(this.playerBank)
	this.playerBank.BindAddr(config.GetString("foreign"))
	this.playerBank.SetMaxSession(0x7FFFFFFF)
}

func (this *mGate) Start() bool {
	this.playerBank.Start()
	return this.Module.Start()
}

func (this *mGate) Stop() {
	this.playerBank.Close()
}

func (this *mGate) assignGameServer() uint32 {
	for a, _ := range this.games {
		return a
	}
	return 0
}
