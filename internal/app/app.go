package app

import (
	adminh "github.com/Kroning/test_shortner/internal/handlers/adminhandler"
)

type app struct {
	name string
	adminh.Page
}

func NewApp(name string) app {
	app := app{ 
		name: name, 
		Page : adminh.NewPage("1111","db"),
	}
	loadConfig() // maybe later
	return app
}

func (myapp *app) Run() {
	myapp.MainInitHandlers()
}

func (myapp *app) Name() string {
  return myapp.name
}

func loadConfig() {
}
