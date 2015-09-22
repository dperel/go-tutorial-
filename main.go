package main

import "net/http"

func main() {
  http.HandleFunc("/", hello) // -> when the root is requested at the port, hello is called.
  http.ListenAndServe(":8080", nil)
}

func hello(w http.ResponseWriter, r *http.Request) { //write with .ResponseWriter, reads the request 
  w.Write([]byte("hello!")) // -> hello returns a byte string "hello"
}
