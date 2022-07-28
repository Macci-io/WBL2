package job

import (
	"Front/src/interfaces"
	"sync"
)

//Job сущность некой работы от которой все работы наследуются
type Job struct {
	workers []interfaces.IWorkable
	toDo    string
}

//NewJob конструктор
func NewJob(ToDo string) Job {
	return Job{toDo: ToDo}
}

//AddWorkers добавляет работников
func (l *Job) AddWorkers(worker ...interfaces.IWorkable) {
	l.workers = append(l.workers, worker...)
}

//StartWork работники начинают работать
func (l *Job) StartWork(HowMuchWork int) {
	wg := sync.WaitGroup{}
	wg.Add(len(l.workers))
	for _, w := range l.workers {
		go w.DoWork(&wg, l.toDo, HowMuchWork)
	}
	wg.Wait()
}

//DoItWork конкретный работник работает
func DoItWork(wg *sync.WaitGroup, job interfaces.IJob, workable ...interfaces.IWorkable) {
	job.AddWorkers(workable...)
	job.StartWork(5)
	wg.Done()
}
