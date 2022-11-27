/*
Starts App for admin interface of shortner.
It will use configs from "config" directory.
*/
package main

import (
	"fmt"
	"net/http"
	"log"
	"os"
	myApp	"github.com/Kroning/test_shortner/internal/app"
)

const appName = "admin" // Appname needs to differ webinterface and redirection services. Also it is used for name of main config (appName+".yml").

// Initializing App, starting server.
func main() {
	// flag.Parse() - don't need yet?

  // Redirecting logs
  f, err := os.OpenFile("logs/"+appName+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
  if err != nil {
	  log.Fatal(err)
  }   
  defer f.Close()
  log.SetOutput(f)

	// Starting app
	fmt.Println("Starting App "+appName)
	log.Println("Starting App")
  app, err := myApp.NewApp(appName)
  if err != nil {
    log.Fatal(err)
  }

	fmt.Println("Running App "+appName)
	app.Run()
  log.Fatal(http.ListenAndServe(":"+app.Cfg.Server.Port, nil)) // handlers and server can go to some App package too
}

