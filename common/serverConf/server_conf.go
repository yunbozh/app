package serverConf

import (
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
)

type ServerCount struct {
}

type MSAddr struct {
	Ip   string
	Port uint16
}

type GSAddr struct {
	IpForServer   string
	PortForServer uint16
	IpForClient   string
	PortForClient uint16
}

type serverConf struct {
	LogLevel string

	MSCount uint16
	GSCount uint16
	LSCount uint16
	DSCount uint16
	SSCount uint16
	RSCount uint16

	MSAddr MSAddr
	GSAddr []GSAddr
}

var (
	conf *serverConf
	once sync.Once
)

func GetServerConf() *serverConf {
	once.Do(func() {
		path, err := filepath.Abs("./server_conf.toml")
		if err != nil {
			panic(err)
		}
		if _, err := toml.DecodeFile(path, &conf); err != nil {
			panic(err)
		}
	})

	return conf
}
