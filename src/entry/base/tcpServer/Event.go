package tcpServer

import (
	"github.com/PanDa30ss/core/event"
)

var EID_ServerSClose int
var EID_ServerSOpen int

func eventInitial() {
	EID_ServerSClose = event.GetEventID()
	EID_ServerSOpen = event.GetEventID()
}
