package main

import (
	"bv_test/songshelper"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	goji "goji.io"
	"goji.io/pat"
)

func hello(w http.ResponseWriter, r *http.Request) {
	name := pat.Param(r, "name")
	fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/hello/:name"), hello)
	mux.HandleFunc(pat.Get("/songs"), songshelper.Index)

	fmt.Printf("Listening at localhost:8000")
	http.ListenAndServe("localhost:8000", mux)
}
