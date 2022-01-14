package main

import (
	"fmt"	
    "net/http"   
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true);
	r.HandleFunc("/", homeHandler)
	http.Handle("/", r)
	http.ListenAndServe(":3000", nil)
}

func homeHandler(w http.ResponseWriter, req *http.Request) {
	// Do someting
	fmt.Fprintf(w, "hello world!")
	vars := mux.Vars(req)
    w.WriteHeader(http.StatusOK)
    fmt.Fprintf(w, "vars: %v\n", vars)
}
