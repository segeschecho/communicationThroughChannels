package main

import (
	"fmt"
	"time"
)

func mainWaitForGoroutinesFinishes(){
	ch := make(chan int, 2)
	chDone := make(chan bool)

	go worker(ch, chDone)

	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)

	fmt.Println("Esperando a que termine el worker...")
	<-chDone  // me quedo esperando a que el worker termine
	fmt.Println("El worker terminó, se termina el proceso principal")
}

func worker(ch chan int, chDone chan bool) {
	for elem := range ch {
		fmt.Println("Procesando elemento:", elem)
		time.Sleep(1 * time.Second)
	}

	// cuando el worker termina, avisa por medio del chan
	fmt.Println("Avisando que se terminó con todo el trabajo.")
	chDone <- true
}


