package databus

type Databus[T any] struct {
	b  map[string]T
	ID []string
}

func New(T any) *Databus[any] {
	var d Databus[any]
	d.b = make(map[string]any)
	return &d
}

func (d *Databus[T]) Set(id string, item T) {
	d.b[id] = item
}

func (d *Databus[T]) Get(id string) any {
	temp := d.b[id]
	delete(d.b, id)
	return temp
}
