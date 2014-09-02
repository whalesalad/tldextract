package main

import (
    "net/http"
    "strings"
)

type Domain struct {
    subdomain, domain, suffix string
}

func query(w http.ResponseWriter, r *http.Request) {
    query := strings.SplitN(r.URL.Path, "/", 3)[2]
    w.Write([]byte(query))
}

func main() {
    http.HandleFunc("/query/", query)
    http.ListenAndServe(":7777", nil)
}