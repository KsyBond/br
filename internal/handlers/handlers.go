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
}

// TODO rm fmt
func Get(d Databus) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("get")
		var p parcel.Parcel
		err := render.Request2Struct(r, &p)
		if err != nil {
			log.Err(err)
		}
		result := d.Get(p.Reciver)
		data, err := json.Marshal(result)
		if err != nil {
			log.Err(err)
		}
		fmt.Fprintf(w, string(data))
	})
}

func Set(d Databus) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("set")
		var p parcel.Parcel
		err := render.Request2Struct(r, &p)
		if err != nil {
			log.Err(err)
		}
		d.Set(p.Reciver, p)
	})
}
