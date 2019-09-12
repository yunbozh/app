package def

// route type
type MsgRouteType uint16

const (
	MSG_ROUTE_TYPE_INVALID MsgRouteType = 0
	MSG_ROUTE_TYPE_CLIENT  MsgRouteType = 1 // 客户端消息
	MSG_ROUTE_TYPE_MS      MsgRouteType = 2 // 主服务器消息
	MSG_ROUTE_TYPE_GS      MsgRouteType = 3 // 网关服务器消息
	MSG_ROUTE_TYPE_DS      MsgRouteType = 4 // 数据库服务器消息
	MSG_ROUTE_TYPE_LS      MsgRouteType = 5 // 逻辑功能服务器消息
	MSG_ROUTE_TYPE_SS      MsgRouteType = 6 // 场景服务器消息
	MSG_ROUTE_TYPE_RS      MsgRouteType = 7 // 关系服务器消息
	MSG_ROUTE_TYPE_COUNT   MsgRouteType = 8
)

func ServerTypeToMsgRouteType(serverType ServerType) MsgRouteType {
	routeType := MSG_ROUTE_TYPE_INVALID

	switch serverType {
	case SERVER_TYPE_MS:
		routeType = MSG_ROUTE_TYPE_MS
	case SERVER_TYPE_GS:
		routeType = MSG_ROUTE_TYPE_GS
	case SERVER_TYPE_DS:
		routeType = MSG_ROUTE_TYPE_DS
	case SERVER_TYPE_LS:
		routeType = MSG_ROUTE_TYPE_LS
	case SERVER_TYPE_SS:
		routeType = MSG_ROUTE_TYPE_SS
	case SERVER_TYPE_RS:
		routeType = MSG_ROUTE_TYPE_RS
	}

	return routeType
}

// 消息区分
const (
	MSG_BASE_INTERVAL uint32 = 1000 // 消息间隔

	MSG_BASE_APP_INSIDE uint32 = 1000 // 服务器内部连接用

	MSG_BASE_C2GS uint32 = 11000
	MSG_BASE_C2LS uint32 = 12000

	MSG_BASE_GS2C  uint32 = 21000
	MSG_BASE_GS2LS uint32 = 22000
)
