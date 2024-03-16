package sio_event

type SocketConnectEvent struct {
	Shopid int `json:"shopid,string"`
}

type SocketDisconnectedEvent struct {
	Shopid int `json:"shopid,string"`
}

type SocketSyncEvent struct {
	Shopid int `json:"shopid,string"`
}
