/*
Package initializes all needed variables. Reads configs, starts dbpool (and test connection).
It also create Page object to start handling urls.
*/
package app

import (
	cfg "github.com/Kroning/test_shortner/internal/config"
	hand "github.com/Kroning/test_shortner/internal/handlers"
	store "github.com/Kroning/test_shortner/internal/storage"

	"context"
	"fmt"
	"log"
)

// All data we need to pass through application
type app struct {
	name string
	hand.Page
	Cfg cfg.Config      // Configs readed from config's files
	Ctx context.Context // Context to db pool
}

// Creates new App object. Reads config's files and saves it at Cfg field.
// Creates context and initializes storage(db, file etc.) using config's data.
// Returnes app object or error (unable to read configs or make connection with storage).
func NewApp(name string) (app, error) {
	app := app{
		name: name,
		Page: hand.NewPage("Shortner", nil, context.Background()),
		Ctx:  context.Background(),
	}

	cfg, err := cfg.ParseConfig(name)
	if err != nil {
		return app, err
	}
	app.Cfg = cfg

	app.Page.Storage, err = app.GetStorage()
	if err != nil {
		return app, err
	}

	return app, nil
}

// Creates new storage object (db or file)
func (myapp *app) GetStorage() (store.Storage, error) {
	var storage store.Storage = nil
	var err error
	if myapp.Cfg.Db.Host != "" {
		storage, err = store.PgConnect(myapp.Ctx, myapp.Cfg)
		if err != nil {
			return nil, err
		}
	} else if myapp.Cfg.Db.File != "" {
	} else {
		return nil, fmt.Errorf("Storage isn't chosed.")
	}
	return storage, nil
}

// Runs application initializing page's handlers
func (myapp *app) Run() {
	log.Println("app Run()")
	switch myapp.name {
	case "admin":
		myapp.MainInitHandlers()
	case "redirect":
		myapp.RedirectInitHandlers()
	}
}

// Returns name of application
func (myapp *app) Name() string {
	return myapp.name
}
