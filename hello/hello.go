package hello

import (
    "fmt"
    "http"
)

func init() {
    http.HandleFunc("/", defaultHandler)
}

func defaultHandler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprint(w, "Hello, world!")
}
