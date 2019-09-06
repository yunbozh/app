package main

import (
	"app/common/consoleCmd"
)

func main() {
	sMasterServer.Init()
	sMasterServer.Run()

	// block
	consoleCmd.ConsoleRun()

	// close
	sMasterServer.Close()
}
