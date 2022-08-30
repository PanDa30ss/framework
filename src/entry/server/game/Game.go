package game

import (
	"entry/base/config"
	"entry/base/proto/pb"
	"entry/base/redisBank"
	"entry/base/tcpClient"

	"sync"

	"github.com/PanDa30ss/core/db/redis"
	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/service"
	"github.com/PanDa30ss/core/tcp"
)

type mGame struct {
	service.Module
	cronC     *tcpClient.ServerC
	gateCs    map[uint32]*tcpClient.ServerC
	players   map[uint32]*player
	gameRedis *redis.Redis
}

var instance *mGame
var once sync.Once

func getInstance() *mGame {
	once.Do(func() {
		instance = makeInstance()
	})
	return instance
}

func makeInstance() *mGame {
	ret := &mGame{}
	ret.gateCs = make(map[uint32]*tcpClient.ServerC)
	ret.players = make(map[uint32]*player)
	return ret
}

func (this *mGame) Init() {
	config.Register(config.GameServer)
	this.gameRedis = redisBank.Register(redisBank.Game)
}

func (this *mGame) Start() bool {

	return this.Module.Start()
}

func (this *mGame) Stop() {
}

func (this *mGame) addPlayer(playerID uint32) *player {
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
