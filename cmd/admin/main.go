package main

import (
	//"fmt"
	"net/http"
	"log"
	//"time"
	myApp	"github.com/Kroning/test_shortner/internal/app"
	//handler "github.com/Kroning/test_shortner/internal/handlers/adminhandler"
	
)

const appName = "admin";

func main() {
	// flag.Parse() - don't need yet?
	app := myApp.NewApp(appName)

	app.Run()
  log.Fatal(http.ListenAndServe(":9990", nil)) // handlers and server can go to some App package too
}

