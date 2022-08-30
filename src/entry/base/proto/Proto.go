package proto
import "github.com/PanDa30ss/core/message"

type Gate2GameHead struct {
	message.Package
	PlayerID	uint32
 }

func MakeGate2GameHead() *Gate2GameHead {
	ret := &Gate2GameHead{}
	ret.PlayerID = 0
	ret.Dump = func(msg *message.Message) {
		msg.Write(ret.PlayerID)
	}
	ret.Load = func(msg *message.Message) {
		msg.Read(&ret.PlayerID)
	}
	return ret
}


type Game2GateHead struct {
	message.Package
	PlayerID	uint32
 	Err	uint16
 }

func MakeGame2GateHead() *Game2GateHead {
	ret := &Game2GateHead{}
	ret.PlayerID = 0
	ret.Err = 0
	ret.Dump = func(msg *message.Message) {
		msg.Write(ret.PlayerID)
		msg.Write(ret.Err)
	}
	ret.Load = func(msg *message.Message) {
		msg.Read(&ret.PlayerID)
		msg.Read(&ret.Err)
	}
	return ret
}




