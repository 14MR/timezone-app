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
	loc, _ := time.LoadLocation(timezone)
	now := time.Now().In(loc) // TODO: this produces runtime panic in Docker for some reason but works

	fmt.Fprintf(w, "time in %s: %s", timezone, now)
}
