package animal

import (
	"fmt"
	"sync"
	"time"
)

type animal struct {
	name      string
	powerFull int
	speed     time.Duration
}

func newAnimal(name string, powerFull int, speed time.Duration) (a *animal) {
	a = new(animal)
	a.powerFull = powerFull
	a.speed = speed
	a.name = name
	return
}

func (a *animal) DoWork(wg *sync.WaitGroup, todo string, count int) {
	defer wg.Done()
	now := time.Now()
	fmt.Println(a.name + " is starting work")
	for i := 0; i < count; i++ {
		time.Sleep(a.speed)
		fmt.Println("\t"+a.name+" is "+todo, "\u001B[;36m", time.Now().Sub(now), "\u001B[0m")
		now = time.Now()
		if i >= a.powerFull {
			fmt.Println(a.name + " is tired")
			return
		}
	}
	fmt.Println(a.name + " is completed " + todo)
}
