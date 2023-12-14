package main

import "net/http"

func readinessHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("200 OK"))
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
}
