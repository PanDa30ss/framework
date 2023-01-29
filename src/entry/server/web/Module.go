package web

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "web"

	service.RegisterModule(moduleName, getInstance())

}
