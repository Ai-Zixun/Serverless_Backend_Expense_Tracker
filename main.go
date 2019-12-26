package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	requestHandler()
}

func requestHandler() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/user/create", userCreate).Methods("POST"); 
	router.HandleFunc("/user/log-in", userLogIn).Methods("POST"); 
	router.HandleFunc("/expense/create", expenseCreate).Methods("POST"); 
	router.HandleFunc("/expense/retrieve", expenseRetrive).Methods("POST"); 
	log.Fatal(http.ListenAndServe(":8081", router)) 
}

// API Endpoint: user_create(user_name, user_password) 
func userCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST END POINT: /user/create")
}

// API Endpoint: user_log_in(user_name, user_password)
func userLogIn(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST END POINT: /user/log-in")
}

// API Endpoint: expense_create(user_api_key, name, timestamp, currency, amount)
func expenseCreate(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST END POINT: /expense/create")
}

// API Endpoint: expense_retrieve(user_api_key, count, bgn_date, end_date, last_return)   
func expenseRetrive(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "POST END POINT: /expense/retrieve")
}

