package main

import (
	"testing"
)

func TestHelloWorldFail(t *testing.T) {
	result := HelloWorld("agung")
	if result != "Hello agung" {
		// To fail the unit test
		t.Fail()
	}
}

func TestHelloWorldFailNow(t *testing.T) {
	result := HelloWorld("agung")
	if result != "Hello agung" {
		// Same like t.Fail()
		// with code after now to be execute
		t.FailNow()
	}
	// Any code here not gonna runing
}

func TestHelloWorldError(t *testing.T) {
	result := HelloWorld("agung")
	if result != "Hello agung" {
		// To fail the unit test
		// print error of the test
		// and call t.Fail()
		t.Error("Error msg")
	}
}

func TestHelloWorldFatal(t *testing.T) {
	result := HelloWorld("agung")
	if result != "Hello agung" {
		// Same like t.Error()
		// with code after now to be execute
		t.Fatal("Error msg")
	}
	// Any code here not gonna runing
}
