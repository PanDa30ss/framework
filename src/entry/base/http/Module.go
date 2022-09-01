package http

import (
	"io/ioutil"

	"github.com/PanDa30ss/core/http"
	"github.com/PanDa30ss/core/service"
)

func init() {

	moduleName := "http"

	service.RegisterModule(moduleName, getInstance())

}

var _ = Register("/example", func(context *http.HttpContext) {
	r := context.GetRequest()
	values := r.URL.Query()
	for k, m := range values {
		for _, v := range m {
			context.Write([]byte(k + ":"))
			context.Write([]byte(v + "\n"))
		}
	}
	context.Finish()
})

var _ = Register("/example_post", func(context *http.HttpContext) {
	r := context.GetRequest()
	defer r.Body.Close()
	result, err := ioutil.ReadAll(r.Body)
	if err != nil {
		context.Finish()
		return
	}
	context.Write(result)
	context.Finish()
})
