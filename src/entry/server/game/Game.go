package game

import (
	"entry/base/config"
	"entry/base/proto/pb"
	"entry/base/tcpClient"
	"sync"

	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/service"
	"github.com/PanDa30ss/core/tcp"
)

type rGame struct {
	service.Module
	cronC   *tcpClient.ServerC
	gateCs  map[uint32]*tcpClient.ServerC
	players map[uint32]*player
}

var instance *rGame
var once sync.Once

func getInstance() *rGame {
	once.Do(func() {
		instance = makeInstance()
	})
	return instance
}

func makeInstance() *rGame {
	ret := &rGame{}
	ret.gateCs = make(map[uint32]*tcpClient.ServerC)
	ret.players = make(map[uint32]*player)
	return ret
}

func (this *rGame) Init() {
	moduleInit()
	config.Register(config.GameServer)
}

func (this *rGame) Start() bool {

	return this.Module.Start()
}

func (this *rGame) Stop() {
}

func (this *rGame) addPlayer(playerID uint32) *player {
	p, ok := this.players[playerID]
	if ok {
		return p
	}
	p = makePlayer(playerID)
	this.players[playerID] = p
	return p
}

func registerCMD(cmd pb.CMD, f func(tcp.ISession, *message.Message) bool) bool {
	return tcpClient.RegisterCMD(cmd, f)
}
