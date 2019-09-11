package main

import (
	"app/def"
	"reflect"
)

func RegisterClientMsg(msgId uint32, msgType reflect.Type, hander def.MsgHandler) bool {
	return sMasterServer.msgRouteHolder.RegisterMsg(def.ROUTE_TYPE_CLIENT, msgId, msgType, hander)
}

func RegisterMSMsg(msgId uint32, msgType reflect.Type, hander def.MsgHandler) bool {
	return sMasterServer.msgRouteHolder.RegisterMsg(def.ROUTE_TYPE_MS, msgId, msgType, hander)
}

func RegisterGSMsg(msgId uint32, msgType reflect.Type, hander def.MsgHandler) bool {
	return sMasterServer.msgRouteHolder.RegisterMsg(def.ROUTE_TYPE_GS, msgId, msgType, hander)
}

func RegisterDSMsg(msgId uint32, msgType reflect.Type, hander def.MsgHandler) bool {
	return sMasterServer.msgRouteHolder.RegisterMsg(def.ROUTE_TYPE_DS, msgId, msgType, hander)
}

func RegisterLSMsg(msgId uint32, msgType reflect.Type, hander def.MsgHandler) bool {
	return sMasterServer.msgRouteHolder.RegisterMsg(def.ROUTE_TYPE_LS, msgId, msgType, hander)
}

func RegisterSSMsg(msgId uint32, msgType reflect.Type, hander def.MsgHandler) bool {
	return sMasterServer.msgRouteHolder.RegisterMsg(def.ROUTE_TYPE_SS, msgId, msgType, hander)
}

func RegisterRSMsg(msgId uint32, msgType reflect.Type, hander def.MsgHandler) bool {
	return sMasterServer.msgRouteHolder.RegisterMsg(def.ROUTE_TYPE_RS, msgId, msgType, hander)
}
