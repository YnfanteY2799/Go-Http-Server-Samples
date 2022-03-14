package main

import (
	"fmt"
	"net/http"
)

func main(){
	
	var ammount int = 1;

	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Hello, you've requested : %s\n like %X times", r.URL.Path, ammount)
		ammount = ammount + 1;
	})
	
	http.ListenAndServe(":80", nil)
}