package handlers

import (
	"encoding/json"
	"fmt"
	"main/internal/parcel"
	"main/internal/utils/render"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Storage interface {
	Set(parcel.Parcel)
	Get(string) parcel.Parcel
}

func Get(s Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p parcel.Parcel
		err := render.JSON(r, p)
		if err != nil {
			log.Err(err)
		}
		result := s.Get(p.Reciver)
		data, err := json.Marshal(result)
		if err != nil {
			log.Err(err)
		}

		fmt.Fprintf(w, string(data))
	})
}

func Set(s Storage) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var p parcel.Parcel
		err := render.JSON(r, p)
		if err != nil {
			log.Err(err)
		}
	})
}
