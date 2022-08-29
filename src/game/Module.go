package game

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "game"

	service.RegisterModule(moduleName, getInstance())

}

func moduleInit() {
	cmdInit()
	eventInit()
}
