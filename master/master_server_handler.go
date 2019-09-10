package main

import "app/common/logger"

func (self *MasterServer) TcpServer_ConnectHandler(connIdx uint32, ip string) {
	logger.Debug("new connection, connIdx: %d, IP: %s", connIdx, ip)
}

func (self *MasterServer) TcpServer_RecvHandler(connIdx uint32, msgId uint32, msg interface{}) {
	logger.Debug("recv msg: %v", msg)

	
}

func (self *MasterServer) TcpServer_CloseHandler(connIdx uint32) {
	logger.Debug("close connection, connIdx: %d", connIdx)
}
