package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_waitGroupSorting(t *testing.T) {
	data := []int{4, 5, 2, 1, 3}

	result := waitGroupSorting(data)
	assert.Equal(t, "", result)
}
