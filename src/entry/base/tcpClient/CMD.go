package tcpClient

import (
	"entry/base/config"
	"entry/base/proto/pb"

	"github.com/PanDa30ss/core/event"
	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/tcp"
)

func cmdInitial() {
	RegisterCMD(
		pb.CMD_RSP_RegisterServer,
		func(s tcp.ISession, msg *message.Message) bool {
			req := &pb.ServerList{}
			msg.Unmarshal(req)
			var lists []*config.Server = nil
			if req.Servers != nil {
				l := len(req.Servers)
				lists = make([]*config.Server, l)
				for i := 0; i < l; i++ {
					lists[i] = &config.Server{}
					lists[i].Load(req.Servers[i])
				}

			}
			event.DispatchEvent(event.MakeEvent(EID_ServerList, lists))

			return true
		})

	RegisterCMD(
		pb.CMD_NTF_RegisterServer,
		func(s tcp.ISession, msg *message.Message) bool {
			req := &pb.Server{}
			msg.Unmarshal(req)
			lists := make([]*config.Server, 1)
			lists[0] = &config.Server{}
			lists[0].Load(req)

			event.DispatchEvent(event.MakeEvent(EID_ServerList, lists))

			return true
		})
}
