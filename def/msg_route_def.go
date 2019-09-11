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

func ServerTypeToRouteType(serverType ServerType) MsgRouteType {
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
