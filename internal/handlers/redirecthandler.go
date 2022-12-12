/*
Redirect if passed links such as /{link_name}
*/
package handlers

import (
	"fmt"
	_ "html/template"
	"log"
	"net/http"
	"regexp"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/pgxpool"
)

var regLink = regexp.MustCompile("^/([a-zA-Z0-9]+)$") // regexp to validate URL

// Main redirect handler
func (p Page) redirectHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Error(w, "Forbidden", 403)
		return
	}

	switch r.Method {
	case "GET":
		aliases := regLink.FindStringSubmatch(r.URL.Path)
		if len(aliases) == 0 {
			log.Println("Wrong url: ", r.URL.Path)
			http.Error(w, "Internal Server Error", 500)
			return
		}
		alias := aliases[1]

		// Looking for link in database
		query := "SELECT url FROM links WHERE deleted_at IS NULL and alias = $1;"
		row := p.Db.QueryRow(p.Ctx, query, alias)
		var url string
		err := row.Scan(&url)
		if err != nil {
			if err == pgx.ErrNoRows {
				http.NotFound(w, r)
				return
			}
			log.Println("Error ", err, " while executing query ", query)
			http.Error(w, "Internal Server Error", 500)
			return
		}
		//fmt.Fprintf(w, "Redirecting to "+url)
		http.Redirect(w, r, url, http.StatusFound)
		return

	default:
		fmt.Fprintf(w, "Sorry, only GET method are supported.")
		http.Error(w, "Forbidden", 403)
		return
	}

}
