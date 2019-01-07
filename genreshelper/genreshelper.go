package genreshelper

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

//Genre : struct describing the table songs
type Genre struct {
	Name        string
	Songs       int
	TotalLength int
}

// Genres : returns a list of genres, with total of songs and lengths
func Genres(w http.ResponseWriter, r *http.Request) {
	db, err := sql.Open("sqlite3", "./jrdd.db")
	checkErr(err)

	query := `SELECT name, count(songs.id) as songs, sum(songs.length) as total_length
	FROM genres LEFT JOIN songs on genres.ID = songs.genre 
	GROUP BY name
	ORDER BY songs DESC`

	rows, err := db.Query(query)
	checkErr(err)
	defer rows.Close()

	var genres []Genre

	for rows.Next() {
		genre := Genre{}
		var totalLength sql.NullInt64
		err = rows.Scan(&genre.Name, &genre.Songs, &totalLength)
		checkErr(err)
		genre.TotalLength = int(totalLength.Int64)
		genres = append(genres, genre)
	}

	data, err := json.Marshal(genres)
	checkErr(err)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, string(data))
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
