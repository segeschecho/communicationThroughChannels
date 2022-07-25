package main

import (
	"fmt"
	"os"
	"time"
)

// La idea es crear dos gorutinas y que una cree un archivo, guarde cosas, y avise
// a la otra gorutina para que la lea y haga algo, el main se queda esperando

//// USar Select ?? en algun lado?

var filename = "test.txt"

func writeFileRoutine(writedFileChan chan bool) {
	f, err := os.Create(filename)
	checkErr(err)
	defer f.Close()

	fmt.Println("[writeFileRoutine] Archivo creado y guardando mensaje...")

	m := "Estamos enviando un mensaje a traves del archivo. Suerte!"
	f.Write([]byte(m))
	time.Sleep(1 * time.Second)

	// una vez finalizado, avisamos que el archivo ya está escrito
	writedFileChan <- true
	fmt.Println("[writeFileRoutine] Mensaje guardado y avisando...")
}

func readFileRoutine(writedFileChan chan bool, done chan bool) {
	// espero a que el archivo a leer este disponible
	fmt.Println("[readFileRoutine] Esperando a que esté disponible el archivo...")
	<-writedFileChan

	fmt.Println("[readFileRoutine] Señal recibida, empezando a leer el archivo...")
	b, err := os.ReadFile(filename)
	checkErr(err)

	fmt.Println("[readFileRoutine] El mensaje leido es: " + string(b))
	done <- true
	fmt.Println("[readFileRoutine] Avisando done")
}

func filesBetweenRoutines(){
	writedFileChan := make(chan bool)
	done := make(chan bool)

	go writeFileRoutine(writedFileChan)
	go readFileRoutine(writedFileChan, done)

	// el main se queda esperando
	fmt.Println("[main] Esperando a que finalicen la gorutinas")
	<-done
	fmt.Println("[main] Las rutinas terminaron...")
}
