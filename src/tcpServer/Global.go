package tcpServer

import (
	"proto/pb"

	"github.com/PanDa30ss/core/event"
	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/tcp"
)

var ServerSID = tcp.GetSessionUID()

var EID_ServerSClose = event.GetEventID()
var EID_ServerSOpen = event.GetEventID()

func RegisterCMD(cmd pb.CMD, f func(tcp.ISession, *message.Message) bool) bool {
	return tcp.RegisterCMD(ServerSID, uint16(cmd), f)
}
