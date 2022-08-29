package http

import "github.com/PanDa30ss/core/http"

func Register(url string, handleFunc http.HttpHandleFunc) bool {
	return getInstance().register(url, handleFunc)
}
