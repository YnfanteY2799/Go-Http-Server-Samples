package main

import (
	"fmt"
	"net/http"
);


 type Route struct {
	uri string;
	action func();
 }

//  func(r *Route) defineRoute{}

 


 var routes = []Route{
	 {
		 uri:"/hello", 
	 	action: func() {},
	},
	{
		uri: "/api/getTotalUsers",
		action: func() {},	
	},
};
 

func main(){

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Welcome to my Website my dude");
	});

	
	fs := http.FileServer(http.Dir("static/"));


	http.Handle("/static/", http.StripPrefix("/static/", fs));


	http.ListenAndServe(":8090", nil);
}