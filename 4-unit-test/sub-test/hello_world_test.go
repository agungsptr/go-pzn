package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubTest(t *testing.T) {
	t.Run("Agung", func(t *testing.T) {
		expected := "Hello agung"
		actual := HelloWorld("agung")
		assert.Equal(t, expected, actual, "Result must be \"Hello agung\"")
	})
	t.Run("Faisal", func(t *testing.T) {
		expected := "Hello faisal"
		actual := HelloWorld("faisal")
		assert.Equal(t, expected, actual, "Result must be \"Hello faisal\"")
	})
}
