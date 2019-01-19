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

	param := "%" + r.URL.Query().Get("query") + "%"
	query := fmt.Sprintf(
		`SELECT songs.id, songs.artist, songs.song, songs.length, genres.name as genre 
		FROM songs inner join genres on songs.genre = genres.id %v`,
		searchQuery(r))

	data := executeSongsQuery(query, db, param, param, param)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}

// IndexLength : returns list of songs ordered by length
func IndexLength(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	checkErr(err)

	queryString, params := buildQueryParams(r)
	query := fmt.Sprintf(
		`SELECT songs.id, songs.artist, songs.song, songs.length, genres.name as genre 
		FROM songs inner join genres on songs.genre = genres.id 
		%v 
		ORDER BY songs.length DESC `,
		queryString)

	data := executeSongsQuery(query, db, params...)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}

// executeSongsQuery : executes the query and returns the list of songs in JSON format
func executeSongsQuery(query string, db *sql.DB, args ...interface{}) []byte {
	rows, err := db.Query(query, args...)
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
		return `where songs.artist like ? 
			or songs.song like ?
			or genres.name like ? `
	}
	return ""
}

// buildQueryParams : returns the formatted string containing the 'where' clause of the song length query
func buildQueryParams(r *http.Request) (string, []interface{}) {
	var queryString []string
	var params []interface{}

	if max := r.URL.Query().Get("max"); max != "" {
		queryString = append(queryString, "songs.length < ?")
		params = append(params, max)
	}
	if min := r.URL.Query().Get("min"); min != "" {
		queryString = append(queryString, "songs.length > ?")
		params = append(params, min)
	}
	if len(params) > 0 {
		return fmt.Sprintf("where %v", strings.Join(queryString, " and ")), params
	}
	return "", params
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
