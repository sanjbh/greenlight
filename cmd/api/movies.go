package main

import (
	"net/http"
	"time"

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

	movie := &data.Movie{
		ID:        id,
		CreatedAt: time.Now(),
		Title:     "Casablanca",
		Runtime:   102,
		Genres:    []string{"drama", "romance", "war"},
		Version:   1,
	}

	if err = app.writeJSON(w, http.StatusOK, envelope{"movie": movie}, nil); err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}

	formatter.FprintF(w, "show the details of movie %d\n", id)
}
