package network

import (
	"net"
)

type ConnSessionIf interface {
	ReadMsg() ([]byte, error)
	WriteMsg(args ...[]byte) error
	LocalAddr() net.Addr
	RemoteAddr() net.Addr
	Close()
	Destroy()

	SetID(int32)
	GetID() int32
	SendMsg(msgId uint32, msg interface{}) error
}
