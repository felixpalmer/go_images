package main

import (
  "io"
  "net/http"

	"github.com/gorilla/mux"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
  io.WriteString(writer, "hello!")
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloHandler)
  staticHandler := http.StripPrefix("/", http.FileServer(http.Dir(".")))
  //staticHandler = NoCache(staticHandler)
	router.PathPrefix("/").Handler(staticHandler)
  http.ListenAndServe("localhost:1234", router)
}
