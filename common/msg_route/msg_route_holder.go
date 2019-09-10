package msg_route

import (
	"app/common/logger"
	"app/def"
	"reflect"
)

type MsgRouteHolder struct {
	// key
	msgRouteMap map[def.MsgRouteType]*MsgRoute
}

func NewMsgRouteHolder() *MsgRouteHolder {
	holder := &MsgRouteHolder{
		msgRouteMap: make(map[def.MsgRouteType]*MsgRoute),
	}

	return holder
}

func (self *MsgRouteHolder) RegisterMsg(routeType def.MsgRouteType, msgId uint32,
	msgType reflect.Type, msgHandler MsgHandler) {

	if route, ok := self.msgRouteMap[routeType]; ok {
		route.RegisterMsg(msgId, msgType, msgHandler)
		return
	}

	if routeType == def.MSG_ROUTE_TYPE_INVALID || routeType >= def.MSG_ROUTE_TYPE_COUNT {
		logger.Error("msg route type no exist, type: %d", routeType)
		return
	}

	route := NewMsgRoute()
	route.RegisterMsg(msgId, msgType, msgHandler)
	self.msgRouteMap[routeType] = route
}

func (self *MsgRouteHolder) RecvMsg(routeType def.MsgRouteType, msgId uint32, msg interface{}) {
	if route, ok := self.msgRouteMap[routeType]; ok {
		msgMeta := route.GetMessageById(msgId)
		if msgMeta == nil {
			logger.Error("msg not register, msgId: %d", msgId)
			return
		}

		if msgMeta.Handler != nil {
			msgMeta.Handler(msg)
		}
	}
}
