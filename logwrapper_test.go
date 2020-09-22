package logwrapper

import (
	"fmt"
	"os"
	"testing"
)

func Setup() {
	fmt.Println("SETTING UP")
}

func TearDown() {
	fmt.Println("TEARING DOWN")
}

// TestMain will run `Setup` and `Teardown` and all the tests in between
func TestMain(m *testing.M) {

	// Run Setup
	Setup()

	// Run all the tests
	returnCode := m.Run()

	// Run teardown
	TearDown()

	// Pass on the exit codes
	os.Exit(returnCode)
}

// TestSetLogLevel is testing that the DefaultLogLevel will be set properly
func TestSetLogLevel(t *testing.T) {
	SetLogLevel("INFO")

	var expected int = 20

	if DefaultLogLevel != expected {
		t.Log("Default log level has not been set correctly")
		t.Fail()
	}

}
