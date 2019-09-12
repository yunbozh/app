package server

import (
	"app/common/logger"
	"app/def"
	"app/network"
)

type ServerStubHolder struct {
	stubList [def.SERVER_TYPE_COUNT][]*ServerStub

	//key为连接ID
	stubMapForTcpServer map[uint32]*ServerStub
	//key为连接ID
	stubMapForTcpClient map[uint32]*ServerStub
}

func NewServerStubHolser() *ServerStubHolder {
	conf := GetServerConf()

	holder := new(ServerStubHolder)
	holder.stubList[def.SERVER_TYPE_MS] = make([]*ServerStub, conf.MSCount+1)
	holder.stubList[def.SERVER_TYPE_GS] = make([]*ServerStub, conf.GSCount+1)
	holder.stubList[def.SERVER_TYPE_DS] = make([]*ServerStub, conf.DSCount+1)
	holder.stubList[def.SERVER_TYPE_LS] = make([]*ServerStub, conf.LSCount+1)
	holder.stubList[def.SERVER_TYPE_SS] = make([]*ServerStub, conf.SSCount+1)
	holder.stubList[def.SERVER_TYPE_RS] = make([]*ServerStub, conf.RSCount+1)

	holder.stubMapForTcpServer = make(map[uint32]*ServerStub)
	holder.stubMapForTcpClient = make(map[uint32]*ServerStub)

	return holder
}

func (self *ServerStubHolder) GetStubByServerUid(serverUid def.ServerUid) *ServerStub {
	if !self.checkServerUid(serverUid) {
		return nil
	}

	return self.stubList[serverUid.ServerType][serverUid.ServerId]
}

func (self *ServerStubHolder) GetStubByConnIdx(peerType network.PeerType, connIdx uint32) *ServerStub {
	if peerType == network.PEER_TYPE_TCP_SERVER {
		if stub, ok := self.stubMapForTcpServer[connIdx]; ok {
			return stub
		}

	} else if peerType == network.PEER_TYPE_TCP_CLIENT {

		if stub, ok := self.stubMapForTcpClient[connIdx]; ok {
			return stub
		}
	}

	return nil
}

func (self *ServerStubHolder) AddStub(peer network.PeerIf, connIdx uint32, serverUid def.ServerUid) bool {
	if !self.checkServerUid(serverUid) {
		return false
	}

	if self.stubList[serverUid.ServerType][serverUid.ServerId] != nil {
		logger.Error("server exist, serverUid: %v", serverUid)
		return false
	}

	if peer.GetPeerType() == network.PEER_TYPE_TCP_SERVER {
		if _, ok := self.stubMapForTcpServer[connIdx]; ok {
			logger.Error("tcp server connIdx exist, conndIdx: %d", connIdx)
			return false
		}

	} else if peer.GetPeerType() == network.PEER_TYPE_TCP_CLIENT {
		if _, ok := self.stubMapForTcpClient[connIdx]; ok {
			logger.Error("tcp client connIdx exist, conndIdx: %d", connIdx)
			return false
		}
	}

	stub := &ServerStub{
		peer:      peer,
		connIdx:   connIdx,
		serverUid: serverUid,
	}

	self.stubList[serverUid.ServerType][serverUid.ServerId] = stub

	if peer.GetPeerType() == network.PEER_TYPE_TCP_SERVER {
		self.stubMapForTcpServer[connIdx] = stub

	} else if peer.GetPeerType() == network.PEER_TYPE_TCP_CLIENT {
		self.stubMapForTcpClient[connIdx] = stub
	}

	return true
}

func (self *ServerStubHolder) checkServerUid(serverUid def.ServerUid) bool {
	if serverUid.ServerType == def.SERVER_TYPE_INVALID || serverUid.ServerType >= def.SERVER_TYPE_COUNT {
		logger.Error("server uid.ServeType error, uid: %v", serverUid)
		return false
	}

	if serverUid.ServerId == def.INVALID_ID || int(serverUid.ServerId) >= len(self.stubList[serverUid.ServerType]) {
		logger.Error("server uid.ServerId error, uid: %v", serverUid)
		return false
	}

	return true
}
