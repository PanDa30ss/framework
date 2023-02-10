package master

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "master"

	depends := []string{"tcpserver"}
	getInstance().SetDepends(depends)

	service.RegisterModule(moduleName, getInstance())

}
