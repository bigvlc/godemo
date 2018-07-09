package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// User struct
type User struct {
	ID        string `json:"id"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Age       string `json:"age"`
}

var users []User

func main() {
	router := mux.NewRouter()
	// sample data
	users = append(users, User{ID: "1", FirstName: "FirstName01", LastName: "LastName01", Age: "28"})
	users = append(users, User{ID: "2", FirstName: "FirstName02", LastName: "LastName02", Age: "33"})
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
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

func getUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	for _, i := range users {
		if i.ID == vars["id"] {
			json.NewEncoder(w).Encode(i)
			return
		}
	}
	// Add Not found 404 if ID does not exists
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&users)
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
