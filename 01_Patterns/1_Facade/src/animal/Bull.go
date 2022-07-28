package animal

import "time"

//Bull животное бык
type Bull struct {
	animal
}

//NewBull конструктор
func NewBull() (b *Bull) {
	return &Bull{*newAnimal("\u001B[1;31mBull\u001B[0m", 3, time.Millisecond*600)}
}
