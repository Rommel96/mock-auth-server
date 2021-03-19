package app

import (
	"net/http"

	"github.com/kataras/iris/v12"
	"github.com/rommel96/mock-auth-server/repository"
)

func initializeRoutes(server *iris.Application) {
	server.Post("/login", login)
	//server.Post("/users", createUsers)
}

func login(c iris.Context) {
	var loginRequest LoginRequest
	if err := c.ReadJSON(&loginRequest); err != nil {
		c.StopWithError(http.StatusBadRequest, err)
		return
	}
	user := repository.FindUser(loginRequest.Email, loginRequest.Passowrd)
	if user == nil {
		c.StopWithJSON(http.StatusNotFound, user)
		return
	}
	c.StopWithJSON(http.StatusOK, user)
}
