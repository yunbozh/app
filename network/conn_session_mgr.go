package network

import (
	"app/common/logger"
	"sync"
	"sync/atomic"
)

type ConnSessionMgr struct {
	sessionMap sync.Map

	connIdx   int32 // 用于产生连接ID
	connCount int32 // 当前连接数量
}

func NewConnSessionMgr() *ConnSessionMgr {
	mgr := &ConnSessionMgr{
		sessionMap: sync.Map{},
		connIdx:    0,
		connCount:  0,
	}

	mgr.SetBaseId(10000)

	return mgr
}

func (self *ConnSessionMgr) SetBaseId(baseId int32) {
	atomic.StoreInt32(&self.connIdx, baseId)
}

func (self *ConnSessionMgr) GenConnIdx() int32 {
	return atomic.AddInt32(&self.connIdx, 1)
}

func (self *ConnSessionMgr) Add(session ConnSessionIf) {
	connIdx := self.GenConnIdx()

	atomic.AddInt32(&self.connCount, 1)

	session.(interface{ SetID(uint32) }).SetID(uint32(connIdx))

	self.sessionMap.Store(connIdx, session)
}

func (self *ConnSessionMgr) Remove(session ConnSessionIf) {
	self.sessionMap.Delete(session.GetID())

	atomic.AddInt32(&self.connCount, -1)
}

// 获取会话数量
func (self *ConnSessionMgr) GetCount() int32 {
	return atomic.LoadInt32(&self.connCount)
}

// 根据ID获取一个会话
func (self *ConnSessionMgr) GetSession(conn_idx int32) ConnSessionIf {
	if v, ok := self.sessionMap.Load(conn_idx); ok {
		return v.(ConnSessionIf)
	}

	return nil
}

// 遍历所有会话并执行回掉函数
func (self *ConnSessionMgr) DispatchSession(cb func(ConnSessionIf) bool) {
	self.sessionMap.Range(func(k, v interface{}) bool {
		return cb(v.(ConnSessionIf))
	})
}

// 关闭所有会话
func (self *ConnSessionMgr) CloseAllSession() {
	self.DispatchSession(func(session ConnSessionIf) bool {
		session.Close()

		return true
	})
}

// 给指定连接ID发送消息
func (self *ConnSessionMgr) SendMsg(conn_idx int32, msgId uint32, msg interface{}) error {
	session := self.GetSession(conn_idx)
	if session == nil {
		logger.Error("connextion not exist, conn_idx: %d", conn_idx)
		return nil
	}

	return session.SendMsg(msgId, msg)
}
