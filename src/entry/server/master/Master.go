package master

import (
	"entry/base/config"
	"entry/base/tcpServer"
	"sync"

	"github.com/PanDa30ss/core/service"
)

type mMaster struct {
	service.Module
	slaves  map[uint32]*tcpServer.ServerS
	roleMap map[uint8]map[uint32]*tcpServer.ServerS
}

var instance *mMaster
var once sync.Once

func getInstance() *mMaster {
	once.Do(func() {
		instance = makeInstance()
	})
	return instance
}

func makeInstance() *mMaster {
	ret := &mMaster{}
	ret.slaves = make(map[uint32]*tcpServer.ServerS)
	ret.roleMap = make(map[uint8]map[uint32]*tcpServer.ServerS)
	return ret
}

func (this *mMaster) Init() {
	config.Register(config.Master)

}

func (this *mMaster) Start() bool {

	return this.Module.Start()
}

func (this *mMaster) Stop() {
}
