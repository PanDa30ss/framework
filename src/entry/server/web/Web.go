package web

import (
	"entry/base/config"
	"entry/base/http"
	"sync"

	"github.com/PanDa30ss/core/service"
)

type mWeb struct {
	service.Module
}

var instance *mWeb
var once sync.Once

func getInstance() *mWeb {
	once.Do(func() {
		instance = makeInstance()
	})
	return instance
}

func makeInstance() *mWeb {
	ret := &mWeb{}
	return ret
}

func (this *mWeb) Init() {
	config.Register(config.Web)

}

func (this *mWeb) Start() bool {

	return this.Module.Start()
}

func (this *mWeb) Initial() {
	initial()
}

func (this *mWeb) Stop() {
}

func (this *mWeb) CheckStart() bool {
	return http.IsStarted()
}
