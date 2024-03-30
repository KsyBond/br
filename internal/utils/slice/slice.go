package slice

import "main/internal/parcel"

func Remove(slice []parcel.Parcel, p parcel.Parcel) []parcel.Parcel {
	return append(slice[:p.ID], slice[p.ID+1:]...)
}
