package main

import "app/common/logger"

var sSessionMgr *SessionMgr

func init() {
	sSessionMgr = &SessionMgr{
		sessionMapByConnIdx:   make(map[uint32]*Session),
		sessionMapBySessionId: make(map[uint32]*Session),
		sessionId:             0,
	}
}

type SessionMgr struct {
	// key: net中的连接ID
	sessionMapByConnIdx map[uint32]*Session

	// key: sessionId
	sessionMapBySessionId map[uint32]*Session

	// 用于产生session id
	sessionId uint32
}

func (self *SessionMgr) GenSessionId() uint32 {
	self.sessionId++
	return self.sessionId
}

func (self *SessionMgr) AddSession(connIdx uint32, ip string) {
	if session := self.GetSessionByConnIdx(connIdx); session != nil {
		logger.Error("session exist, connIdx: %d", connIdx)
		return
	}

	sessionId := self.GenSessionId()
	session := NewSession(connIdx, sessionId, ip)

	self.sessionMapByConnIdx[connIdx] = session
	self.sessionMapBySessionId[sessionId] = session
}

func (self *SessionMgr) GetSessionByConnIdx(connIdx uint32) *Session {
	if session, ok := self.sessionMapByConnIdx[connIdx]; ok {
		return session
	}

	return nil
}

func (self *SessionMgr) GetSessionBySessionId(sessionId uint32) *Session {
	if session, ok := self.sessionMapBySessionId[sessionId]; ok {
		return session
	}

	return nil
}

func (self *SessionMgr) RemoveSessionByConnIdx(connIdx uint32) {
	if session, ok := self.sessionMapByConnIdx[connIdx]; ok {
		delete(self.sessionMapByConnIdx, connIdx)
		delete(self.sessionMapBySessionId, session.GetSessionId())
	}
}

func (self *SessionMgr) RemoveSessionBySessionId(sessionId uint32) {
	if session, ok := self.sessionMapBySessionId[sessionId]; ok {
		delete(self.sessionMapBySessionId, sessionId)
		delete(self.sessionMapByConnIdx, session.GetConnIdx())
	}
}
