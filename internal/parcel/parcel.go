package parcel

type Parcel struct {
	ID      int    `json:"id"`
	Data    []byte `json:"data"`
	Sender  string `sender:"sender"`
	Reciver string `reciver:"reciver"`
}
