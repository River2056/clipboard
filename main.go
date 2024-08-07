package main

import (
	"flag"
	"fmt"
	"net/http"
)

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

func runServer() {
	store := make(map[string]string)

	http.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		store["test key"] = "test value"
	})

	http.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(store["test key"]))
	})

	port := 9001
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
	handleErr(err)
}

func main() {
	mode := flag.String("m", "server", "mode to activate; server or client")
	flag.Parse()

	switch *mode {
	case "server":
		runServer()
	}
}
