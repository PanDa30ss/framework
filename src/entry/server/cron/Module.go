package cron

import "github.com/PanDa30ss/core/service"

func init() {

	moduleName := "cron"

	service.RegisterModule(moduleName, getInstance())

}
