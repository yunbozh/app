package server

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
	msgType reflect.Type, msgHandler def.MsgHandler) bool {

	if route, ok := self.msgRouteMap[routeType]; ok {
		route.RegisterMsg(msgId, msgType, msgHandler)
		return false
	}

	if routeType == def.MSG_ROUTE_TYPE_INVALID || routeType >= def.MSG_ROUTE_TYPE_COUNT {
		logger.Error("msg route type no exist, type: %d", routeType)
		return false
	}

	route := NewMsgRoute()
	route.RegisterMsg(msgId, msgType, msgHandler)
	self.msgRouteMap[routeType] = route

	return true
}

func (self *MsgRouteHolder) RouteMsg(stub ServerStubIf, connIdx uint32, msgId uint32, msg interface{}) {
	routeType := def.ServerTypeToRouteType(stub.GetServerType())

	if route, ok := self.msgRouteMap[routeType]; ok {

		msgMeta := route.GetMessageById(msgId)
		if msgMeta == nil {
			logger.Error("msg not register, msgId: %d", msgId)
			return
		}

		if msgMeta.Handler != nil {
			msgMeta.Handler(stub, msg)
		}

	} else {

		logger.Error("msg route not exist, route type: %d", routeType)
	}
}
