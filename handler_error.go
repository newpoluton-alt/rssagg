package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	respondWithError(w, 400, "Some Critical Server Error")
}
