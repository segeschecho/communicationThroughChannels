package main

import "testing"

func Test_goRoutinesWaitMainToStart2(t *testing.T) {
	dataToSend := []string{"1", "2", "3", "4"}

	goRoutinesWaitMainToStart2(dataToSend)
}
