package server

import (
	"clipboard/utils"
	"encoding/json"
	"fmt"
	"net/http"
)

func RunServer() {
	mux := http.NewServeMux()
	store := make(map[string]string)

	mux.HandleFunc("/set", func(w http.ResponseWriter, r *http.Request) {
		var j interface{}
		err := json.NewDecoder(r.Body).Decode(&j)
		utils.HandleErr(err)
		defer r.Body.Close()

		jMap := j.(map[string]interface{})

		for key, value := range jMap {
			store[key] = value.(string)
		}
	})

	mux.HandleFunc("/get/{key}", func(w http.ResponseWriter, r *http.Request) {
		storeKey := r.PathValue("key")

		w.Write([]byte(store[storeKey]))
	})

	port := 9001
	err := http.ListenAndServe(fmt.Sprintf(":%v", port), mux)
	utils.HandleErr(err)
}
