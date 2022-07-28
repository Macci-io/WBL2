package interfaces

import "sync"

//IWorkable интерфейс работника для выполнения работы
type IWorkable interface {
	DoWork(wg *sync.WaitGroup, todo string, count int)
}
