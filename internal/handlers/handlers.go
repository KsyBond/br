package handlers

import (
	"encoding/json"
	"fmt"
	"main/internal/parcel"
	"main/internal/utils/render"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Databus interface {
	Set(string, any)
	Get(string) any
	All() any
}

// DefaultHandler is for handling everything that is not a match
func DefaultHandler(w http.ResponseWriter, r *http.Request) {
	log.Info().Msgf("DefaultHandler Serving:", r.URL.Path, "from", r.Host, "with method", r.Method)
	w.WriteHeader(http.StatusNotFound)
	Body := r.URL.Path + " is not supported.\n"
	fmt.Fprintf(w, "%s", Body)
}

func All(d databus) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWritter, r *http.Request){
		result := d.
	})
}

func Get(d Databus) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p parcel.Parcel
		err := render.Request2Struct(r, &p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Err(err)
		}
		result := d.Get(p.Reciver)
		data, err := json.Marshal(result)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			log.Err(err)
		}
		fmt.Fprintf(w, string(data))
	})
}

func Set(d Databus) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p parcel.Parcel
		err := render.Request2Struct(r, &p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Err(err)
		}
		d.Set(p.Reciver, p)
	})
}
