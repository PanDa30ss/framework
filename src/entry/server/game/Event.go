package game

import (
	"entry/base/config"
	"entry/base/tcpClient"
	"fmt"

	"github.com/PanDa30ss/core/event"
)

func eventInitial() {

	event.RegisterEventHandler(
		tcpClient.EID_ServerList,
		func(params ...interface{}) {
			fmt.Println("11111111111")
			list := params[0].([]*config.Server)
			for i := 0; i < len(list); i++ {
				if list[i].Roles[config.Cron] {
					getInstance().cronC = tcpClient.AddClient(list[i])
				}
				if list[i].Roles[config.Gate] {
					getInstance().gateCs[list[i].ServerID] = tcpClient.AddClient(list[i])
				}
			}
		})
}
