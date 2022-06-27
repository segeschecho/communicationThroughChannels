package main

import (
	"fmt"
	"time"
)

// la idea es armar una serie de workers que procesen información y una cola de tareas


func queueToProcess() {
	data := []string{"1", "2", "3", "4", "5", "6"}

	dataChan := make(chan []string)
	doneChan := make(chan bool)

	go queue(dataChan, doneChan)

	dataChan <- data

	<- doneChan
	println("Terminado el proceso de toda la información")
}

func queue(dataChan chan []string, doneChan chan bool) {
	// La cola de mensajes tiene dos workers solamente
	elemChan1, elemChan2 := make(chan string), make(chan string)

	w1 := WorkerForQueue{Name: "worker1", TimeToProcess: 1}
	w2 := WorkerForQueue{Name: "worker2", TimeToProcess: 2}

	go w1.workForQueue(elemChan1)
	go w2.workForQueue(elemChan2)

	ss := <- dataChan

	for _, s := range ss {
		for {
			if w1.Ready {
				elemChan1 <- s
				break
			}

			if w2.Ready {
				elemChan2 <- s
				break
			}
			// si los dos workers estan ocupados, espero un rato y vuelvo a preguntar
			time.Sleep(10 * time.Millisecond)
		}
	}

	doneChan <- true
}

type WorkerForQueue struct {
	Ready bool
	Name string
	TimeToProcess int
}

func (w *WorkerForQueue) workForQueue(element chan string) {

	for {
		w.Ready = true
		e := <-element
		w.Ready = false

		fmt.Println("Worker: ", w.Name, " procesa: ", e)
		time.Sleep(1 * time.Second)
	}
}