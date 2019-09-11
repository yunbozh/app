package def

import "app/common/server"

// server type
type ServerType uint16

const (
	SERVER_TYPE_INVALID ServerType = 0
	SERVER_TYPE_MS      ServerType = 1 // 主服务器
	SERVER_TYPE_GS      ServerType = 2 // 网关服务器
	SERVER_TYPE_DS      ServerType = 3 // 数据库服务器
	SERVER_TYPE_LS      ServerType = 4 // 逻辑功能服务器
	SERVER_TYPE_SS      ServerType = 5 // 场景服务器
	SERVER_TYPE_RS      ServerType = 6 // 关系服务器
	SERVER_TYPE_COUNT   ServerType = 7
)

const (
	INVALID_ID    = 0
	INVALID_VALUE = 0

	// 同一类型服务器最大数量
	SERVER_MAX_COUNT = 5
)

type ServerUid struct {
	// server type
	ServerType ServerType
	// server id
	ServerId uint16
}

type ClientUid struct {
	// server Id
	ServerId uint16
	// session id
	SessionId uint32
}

// 消息处理函数
type MsgHandler func(server.ServerStubIf, interface{})
