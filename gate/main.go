package main

import (
	"app/common/consoleCmd"
)

func main() {
	sGateServer.Init()
	sGateServer.Run()

	// block
	consoleCmd.ConsoleRun()

	// close
	sGateServer.Close()
}
