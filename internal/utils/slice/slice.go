package slice

import "main/internal/parcel"

func Remove(slice []parcel.Parcel, p parcel.Parcel) []parcel.Parcel {
	return append(slice[:p.ID], slice[p.ID+1:]...)
}

func RemoveIndex(slice []interface{}, index int) []interface{} {
	return append(slice[:index], slice[index+1:]...)
}
