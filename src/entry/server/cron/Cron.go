package cron

import (
	"entry/base/config"
	"entry/base/tcpServer"
	"sync"
	"time"

	"github.com/PanDa30ss/core/service"
	"github.com/PanDa30ss/core/timeUtil"
)

type mCron struct {
	service.Module
	games  map[uint32]*tcpServer.ServerS
	ticker *timeUtil.Ticker
}

var instance *mCron
var once sync.Once

func getInstance() *mCron {
	once.Do(func() {
		instance = makeInstance()
	})
	return instance
}

func makeInstance() *mCron {
	ret := &mCron{}
	ret.games = make(map[uint32]*tcpServer.ServerS)
	return ret
}

func (this *mCron) Init() {
	config.Register(config.Cron)
}

func (this *mCron) Start() bool {
	this.ticker = timeUtil.MakeTicker(1*time.Second, runTest, this)
	return this.Module.Start()
}

func (this *mCron) Stop() {
	this.ticker.Stop()
}

func runTest(params ...interface{}) {
	// this := params[0].(*rCron)
	// pkg := &pb.Test{}
	// for _, gs := range this.games {
	// 	gs.SendPBMessage(pb.CMD_Test11, pkg)
	// }
}
