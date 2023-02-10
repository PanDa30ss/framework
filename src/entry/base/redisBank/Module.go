package redisBank

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "redisbank"

	depends := []string{}
	getInstance().SetDepends(depends)

	service.RegisterModule(moduleName, getInstance())

}
