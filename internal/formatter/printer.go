package formatter

import (
	"fmt"
	"log"
	"net/http"
)

func FprintF(w http.ResponseWriter, params ...any) {
	var err error
	if len(params) == 2 {
		_, err = fmt.Fprintf(w, params[0].(string), params[1])
	} else {
		_, err = fmt.Fprintln(w, params[0].(string))
	}
	if err != nil {
		// http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err.Error())
	}
}
