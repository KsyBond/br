package parcel

type Parcel struct {
	ID      int    `json:"id,omitempty"`
	Data    []byte `json:"data"`
	Sender  string `json:"sender"`
	Reciver string `json:"reciver"`
}
