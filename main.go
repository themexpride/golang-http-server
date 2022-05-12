package main

import (
	"fmt"
  "log"
  "net/http"
)

func emailHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Println("POST request successful")

    email := r.FormValue("email")
    fmt.Fprintf(w, "Successfully signed up: %s \n", email)
}


func formHandler(w http.ResponseWriter, r *http.Request) {
    if err := r.ParseForm(); err != nil {
        fmt.Fprintf(w, "ParseForm() err: %v", err)
        return
    }
    fmt.Println("POST request successful")

    name := r.FormValue("name")
    address := r.FormValue("address")
    fmt.Fprintf(w, "Name: %s \n", name)
    fmt.Fprintf(w, "Address: %s \n", address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found", http.StatusNotFound)
        return
    }
    if r.Method != "GET" {
        http.Error(w, "method is not supported", http.StatusNotFound)
        return
    }
    fmt.Println("HELLO request successful")
    fmt.Fprintf(w, "hello!")
}

func main() {
  	fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer)
    http.HandleFunc("/form", formHandler)
    http.HandleFunc("/hello", helloHandler)
    http.HandleFunc("/email", emailHandler)
  
    fmt.Println("Starting server at port 8080")
    if err := http.ListenAndServe(":8080", nil); err != nil {
      log.Fatal(err)
    }
}
