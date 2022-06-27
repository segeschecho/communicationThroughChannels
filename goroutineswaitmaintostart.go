package main

import (
	"fmt"
	"time"
)

// La idea de este ejercicio es que el main cree una gorutina y que esta ultima
// espere a procesar información cuando este lista. La main routina avisa cuando empezar

func closeDataWhenReady(ready chan bool, data chan string) {
	<-ready
	close(data)
}

func waitToStart(ready chan bool, data chan string, errorChan chan error){
	var dataArr []string

	go closeDataWhenReady(ready, data)

	for {
		d, openChan := <-data

		if !openChan {
			fmt.Println("Se cerro el channel, entonces se sale del ciclo")
			break
		}

		dataArr = append(dataArr, d)
		fmt.Println("Agregado: ", d)

		errorChan <- nil

	}

	fmt.Println("Ahora que main dijo que está Ready, se procesa la información")
	fmt.Println("Informacion recibida:", dataArr)
}

func goRoutinesWaitMainToStart() {
	dataToSend := []string{"1", "2", "3", "4"}

	readyChan := make(chan bool)
	dataChan := make(chan string)
	errChan := make(chan error)

	go waitToStart(readyChan, dataChan, errChan)

	for i := 0; i < len(dataToSend); i++ {
		dataChan <- dataToSend[i]

		err := <-errChan
		if err != nil {
			panic("Error")
		}

		time.Sleep(1 * time.Second)  // esperamos un tiempo para ver que funcione todo ok
	}

	fmt.Println("Avisando que esta Ready")
	readyChan <- true

	time.Sleep(5 * time.Second)
	fmt.Println("Finalizando main")
}


