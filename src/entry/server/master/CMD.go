package master

import (
	"entry/base/config"
	"entry/base/proto/pb"
	"entry/base/tcpServer"

	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/tcp"
)

var _ = tcpServer.RegisterCMD(
	pb.CMD_REQ_RegisterServer,
	func(s tcp.ISession, msg *message.Message) bool {
		session := s.(*tcpServer.ServerS)
		req := &pb.Server{}
		msg.Unmarshal(req)
		session.Data = &config.Server{}
		session.Data.Load(req)
		slaves := getInstance().slaves
		roleMap := getInstance().roleMap
		//加入管理
		slaves[session.GetServerID()] = session
		for key, _ := range session.Data.Roles {
			m, ok := roleMap[key]
			if !ok {
				m = make(map[uint32]*tcpServer.ServerS)
				roleMap[key] = m
			}
			m[session.GetServerID()] = session
		}
		//回调 和 推送
		rsp := &pb.ServerList{}
		rsp.Servers = make([]*pb.Server, len(slaves))
		index := 0
		for serverID, serverS := range slaves {
			rsp.Servers[index] = &pb.Server{}
			serverS.Data.Dump(rsp.Servers[index])
			index++
			if serverID == session.Data.ServerID {
				continue
			}
			serverS.SendPBMessage(pb.CMD_NTF_RegisterServer, req)
		}
		session.SendPBMessage(pb.CMD_RSP_RegisterServer, rsp)
		return true
	})
