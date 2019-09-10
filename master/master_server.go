package main

import (
	"app/common/logger"
	"app/common/server"
	"app/def"
	"app/network"
	"runtime/debug"
)

var (
	serverName    = "ms"
	sMasterServer *MasterServer
)

func init() {
	sMasterServer = &MasterServer{}
}

type MasterServer struct {
	server.ServerInfo
	mainLoop *server.MainLoop

	tcpServer *network.TCPServer
}

func (self *MasterServer) Init() {
	serverId := server.GetCmdLineArgs().ServerId
	conf := server.GetServerConf()

	if serverId <= def.INVALID_ID || serverId > uint(conf.MSCount) {
		logger.Error("invalid server id: %d", serverId)
		return
	}

	// 初始serverUid
	self.SetServerUid(def.SERVER_TYPE_MASTER, uint16(serverId))

	// 初始主循环
	self.mainLoop = server.NewMainLoop(self.Update)

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

	self.mainLoop.Start()
}

func (self *MasterServer) Close() {
	self.mainLoop.Stop()

	self.tcpServer.Close()
}

func (self *MasterServer) Update() {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("%v", err)
			logger.Error("%s", debug.Stack())
		}
	}()

	//now := time.Now().UnixNano() / 1e6


	// github.com/op/go-logging
	// github.com/phachon/go-logger
	// github.com/wonderivan/logger
	// github.com/gxlog/gxlog
	// github.com/ngaut/log
	// github.com/sdbaiguanghe/glog
}
