package gate

import (
	"config"
	"sync"
	"tcpServer"

	"github.com/PanDa30ss/core/service"
)

type rGate struct {
	service.Module
	playerBank *playerBank
	games      map[uint32]*tcpServer.ServerS
	players    map[uint32]*player
}

var instance *rGate
var once sync.Once

func getInstance() *rGate {
	once.Do(func() {
		instance = makeInstance()
	})
	return instance
}

func makeInstance() *rGate {
	ret := &rGate{}
	ret.games = make(map[uint32]*tcpServer.ServerS)
	ret.playerBank = makePlayerBank()
	ret.players = make(map[uint32]*player)
	return ret
}

func (this *rGate) Init() {
	moduleInit()
	config.Register(config.Gate)
	this.playerBank.Init(this.playerBank)
	this.playerBank.BindAddr(config.GetString("foreign"))
	this.playerBank.SetMaxSession(0x7FFFFFFF)
}

func (this *rGate) Start() bool {
	this.playerBank.Start()
	return this.Module.Start()
}

func (this *rGate) Stop() {
	this.playerBank.Close()
}

func (this *rGate) assignGameServer() uint32 {
	for a, _ := range this.games {
		return a
	}
	return 0
}
