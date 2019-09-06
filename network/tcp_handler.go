package network

type OnConnectHandler func(connIdx int32)
type OnRecvHandler func(connIdx int32, msgId int32, msg interface{})
type OnCloseHandler func(connIdx int32)

type TCPServerOptions struct {
	Ip               string
	Port             uint16
	OnConnectHandler OnConnectHandler
	OnRecvHandler    OnRecvHandler
	OnCloseHandler   OnCloseHandler
}

type TCPClientOptions struct {
	Ip               string
	Port             uint16
	ConnNum          uint16
	OnConnectHandler OnConnectHandler
	OnRecvHandler    OnRecvHandler
	OnCloseHandler   OnCloseHandler
}
