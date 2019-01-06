package songshelper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//Song : struct describing the table songs
type Song struct {
	ID     int
	Artist string
	Song   string
	Genre  string
	Length int
}

//Index : handler of songs index
func Index(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	checkErr(err)

	query := `SELECT songs.id, songs.artist, songs.song, songs.length, genres.name as genre 
	FROM songs inner join genres on songs.genre = genres.id `

	if r.URL.Query().Get("query") != "" {
		query = searchQuery(query, r.URL.Query().Get("query"))
	}

	rows, err := db.Query(query)
	checkErr(err)

	var songs []Song

	for rows.Next() {
		song := Song{}
		err = rows.Scan(&song.ID, &song.Artist, &song.Song, &song.Length, &song.Genre)
		checkErr(err)
		songs = append(songs, song)
	}

	rows.Close()

	data, err := json.Marshal(songs)
	checkErr(err)
	fmt.Fprint(w, string(data))
}

// IndexLength : returns list of songs ordered by length
func IndexLength(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	checkErr(err)

	query := fmt.Sprintf(
		`SELECT songs.id, songs.artist, songs.song, songs.length, genres.name as genre 
		FROM songs inner join genres on songs.genre = genres.id 
		%v 
		ORDER BY songs.length DESC `,
		buildQueryParams(r))

	rows, err := db.Query(query)
	checkErr(err)

	var songs []Song

	for rows.Next() {
		song := Song{}
		err = rows.Scan(&song.ID, &song.Artist, &song.Song, &song.Length, &song.Genre)
		checkErr(err)
		songs = append(songs, song)
	}

	rows.Close()

	data, err := json.Marshal(songs)
	checkErr(err)
	fmt.Fprint(w, string(data))
}

func searchQuery(query string, param string) string {
	return fmt.Sprintf(
		`%v where songs.artist like "%%%v%%" 
		or songs.song like "%%%v%%"
		or genres.name like "%%%v%%" `,
		query, param, param, param)
}

func buildQueryParams(r *http.Request) string {

	var params []string

	if r.URL.Query().Get("max") != "" {
		params = append(params, fmt.Sprintf("songs.length < %v", r.URL.Query().Get("max")))
	}
	if r.URL.Query().Get("min") != "" {
		params = append(params, fmt.Sprintf("songs.length > %v", r.URL.Query().Get("min")))
	}
	if len(params) > 0 {
		return fmt.Sprintf("where %v", strings.Join(params, " and "))
	}
	return ""
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
