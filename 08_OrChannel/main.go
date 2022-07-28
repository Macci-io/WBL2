package main

import (
	"fmt"
	"reflect"
	"time"
)

/*
Реализовать функцию, которая будет объединять один или более done-каналов в single-канал,
	если один из его составляющих каналов закроется.

Очевидным вариантом решения могло бы стать выражение при использовании select,
	которое бы реализовывало эту связь, однако иногда неизвестно общее число done-каналов,
	с которыми вы работаете в рантайме. В этом случае удобнее использовать вызов единственной функции,
	которая, приняв на вход один или более or-каналов, реализовывала бы весь функционал.


Определение функции:
var or func(channels ...<- chan interface{}) <- chan interface{}
Пример использования функции:
sig := func(after time.Duration) <- chan interface{} {
    c := make(chan interface{})
    go func() {
        defer close(c)
        time.Sleep(after)
	}()
	return c
}

start := time.Now()
<-or (
    sig(2*time.Hour),
    sig(5*time.Minute),
    sig(1*time.Second),
    sig(1*time.Hour),
    sig(1*time.Minute),
)

fmt.Printf(“fone after %v”, time.Since(start))


*/

func main() {
	sig := func(after time.Duration) <-chan interface{} {
		c := make(chan interface{})
		go func() {
			defer close(c)
			time.Sleep(after)
		}()
		return c
	}

	// Способ 1 через reflect.Select()
	start := time.Now()
	<-or1(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Second*5),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("fone after %v\n", time.Since(start))

	// Способ 2 через цикл и select
	start = time.Now()
	<-or2(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Second*5),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Printf("ftwo after %v\n", time.Since(start))
}

func or1(channels ...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{})
	var cases = make([]reflect.SelectCase, 0, len(channels))
	for _, c := range channels {
		// Заполняем слайс каналами
		cases = append(cases, reflect.SelectCase{
			// Dir Обозначаем канал как приемный
			Dir: reflect.SelectRecv,
			// Записываем сам канал в структуру
			Chan: reflect.ValueOf(c),
		})
	}
	go func() {
		// reflect.Select будет заблокирована пока не появятся данные хотя бы на одном канале
		reflect.Select(cases)
		res <- struct{}{}
	}()
	return res
}

func or2(channels ...<-chan interface{}) <-chan interface{} {
	res := make(chan interface{})
	go func() {
		length := len(channels)
		if length == 0 {
			res <- struct{}{}
			return
		}
		for {
			for i := 0; i < length; i++ {
				select {
				case <-channels[i]:
					res <- struct{}{}
					return
				default:
				}
			}
		}
	}()
	return res
}
