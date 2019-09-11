package server

import "app/def"

type ServerStubIf interface {
	Send(msgId uint32, msg interface{})
	GetServerType() def.ServerType
}
