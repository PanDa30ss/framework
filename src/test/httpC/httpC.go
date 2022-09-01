package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	// "time"

	"github.com/PanDa30ss/core/http"
	"github.com/PanDa30ss/core/service"
	// "github.com/PanDa30ss/core/timeUtil"
)

func main() {

	service.Init()
	if !service.Start() {
		return
	}

	for i := 0; i < 10240; i++ {
		// for i := 0; i < 1; i++ {
		http.Get("http://127.0.0.1:9999/example?id="+strconv.Itoa(i), callback)
		http.Post("http://127.0.0.1:9999/example_post", "id="+strconv.Itoa(i), "", callback)
	}

	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt, os.Kill, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL, syscall.SIGHUP)
	<-sigs
}

func callback(result *http.HttpResult, params ...interface{}) {

	fmt.Println(result.Err)
	fmt.Println(result.Result)
}
