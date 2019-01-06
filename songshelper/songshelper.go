package songshelper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

//Song : struct describing the table songs
type Song struct {
	ID     int
	Artist string
	Song   string
	Genre  int
	Length int
}

//Index : handler of songs index
func Index(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	checkErr(err)

	rows, err := db.Query("SELECT * FROM songs")
	checkErr(err)

	var songs []Song

	for rows.Next() {
		song := Song{}
		err = rows.Scan(&song.ID, &song.Artist, &song.Song, &song.Genre, &song.Length)
		checkErr(err)
		songs = append(songs, song)
	}

	rows.Close()

	data, err := json.Marshal(songs)
	checkErr(err)
	fmt.Fprint(w, string(data))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
