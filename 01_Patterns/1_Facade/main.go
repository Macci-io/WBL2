package main

import (
	"Front/src/animal"
	"Front/src/interfaces"
	"Front/src/job"
	"math/rand"
	"sync"
	"time"
)

/*
	Реализовать паттерн «фасад».
	Объяснить применимость паттерна, его плюсы и минусы,
	а также реальные примеры использования данного примера на практике.
*/

/*
	Фасад позволяет скрыть сложные части подсистемы,
	давая доступ только к необходимым рычагам,
	тем самым упрощая использование системы для клиента,

Преимущества:
	* Изолирует клиентов от компонентов сложной подсистемы.
Недостатки:
	* Фасад рискует стать божественным объектом, привязанным ко всем классам программы.
		Божественный_объект - описывающий объект, который хранит в себе «слишком много» или делает «слишком много».
*/

func main() {
	//Использование фасада
	NewFacadeWorker(3).Run()
}

//FacadeWorker структура реализующая паттерн Фасад
type FacadeWorker struct {
	workers []interfaces.IWorkable
	jobs    []interfaces.IJob
}

//NewFacadeWorker конструктор для фасада принимающая в качестве параметра количество работников
func NewFacadeWorker(workers int) (fw *FacadeWorker) {
	fw = new(FacadeWorker)
	for i := 0; i < workers; i++ {
		rand.Seed(time.Now().UnixNano())
		switch rand.Intn(3) {
		case 0:
			fw.jobs = append(fw.jobs, job.NewDancer())
		case 1:
			fw.jobs = append(fw.jobs, job.NewLoader())
		case 2:
			fw.jobs = append(fw.jobs, job.NewPoster())
		}
	}
	for i := 0; i < workers; i++ {
		rand.Seed(time.Now().UnixNano())
		switch rand.Intn(3) {
		case 0:
			fw.workers = append(fw.workers, animal.NewBull())
		case 1:
			fw.workers = append(fw.workers, animal.NewMonkey())
		case 2:
			fw.workers = append(fw.workers, animal.NewElephant())
		}
	}
	return
}

//Run Запускаем работу
func (f FacadeWorker) Run() {
	wg := sync.WaitGroup{}
	if len(f.workers) >= len(f.jobs) {
		wg.Add(len(f.jobs))
		for i := range f.jobs {
			go job.DoItWork(&wg, f.jobs[i], f.workers[i])
		}
	} else {
		wg.Add(len(f.workers))
		for i := range f.workers {
			go job.DoItWork(&wg, f.jobs[i], f.workers[i])
		}
	}
	wg.Wait()
}
