package webServer

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func StartServer() {
	//rktServer := http.ListenAndServe(":8081", nil)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":8081", nil))

}
