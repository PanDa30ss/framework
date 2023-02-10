package game

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "game"

	depends := []string{"redisbank", "tcpclient"}
	getInstance().SetDepends(depends)

	service.RegisterModule(moduleName, getInstance())

}
