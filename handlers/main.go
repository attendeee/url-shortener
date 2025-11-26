package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/attendeee/url-shortener/storage/db"
	"github.com/attendeee/url-shortener/storage/lite"
	"github.com/attendeee/url-shortener/utils"
	"github.com/gorilla/mux"
	"github.com/mattn/go-sqlite3"
)

var All *mux.Router

func Init() {
	All = mux.NewRouter()
	http.Handle("/", utils.CorsMiddleware(All))

	All.HandleFunc("/view/create", func(w http.ResponseWriter, r *http.Request) {})
	All.HandleFunc("/view/urls", func(w http.ResponseWriter, r *http.Request) {})

	All.HandleFunc("/get", func(w http.ResponseWriter, r *http.Request) {
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

	}).Methods("GET")

	All.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		var url db.Url

		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&url)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Fatalln("[1]Create url error: ", err)
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

	}).Methods("POST")

	All.HandleFunc("/url/{url}", func(w http.ResponseWriter, r *http.Request) {
		urlParam := mux.Vars(r)["url"]

		url, err := lite.Query.GetByShorthand(context.Background(), urlParam)
		if err != nil {
			log.Println("Redirect url error: ", err)
			http.Redirect(w, r, "/read-urls", http.StatusSeeOther)
			return
		}

		http.Redirect(w, r, url.Longhand, http.StatusFound)

	}).Methods("GET")
}
