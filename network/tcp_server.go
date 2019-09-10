package network

import (
	"app/common/logger"
	"app/network/protobuf"
	"fmt"
	"net"
	"sync"
	"time"
)

const (
	// 最大连接数据
	TCP_CONNECT_MAX_COUNT = 10000
)

type TCPServer struct {
	options        TCPServerOptions

	Addr string

	ln   net.Listener
	lnWG sync.WaitGroup

	processor ProcessorIf
	msgParser *MsgParser

	connMgr *ConnSessionMgr
}

func NewTCPServer(options *TCPServerOptions) *TCPServer {
	tcpServer := new(TCPServer)
	tcpServer.options.Ip = options.Ip
	tcpServer.options.Port = options.Port
	tcpServer.options.OnConnectHandler = options.OnConnectHandler
	tcpServer.options.OnRecvHandler = options.OnRecvHandler
	tcpServer.options.OnCloseHandler = options.OnCloseHandler

	tcpServer.Addr = fmt.Sprintf("%s:%d", options.Ip, options.Port)

	return tcpServer
}

func (self *TCPServer) OnConnectHandler(connIdx uint32, ip string) {
	self.options.OnConnectHandler(connIdx, ip)
}

func (self *TCPServer) OnRecvHandler(connIdx uint32, msgId uint32, msg interface{}) {
	self.options.OnRecvHandler(connIdx, msgId, msgId)
}

func (self *TCPServer) OnCloseHandler(connIdx uint32) {
	self.options.OnCloseHandler(connIdx)
}

func (self *TCPServer) GetMsgParser() *MsgParser {
	return self.msgParser
}

func (self *TCPServer) GetProcessor() ProcessorIf {
	return self.processor
}

func (self *TCPServer) Start() {
	self.init()
	go self.run()

	// TODO 测试代码
	//go func() {
	//	data := &test.Phone{
	//		MsgId:  1001,
	//		Type:   1,
	//		Number: 10,
	//		Name:   "billy",
	//	}
	//
	//	for {
	//		time.Sleep(1 * time.Second)
	//
	//		self.connMgr.DispatchSession(func(session ConnSessionIf) bool {
	//			time.Sleep(1 * time.Millisecond)
	//			session.SendMsg(data.GetMsgId(), data)
	//
	//			return true
	//		})
	//	}
	//}()
}

func (self *TCPServer) Close() {
	self.ln.Close()
	self.lnWG.Wait()

	self.connMgr.CloseAllSession()
}

func (self *TCPServer) init() {
	ln, err := net.Listen("tcp", self.Addr)
	if err != nil {
		logger.Error("listen error, %v", err)
		return
	}

	self.ln = ln

	self.msgParser = NewMsgParser()
	self.processor = protobuf.NewProcessor()

	self.connMgr = NewConnSessionMgr()
}

func (self *TCPServer) run() {
	self.lnWG.Add(1)
	defer self.lnWG.Done()

	// 开始监听端口
	logger.Debug("server start, listen port: %d", self.options.Port)

	for {
		conn, err := self.ln.Accept()
		if err != nil {
			logger.Error("listener accept error, %v", err)

			if ne, ok := err.(net.Error); ok && ne.Temporary() {
				time.Sleep(10 * time.Millisecond)
				continue
			}

			break
		}

		if self.connMgr.GetCount() >= TCP_CONNECT_MAX_COUNT {
			conn.Close()

			logger.Error("max count conections")
			time.Sleep(10 * time.Millisecond)
			continue
		}

		go self.newConnect(conn)
	}
}

func (self *TCPServer) newConnect(conn net.Conn) {
	session := newConnSession(conn, self)
	self.connMgr.Add(session)
	logger.Debug("new connection, address: %s, connIdx: %d", conn.RemoteAddr().String(), session.GetID())

	session.Run()

	logger.Debug("close connection, address: %s, connIdx: %d", conn.RemoteAddr().String(), session.GetID())
	self.connMgr.Remove(session)
}
