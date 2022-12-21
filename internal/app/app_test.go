package app

import (
	"context"
	hand "github.com/Kroning/test_shortner/internal/handlers"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"os"
	"reflect"
	"testing"
	//"github.com/Kroning/test_shortner/internal/config"
	"github.com/jackc/pgx/v5/pgxpool"
)

func TestRun(t *testing.T) {
	app := &app{
		Page: hand.NewPage("Shortner", nil, context.Background()),
		Ctx:  context.Background(),
	}
	app.Run()
	//require.NoError(t, err)
}

// Don't know how to test pgxpool.New. Leave it for a while
/*func TestGetPool(t *testing.T) {
	app := &app{
		Page: hand.NewPage("Shortner", nil, context.Background()),
		Ctx: context.Background(),
		Cfg: config.Config{},
	}
	app.Cfg.Db.Username = "1@@@@/@.@.@2#;//"
	_, err := app.GetPool()
	require.NoError(t, err)
}*/

func TestNewApp(t *testing.T) {
	os.Chdir("../../")
	t.Run("Test-admin", func(t *testing.T) {
		//os.Chdir("../../")
		app, err := NewApp("admin")
		require.NoError(t, err)
		assert.Equal(t, "Shortner", app.Page.Title)

		if reflect.ValueOf(hand.Page{}).Type() != reflect.ValueOf(app.Page).Type() {
			t.Errorf("app.Page type wrong: %s ; must be: %s", reflect.ValueOf(app.Page).Type(), reflect.ValueOf(&hand.Page{}).Type())
		}

		if reflect.ValueOf(app.Page.Db).Type() != reflect.ValueOf(app.Page.Db).Type() {
			t.Errorf("app.Page.Db type wrong: %s ; must be: %s", reflect.ValueOf(app.Page.Db).Type(), reflect.ValueOf(&pgxpool.Pool{}).Type())
		}
	})

	// Though the test are same there may be different problems (like configs)
	t.Run("Test-redirect", func(t *testing.T) {
		//os.Chdir("../../")
		app, err := NewApp("redirect")
		require.NoError(t, err)
		assert.Equal(t, "Shortner", app.Page.Title)

		if reflect.ValueOf(hand.Page{}).Type() != reflect.ValueOf(app.Page).Type() {
			t.Errorf("app.Page type wrong: %s ; must be: %s", reflect.ValueOf(app.Page).Type(), reflect.ValueOf(&hand.Page{}).Type())
		}

		if reflect.ValueOf(app.Page.Db).Type() != reflect.ValueOf(app.Page.Db).Type() {
			t.Errorf("app.Page.Db type wrong: %s ; must be: %s", reflect.ValueOf(app.Page.Db).Type(), reflect.ValueOf(&pgxpool.Pool{}).Type())
		}
	})

	t.Run("Test-badname", func(t *testing.T) {
		_, err := NewApp("bad")
		require.Error(t, err)
	})

	t.Run("Test-bad", func(t *testing.T) {
		_, err := NewApp("admin_example")
		require.Error(t, err)
	})
}
