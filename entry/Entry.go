package main

import (
	"config"
	log "core/logManager"
	"core/service"
	_ "moduleMgr"
	"os"
	"os/signal"
	"syscall"

	_ "net/http/pprof"
)

func main() {
	// go func() {
	// 	http.ListenAndServe("0.0.0.0:6060", nil)
	// }()
	// bytesBuffer := bytes.NewBuffer([]byte{})
	// for i := 0; i < 10; i++ {
	// 	bytesBuffer.WriteByte(0)
	// }
	// fmt.Println(bytesBuffer.Bytes())
	// bytesBuffer.Bytes()[0] = 1
	// fmt.Println(bytesBuffer.Bytes())
	// var a uint32 = 1
	// binary.BigEndian.PutUint32(bytesBuffer.Bytes()[5:], a)
	// fmt.Println(bytesBuffer.Bytes())
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
