package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntToStringSlice(t *testing.T) {
	input := []int32{443, 80, 21, 8080}
	want := []string{"443", "80", "21", "8080"}

	assert.ElementsMatch(t, want, IntToStringSlice(input))
}
