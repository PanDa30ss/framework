package http

import (
	"github.com/PanDa30ss/core/service"
)

func init() {

	moduleName := "http"

	service.RegisterModule(moduleName, getInstance())

}
