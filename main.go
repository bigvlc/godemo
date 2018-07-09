package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type user struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       string `json:"age"`
}

type users []user

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", index).Methods("GET")
	router.HandleFunc("/user", getUsers).Methods("GET")
	router.HandleFunc("/user/{id:[0-9]+}", getUser).Methods("GET")
	router.HandleFunc("/user", createUser).Methods("POST")
	router.HandleFunc("/user/{id:[0-9]+}", updateUser).Methods("PUT")
	router.HandleFunc("/user/{id:[0-9]+}", deleteUser).Methods("DELETE")
	router.Use(loggingMiddleware)
	log.Fatal(http.ListenAndServe(":"+os.Getenv("DEMO_PORT"), router))
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.RequestURI, r.Method, r.UserAgent())
		next.ServeHTTP(w, r)
	})
}

func index(w http.ResponseWriter, r *http.Request) {
	i := "Welcome to Demo API"
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(i)
}

func getUsers(w http.ResponseWriter, r *http.Request) {
	u := users{
		user{
			ID:        "1",
			FirstName: "FirstName01",
			LastName:  "LastName01",
			Age:       "28",
		},
		user{
			ID:        "2",
			FirstName: "FirstName02",
			LastName:  "LastName02",
			Age:       "33",
		},
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(u)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Fprintln(w, "Here is user with ID:", userID)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Not Implemented")
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Fprintln(w, "Updated user with ID:", userID)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID := vars["id"]
	fmt.Fprintln(w, "Deleted user with ID:", userID)
}
