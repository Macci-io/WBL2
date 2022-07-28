package main

import (
	"fmt"
	"time"
)

/*
Описание:
	Цепочка обязанностей — это поведенческий паттерн проектирования,
	который позволяет передавать запросы последовательно по цепочке обработчиков.
	Каждый последующий обработчик решает, может ли он обработать запрос сам
	и стоит ли передавать запрос дальше по цепи.
Преимущества:
	1. Уменьшает зависимость между клиентом и обработчиками.
	2. Реализует принцип единственной обязанности.
	3. Реализует принцип открытости/закрытости.
Недостатки:
	1. Запрос может остаться никем не обработанным.
*/

/*
	Цепочка ответственностей может быть как линейным, так и разветвленным.
	Ниже будет пример линейной односвязной цепочки на примере структуры Warrior
*/

func main() {
	warrior := NewWarrior("Прометей", 700, 220)
	warrior.Next = NewWarrior("Ахилес", 500, 350)
	warrior.Next.Next = NewWarrior("Геркулес", 1000, 500)

	enemy := NewEnemy("Змей Горыныч", 2400, 490)

	//Здесь будет использоваться цепочка ответственностей,
	//если первый воин не справится с врагом то на смену ему придет следующий воин
	StartWar(warrior, enemy)
}

//IHero интерфейс персонажа воина
type IHero interface {
	Attack(hero IHero)
	TakeDamage(hero IHero)
	GetInfo() string
	GetName() string
	IsAlive() bool
	GetPowerAttack() int
}

//Hero персонаж воин
type Hero struct {
	name        string
	health      int
	powerAttack int
}

//NewHero новый персонаж (конструктор)
func NewHero(name string, health, powerAttack int) *Hero {
	return &Hero{name: name, health: health, powerAttack: powerAttack}
}

//GetPowerAttack геттер силы атаки
func (h *Hero) GetPowerAttack() int {
	return h.powerAttack
}

//Attack атакует
func (h *Hero) Attack(hero IHero) {
	hero.TakeDamage(h)
	fmt.Println(h.name + " атаковал " + hero.GetName() + ".\n\t" + hero.GetInfo())
	time.Sleep(time.Second)
}

//TakeDamage получить урон
func (h *Hero) TakeDamage(hero IHero) {
	//fmt.Println(h.GetName() + " получил урон от " + hero.GetName())
	h.health -= hero.GetPowerAttack()
	//fmt.Println(h.GetInfo())
}

//GetInfo получить информацию
func (h *Hero) GetInfo() string {
	if h.IsAlive() {
		return fmt.Sprintf("%s, осталось здоровья: %d", h.GetName(), h.health)
	}
	return fmt.Sprintf("%s, погиб в бою", h.GetName())
}

//GetName геттер имени
func (h *Hero) GetName() string {
	return h.name
}

//IsAlive жив ли персонаж
func (h *Hero) IsAlive() bool {
	return h.health > 0
}

// Enemy структура врага
type Enemy struct{ Hero }

// NewEnemy конструктор врага
func NewEnemy(name string, health, powerAttack int) *Enemy {
	return &Enemy{*NewHero(name, health, powerAttack)}
}

// Warrior структура воинов
type Warrior struct {
	*Hero
	Next *Warrior
}

//NewWarrior конструктор
func NewWarrior(name string, health, powerAttack int) *Warrior {
	return &Warrior{NewHero(name, health, powerAttack), nil}
}

//StartWar точка входа в войну
func StartWar(warrior *Warrior, enemy *Enemy) {
	if enemy == nil {
		return
	}
	i := 1
	for warrior != nil {
		fmt.Printf("\033[0;31m------------------%d-й Раунд------------------\u001B[0m\n", i)
		i++
		//Warrior атакует врага
		warrior.Attack(enemy)
		if !enemy.IsAlive() {
			break
		}

		//Enemy атакует воина
		enemy.Attack(warrior)

		// Если воин погибает, то на смену ему приходит другой воин
		// согласно паттерну цепочки ответственностей
		if !warrior.IsAlive() {
			warrior = warrior.Next
		}
	}
	if !enemy.IsAlive() {
		fmt.Println("\t\tВраг побежден!")
		fmt.Println("Победитель:", warrior.GetInfo())
	} else if warrior != nil {
		fmt.Println("\t\tВоины получили поражение!")
		fmt.Println("Победитель:", enemy.GetInfo())
	}
}
