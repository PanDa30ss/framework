package web

import (
	"entry/base/http"
	"os/exec"

	. "github.com/PanDa30ss/core/http"
)

var _ = http.Register("/version/master", func(context *HttpContext) {
	cmdUp := exec.Command("svn", "up", `/data/P03/master`)
	_, errUp := cmdUp.CombinedOutput()
	if errUp != nil {
		context.Write([]byte("svn up fail"))
		context.Finish()
		return
	}
	cmd := exec.Command("svn", "log", `/data/P03/master`)
	out, err := cmd.CombinedOutput()
	if err != nil {
		context.Write([]byte("svn show log fail"))
		context.Finish()
		return
	}
	context.Write([]byte(out))
	context.Finish()

})

var _ = http.Register("/version/alpha", func(context *HttpContext) {
	cmdUp := exec.Command("svn", "up", `/data/P03/alpha`)
	_, errUp := cmdUp.CombinedOutput()
	if errUp != nil {
		context.Write([]byte("svn up fail"))
		context.Finish()
		return
	}
	cmd := exec.Command("svn", "log", `/data/P03/alpha`)
	out, err := cmd.CombinedOutput()
	if err != nil {
		context.Write([]byte("svn show log fail"))
		context.Finish()
		return
	}
	context.Write([]byte(out))
	context.Finish()

})
