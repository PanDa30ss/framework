package redisBank

import (
	"entry/base/config"
	"sync"

	. "github.com/PanDa30ss/core/db/redis"
	"github.com/PanDa30ss/core/service"
)

const (
	Game  = iota // value --> 0
	Login        // value --> 1
)

type mRedisBank struct {
	service.Module
	dbs map[int]*Redis
}

var instance *mRedisBank
var once sync.Once

func getInstance() *mRedisBank {
	once.Do(func() {
		instance = makeInstance()
	})
	return instance
}

func makeInstance() *mRedisBank {
	ret := &mRedisBank{}
	ret.dbs = make(map[int]*Redis)
	return ret
}

func (this *mRedisBank) Init() {
}

func (this *mRedisBank) Start() bool {
	for k, r := range this.dbs {
		url, con, idle := config.GetRedis(k)
		r.SetUrl(url)
		r.SetMaxConn(con)
		r.SetMaxIdle(idle)
		r.Open()
	}
	return this.Module.Start()
}

func (this *mRedisBank) Stop() {
	for _, r := range this.dbs {
		r.Close()
	}
}

func Register(key int) *Redis {
	r, ok := getInstance().dbs[key]
	if ok {
		return r
	}
	r = &Redis{}
	r.InitDB()
	getInstance().dbs[key] = r
	return r
}
