package tcpServer

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "tcpserver"

	depends := []string{}
	getInstance().SetDepends(depends)

	service.RegisterModule(moduleName, getInstance())

}
