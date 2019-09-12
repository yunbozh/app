package main

import (
	"app/common/logger"
	"app/def"
	"app/network"
	"app/pb/s2s"
	"github.com/gogo/protobuf/proto"
)

func (self *GateServer) TcpServerForServer_ConnectHandler(connIdx uint32, ip string) {
	defer Recover()
	logger.Debug("new connection, connIdx: %d, IP: %s", connIdx, ip)
}

func (self *GateServer) TcpServerForServer_RecvHandler(connIdx uint32, msgId uint32, msg []byte) {
	defer Recover()
	logger.Debug("recv msg: %v", msg)
}

func (self *GateServer) TcpServerForServer_CloseHandler(connIdx uint32) {
	defer Recover()
	logger.Debug("close connection, connIdx: %d", connIdx)
}

func (self *GateServer) TcpServerForClient_ConnectHandler(connIdx uint32, ip string) {
	defer Recover()
	logger.Debug("new connection, connIdx: %d, IP: %s", connIdx, ip)

	sSessionMgr.AddSession(connIdx, ip)
}

func (self *GateServer) TcpServerForClient_RecvHandler(connIdx uint32, msgId uint32, msg []byte) {
	defer Recover()
	logger.Debug("recv msg: %v", msg)
}

func (self *GateServer) TcpServerForClient_CloseHandler(connIdx uint32) {
	defer Recover()
	logger.Debug("close connection, connIdx: %d", connIdx)

	sSessionMgr.RemoveSessionByConnIdx(connIdx)
}

func (self *GateServer) TcpClient_ConnectHandler(connIdx uint32, ip string) {
	defer Recover()
	logger.Debug("new connection, connIdx: %d", connIdx)

	req := &s2s.ReqEnter{
		MsgId:      s2s.MSGID_S2S_REQ_ENTER,
		ServerType: uint32(self.GetServerUid().ServerType),
		ServerId:   uint32(self.GetServerUid().ServerId),
	}

	self.tcpClient.Send(connIdx, uint32(req.GetMsgId()), req)
}

func (self *GateServer) TcpClient_RecvHandler(connIdx uint32, msgId uint32, msg []byte) {
	defer Recover()
	logger.Debug("recv msg: %v", msg)

	// 内部连接消息
	if msgId < def.MSG_BASE_APP_INSIDE+def.MSG_BASE_INTERVAL {
		switch s2s.MSGID(msgId) {

		case s2s.MSGID_S2S_REP_ENTER:
			reqEnter := s2s.ReqEnter{}
			proto.Unmarshal(msg, &reqEnter)

			serverUid := def.ServerUid{
				ServerType: def.ServerType(reqEnter.ServerType),
				ServerId:   uint16(reqEnter.ServerId),
			}

			if !self.serverStubHolder.AddStub(self.tcpClient, connIdx, serverUid) {
				return
			}
		}
		return
	}

	stub := self.serverStubHolder.GetStubByConnIdx(network.PEER_TYPE_TCP_SERVER, connIdx)
	if stub == nil {
		logger.Error("get tcp server stub fail, connIdx: %d", connIdx)
		return
	}

	self.msgRouteHolder.RouteMsg(stub, connIdx, msgId, msg)
}

func (self *GateServer) TcpClient_CloseHandler(connIdx uint32) {
	defer Recover()
	logger.Debug("close connection, connIdx: %d", connIdx)
}
