package main

import (
	"github.com/rommel96/mock-auth-server/app"
	"github.com/rommel96/mock-auth-server/startup"
)

func main() {
	startup.LoadData()
	app.StartServer()
}
