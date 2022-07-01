package main

import (
"fmt"
"io/ioutil"
"net/http"
"time"
)

var addr = "127.0.0.1:1111"

type holaHandler struct {
	MaxRequests    int
	ActualRequests int
	Ready          chan bool
}

func (h *holaHandler) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("Hola para vos tambien"))
	checkErr(err)

	h.ActualRequests++

	// Cuando se recibe la cantidad de request maximo, se avisa que todo esta listo
	if h.ActualRequests > h.MaxRequests {
		h.Ready <- true
	}
}

func (h *holaHandler) Handler(rw http.ResponseWriter, r *http.Request) {}

func client() {
	client := http.Client{}

	time.Sleep(1 * time.Second)

	// Envia una cantidad de request al server
	for i := 0; i < 10; i++ {
		resp, err := client.Get("http://" + addr + "/hola")

		checkErr(err)
		defer resp.Body.Close() // hay que cerrar el body, por que??

		body, err := ioutil.ReadAll(resp.Body)
		checkErr(err)

		fmt.Println(string(body))
		time.Sleep(1 * time.Second)
	}
}

func server (ready chan bool) {
	mux := new(http.ServeMux)

	h := &holaHandler{MaxRequests: 3, Ready: ready}

	mux.Handle("/hola", h)

	err := http.ListenAndServe(addr, mux)
	checkErr(err)
}

func serverHttp() {
	ready := make(chan bool)

	go server(ready)
	go client()

	<-ready   // cuando está todo listo se termina el programa
	fmt.Println("Se cierra todo")
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
