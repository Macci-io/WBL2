package animal

import "time"

//Elephant животное слон
type Elephant struct {
	animal
}

//NewElephant конструктор
func NewElephant() (e *Elephant) {
	return &Elephant{*newAnimal("\u001B[1;35mElephant\u001B[0m", 6, time.Second)}
}
