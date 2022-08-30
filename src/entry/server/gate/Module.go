package gate

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "gate"

	service.RegisterModule(moduleName, getInstance())

}

func moduleInit() {
	cmdInit()
	eventInit()
	playerCMDInit()
}
