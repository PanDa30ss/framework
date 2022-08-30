package proto

// import "core/message"

// const (	//错误码
// 	ERROR_OK = iota
// 	ERROR_FAIL
// 	ERROR_ANOTHER_LOGIN
// )
// type ServerData struct {
// 	message.Package
// 	ServerID	uint32
//  	Roles	[]uint8
//  	Internal	string
//  }

// func MakeServerData() *ServerData {
// 	ret := &ServerData{}
// 	ret.ServerID = 0
// 	ret.Roles = nil
// 	ret.Internal = ""
// 	ret.Dump = func(msg *message.Message) {
// 		msg.Write(ret.ServerID)
// 		func() {
// 			l := uint16(len(ret.Roles))
// 			msg.Write(l)
// 			for i := 0; i < int(l); i++ {
// 				msg.Write(ret.Roles[i])
// 			}
// 		}()
// 		msg.WriteString(ret.Internal)
// 	}
// 	ret.Load = func(msg *message.Message) {
// 		msg.Read(&ret.ServerID)
// 		func() {
// 			var l uint16
// 			msg.Read(&l)
// 			ret.Roles = make([]uint8, l)
// 			for i := 0; i < int(l); i++ {
// 				ret.Roles[i] = 0
// 				msg.Read(&ret.Roles[i])
// 			}
// 		}()
// 		msg.ReadString(&ret.Internal)
// 	}
// 	return ret
// }

// var S_S_REGISTERSERVER_ACK uint16 = 1

// type P_S_S_RegisterServer_Ack struct {
// 	message.Package
// 	Servers	[]*ServerData
//  }

// func MakeP_S_S_RegisterServer_Ack() *P_S_S_RegisterServer_Ack {
// 	ret := &P_S_S_RegisterServer_Ack{}
// 	ret.Servers = nil
// 	ret.Dump = func(msg *message.Message) {
// 		func() {
// 			l := uint16(len(ret.Servers))
// 			msg.Write(l)
// 			for i := 0; i < int(l); i++ {
// 				msg.WritePackage(ret.Servers[i])
// 			}
// 		}()
// 	}
// 	ret.Load = func(msg *message.Message) {
// 		func() {
// 			var l uint16
// 			msg.Read(&l)
// 			ret.Servers = make([]*ServerData, l)
// 			for i := 0; i < int(l); i++ {
// 				ret.Servers[i] = MakeServerData()
// 				msg.ReadPackage(ret.Servers[i])
// 			}
// 		}()
// 	}
// 	return ret
// }

// var S_S_REGISTERSERVER_NTF uint16 = 2

// type P_S_S_RegisterServer_Ntf struct {
// 	message.Package
// 	Server	*ServerData
//  }

// func MakeP_S_S_RegisterServer_Ntf() *P_S_S_RegisterServer_Ntf {
// 	ret := &P_S_S_RegisterServer_Ntf{}
// 	ret.Server = nil
// 	ret.Dump = func(msg *message.Message) {
// 		msg.WritePackage(ret.Server)
// 	}
// 	ret.Load = func(msg *message.Message) {
// 		ret.Server = MakeServerData()
// 		msg.ReadPackage(ret.Server)
// 	}
// 	return ret
// }

// var S_S_REGISTERSERVER_REQ uint16 = 3

// type P_S_S_RegisterServer_Req struct {
// 	message.Package
// 	Server	*ServerData
//  }

// func MakeP_S_S_RegisterServer_Req() *P_S_S_RegisterServer_Req {
// 	ret := &P_S_S_RegisterServer_Req{}
// 	ret.Server = nil
// 	ret.Dump = func(msg *message.Message) {
// 		msg.WritePackage(ret.Server)
// 	}
// 	ret.Load = func(msg *message.Message) {
// 		ret.Server = MakeServerData()
// 		msg.ReadPackage(ret.Server)
// 	}
// 	return ret
// }
