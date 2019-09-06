package main

import "app/def"

type Session struct {
	// net中的连接ID
	connIdx int32
	// uid
	sessionUid int64

	// 客户端ip
	clientIP string
	// client uid (又severid 和 session_uid 组成)
	clientUid def.ClientUid
}
