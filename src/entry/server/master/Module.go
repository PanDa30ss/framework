package master

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "master"

	service.RegisterModule(moduleName, getInstance())

}
