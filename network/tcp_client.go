package network

import (
	"app/network/protobuf"
	"fmt"
	"net"
	"sync"
	"time"
)

type TCPClient struct {
	options TCPClientOptions

	sync.Mutex

	Addr    string
	ConnNum uint16

	connWG sync.WaitGroup

	msgParser *MsgParser
	processor ProcessorIf
	closeFlag bool

	connMgr *ConnSessionMgr
}

func NewTCPClient(options *TCPClientOptions) *TCPClient {
	tcpClient := new(TCPClient)
	tcpClient.options.Ip = options.Ip
	tcpClient.options.Port = options.Port
	tcpClient.options.ConnNum = options.ConnNum
	tcpClient.options.OnConnectHandler = options.OnConnectHandler
	tcpClient.options.OnRecvHandler = options.OnRecvHandler
	tcpClient.options.OnCloseHandler = options.OnCloseHandler

	tcpClient.ConnNum = options.ConnNum
	tcpClient.Addr = fmt.Sprintf("%s:%d", options.Ip, options.Port)

	return tcpClient
}

func (self *TCPClient) OnConnectHandler(connIdx int32) {
	self.options.OnConnectHandler(connIdx)
}

func (self *TCPClient) OnRecvHandler(connIdx int32, msgId int32, msg interface{}) {
	self.options.OnRecvHandler(connIdx, msgId, msg)
}

func (self *TCPClient) OnCloseHandler(connIdx int32) {
	self.options.OnCloseHandler(connIdx)
}

func (self *TCPClient) GetMsgParser() *MsgParser {
	return self.msgParser
}

func (self *TCPClient) GetProcessor() ProcessorIf {
	return self.processor
}

func (self *TCPClient) Start() {
	self.init()

	for i := uint16(0); i < self.ConnNum; i++ {
		time.Sleep(1 * time.Millisecond)
		go self.newConnect(self.Addr, nil)
	}

	// TODO 测试代码
	//go func() {
	//	data := &test.Phone{
	//		MsgId:  1002,
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

func (self *TCPClient) init() {
	self.Lock()
	defer self.Unlock()

	self.closeFlag = false

	self.msgParser = NewMsgParser()
	self.processor = protobuf.NewProcessor()

	self.connMgr = NewConnSessionMgr()
}

func (self *TCPClient) newConnect(addr string, id chan<- int32) {
	self.connWG.Add(1)
	defer self.connWG.Done()

	self.Lock()
	if self.closeFlag {
		self.Unlock()
		return
	}
	self.Unlock()

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		logger.Debugf("connect to %s error: %v", self.Addr, err)
		return
	}

	session := newConnSession(conn, self)
	self.connMgr.Add(session)
	logger.Debugf("new connection, address: %s", conn.RemoteAddr().String())

	if id != nil {
		id <- session.GetID()
	}

	session.Run()

	logger.Debugf("close connection, address: %s", conn.RemoteAddr().String())
	self.connMgr.Remove(session)

}

func (self *TCPClient) Connect(addr string) (int32, bool) {
	id := make(chan int32, 1)

	go self.newConnect(addr, id)

	select {
	case connIdx := <-id:
		return connIdx, true

	case <-time.After(5 * time.Second):
		return 0, false
	}
}

func (self *TCPClient) Close() {
	self.Lock()
	self.closeFlag = true
	self.Unlock()

	self.connMgr.CloseAllSession()

	self.connWG.Wait()
}
