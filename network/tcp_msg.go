package network

import (
	"encoding/binary"
	"errors"
	"io"
)

// --------------
// | len | data |
// --------------

const (
	// 用2个字节来表示消息长度
	MSG_LEN = 2
	// 消息最大长度
	MSG_MAX_LEN = 65535

)

type MsgParser struct {
}

func NewMsgParser() *MsgParser {
	p := new(MsgParser)
	return p
}

// goroutine safe
func (self *MsgParser) Read(conn *ConnSession) ([]byte, error) {
	msgLenBuf := make([]byte, MSG_LEN)

	if _, err := io.ReadFull(conn, msgLenBuf); err != nil {
		return nil, err
	}

	msgLen := binary.BigEndian.Uint16(msgLenBuf)

	if msgLen > MSG_MAX_LEN {
		return nil, errors.New("message too long")
	}

	// data
	msgData := make([]byte, msgLen)
	if _, err := io.ReadFull(conn, msgData); err != nil {
		return nil, err
	}

	return msgData, nil
}

// goroutine safe
func (self *MsgParser) Write(conn *ConnSession, args ...[]byte) error {
	var msgLen uint16

	for i := 0; i < len(args); i++ {
		msgLen += uint16(len(args[i]))
	}

	// check len
	if msgLen > MSG_MAX_LEN {
		return errors.New("message too long")
	}

	msgBuf := make([]byte, MSG_LEN+msgLen)

	binary.BigEndian.PutUint16(msgBuf, msgLen)

	index := MSG_LEN
	for _, arg := range args {
		copy(msgBuf[index:], arg)
		index += len(arg)
	}

	conn.Write(msgBuf)

	return nil
}
