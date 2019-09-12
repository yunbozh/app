package protobuf

import (
	"app/common/logger"
	"app/pb/c2s"
	"encoding/binary"
	"errors"
	"github.com/golang/protobuf/proto"
)

// -------------------------
// | id | protobuf message |
// -------------------------

const (
	MSG_ID_LEN = 4
)

type Processor struct {
	msgRouter map[uint32]MsgHandler
}

type MsgHandler func([]byte)

func NewProcessor() *Processor {
	p := new(Processor)
	p.msgRouter = make(map[uint32]MsgHandler)
	return p
}

func (self *Processor) Register(msgId uint32, handler MsgHandler) {
	if _, ok := self.msgRouter[msgId]; ok {
		logger.Error("message %s is already registered", msgId)
		return
	}

	self.msgRouter[msgId] = handler
}

// goroutine safe
func (self *Processor) Route(msg interface{}) error {
	msgInfo, ok := msg.(*c2s.ReqLogin)
	if !ok {
		return errors.New("msg type error")
	}

	logger.Debug("recv msg: %v", msgInfo)

	return nil
}

// goroutine safe
func (self *Processor) Unmarshal(data []byte) (interface{}, error) {
	if len(data) < MSG_ID_LEN {
		return nil, errors.New("msg data too short")
	}

	//msgId := binary.BigEndian.Uint32(data)
	msg := new(c2s.ReqLogin)
	err := proto.UnmarshalMerge(data[4:], msg)

	return msg, err
}

// goroutine safe
func (self *Processor) Marshal(id uint32, msg interface{}) ([][]byte, error) {
	// 消息ID
	msgId := make([]byte, MSG_ID_LEN)
	binary.BigEndian.PutUint32(msgId, id)

	// 消息内容
	data, err := proto.Marshal(msg.(proto.Message))
	return [][]byte{msgId, data}, err
}
