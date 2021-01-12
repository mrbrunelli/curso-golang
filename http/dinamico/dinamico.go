package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func welcome(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "->", r.URL)
	j := `{"message":"welcome","routes":"/hora-certa"}`
	fmt.Fprint(w, j)
}

func horaCerta(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Method, "->", r.URL)
	s := time.Now().Format("02/01/2006 03:04:05")
	fmt.Fprintf(w, "<h1>Hora Certa: %s</h1>", s)
}

func main() {
	http.HandleFunc("/", welcome)
	http.HandleFunc("/hora-certa", horaCerta)
	log.Println("Executanto...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
