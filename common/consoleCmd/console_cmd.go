package consoleCmd

import (
	"bufio"
	"fmt"
	"os"
	"os/signal"
	"strings"
)

type ConsoleCmd struct {
}

func (self *ConsoleCmd) Run() {
	sig := make(chan os.Signal, 1)
	cmd := make(chan string, 1)
	go self.readSignal(sig)
	go self.readCmd(cmd)

	select {
	case s := <-sig:
		fmt.Printf("server closing down (signal: %v) \n", s)
	case s := <-cmd:
		fmt.Printf("server closing down (cmd: %v) \n", s)
	}
}

func (self *ConsoleCmd) readSignal(c chan os.Signal) {
	signal.Notify(c, os.Interrupt, os.Kill)
}

func (self *ConsoleCmd) readCmd(c chan string) {
	for {
		// 从标准输入读取字符串，以\n为分割
		text, err := bufio.NewReader(os.Stdin).ReadString('\n')
		if err != nil {
			break
		}

		// 去掉读入内容的空白符
		text = strings.TrimSpace(text)

		if text == "shutdown" {
			c <- text
			break
		}
	}
}

func ConsoleRun() {
	console := new(ConsoleCmd)
	console.Run()
}
