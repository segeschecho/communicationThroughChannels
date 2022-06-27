package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func raceCondition() {
	sum := 0
	//wg := new(sync.WaitGroup)  // tambien funciona, agregar un parametro a race1 para recibirlo
	//var wg sync.WaitGroup      // tambien funciona, agregar un parametro a race1 para recibirlo

	wg.Add(10)
	for i := 0; i < 10 ; i++ {
		go race1(&sum)
	}

	fmt.Println("Numero de Go rutinas:", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Numero de Gorutinas al final:", runtime.NumGoroutine())
	fmt.Println("Suma = ", sum)
}

func race1(sum *int) {
	for i := 0; i < 100; i++ {
		*sum += 1
		time.Sleep(10 * time.Millisecond)
	}

	wg.Done()  // o se puede hacer un defer wg.Done()
}