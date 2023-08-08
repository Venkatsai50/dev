package main

import "net/http"


func handleerr(w http.ResponseWriter, r *http.Request) {
	Respondtoerror(w, 400, "Something went wrong")
} 