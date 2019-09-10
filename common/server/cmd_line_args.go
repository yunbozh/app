package server

import (
	"flag"
)

var (
	cmdLineArgs *CmdLineArgs
)

func init() {
	cmdLineArgs = new(CmdLineArgs)
	flag.UintVar(&cmdLineArgs.ServerId, "id", 1, "server id")
	flag.Parse()
}

type CmdLineArgs struct {
	ServerId uint
}

func GetCmdLineArgs() *CmdLineArgs {
	return cmdLineArgs
}
