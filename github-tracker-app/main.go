package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/gorilla/mux"
)

func postHandler(w http.ResponseWriter, r *http.Request){
	fmt.Println("Received Post Request")

	defer r.Body.Close()

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading Request")
		return
	}

	fmt.Println("Received Body: ", string(body))
}

func main(){
	router := mux.NewRouter()

	router.HandleFunc("/hello",postHandler).Methods("POST")

	fmt.Println("Server esta corrindo no porto 8080")
	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println(err.Error())
	}
}