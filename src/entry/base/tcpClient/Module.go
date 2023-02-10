package tcpClient

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "tcpclient"

	depends := []string{}
	getInstance().SetDepends(depends)

	service.RegisterModule(moduleName, getInstance())

}
