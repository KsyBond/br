package storage

import (
	"main/internal/parcel"
	"main/internal/queue"
)

type Storage struct {
	Queue queue.Queue
}

func (s *Storage) Set(item parcel.Parcel) {
	s.Queue.Enqueue(item)
}

func (s *Storage) Get(reciver string) parcel.Parcel {
	return s.Queue.Find(reciver)
}
