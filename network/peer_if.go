package network

type Peer interface {
	Start()
	Close()
	GetMsgParser() *MsgParser
	GetProcessor() ProcessorIf
	OnConnectHandler(connIdx int32)
	OnRecvHandler(connIdx int32, msgId int32, msg interface{})
	OnCloseHandler(connIdx int32)
}
