package network

// peer type
type PeerType uint16

const (
	PEER_TYPE_TCP_SERVER  PeerType = 1
	PEER_TYPE_TCP_CLIENT  PeerType = 2
	PEER_TYPE_UDP_SERVER  PeerType = 3
	PEER_TYPE_UDP_CLIENT  PeerType = 4
	PEER_TYPE_WS_SERVER   PeerType = 5
	PEER_TYPE_WS_CLIENT   PeerType = 6
	PEER_TYPE_HTTP_SERVER PeerType = 7
	PEER_TYPE_HTTP_CLIENT PeerType = 8
)
