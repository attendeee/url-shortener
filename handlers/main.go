package handlers

import (
	"net/http"

	"github.com/attendeee/url-shortener/utils"
	"github.com/gorilla/mux"
)

var All *mux.Router

func Init() {
	All = mux.NewRouter()
	http.Handle("/", utils.CorsMiddleware(All))

	All.HandleFunc("/view/create", func(w http.ResponseWriter, r *http.Request) {})
	All.HandleFunc("/view/urls", func(w http.ResponseWriter, r *http.Request) {})

	All.HandleFunc("/get", getAllUrls).Methods("GET")

	All.HandleFunc("/create", createShorthand).Methods("POST")

	All.HandleFunc("/url/{url}", redirectByUrlParam).Methods("GET")
}
