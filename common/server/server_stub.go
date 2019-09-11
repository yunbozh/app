package server

import (
	"app/def"
	"app/network"
)

type ServerStub struct {
	// owner peer
	peer network.PeerIf

	// net中的连接ID
	connIdx uint32
	// server Uid
	serverUid def.ServerUid
}

func (self *ServerStub) Send(msgId uint32, msg interface{}) {
	self.peer.Send(self.connIdx, msgId, msg)
}

func (self *ServerStub) GetServerType() def.ServerType {
	return self.serverUid.ServerType
}
