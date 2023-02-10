package http

import (
	"github.com/PanDa30ss/core/service"
)

func init() {

	moduleName := "http"

	depends := []string{}
	getInstance().SetDepends(depends)

	service.RegisterModule(moduleName, getInstance())

}
