package interfaces

//IJob интерфейс работы для выполнения работниками
type IJob interface {
	AddWorkers(...IWorkable)
	StartWork(int)
}
