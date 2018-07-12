package main

import (
	"fmt"
	"net/http"
	"os"

	"google.golang.org/appengine"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Redirect(w, r, "/", http.StatusFound)
		return
	}

	fmt.Fprintln(w, "Hello, Gopher Network!")
	fmt.Fprintln(w, os.Getenv("TEST"))
}

func main() {
	http.HandleFunc("/", indexHandler)
	appengine.Main()
}
