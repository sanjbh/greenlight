package main

import (
	"net/http"

	"greenlight.sanjbh.net/internal/data"
	"greenlight.sanjbh.net/internal/formatter"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	formatter.FprintF(w, "Create a new movie")
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		http.NotFound(w, r)
	}

	movie := data.Movie {
		ID: id,
		CreatedAt: ,
	}

	formatter.FprintF(w, "show the details of movie %d\n", id)
}
