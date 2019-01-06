package main

import (
	"bv_test/genreshelper"
	"bv_test/songshelper"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
	goji "goji.io"
	"goji.io/pat"
)

func main() {
	mux := goji.NewMux()
	mux.HandleFunc(pat.Get("/songs"), songshelper.Index)
	mux.HandleFunc(pat.Get("/songs/byLength"), songshelper.IndexLength)
	mux.HandleFunc(pat.Get("/genres"), genreshelper.Genres)

	fmt.Printf("Listening at localhost:8000")
	http.ListenAndServe("localhost:8000", mux)
}
