package main

import (
	"app/def"
)

type Session struct {
	// net中的连接ID
	connIdx uint32
	// session id
	sessionId uint32

	// 客户端ip
	clientIP string
	// client uid (由severId 和 sessionId 组成)
	clientUid def.ClientUid
}

func NewSession(conndIdx, sessionId uint32, ip string) *Session {
	session := &Session{
		connIdx:   conndIdx,
		sessionId: sessionId,
		clientIP:  ip,
		clientUid: def.ClientUid{
			ServerId:  sGateServer.GetServerUid().ServerId,
			SessionId: sessionId,
		},
	}

	return session
}

//
// get
//

func (self *Session) GetConnIdx() uint32 {
	return self.connIdx
}

func (self *Session) GetSessionId() uint32 {
	return self.sessionId
}

func (self *Session) GetClientIP() string {
	return self.clientIP
}

func (self *Session) GetClientUid() def.ClientUid {
	return self.clientUid
}
