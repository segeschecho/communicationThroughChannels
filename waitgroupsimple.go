package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func waitGroupSimple() {
	wg := new(sync.WaitGroup)
	wg.Add(2)

	go runner1(wg)
	go runner2(wg)

	fmt.Println("Numero de go rutinas:", runtime.NumGoroutine())

	wg.Wait()  // Esperamos a que todas las gorutinas terminen

	fmt.Println("Programa terminado")
	fmt.Println("Numero de go rutinas:", runtime.NumGoroutine())
}

func runner1(wg *sync.WaitGroup){
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("Runner 1")
}

func runner2(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Runner 2")
}