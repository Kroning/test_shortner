/*
Page structure, shared methods and routing.
*/
package handlers

import (
	"context"
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
}

// Returns new page object with title and db handler provided
func NewPage(t string, db *pgxpool.Pool, ctx context.Context) Page {
	return Page{
		Title: t,
		Db:    db,
		Ctx:   ctx,
	}
}

func (p Page) MainInitHandlers() {
	http.HandleFunc("/", p.mainHandler)
	http.HandleFunc("/save", p.saveHandler)
}

func (p Page) RedirectInitHandlers() {
	http.HandleFunc("/", p.redirectHandler)
}
