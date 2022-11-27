/*
Page structure and shared method.
*/
package handlers

import (
  "github.com/jackc/pgx/v5/pgxpool"
)

const tmplPath = "views"

// Page contains Title and db handler
type Page struct {
    Title string
    Db    *pgxpool.Pool
}

// Returns new page object with title and db handler provided
func NewPage(t string, db *pgxpool.Pool) Page {
  return Page{
    Title : t,
    Db :  db,
  }
}

