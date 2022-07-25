package main

import (
	"fmt"
	"time"
)

func goRoutinesWaitMainToStart2(dataToSend []string)  {

	readyChan := make(chan bool)
	dataChan := make(chan string)
	errChan := make(chan error)

	go waitToStart2(readyChan, dataChan, errChan)

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

// La idea es usar select para poder solucionar el problema
func waitToStart2(readyChan chan bool, dataChan chan string, errChan chan error) {
	var receivedData []string

	for {
		fmt.Println("Ciclo")
		time.Sleep(1 * time.Second)
		select {
			case <- readyChan:
				// Nos avisan que está ready el envío de la información entonces se sale
				fmt.Println("Recibida señal de Ready")
				break
			case d := <- dataChan:   // puede ser default tambien
				// en el caso default se espera por información
				fmt.Println(d)
				receivedData = append(receivedData, d)
				errChan <- nil
		}
	}

	fmt.Println("Ahora que main dijo que está Ready, se procesa la información")
	fmt.Println("Informacion recibida:", receivedData)
}