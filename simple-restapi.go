package main

import (
	"fmt"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "hello kanchan")
}

func main() {
	handleRequests()
}


func handleRequests() {
	http.HandleFunc("/",homePage)
	fmt.Println(http.ListenAndServe(":8081",nil))
}	