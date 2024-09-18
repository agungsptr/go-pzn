package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	expected := "Hello agung"
	actual := HelloWorld("agung")
	assert.Equal(t, expected, actual, "Result must be \"Hello agung\"")
}

func TestHelloWorldRequire(t *testing.T) {
	expected := "Hello agung"
	actual := HelloWorld("agung")
	require.Equal(t, expected, actual, "Result must be \"Hello agung\"")
}

func TestMain(m *testing.M) {
	// To manipulate how the unit test will run, use this function.
	// Function name must ne TestMain.
	// In this function we can use to define what should run before
	// and after unit test running.
	// This TestMain function ONLY will execute ONCE per package.

	// Before unit test running
	fmt.Print("BEFORE UNIT TEST RUN\n")

	m.Run()

	// After all unit test done
	fmt.Print("AFTER ALL UNIT TEST DONE\n")
}
