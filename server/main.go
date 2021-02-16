package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/szymslo/betting-site/matches"
)

const port string = ":5000"

func main() {

	// init mux router
	r := mux.NewRouter()

	// set route handlers (endpoints)
	r.HandleFunc("/api/matches", matches.GetMatches).Methods("GET")
	r.HandleFunc("/api/match/{id}", matches.GetMatch).Methods("GET")
	r.HandleFunc("/api/match", matches.CreateMatch).Methods("POST")
	r.HandleFunc("/api/match/{id}", matches.UpdateMatch).Methods("POST")
	r.HandleFunc("/api/match/{id}", matches.DeleteMatch).Methods("DELETE")

	// run server
	fmt.Println("Server listening at", port)
	log.Fatal(http.ListenAndServe(port, r))

}
