package main

import (
	"os"

	"github.com/istherepie/logwrapper"
)

func main() {

	// Logging things
	logwrapper.Trace("Tracing stuff...")
	logwrapper.Debug("Some logs for the ops guys...")
	logwrapper.Info("Informing everyone about things...")
	logwrapper.Warning("I am soooo warning you!!!!")
	logwrapper.Error("This really does not work!")

	// Set log level
	logwrapper.SetLogLevel("INFO")

	logwrapper.Debug("You should no longer see this log...")
	logwrapper.Info("This one you should see however...")

	// Log to file
	file, err := os.OpenFile("logfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)

	if err != nil {
		logwrapper.Fatal(err)
	}

	defer file.Close()

	logwrapper.SetOutput(file)
	logwrapper.Info("Logging to logfile.txt")

}
