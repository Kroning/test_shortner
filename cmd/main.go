/*
Usage: go run cmd/admin/main.go -app [admin|redirect]
Starts App for admin interface of shortner.
It will use configs from "config" directory.
*/
package main

import (
	"errors"
	"flag"
	"fmt"
	myApp "github.com/Kroning/test_shortner/internal/app"
	"log"
	"net/http"
	"os"
)

var (
	appName    string    // Appname needs to differ admin and redirection services. Also it is used for name of main config (appName+".yml").
	logPath    = "logs/" // for testing purpose
	ErrNoFlags = "Provide app name: [admin|redirect]"
)

// Starts main "start()" function
func main() {
	// Small main for not to test it

	flag.StringVar(&appName, "app", "", "application name [admin|redirect]")
	err := parseFlags()
	if err != nil {
		log.Fatal(err)
	}

	port, err := start()
	if err != nil {
		log.Fatal(err)
	}
	log.Fatal(http.ListenAndServe(":"+port, nil)) // handlers and server can go to some App package too
}

// Parses flags
// Currently only app name - global var
func parseFlags() error {
	flag.Parse()
	if appName == "" {
		return errors.New(ErrNoFlags)
	}
	return nil
}

// Initializing App, starting server.
func start() (string, error) {
	// Redirecting logs
	f, err := os.OpenFile(logPath+appName+".log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return "", err
	}
	defer f.Close()
	log.SetOutput(f)

	// Starting app
	fmt.Println("Starting App " + appName)
	log.Println("Starting App")
	app, err := myApp.NewApp(appName)
	if err != nil {
		return "", err
	}

	fmt.Println("Running App " + appName + " at :" + app.Cfg.Server.Port)
	app.Run()
	return app.Cfg.Server.Port, nil
}

func GetAppName() string {
	return appName
}
