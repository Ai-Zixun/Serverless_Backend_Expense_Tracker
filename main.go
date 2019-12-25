package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type Article struct {
	Title string `json:"title"`
	Desc string `json:"desc"`
	Content string `json:"content"`
}


func main() {
	requestHandler()
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := []Article{
		Article{Title: "Title 1", Desc: "Description 1", Content: "Content 1"}, 
	}


	fmt.Println("Endpoint: All Article")
	json.NewEncoder(w).Encode(articles); 
}

func homepage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpint Hit")
}

func requestHandler() {
	http.HandleFunc("/", homepage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":8081", nil)) 
}

// API Endpoint: creat_user(user_name, user_password) 