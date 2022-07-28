package animal

import "time"

//Monkey животное мартышка
type Monkey struct {
	animal
}

//NewMonkey конструктор
func NewMonkey() (e *Monkey) {
	return &Monkey{*newAnimal("\u001B[;32mMonkey\u001B[0m", 4, time.Millisecond*500)}
}
