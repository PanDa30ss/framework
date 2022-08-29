package http

import (
	"github.com/PanDa30ss/core/http"
	"github.com/PanDa30ss/core/service"
)

func init() {

	moduleName := "http"

	service.RegisterModule(moduleName, getInstance())

}

func moduleInit() {
	Register("/example", func(context *http.HttpContext) {
		context.Write([]byte("aaaa"))
		context.Finish()
	})
}
