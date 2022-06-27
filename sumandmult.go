package main

import (
	"fmt"
	"time"
)

func sumAndMulti() {
	chOperations := make(chan int, 2)
	chResults := make(chan int, 1)

	go sumCh(chOperations, 1)
	go multCh(chOperations, chResults, 2)

	time.Sleep(1 * time.Second)

	for result := range chResults {
		fmt.Println("Resultado: ", result)
		time.Sleep(1 * time.Second)
	}
}

// funcion que suma un entero a una serie de valores
func sumCh(chOperations chan int, toAdd int) {
	for i := 0; i < 4 ; i++ {
		chOperations <- i + toAdd
		fmt.Println("Suma: ", i, "+", toAdd)
		time.Sleep(1 * time.Second)
	}
	// Cuando termino de usar el channel, lo cierro porque no se va a asignar mas nada
	close(chOperations)
}

// Funcion que multiplica un valor a una serie de valores recibidos por un channel
func multCh(chOperations chan int, chResults chan int, toMult int) {
	// cuando haya elementos en el channel, hacer la multiplicacion y enviarlo al resultado
	for elem := range chOperations {
		chResults <- elem * toMult
		fmt.Println("Multiplica: ", elem, "*", toMult)
		time.Sleep(1 * time.Second)
	}
	close(chResults)
}

