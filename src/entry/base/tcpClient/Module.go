package tcpClient

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "tcpclient"

	service.RegisterModule(moduleName, getInstance())

}
