package main

import (
	"app/common/serverConf"
	"app/network"
	"runtime/debug"
	"time"
)

var (
	serverName    = "ms"
	sMasterServer *MasterServer
)

func init() {
	sMasterServer = &MasterServer{}
}

type MasterServer struct {
	tcpServer *network.TCPServer

	mainLoopTimer *time.Timer
}

func (self *MasterServer) Init() {
	conf := serverConf.GetServerConf()

	self.tcpServer = network.NewTCPServer(&network.TCPServerOptions{
		Ip:               conf.MSAddr.Ip,
		Port:             conf.MSAddr.Port,
		OnConnectHandler: self.TcpServer_ConnectHandler,
		OnRecvHandler:    self.TcpServer_RecvHandler,
		OnCloseHandler:   self.TcpServer_CloseHandler,
	})
}

func (self *MasterServer) Run() {
	self.tcpServer.Start()

	self.MainLoop()
}

func (self *MasterServer) Close() {
	self.mainLoopTimer.Stop()

	self.tcpServer.Close()
}

func (self *MasterServer) MainLoop() {
	self.mainLoopTimer = time.AfterFunc(1000*time.Millisecond, self.MainLoop)

	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("%v", err)
			logger.Errorf("%s", debug.Stack())
		}
	}()

	self.Update(time.Now().UnixNano() / 1e6)
}

func (self *MasterServer) Update(now int64) {

	// github.com/op/go-logging
	// github.com/phachon/go-logger
	// github.com/wonderivan/logger
	// github.com/gxlog/gxlog
	// github.com/ngaut/log
}
