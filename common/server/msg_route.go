package server

import (
	"app/common/logger"
	"app/def"
	"reflect"
)

type MsgRoute struct {
	// key：消息ID
	msgMapById map[uint32]*MsgMeta
	// key：消息类型
	msgMapByType map[reflect.Type]*MsgMeta
}

func NewMsgRoute() *MsgRoute {
	msgRoute := new(MsgRoute)
	msgRoute.msgMapById = make(map[uint32]*MsgMeta)
	msgRoute.msgMapByType = make(map[reflect.Type]*MsgMeta)

	return msgRoute
}

// 消息注册
func (self *MsgRoute) RegisterMsg(msgId uint32, msgType reflect.Type, msgHandler def.MsgHandler) bool {
	if msgId == 0 || msgType == nil || msgHandler == nil {
		return false
	}

	// 类型统一为非指针类型
	if msgType.Kind() == reflect.Ptr {
		msgType = msgType.Elem()
	}

	if _, ok := self.msgMapById[msgId]; ok {
		// 消息已经注册
		logger.Error("repeat message register by msgId: %d", msgId)
		return false
	}

	if _, ok := self.msgMapByType[msgType]; ok {
		// 消息已经注册
		logger.Error("repeat message register by msgName: %s", msgType.Name())
		return false
	}

	msg := &MsgMeta{
		Id:      msgId,
		Type:    msgType,
		Handler: msgHandler,
	}

	self.msgMapById[msg.Id] = msg
	self.msgMapByType[msg.Type] = msg

	return true
}

// 根据消息ID获取元消息
func (self *MsgRoute) GetMessageById(msgId uint32) *MsgMeta {
	if v, ok := self.msgMapById[msgId]; ok {
		return v
	}

	return nil
}

// 根据消息类型获取元消息
func (self *MsgRoute) GetMessageByType(msgType reflect.Type) *MsgMeta {
	if msgType.Kind() == reflect.Ptr {
		msgType = msgType.Elem()
	}

	if v, ok := self.msgMapByType[msgType]; ok {
		return v
	}

	return nil
}

// 根据消息对象获取元消息
func (self *MsgRoute) GetMessageByMsg(msg interface{}) *MsgMeta {
	return self.GetMessageByType(reflect.TypeOf(msg))
}
