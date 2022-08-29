package master

import (
	"tcpServer"

	"github.com/PanDa30ss/core/event"
)

func eventInit() {
	event.RegisterEventHandler(
		tcpServer.EID_ServerSClose,
		func(params ...interface{}) {
			session := params[0].(*tcpServer.ServerS)
			s, ok := getInstance().slaves[session.GetServerID()]
			if !ok {
				return
			}
			if s != session {
				return
			}
			delete(getInstance().slaves, session.GetServerID())
		})
}
