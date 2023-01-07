/*
Redirect if passed links such as /{link_name}
*/
package handlers

import (
	"fmt"
	"log"
	"net/http"
	"regexp"

	store "github.com/Kroning/test_shortner/internal/storage"
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
		url, err := p.Storage.CheckLinkExistance(p.Ctx, alias)
		if err != nil {
			if err == store.LinkNotExists {
				http.NotFound(w, r)
				return
			}
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
