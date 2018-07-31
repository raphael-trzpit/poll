package main

import (
	"fmt"
	"net/http"
	"os"

	"google.golang.org/appengine"
	"github.com/raphael-trzpit/poll/internal"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Fprintln(w, internal.Str)
	fmt.Fprintln(w, os.Getenv("TEST"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}
