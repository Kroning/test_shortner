/*
Package initializes all needed variables. Reads configs, starts dbpool (and test connection).
It also create Page object to start handling urls.
*/
package app

import (
	cfg "github.com/Kroning/test_shortner/internal/config"
	hand "github.com/Kroning/test_shortner/internal/handlers"

	"context"
	"fmt"
	"log"
	_ "os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// All data we need to pass through application
type app struct {
	name string
	hand.Page
	Cfg cfg.Config      // Configs readed from config's files
	Ctx context.Context // Context to db pool
}

// Creates new App object. Reads config's files and saves it at Cfg field.
// Creates context and initializes dbpool withit using config's data.
// Returnes app object or error (unable to read configs or make connection).
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

	_, err = app.GetPool()
	if err != nil {
		return app, err
	}

	return app, nil
}

// Creates postgres db pool
// Test connection with Acquire
func (myapp *app) GetPool() (*pgxpool.Pool, error) {
	db := myapp.Cfg.Db
	dburl := "postgres://" + db.Username + ":" + db.Password + "@" + db.Host + ":" + db.Port + "/" + db.Dbname
	dbpool, err := pgxpool.New(myapp.Page.Ctx, dburl)
	if err != nil {
		return nil, err
	}
	//defer dbpool.Close() - No need actually

	myapp.Page.Db = dbpool

	// In container DB can start a few seconds.
	// Docker with "depends_on" wait for container, but not DB.
	// This is workaround for start up.
	cnt := 0
	for true {
		_, err = myapp.Page.Db.Acquire(myapp.Page.Ctx)
		if err != nil {
			cnt++
			if cnt > 5 {
				return nil, err
			}
			fmt.Println("No connect to database, attempt ", cnt)
			time.Sleep(2 * time.Second)
			continue
		}
		fmt.Println("DB connection succesfull")
		break
	}

	return dbpool, err
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
