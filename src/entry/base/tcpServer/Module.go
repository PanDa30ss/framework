package tcpServer

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "tcpserver"

	service.RegisterModule(moduleName, getInstance())

}
