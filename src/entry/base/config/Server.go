package config

import "entry/base/proto/pb"

const (
	Master = iota
	Cron
	GameServer
	Gate
	Web
	Count
)

type Server struct {
	ServerID uint32
	Internal string
	Roles    map[uint8]bool
}

func (this *Server) Load(src *pb.Server) {
	this.ServerID = src.GetServerID()
	this.Internal = src.GetInternal()
	var rs = src.GetRoles()
	this.Roles = make(map[uint8]bool)
	for i := 0; i < len(rs); i++ {
		this.Roles[rs[i]] = true
	}
}

func (this *Server) Dump(src *pb.Server) {
	src.ServerID = &this.ServerID
	src.Internal = &this.Internal
	src.Roles = make([]byte, len(this.Roles))
	index := 0
	for key, _ := range this.Roles {
		src.Roles[index] = key
		index++
	}
}
