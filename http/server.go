package http

import (
	"fmt"
	"log"
	"net/http"
)



func Handler(w http.ResponseWriter, r *http.Request){
	fmt.Printf("reuqest method %s\n", r.Method)
	fmt.Printf("reuqest host%s\n", r.Host)
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")


	fmt.Fprintf(w, "test")
}


func Router(){
	fmt.Println("run http-server http://localhost:8000")
	http.HandleFunc("/", Handler)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
