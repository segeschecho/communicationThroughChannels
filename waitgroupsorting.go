package main

import (
	"fmt"
	"sync"
	"time"
)

// La idea es que se disparen gorutinas y que estas se ejecuten en orden segun el valor recibido
func waitGroupSorting(data []int) string {
	wg := new(sync.WaitGroup)
	wg.Add(len(data))

	for _, v := range data {
		go func(v int, wgg *sync.WaitGroup) {
			time.Sleep(time.Duration(v) * time.Millisecond)
			fmt.Println(v)
			wgg.Done()
		}(v, wg)
	}

	wg.Wait()

	return ""
}
