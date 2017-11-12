package main

import (
	"fmt"
	"net/http"
)

func AuthMock(w http.ResponseWriter, r *http.Request) {
    // Write the album details as plain text to the client.
    fmt.Fprintf(w, "200")
}

func main() {
    http.HandleFunc("/", AuthMock)
    http.ListenAndServe(":7878", nil)
}
