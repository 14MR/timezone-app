package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", TimezoneServer)
	http.ListenAndServe(":8080", nil)
}

func TimezoneServer(w http.ResponseWriter, r *http.Request) {
	timezone := r.URL.Path[1:]
	if timezone == "" {
		timezone = "UTC"
	}
	loc, err := time.LoadLocation(timezone)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	now := time.Now().In(loc)

	fmt.Fprintf(w, "time in %s: %s", timezone, now)
}
