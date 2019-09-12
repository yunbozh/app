package def

type ServerStubIf interface {
	Send(msgId uint32, msg interface{})
	GetServerType() ServerType
}
