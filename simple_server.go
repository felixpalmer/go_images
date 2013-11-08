package main

import (
  "io"
  "net/http"

	"github.com/gorilla/mux"
)

func HelloHandler(writer http.ResponseWriter, request *http.Request) {
  io.WriteString(writer, "hello!")
}

func NoCacheDecorator(h http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")
    h.ServeHTTP(w, r)
  })
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/hello", HelloHandler)
  staticHandler := http.StripPrefix("/", http.FileServer(http.Dir(".")))
  staticHandler = NoCacheDecorator(staticHandler)
	router.PathPrefix("/").Handler(staticHandler)
  http.ListenAndServe("localhost:1234", router)
}
