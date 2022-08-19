package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//crea swith para ver que tipo de peticiones
		switch r.Method {
		case "GET":
			fmt.Fprintf(w, "GET")
		case "POST":
			fmt.Fprintf(w, "POST")
		case "PUT":
			fmt.Fprintf(w, "PUT")
		case "DELETE":
			fmt.Fprintf(w, "DELETE")
		default:
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, "UNKNOWN")
		}

	})
	srv := http.Server{
		Addr: ":8080",
	}
	error := srv.ListenAndServe()
	if error != nil {
		log.Fatal(error)
	}
}
