package main

import (
	"log"
	"net/http"

	"rest-api/rest/counter"
	"rest-api/rest/email"
	"rest-api/rest/substr"
	"rest-api/rest/user"

	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()

	router.HandleFunc("/rest/substr/find", substr.GetSubStr)
	router.HandleFunc("/rest/email/check", email.Email)
	//REDIS
	router.HandleFunc("/rest/counter/add/{i}", counter.CounterAdd).Methods("POST")
	router.HandleFunc("/rest/counter/sub/{i}", counter.CounterSub).Methods("POST")
	router.HandleFunc("/rest/counter/val", counter.CounterVal).Methods("GET")
	//POSTGRES
	router.HandleFunc("/rest/user", user.CreateUser).Methods("POST")
	router.HandleFunc("/rest/user/{id}", user.GetUser).Methods("GET")
	router.HandleFunc("/rest/user/{id}", user.UpdateUser).Methods("PUT")
	router.HandleFunc("/rest/user/{id}", user.DeleteUser).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":8000", router))
}
