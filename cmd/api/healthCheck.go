package main

import (
	"net/http"
)

func (app *application) healthCheckHandler(w http.ResponseWriter, r *http.Request) {

	env := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}
	/* data := map[string]string{
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}
	*/
	/* js, err := json.Marshal(data)
	if err != nil {
		app.logger.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')
	w.Header().Set("Content-Type", "application/json")
	w.Write(js) */

	/* js := `{"status": "available", "environment": %q, "version": %q}`
	js = fmt.Sprintf(js, app.config.env, version)

	w.Header().Set("Content-Type", "application/json")

	w.Write([]byte(js)) */

	if err := app.writeJSON(w, http.StatusOK, env, nil); err != nil {
		app.logger.Println(err)
		http.Error(w, "The server encountered a problem and could not process your request", http.StatusInternalServerError)
	}
}
