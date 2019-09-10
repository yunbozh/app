package def

// msg route type
type MsgRouteType uint16

const (
	MSG_ROUTE_TYPE_INVALID MsgRouteType = 0
	MSG_ROUTE_TYPE_CLIENT  MsgRouteType = 1
	MSG_ROUTE_TYPE_MS      MsgRouteType = 2
	MSG_ROUTE_TYPE_GS      MsgRouteType = 3
	MSG_ROUTE_TYPE_LS      MsgRouteType = 4
	MSG_ROUTE_TYPE_SS      MsgRouteType = 5
	MSG_ROUTE_TYPE_DS      MsgRouteType = 6
	MSG_ROUTE_TYPE_RS      MsgRouteType = 7
	MSG_ROUTE_TYPE_COUNT   MsgRouteType = 8
)
