package redisBank

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "redisBank"

	service.RegisterModule(moduleName, getInstance())

}
