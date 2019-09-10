package server

import "time"

var MAIN_LOOP_TIME = 100 * time.Millisecond

type MainLoop struct {
	timer *time.Timer

	loopHandler func()
}

func NewMainLoop(handler func()) *MainLoop {
	mainLoop := new(MainLoop)
	mainLoop.loopHandler = handler

	return mainLoop
}

func (self *MainLoop) Start() {
	self.timer = time.AfterFunc(MAIN_LOOP_TIME, self.loop)
}

func (self *MainLoop) Stop() {
	self.timer.Stop()
}

func (self *MainLoop) loop() {
	self.loopHandler()

	self.timer = time.AfterFunc(MAIN_LOOP_TIME, self.loop)
}
