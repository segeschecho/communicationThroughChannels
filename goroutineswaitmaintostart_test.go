package main

import "testing"

func Test_goRoutinesWaitMainToStart(t *testing.T) {
	dataToSend := []string{"1", "2", "3", "4"}

	goRoutinesWaitMainToStart(dataToSend)
}
