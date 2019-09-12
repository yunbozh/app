package main

import (
	"app/common/logger"
	"app/common/server"
	"app/def"
	"app/network"
)

var (
	serverName  = "gs"
	sGateServer *GateServer
)

func init() {
	sGateServer = &GateServer{}
}

type GateServer struct {
	server.ServerInfo
	mainLoop *server.MainLoop

	tcpServerForClient *network.TCPServer
	tcpServerForServer *network.TCPServer
	tcpClient          *network.TCPClient

	serverStubHolder *server.ServerStubHolder
	msgRouteHolder *server.MsgRouteHolder
}

func (self *GateServer) Init() {
	serverId := server.GetCmdLineArgs().ServerId
	conf := server.GetServerConf()

	if serverId <= def.INVALID_ID || serverId > uint(conf.MSCount) {
		logger.Error("invalid server id: %d", serverId)
		return
	}

	// 初始serverUid
	self.SetServerUid(def.SERVER_TYPE_GS, uint16(serverId))

	// 初始主循环
	self.mainLoop = server.NewMainLoop(self.Update)

	self.tcpServerForServer = network.NewTCPServer(&network.TCPServerOptions{
		Ip:               conf.GSAddr[0].IpForServer,
		Port:             conf.GSAddr[0].PortForServer,
		OnConnectHandler: self.TcpServerForServer_ConnectHandler,
		OnRecvHandler:    self.TcpServerForServer_RecvHandler,
		OnCloseHandler:   self.TcpServerForServer_CloseHandler,
	})

	self.tcpServerForClient = network.NewTCPServer(&network.TCPServerOptions{
		Ip:               conf.GSAddr[0].IpForClient,
		Port:             conf.GSAddr[0].PortForClient,
		OnConnectHandler: self.TcpServerForClient_ConnectHandler,
		OnRecvHandler:    self.TcpServerForClient_RecvHandler,
		OnCloseHandler:   self.TcpServerForClient_CloseHandler,
	})

	self.tcpClient = network.NewTCPClient(&network.TCPClientOptions{
		Ip:               conf.MSAddr.Ip,
		Port:             conf.MSAddr.Port,
		ConnNum:          1,
		OnConnectHandler: self.TcpClient_ConnectHandler,
		OnRecvHandler:    self.TcpClient_RecvHandler,
		OnCloseHandler:   self.TcpClient_CloseHandler,
	})

	// 服务器连接stub
	self.serverStubHolder = server.NewServerStubHolser()

	// 消息路由
	self.msgRouteHolder = server.NewMsgRouteHolder()
}

func (self *GateServer) Run() {
	self.tcpServerForServer.Start()
	self.tcpServerForClient.Start()
	self.tcpClient.Start()

	self.mainLoop.Start()
}

func (self *GateServer) Close() {
	self.mainLoop.Stop()

	self.tcpServerForServer.Close()
	self.tcpServerForClient.Close()
	self.tcpClient.Close()
}

func (self *GateServer) Update() {
	defer Recover()

	//now := time.Now().UnixNano() / 1e6
}
