package main

import (
	"app/common/logger"
	"app/def"
	"app/network"
	"app/pb/s2s"
	"github.com/golang/protobuf/proto"
)

func (self *MasterServer) TcpServer_ConnectHandler(connIdx uint32, ip string) {
	defer Recover()
	logger.Debug("new connection, connIdx: %d, IP: %s", connIdx, ip)
}

func (self *MasterServer) TcpServer_RecvHandler(connIdx uint32, msgId uint32, msg []byte) {
	defer Recover()
	logger.Debug("recv msg: %v", msg)

	// 内部连接消息
	if msgId < def.MSG_BASE_APP_INSIDE+def.MSG_BASE_INTERVAL {
		switch s2s.MSGID(msgId) {

		case s2s.MSGID_S2S_REQ_ENTER:
			req := s2s.ReqEnter{}
			proto.Unmarshal(msg, &req)

			serverUid := def.ServerUid{
				ServerType: def.ServerType(req.ServerType),
				ServerId:   uint16(req.ServerId),
			}

			if !self.serverStubHolder.AddStub(self.tcpServer, connIdx, serverUid) {
				return
			}

			// 返回
			rep := &s2s.RepEnter{
				MsgId:      s2s.MSGID_S2S_REP_ENTER,
				ServerType: uint32(self.GetServerUid().ServerType),
				ServerId:   uint32(self.GetServerUid().ServerId),
			}
			self.tcpServer.Send(connIdx, uint32(rep.GetMsgId()), rep)

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

func (self *MasterServer) TcpServer_CloseHandler(connIdx uint32) {
	defer Recover()
	logger.Debug("close connection, connIdx: %d", connIdx)

}
