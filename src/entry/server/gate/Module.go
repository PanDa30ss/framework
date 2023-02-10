package gate

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "gate"

	depends := []string{"tcpserver", "tcpclient"}
	getInstance().SetDepends(depends)

	service.RegisterModule(moduleName, getInstance())

}
