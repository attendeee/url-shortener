package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"text/template"

	"github.com/attendeee/url-shortener/storage/db"
	"github.com/attendeee/url-shortener/storage/lite"
	"github.com/gorilla/mux"
	"github.com/mattn/go-sqlite3"
)

var addUrlTemplate *template.Template = template.Must(template.ParseFiles("./templates/add-url.html"))
var getUrlsTemplate *template.Template = template.Must(template.ParseFiles("./templates/get-urls.html"))

func viewCreate(w http.ResponseWriter, r *http.Request) {
	addUrlTemplate.Execute(w, nil)

}

func viewGetAllUrls(w http.ResponseWriter, r *http.Request) {
	urls, err := lite.Query.GetAll(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("Get all url on view error: ", err)
	}

	getUrlsTemplate.Execute(w, urls)

}

func notFoundUrlHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/view/get-urls", http.StatusSeeOther)
}

func getAllUrls(w http.ResponseWriter, r *http.Request) {
	urls, err := lite.Query.GetAll(context.Background())
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("Get all url error: ", err)
	}

	urlsResponse, err := json.Marshal(urls)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("Get all url error: ", err)
	}

	w.WriteHeader(http.StatusFound)
	w.Write([]byte(urlsResponse))

}

func createShorthand(w http.ResponseWriter, r *http.Request) {
	var url db.Url

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&url)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Fatalln("[1]Create url error: ", err)
	}

	if len(url.Longhand) == 0 || len(url.Shorthand) == 0 {
		w.WriteHeader(http.StatusBadRequest)
	}

	err = lite.Query.CreateShorthand(
		context.Background(),
		db.CreateShorthandParams{Longhand: url.Longhand, Shorthand: url.Shorthand},
	)

	if err != nil {
		if errors.Unwrap(err) == errors.Unwrap(sqlite3.ErrConstraint) {
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte("Shorthand already exists"))
			return
		}

		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalln("[2]Create url error: ", err)
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(url)

}

func redirectByUrlParam(w http.ResponseWriter, r *http.Request) {
	urlParam := mux.Vars(r)["url"]

	url, err := lite.Query.GetByShorthand(context.Background(), urlParam)
	if err != nil {
		if errors.Unwrap(err) != errors.Unwrap(sql.ErrNoRows) {
			log.Println("Redirect url error: ", err)
		}

		http.Redirect(w, r, "/view/get-urls", http.StatusSeeOther)
		return
	}

	http.Redirect(w, r, url.Longhand, http.StatusFound)

}
