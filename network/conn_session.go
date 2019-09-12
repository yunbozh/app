package network

import (
	"app/common/logger"
	"app/network/protobuf"
	"encoding/binary"
	"errors"
	"net"
	"sync"
)

const (
	CHAN_WRITE_NUM = 100
)

type ConnSession struct {
	sync.Mutex

	connIdx   uint32 // 连接ID
	closeFlag bool

	conn      net.Conn
	connWG    sync.WaitGroup
	writeChan chan []byte

	// 消息解析
	msgParser *MsgParser
	processor ProcessorIf

	peer PeerIf
}

func newConnSession(conn net.Conn, peer PeerIf) *ConnSession {
	session := new(ConnSession)
	session.conn = conn
	session.writeChan = make(chan []byte, CHAN_WRITE_NUM)
	session.msgParser = peer.GetMsgParser()
	session.processor = peer.GetProcessor()
	session.peer = peer
	return session
}

func (self *ConnSession) Run() {
	self.connWG.Add(2)

	go self.recvLoop()
	go self.sendLoop()

	self.peer.OnConnectHandler(self.GetID(), self.RemoteAddr().String())
	self.connWG.Wait()
	self.peer.OnCloseHandler(self.GetID())
}

func (self *ConnSession) Close() {
	self.Lock()
	defer self.Unlock()

	if self.closeFlag {
		return
	}

	self.doWrite(nil)
	self.closeFlag = true
}

// goroutine write
func (self *ConnSession) sendLoop() {
	for data := range self.writeChan {
		if data == nil {
			break
		}

		_, err := self.conn.Write(data)
		if err != nil {
			break
		}
	}

	self.conn.Close()
	self.connWG.Done()
}

// goroutine read
func (self *ConnSession) recvLoop() {
	for {
		data, err := self.ReadMsg()
		if err != nil {
			logger.Error("read msg error: %v \n", err)
			break
		}

		if len(data) < protobuf.MSG_ID_LEN {
			logger.Error("read msg error, msg data too short")
			break
		}

		msgId := binary.BigEndian.Uint32(data)
		self.peer.OnRecvHandler(self.GetID(), msgId, data[protobuf.MSG_ID_LEN:])

	}

	self.Close()
	self.connWG.Done()
}

func (self *ConnSession) doDestroy() {
	self.conn.(*net.TCPConn).SetLinger(0)
	self.conn.Close()

	if !self.closeFlag {
		close(self.writeChan)
		self.closeFlag = true
	}
}

func (self *ConnSession) Destroy() {
	self.Lock()
	defer self.Unlock()

	self.doDestroy()
}

func (self *ConnSession) doWrite(msg []byte) {
	if len(self.writeChan) == cap(self.writeChan) {
		logger.Error("close conn: channel full")
		self.doDestroy()
		return
	}

	self.writeChan <- msg
}

func (self *ConnSession) Write(msg []byte) {
	self.Lock()
	defer self.Unlock()

	if self.closeFlag || msg == nil {
		return
	}

	self.doWrite(msg)
}

func (self *ConnSession) Read(msg []byte) (int, error) {
	return self.conn.Read(msg)
}

func (self *ConnSession) LocalAddr() net.Addr {
	return self.conn.LocalAddr()
}

func (self *ConnSession) RemoteAddr() net.Addr {
	return self.conn.RemoteAddr()
}

func (self *ConnSession) ReadMsg() ([]byte, error) {
	return self.msgParser.Read(self)
}

func (self *ConnSession) WriteMsg(args ...[]byte) error {
	return self.msgParser.Write(self, args...)
}

// 设定连接ID
func (self *ConnSession) SetID(id uint32) {
	self.connIdx = id
}

// 获取ID
func (self *ConnSession) GetID() uint32 {
	return self.connIdx
}

// 发送消息
func (self *ConnSession) SendMsg(msgId uint32, msg interface{}) error {
	buf, err := self.processor.Marshal(msgId, msg)
	if err != nil {
		return errors.New("marshal error")
	}

	return self.WriteMsg(buf...)
}
