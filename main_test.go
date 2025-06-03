package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {

}

func TestDivde(t *testing.T) {
	result, err := division(10, 2)
	assert.Nil(t, err)
	assert.Equal(t, 5.0, result)
	assert.ErrorContains(t, err, "除数不能为0")
}

func division(a, b int) (int, error) {
	return a / b, nil
}
