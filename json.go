package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func Respondtojson(w http.ResponseWriter, code int, payload interface{}) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}

func Respondtoerror(w http.ResponseWriter, code int, message string) {

	if(code>499){
		log.Fatal("Error with 5xx ")
	}
	type erroemsg struct{
		Error string `json:"error"`;
	}
	Respondtojson(w, code, erroemsg{Error:message})
}
