package service

import (
	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// NewServer configures and returns a Server.
func NewServer() *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	negroniMiddleware := negroni.Classic()
	mux := mux.NewRouter()

	initRoutes(mux, formatter)
	negroniMiddleware.UseHandler(mux)

	return negroniMiddleware
}

func initRoutes(mx *mux.Router, formatter *render.Render) {
	mx.HandleFunc("/", rootHandler(formatter)).Methods("GET")
	mx.HandleFunc("/skus/{sku}", getFullfillmentStatusHandler(formatter)).Methods("GET")
}