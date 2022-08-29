package tcpServer

import (
	"config"
	"proto/pb"

	"github.com/PanDa30ss/core/event"
	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/tcp"
)

func cmdInit() {

	RegisterCMD(
		pb.CMD_REQ_RegisterRole,
		func(s tcp.ISession, msg *message.Message) bool {
			session := s.(*ServerS)
			req := &pb.Server{}
			msg.Unmarshal(req)
			session.Data = &config.Server{}
			session.Data.Load(req)
			event.DispatchEvent(event.MakeEvent(EID_ServerSOpen, session))
			return true
		})

}
