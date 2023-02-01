package tcpClient

import (
	"github.com/PanDa30ss/core/event"
)

var EID_ServerList int
var EID_ServerCOpen int

func eventInitial() {
	EID_ServerList = event.GetEventID()
	EID_ServerCOpen = event.GetEventID()
}
