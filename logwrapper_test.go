package logwrapper

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"testing"
)

// We need a buffer to write to
// In order to test log output
var buf bytes.Buffer

func Setup() {
	fmt.Println("SETTING UP")

	// Set output to the buffer
	SetOutput(&buf)
}

func TearDown() {
	fmt.Println("TEARING DOWN")
	SetOutput(os.Stdout)
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

// TestInvalidLogLevel is testing that `SetLogLevel` will throw an error on an invalid log level
func TestInvalidLogLevel(t *testing.T) {
	err := SetLogLevel("BOGUS")

	if err == nil {
		t.Log("An error should have been thrown")
		t.Fail()
	}

}

// TestInvalidLogLevel is testing that `SetLogLevel` will throw an error on an invalid log level
func TestLogFormat(t *testing.T) {

	// Log test message
	var message string = "this is a log entry"

	// Create log entry
	Info(message)

	// The log format should look like this:
	// INFO 2020/09/22 17:34:16 INFO this is a log entry
	var pattern string = fmt.Sprintf("[A-Z]{1,4} [0-9]{1,4}/[0-9]{1,2}/[0-9]{1,2} [0-9]{1,2}:[0-9]{1,2}:[0-9]{1,2} %s", message)
	r, err := regexp.Compile(pattern)

	if err != nil {
		t.Log(err)
		t.Fail()
	}

	matched := r.MatchString(buf.String())
	t.Log(matched)

	if matched != true {
		t.Log("The log entry did not match the regex pattern")

		// Print the log output
		t.Log(buf.String())
		t.Fail()
	}
}

func TestLogLevelFilter(t *testing.T) {

	// WORKAROUND: Reset buffer
	buf.Reset()

	// Set log level
	SetLogLevel("INFO")

	// Log test message
	var message string = "this is a debug log entry"

	// Create debug log entry
	Debug(message)

	if len(buf.Bytes()) != 0 {
		t.Log("Log output should not be received in the buffer")
		t.Fail()
	}
}
