package network

type PeerIf interface {
	Start()
	Close()
	GetMsgParser() *MsgParser
	GetProcessor() ProcessorIf
	OnConnectHandler(connIdx uint32, ip string)
	OnRecvHandler(connIdx uint32, msgId uint32, msg interface{})
	OnCloseHandler(connIdx uint32)

	Send(connIdx uint32, msgId uint32, msg interface{})
	GetPeerType() PeerType
}
