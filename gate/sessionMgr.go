package main

type SessionMgr struct {
	// key: net中的连接ID
	sessionMapByConnIdx map[int32]*Session

	// key: session uid
	sessionMapByUid map[int64]*Session


}
