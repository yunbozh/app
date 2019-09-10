package server

import "app/def"

type ServerInfo struct {
	uid  def.ServerUid
	stat uint16
}

func (self *ServerInfo) SetServerUid(serverType def.ServerType, serverId uint16) {
	self.uid.ServerType = serverType
	self.uid.ServerId = serverId
}

func (self *ServerInfo) GetServerUid() def.ServerUid {
	return self.uid
}
