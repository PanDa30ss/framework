package tcpClient

import (
	"config"
	"proto/pb"

	"github.com/PanDa30ss/core/event"
	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/tcp"
)

var ServerCID = tcp.GetSessionUID()

var EID_ServerList = event.GetEventID()
var EID_ServerCOpen = event.GetEventID()

func RegisterCMD(cmd pb.CMD, f func(tcp.ISession, *message.Message) bool) bool {
	return tcp.RegisterCMD(ServerCID, uint16(cmd), f)
}

func AddClient(data *config.Server) *ServerC {
	sc, ok := getInstance().serverCs[data.ServerID]
	if !ok {
		sc = newServerC(data)
		getInstance().serverCs[data.ServerID] = sc
		getInstance().bank.AddConnectSession(sc)
	}
	return sc
}
