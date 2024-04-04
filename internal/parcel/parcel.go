package parcel

type Parcel struct {
	ID      int    `json:"id,omitempty"`
	Data    []byte `json:"data"`
	Sender  string `sender:"sender"`
	Reciver string `reciver:"reciver"`
}
