package src

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitServer() {
	var mux = mux.NewRouter()

	http.ListenAndServe(":8080", mux)
}
