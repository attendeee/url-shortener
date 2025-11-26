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

	All.NotFoundHandler = http.HandlerFunc(notFoundUrlHandler)

	All.HandleFunc("/view/add-url", viewCreate).Methods("GET")
	All.HandleFunc("/view/get-urls", viewGetAllUrls).Methods("GET")

	All.HandleFunc("/get", getAllUrls).Methods("GET")

	All.HandleFunc("/create", createShorthand).Methods("POST")

	All.HandleFunc("/url/{url}", redirectByUrlParam).Methods("GET")
}
