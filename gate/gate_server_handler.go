package main

func (self *GateServer) TcpServerForServer_ConnectHandler(connIdx int32) {
	logger.Debugf("new connection, connIdx: %d", connIdx)
}

func (self *GateServer) TcpServerForServer_RecvHandler(connIdx int32, msgId int32, msg interface{}) {
	logger.Debugf("recv msg: %v", msg)
}

func (self *GateServer) TcpServerForServer_CloseHandler (connIdx int32) {
	logger.Debugf("close connection, connIdx: %d", connIdx)
}

func (self *GateServer) TcpServerForClient_ConnectHandler(connIdx int32) {
	logger.Debugf("new connection, connIdx: %d", connIdx)
}

func (self *GateServer) TcpServerForClient_RecvHandler(connIdx int32, msgId int32, msg interface{}) {
	logger.Debugf("recv msg: %v", msg)
}

func (self *GateServer) TcpServerForClient_CloseHandler (connIdx int32) {
	logger.Debugf("close connection, connIdx: %d", connIdx)
}

func (self *GateServer) TcpClient_ConnectHandler(connIdx int32) {
	logger.Debugf("new connection, connIdx: %d", connIdx)
}

func (self *GateServer) TcpClient_RecvHandler(connIdx int32, msgId int32, msg interface{}) {
	logger.Debugf("recv msg: %v", msg)
}

func (self *GateServer) TcpClient_CloseHandler (connIdx int32) {
	logger.Debugf("close connection, connIdx: %d", connIdx)
}