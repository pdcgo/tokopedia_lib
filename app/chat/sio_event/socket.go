package sio_event

type SocketConnectEvent struct {
	Shopid int `json:"shopid"`
}

type SocketDisconnectedEvent struct {
	Shopid int `json:"shopid"`
}

type SocketSyncEvent struct {
	Shopid int `json:"shopid"`
}
