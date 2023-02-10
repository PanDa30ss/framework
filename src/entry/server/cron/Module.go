package cron

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "cron"

	depends := []string{"tcpserver", "tcpclient"}
	getInstance().SetDepends(depends)

	service.RegisterModule(moduleName, getInstance())

}
