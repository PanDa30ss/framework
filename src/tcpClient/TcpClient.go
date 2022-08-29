package tcpClient

import (
	"config"
	"sync"

	"github.com/PanDa30ss/core/service"
	"github.com/PanDa30ss/core/tcp"
)

type mTcpClient struct {
	service.Module
	masterC  *MasterC
	bank     *tcp.SessionCBank
	serverCs map[uint32]*ServerC
}

var instance *mTcpClient
var once sync.Once

func getInstance() *mTcpClient {
	once.Do(func() {
		instance = makeInstance()
	})
	return instance
}

func makeInstance() *mTcpClient {
	ret := &mTcpClient{}
	ret.bank = &tcp.SessionCBank{}
	ret.serverCs = make(map[uint32]*ServerC)
	return ret
}

func (this *mTcpClient) Init() {
	moduleInit()
	this.bank.Init()
}

func (this *mTcpClient) Start() bool {
	m := false
	roles := config.GetInstance().Server.Roles
	for key, _ := range roles {
		if key != config.Master {
			m = true
			break
		}
	}
	if m {
		this.masterC = newMasterC(config.GetString("master"))
		this.bank.AddConnectSession(this.masterC)
	}
	this.bank.Start()
	return this.Module.Start()
}

func (this *mTcpClient) Stop() {
	this.bank.Close()
}
