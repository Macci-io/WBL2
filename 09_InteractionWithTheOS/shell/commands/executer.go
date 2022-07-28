package commands

import (
	"fmt"
	"log"
	"sync"
	"syscall"
)

//ICommand интерфейс для команд
type ICommand interface {
	Run(*sync.WaitGroup) error
	Writer() int
	SetWriter(int)
	Reader() int
	SetReader(int)
	Pid() uintptr
}

//ExecuteAll исполнитель команд
func ExecuteAll(executable []ICommand) {
	var (
		ok    error
		group sync.WaitGroup
	)

	for _, e := range executable {
		group.Add(1)
		if ok = e.Run(&group); ok != nil {
			if ok = fmt.Errorf("%v", ok); ok != nil {
				log.Fatal(ok)
			}
		} else if e.Pid() > 0 {
			if _, ok = syscall.Wait4(int(e.Pid()), nil, 0, nil); ok != nil {
				log.Fatal(ok)
			}
		}
		group.Wait()
	}
}
