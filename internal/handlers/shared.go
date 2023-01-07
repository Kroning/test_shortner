/*
Page structure, shared methods and routing.
*/
package handlers

import (
	"context"
	store "github.com/Kroning/test_shortner/internal/storage"
	"github.com/jackc/pgx/v5/pgxpool"
	"net/http"
)

const tmplPath = "views"

// Page contains Title and db handler
type Page struct {
	Title   string
	Message string
	Db      *pgxpool.Pool
	Ctx     context.Context
	Result  [][]string
	Storage store.Storage
}

// Returns new page object with title and db handler provided
func NewPage(t string, Storage store.Storage, ctx context.Context) Page {
	return Page{
		Title:   t,
		Storage: Storage,
		Ctx:     ctx,
	}
}

func (p Page) MainInitHandlers() {
	http.HandleFunc("/", p.mainHandler)
	http.HandleFunc("/save", p.saveHandler)
}

func (p Page) RedirectInitHandlers() {
	http.HandleFunc("/", p.redirectHandler)
}
