package master

import (
	"config"
	"sync"
	"tcpServer"

	"github.com/PanDa30ss/core/service"
)

type rMaster struct {
	service.Module
	slaves  map[uint32]*tcpServer.ServerS
	roleMap map[uint8]map[uint32]*tcpServer.ServerS
}

var instance *rMaster
var once sync.Once

func getInstance() *rMaster {
	once.Do(func() {
		instance = makeInstance()
	})
	return instance
}

func makeInstance() *rMaster {
	ret := &rMaster{}
	ret.slaves = make(map[uint32]*tcpServer.ServerS)
	ret.roleMap = make(map[uint8]map[uint32]*tcpServer.ServerS)
	return ret
}

func (this *rMaster) Init() {
	moduleInit()
	config.Register(config.Master)
}

func (this *rMaster) Start() bool {

	return this.Module.Start()
}

func (this *rMaster) Stop() {
}
