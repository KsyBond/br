package queue

import (
	"main/internal/parcel"
	"main/internal/utils/slice"
)

type Queue struct {
	items []parcel.Parcel
}

func (q *Queue) Enqueue(item parcel.Parcel) {
	q.items = append(q.items, item)
}

func (q *Queue) Dequeue() parcel.Parcel {
	if len(q.items) == 0 {
		return parcel.Parcel{}
	}
	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Find(reciver string) parcel.Parcel {
	var p parcel.Parcel
	for i := 0; i < len(q.items); i++ {
		if q.items[i].Reciver == reciver {
			p = q.items[i]
		}
	}
	q.items = slice.Remove(q.items, p)
	return p
}
