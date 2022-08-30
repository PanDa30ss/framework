package config

import (
	"io/ioutil"
	"os"
	"sync"

	log "github.com/PanDa30ss/core/logManager"

	"github.com/spyzhov/ajson"
)

var debugConfigPath = "../../conf"
var defaultconfigPath = "../conf"
var configFile = "/config.json"
var moduleFile = "/module.json"
var instance *config
var once sync.Once

func GetInstance() *config {
	once.Do(func() {
		instance = &config{}
	})
	return instance
}

type config struct {
	data    *ajson.Node
	modules map[string]bool
	Server  *Server
}

func (this *config) Load(argPath string) (ret bool) {

	ret = false
	if argPath == "" {
		argPath = defaultconfigPath
	}
	file, err := os.Open(argPath + configFile)
	if err != nil {
		argPath = debugConfigPath

		file, err = os.Open(argPath + configFile)
		if err != nil {
			return
		}
	}

	defer file.Close()
	data, err1 := ioutil.ReadAll(file)
	if err1 != nil {
		return
	}

	ret = this.loadData(&data)
	return
}

func (this *config) loadData(data *[]byte) bool {
	root, err1 := ajson.Unmarshal(*data)
	if err1 != nil {
		log.Error("解析config.json 失败")
		return false
	}
	this.data = root
	ret := this.init()
	if !ret {
		log.Error("解析config.json 失败")
		return false
	}
	return true
}

func (this *config) init() (ret bool) {
	defer func() {
		if e := recover(); e != nil {
			log.Error(e)
			ret = false
		}
	}()
	ret = false
	if this.data == nil {
		return
	}

	this.modules = make(map[string]bool)
	this.initBaseModule()
	for _, object := range this.data.MustKey("modules").MustArray() {
		this.modules[object.MustString()] = true
	}
	this.Server = &Server{}
	this.Server.Roles = make(map[uint8]bool)
	this.Server.Internal = GetString("internal")
	this.Server.ServerID = uint32(GetInt("serverID"))
	if this.Server.ServerID <= 0 {
		return
	}
	ret = true
	return
}
func (this *config) initBaseModule() {
	this.modules["tcpserver"] = true
	this.modules["tcpclient"] = true
	this.modules["http"] = true
}

func Register(role int) {
	GetInstance().Server.Roles[uint8(role)] = true
}

func GetModules() map[string]bool {
	return GetInstance().modules
}

func GetString(key string) string {
	defer func() {
		if e := recover(); e != nil {
			log.Error(e)
		}
	}()
	return GetInstance().data.MustKey(key).MustString()
}

func GetInt(key string) int {
	defer func() {
		if e := recover(); e != nil {
			log.Error(e)
		}
	}()
	ret := 0
	ret = int(GetInstance().data.MustKey(key).MustNumeric())
	return ret
}

func GetBool(key string) bool {
	defer func() {
		if e := recover(); e != nil {
			log.Error(e)
		}
	}()
	return GetInstance().data.MustKey(key).MustBool()
}
