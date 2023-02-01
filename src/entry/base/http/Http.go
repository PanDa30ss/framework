package http

import (
	"entry/base/config"
	"sync"

	"github.com/PanDa30ss/core/http"
	"github.com/PanDa30ss/core/service"
)

type mHttp struct {
	service.Module
	server *http.HttpServer
}

var instance *mHttp
var once sync.Once

func getInstance() *mHttp {
	once.Do(func() {
		instance = makeInstance()
	})
	return instance
}

func makeInstance() *mHttp {
	ret := &mHttp{}
	ret.server = http.MakeHttpServer()
	return ret
}

func (this *mHttp) Init() {
	this.server.Init(config.GetString("http"))
}

func (this *mHttp) Start() bool {
	if !this.server.Start() {
		return false
	}
	return this.Module.Start()
}

func (this *mHttp) Initial() {
	initial()
}

func IsStarted() bool {
	return getInstance().IsStarted()
}

func (this *mHttp) register(url string, handleFunc http.HttpHandleFunc) bool {
	return this.server.Register(url, handleFunc)
}

func Register(url string, handleFunc http.HttpHandleFunc) bool {
	return getInstance().register(url, handleFunc)
}
