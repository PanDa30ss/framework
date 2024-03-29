package game

import (
	. "entry/base/proto/pb"
	"entry/base/tcpClient"
	"fmt"

	"github.com/PanDa30ss/core/db/redis"
	"github.com/PanDa30ss/core/message"
	"github.com/PanDa30ss/core/tcp"
	. "github.com/garyburd/redigo/redis"
)

func cmdInitial() {
	registerCMD(
		CMD_REQ_Test,
		func(s tcp.ISession, msg *message.Message) bool {
			playerID := msg.ReadUInt32()
			pkg := &Test{}
			msg.Unmarshal(pkg)
			p := getInstance().players[playerID]
			p.sendPBMessage(CMD_RSP_Test, 0, pkg)
			getInstance().gameRedis.Query(redis.MakeRedisCommand(redisTest), "hgetall", "NAME_{00}")
			return true
		})

	registerCMD(
		CMD_REQ_EnterGame,
		func(s tcp.ISession, msg *message.Message) bool {
			session := s.(*tcpClient.ServerC)
			pkg := &EnterGame{}
			msg.Unmarshal(pkg)

			p := getInstance().addPlayer(pkg.GetTestID())
			p.gate = session.Data.ServerID
			p.sendPBMessage(CMD_RSP_EnterGame, 0, pkg)

			return true
		})

}

func redisTest(result *redis.RedisResult, params ...interface{}) {
	a, err := IntMap(result.Result, nil)
	fmt.Println(err)
	fmt.Println(a)
}
