package msg_route

import (
	"path"
	"reflect"
	"strings"
)

// 消息处理函数
type MsgHandler func(interface{})

type MsgMeta struct {
	Id      uint32       // 消息ID
	Type    reflect.Type // 消息类型
	Handler MsgHandler   // 消息处理函数
}

// 消息全名
func (self *MsgMeta) FullName() string {
	var sb strings.Builder
	sb.WriteString(path.Base(self.Type.PkgPath()))
	sb.WriteString(".")
	sb.WriteString(self.Type.Name())

	return sb.String()
}
