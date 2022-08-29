package cron

import (
	"config"
	"sync"
	"tcpServer"
	"time"

	"github.com/PanDa30ss/core/service"
	"github.com/PanDa30ss/core/timeUtil"
)

type rCron struct {
	service.Module
	games  map[uint32]*tcpServer.ServerS
	ticker *timeUtil.Ticker
}

var instance *rCron
var once sync.Once

func getInstance() *rCron {
	once.Do(func() {
		instance = makeInstance()
	})
	return instance
}

func makeInstance() *rCron {
	ret := &rCron{}
	ret.games = make(map[uint32]*tcpServer.ServerS)
	return ret
}

func (this *rCron) Init() {
	config.Register(config.Cron)
	moduleInit()
}

func (this *rCron) Start() bool {
	this.ticker = timeUtil.MakeTicker(1*time.Second, runTest, this)
	return this.Module.Start()
}

func (this *rCron) Stop() {
	this.ticker.Stop()
}

func runTest(params ...interface{}) {
	// this := params[0].(*rCron)
	// pkg := &pb.Test{}
	// for _, gs := range this.games {
	// 	gs.SendPBMessage(pb.CMD_Test11, pkg)
	// }
}
