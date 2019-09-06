package main

import (
	"app/common/serverConf"
	"app/def"
	"app/network"
	"runtime/debug"
	"time"
)

var (
	serverName = "gs"
	sGateServer *GateServer
)

func init() {
	sGateServer = &GateServer{}
}

type GateServer struct {
	serverUid def.ServerUid

	tcpServerForClient *network.TCPServer
	tcpServerForServer *network.TCPServer

	tcpClient *network.TCPClient

	mainLoopTimer *time.Timer
}

func (self *GateServer) Init() {
	conf := serverConf.GetServerConf()

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
}

func (self *GateServer) Run() {
	self.tcpServerForServer.Start()
	self.tcpServerForClient.Start()
	self.tcpClient.Start()

	self.MainLoop()
}

func (self *GateServer) Close() {
	self.mainLoopTimer.Stop()

	self.tcpServerForServer.Close()
	self.tcpServerForClient.Close()
	self.tcpClient.Close()
}

func (self *GateServer) MainLoop() {
	self.mainLoopTimer = time.AfterFunc(100*time.Millisecond, self.MainLoop)

	defer func() {
		if err := recover(); err != nil {
			logger.Errorf("%v", err)
			logger.Errorf("%s", debug.Stack())
		}
	}()

	self.Update(time.Now().UnixNano() / 1e6)
}

func (self *GateServer) Update(now int64) {

}
