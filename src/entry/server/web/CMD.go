package web

import (
	"entry/base/http"
	"os/exec"

	. "github.com/PanDa30ss/core/http"
)

func cmdInitial() {
	http.Register("/version/master", func(context *HttpContext) {
		cmdUp := exec.Command("svn", "up", `/sda1/home/sjl4610/test/publish/master`)
		_, errUp := cmdUp.CombinedOutput()
		if errUp != nil {
			context.Write([]byte("svn up fail"))
			context.Finish()
			return
		}
		cmd := exec.Command("svn", "log", `/sda1/home/sjl4610/test/publish/master`)
		out, err := cmd.CombinedOutput()
		if err != nil {
			context.Write([]byte("svn show log fail"))
			context.Finish()
			return
		}
		context.Write([]byte(out))
		context.Finish()

	})

	http.Register("/version/alpha", func(context *HttpContext) {
		cmdUp := exec.Command("svn", "up", `/sda1/home/sjl4610/test/publish/alpha`)
		_, errUp := cmdUp.CombinedOutput()
		if errUp != nil {
			context.Write([]byte("svn up fail"))
			context.Finish()
			return
		}
		cmd := exec.Command("svn", "log", `/sda1/home/sjl4610/test/publish/alpha`)
		out, err := cmd.CombinedOutput()
		if err != nil {
			context.Write([]byte("svn show log fail"))
			context.Finish()
			return
		}
		context.Write([]byte(out))
		context.Finish()

	})
}
