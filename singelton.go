package main

import "fmt"

type SingletonCounter struct {
	count int
}

var instance *SingletonCounter

func GetInstance() *SingletonCounter {
	if instance == nil {
		instance = &SingletonCounter{}
	}
	return instance
}

func (s *SingletonCounter) Increment() {
	s.count++
}

func (s *SingletonCounter) GetCount() int {
	return s.count
}

func main() {
	counter1 := GetInstance()
	counter1.Increment()
	fmt.Println("Счетчик 1:", counter1.GetCount()) // Выводит "Счетчик 1: 1"

	counter2 := GetInstance()
	counter2.Increment()
	fmt.Println("Счетчик 2:", counter2.GetCount()) // Выводит "Счетчик 2: 2"

	// counter1 и counter2 - это один и тот же объект.
}
