package main

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/require"
)

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
