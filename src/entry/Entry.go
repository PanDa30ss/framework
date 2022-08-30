package main

import (
	"entry/base/config"

	log "github.com/PanDa30ss/core/logManager"

	_ "entry/base/modules"
	_ "entry/server/modules"

	"os"
	"os/signal"
	"syscall"

	"github.com/PanDa30ss/core/service"
	// _ "net/http/pprof"
)

func main() {

	if !loadArgs() {
		log.Error("entry args error")
		return
	}

	service.Init()

	if !service.SelectModule(config.GetModules()) {
		log.Error("lack of module")
		return
	}
	if !service.Start() {
		return
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)
	<-sigs
	service.Stop()
}

func loadArgs() bool {
	argPath := ""
	version := "debug"
	length := len(os.Args)
	switch length {
	case 2:
		argPath = os.Args[1]
	case 3:
		argPath = os.Args[1]
		version = os.Args[2]
	default:
	}
	if !config.GetInstance().Load(argPath) {
		log.Error("read config.json fail")
		return false
	}
	log.SetLevel(version)
	return true
}
