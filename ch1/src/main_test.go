package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func add(a, b int) int {
	return a + b
}

func TestAdd(t *testing.T) {
	t.Log("Testing add function")
	result := add(11, 22)
	assert.Equal(t, result, 33)
}
