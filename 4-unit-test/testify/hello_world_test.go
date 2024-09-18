package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHelloWorldAssert(t *testing.T) {
	// If unit test is fail, assert will call t.Fail()
	// So code after the assert will continue to execute
	expected := "Hello agung"
	actual := HelloWorld("agung")
	assert.Equal(t, expected, actual, "Result must be \"Hello agung\"")
}

func TestHelloWorldRequire(t *testing.T) {
	// If unit test is fail, require will call t.FailNow()
	// So code after the assert will NOT continue to execute
	expected := "Hello agung"
	actual := HelloWorld("agung")
	require.Equal(t, expected, actual, "Result must be \"Hello agung\"")
}

func TestHelloWorldSkip(t *testing.T) {
	// To skip unit test
	// For example the unit test can't run in macOS
	if runtime.GOOS == "darwin" {
		t.Skip("Unit test can't run in macOS")
	}

	expected := "Hello agung"
	actual := HelloWorld("agung")
	require.Equal(t, expected, actual, "Result must be \"Hello agung\"")
}
