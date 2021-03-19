package app

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/accesslog"
	"github.com/kataras/iris/v12/middleware/requestid"
	"github.com/rommel96/mock-auth-server/config"
)

func StartServer() {
	server := iris.New()
	/*Config server*/
	ac := accesslog.New(nil)
	ac.SetFormatter(config.NewCustomFormatter(' ', "-"))
	server.UseRouter(ac.Handler)
	server.UseRouter(requestid.New())
	initializeRoutes(server)
	// /*---------*/
	server.Get("/", index)
	server.Listen(config.GetPortServer())
}

func index(c iris.Context) {
	c.JSON(iris.Map{
		"status":  "OK",
		"message": "Welcome to Mock Authentication Server",
		"author":  "@Rommel96",
	})
}
