package web

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "web"

	depends := []string{"http"}
	getInstance().SetDepends(depends)

	service.RegisterModule(moduleName, getInstance())

}
