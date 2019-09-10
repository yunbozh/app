package def

// server type
type ServerType uint8

const (
	SERVER_TYPE_INVALID  ServerType = 0
	SERVER_TYPE_MASTER   ServerType = 1 // 主服务器
	SERVER_TYPE_GATE     ServerType = 2 // 网关服务器
	SERVER_TYPE_DATABASE ServerType = 3 // 数据库服务器
	SERVER_TYPE_LOGIC    ServerType = 4 // 逻辑功能服务器
	SERVER_TYPE_SCENE    ServerType = 5 // 场景服务器
	SERVER_TYPE_RELATION ServerType = 6 // 关系服务器
	SERVER_TYPE_COUNT    ServerType = 7
)

const (
	INVALID_ID = 0
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