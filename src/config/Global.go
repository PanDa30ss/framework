package config

import (
	log "github.com/PanDa30ss/core/logManager"
)

const (
	Master     = iota // value --> 0
	Cron              // value --> 1
	GameServer        // value --> 2
	Gate              // value --> 3
	Count             // value --> 3
)

func GetInstance() *config {
	once.Do(func() {
		instance = &config{}
	})
	return instance
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
