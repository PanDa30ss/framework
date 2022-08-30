package tcpServer

import (
	"entry/base/config"
	"sync"

	"github.com/PanDa30ss/core/service"
)

type mTcpServer struct {
	service.Module
	server *serverSBank
}

var instance *mTcpServer
var once sync.Once

func getInstance() *mTcpServer {
	once.Do(func() {
		instance = makeInstance()
	})
	return instance
}

func makeInstance() *mTcpServer {
	ret := &mTcpServer{}
	ret.server = makeServerSBank()
	return ret
}

func (this *mTcpServer) Init() {
	this.server.Init(this.server)
	this.server.BindAddr(config.GetString("internal"))
	this.server.SetMaxSession(0x7FFFFFFF)
}

func (this *mTcpServer) Start() bool {
	this.server.Start()
	return this.Module.Start()
}

func (this *mTcpServer) Stop() {
	this.server.Close()
}
