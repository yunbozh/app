package main

func (self *MasterServer) TcpServer_ConnectHandler(connIdx int32) {
	logger.Debugf("new connection, connIdx: %d", connIdx)
}

func (self *MasterServer) TcpServer_RecvHandler(connIdx int32, msgId int32, msg interface{}) {
	logger.Debugf("recv msg: %v", msg)
}

func (self *MasterServer) TcpServer_CloseHandler (connIdx int32) {
	logger.Debugf("close connection, connIdx: %d", connIdx)
}
