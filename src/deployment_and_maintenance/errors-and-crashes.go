package main

import (
	"net/http"
	"go-releated/src/logs"
	"html/template"
	"fmt"
	"log"
)

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello astaxie!") // send data to client side
}

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		sayhelloName(w, r)
		return
	}
	NotFound404(w, r)
	return
}

func NotFound404(w http.ResponseWriter, r *http.Request) {
	logs.Logger.Critical("page not found")       // error logging
	t, _ := template.ParseFiles("tmpl/404.html") // parse the template file
	ErrorInfo := "File not found"                // Get the current user information
	t.Execute(w, ErrorInfo)                      // execute the template merger operation
}

func main() {
	http.HandleFunc("/hello", ServeHTTP)          // set router
	err := http.ListenAndServe(":9090", nil) // set listen port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
