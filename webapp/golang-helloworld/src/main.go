package main

import (
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", greetApiHandler)
	http.ListenAndServe(":5000", handlers.CombinedLoggingHandler(os.Stdout, mux))
}

func greetApiHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.WriteHeader(200)
		w.Write([]byte("Hello World!"))
	}
}
