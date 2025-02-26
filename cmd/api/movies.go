package main

import (
	"fmt"
	"net/http"
	"time"

	"greenlight.sanjbh.net/internal/data"
)

func (app *application) createMovieHandler(w http.ResponseWriter, r *http.Request) {
	var data struct {
		Title   string       `json:"title"`
		Year    int32        `json:"year"`
		Runtime data.Runtime `json:"runtime"`
		Genres  []string     `json:"genres"`
	}

	/* if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	} */

	if err := app.readJSON(w, r, &data); err != nil {
		app.badRequestResponse(w, r, err)
		return
	}

	fmt.Fprintf(w, "%+v\n", data)
}

func (app *application) showMovieHandler(w http.ResponseWriter, r *http.Request) {

	id, err := app.readIDParam(r)
	if err != nil {
		// http.NotFound(w, r)
		app.notFoundResponse(w, r)
		return
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
		// app.logger.Println(err)
		// http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
		app.serverErrorResponse(w, r, err)
	}
}
