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

	query := fmt.Sprintf(
		`SELECT songs.id, songs.artist, songs.song, songs.length, genres.name as genre 
		FROM songs inner join genres on songs.genre = genres.id %v`,
		searchQuery(r))

	data := executeSongsQuery(query, db)
	w.Header().Set("Content-Type", "application/json")
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

	data := executeSongsQuery(query, db)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}

// executeSongsQuery : executes the query and returns the list of songs in JSON format
func executeSongsQuery(query string, db *sql.DB) []byte {
	rows, err := db.Query(query)
	checkErr(err)
	defer rows.Close()

	var songs []Song

	for rows.Next() {
		song := Song{}
		err = rows.Scan(&song.ID, &song.Artist, &song.Song, &song.Length, &song.Genre)
		checkErr(err)
		songs = append(songs, song)
	}

	data, err := json.Marshal(songs)
	checkErr(err)
	return data
}

// searchQuery : returns the formatted string containing the 'where' clause of the index search query
func searchQuery(r *http.Request) string {
	if param := r.URL.Query().Get("query"); param != "" {
		return fmt.Sprintf(
			`where songs.artist like "%%%v%%" 
			or songs.song like "%%%v%%"
			or genres.name like "%%%v%%" `,
			param, param, param)
	}
	return ""
}

// buildQueryParams : returns the formatted string containing the 'where' clause of the song length query
func buildQueryParams(r *http.Request) string {
	var params []string

	if max := r.URL.Query().Get("max"); max != "" {
		params = append(params, fmt.Sprintf("songs.length < %v", max))
	}
	if min := r.URL.Query().Get("min"); min != "" {
		params = append(params, fmt.Sprintf("songs.length > %v", min))
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
