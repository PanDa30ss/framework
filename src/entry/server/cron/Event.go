package cron

import (
	"entry/base/config"
	"entry/base/tcpServer"

	"github.com/PanDa30ss/core/event"
)

var _ = event.RegisterEventHandler(
	tcpServer.EID_ServerSClose,
	func(params ...interface{}) {
		session := params[0].(*tcpServer.ServerS)

		if session.Data.Roles[config.GameServer] {
			s, ok := getInstance().games[session.GetServerID()]
			if !ok {
				return
			}
			if s != session {
				return
			}
			delete(getInstance().games, session.GetServerID())
		}

	})

var _ = event.RegisterEventHandler(
	tcpServer.EID_ServerSOpen,
	func(params ...interface{}) {
		session := params[0].(*tcpServer.ServerS)
		if session.Data.Roles[config.GameServer] {
			getInstance().games[session.GetServerID()] = session
		}
	})
