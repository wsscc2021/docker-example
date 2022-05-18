package main

import (
	"math/rand"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthApiHandler)
	mux.HandleFunc("/stress/cpu", stressCpuApiHandler)
	http.ListenAndServe(":5000", handlers.CombinedLoggingHandler(os.Stdout, mux))
}

func healthApiHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}
}

func stressCpuApiHandler(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case "GET":
		go sub()
		go sub()
		go sub()
		go sub()
		w.WriteHeader(200)
		w.Write([]byte("cpu loaded..."))
	}
}

func sub() {
	for i := 1; i <= 10000; i++ {
		x := rand.Float64()
		x *= x
	}
}
