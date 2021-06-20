package http

import (
	"fmt"
	"log"
	"net/http"
)


func configHandler(fn func(http.ResponseWriter, *http.Request)) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request){
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
		fn(w, r)
	}
}

func Handler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("reuqest method %s\n", r.Method)
	fmt.Printf("reuqest host%s\n", r.Host)
	_, err := fmt.Fprintf(w, "test")
	if err != nil {
		return
	}
}


func Router(){
	fmt.Println("run http-server http://localhost:8000")
	http.HandleFunc("/", configHandler(Handler))
	log.Fatal(http.ListenAndServe(":8000", nil))
}
