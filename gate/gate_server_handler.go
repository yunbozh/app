package main

import "app/common/logger"

func (self *GateServer) TcpServerForServer_ConnectHandler(connIdx uint32, ip string) {
	logger.Debug("new connection, connIdx: %d, IP: %s", connIdx, ip)
}

func (self *GateServer) TcpServerForServer_RecvHandler(connIdx uint32, msgId uint32, msg interface{}) {
	logger.Debug("recv msg: %v", msg)
}

func (self *GateServer) TcpServerForServer_CloseHandler(connIdx uint32) {
	logger.Debug("close connection, connIdx: %d", connIdx)
}

func (self *GateServer) TcpServerForClient_ConnectHandler(connIdx uint32, ip string) {
	logger.Debug("new connection, connIdx: %d, IP: %s", connIdx, ip)

	sSessionMgr.AddSession(connIdx, ip)
}

func (self *GateServer) TcpServerForClient_RecvHandler(connIdx uint32, msgId uint32, msg interface{}) {
	logger.Debug("recv msg: %v", msg)
}

func (self *GateServer) TcpServerForClient_CloseHandler(connIdx uint32) {
	logger.Debug("close connection, connIdx: %d", connIdx)

	sSessionMgr.RemoveSessionByConnIdx(connIdx)
}

func (self *GateServer) TcpClient_ConnectHandler(connIdx uint32, ip string) {
	logger.Debug("new connection, connIdx: %d", connIdx)
}

func (self *GateServer) TcpClient_RecvHandler(connIdx uint32, msgId uint32, msg interface{}) {
	logger.Debug("recv msg: %v", msg)
}

func (self *GateServer) TcpClient_CloseHandler(connIdx uint32) {
	logger.Debug("close connection, connIdx: %d", connIdx)
}
