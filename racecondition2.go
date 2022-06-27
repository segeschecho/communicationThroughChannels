package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var sum int32

func raceCondition2() {
	wg :=new(sync.WaitGroup)
	wg.Add(3)

	go race2("Hola", wg)
	go race2("que", wg)
	go race2("Talco", wg)

	fmt.Println("Num Gorutinas: ", runtime.NumGoroutine())
	wg.Wait()
	fmt.Println("Terminaron todas las gorutinas, numero actual: ", runtime.NumGoroutine())

	fmt.Println("Resultado de la suma: ", sum)
}

func race2(s string, wg *sync.WaitGroup) {
	fmt.Println("Procesando: ", s)

	sum = int32(len(s))
	runtime.Gosched()  // cede el lugar a otra gorutina.
	time.Sleep(1 * time.Second)

	wg.Done() // defer wg.Done()
}
