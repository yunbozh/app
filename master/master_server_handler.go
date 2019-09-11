package main

import (
	"app/common/logger"
	"app/network"
	"runtime/debug"
)

func (self *MasterServer) TcpServer_ConnectHandler(connIdx uint32, ip string) {
	logger.Debug("new connection, connIdx: %d, IP: %s", connIdx, ip)

	defer func() {
		if err := recover(); err != nil {
			logger.Error("%v", err)
			logger.Error("%s", debug.Stack())
		}
	}()
}

func (self *MasterServer) TcpServer_RecvHandler(connIdx uint32, msgId uint32, msg interface{}) {
	logger.Debug("recv msg: %v", msg)

	defer func() {
		if err := recover(); err != nil {
			logger.Error("%v", err)
			logger.Error("%s", debug.Stack())
		}
	}()

	// 内部连接消息
	if msgId < 1000 {

		return
	}

	stub := self.serverStubHolder.GetStubByConnIdx(network.PEER_TYPE_TCP_SERVER, connIdx)
	if stub == nil {
		logger.Error("get tcp server stub fail, connIdx: %d", connIdx)
		return
	}

	self.msgRouteHolder.RouteMsg(stub, connIdx, msgId, msg)
}

func (self *MasterServer) TcpServer_CloseHandler(connIdx uint32) {
	logger.Debug("close connection, connIdx: %d", connIdx)

	defer func() {
		if err := recover(); err != nil {
			logger.Error("%v", err)
			logger.Error("%s", debug.Stack())
		}
	}()
}
