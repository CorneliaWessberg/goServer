package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	defaultPort := "8080"
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", cors(fileServer))

	fmt.Printf("Server starting at port %s\n", defaultPort)
	if err := http.ListenAndServe((":" + defaultPort), nil); err != nil {
		log.Fatal(err)
	}
}

func cors(fs http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// our CORS stuff
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
		fs.ServeHTTP(w, r)
	}
}
