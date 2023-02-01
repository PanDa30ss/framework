package tcpClient

import (
	"entry/base/config"
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

func (this *mTcpClient) Initial() {
	initial()
}

func (this *mTcpClient) Stop() {
	this.bank.Close()
}

func AddClient(data *config.Server) *ServerC {
	sc, ok := getInstance().serverCs[data.ServerID]
	if !ok {
		sc = newServerC(data)
		getInstance().serverCs[data.ServerID] = sc
		getInstance().bank.AddConnectSession(sc)
	}
	return sc
}

func IsStarted() bool {
	return getInstance().IsStarted()
}
