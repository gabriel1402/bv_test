package songshelper

import (
	"net/http"
	"net/http/httptest"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestIndex(t *testing.T) {
	req, err := http.NewRequest("GET", "/songs", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(Index)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func TestIndexLength(t *testing.T) {
	req, err := http.NewRequest("GET", "/songs/byLength", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(IndexLength)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}

func Test_searchQuery(t *testing.T) {
	type args struct {
		r *http.Request
	}
	request, _ := http.NewRequest("GET", "songs?query=Color", nil)
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test case 1",
			args{request},
			`where songs.artist like "%Color%" 
			or songs.song like "%Color%"
			or genres.name like "%Color%" `,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchQuery(tt.args.r); got != tt.want {
				t.Errorf("searchQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_buildQueryParams(t *testing.T) {
	type args struct {
		r *http.Request
	}
	request, _ := http.NewRequest("GET", "songs/byLength?max=200&min=100", nil)
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			"Test case 1",
			args{request},
			`where songs.length < 200 and songs.length > 100`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := buildQueryParams(tt.args.r); got != tt.want {
				t.Errorf("buildQueryParams() = %v, want %v", got, tt.want)
			}
		})
	}
}
