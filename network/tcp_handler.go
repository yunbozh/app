package network

type OnConnectHandler func(connIdx uint32, ip string)
type OnRecvHandler func(connIdx uint32, msgId uint32, msg []byte)
type OnCloseHandler func(connIdx uint32)

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
