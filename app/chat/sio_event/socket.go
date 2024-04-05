package sio_event

type SocketConnectEvent struct {
	Shopid int `json:"shopid,string"`
}

func NewSocketConnectEvent(shopid int) *SocketConnectEvent {
	return &SocketConnectEvent{
		Shopid: shopid,
	}
}

type SocketDisconnectedEvent struct {
	Shopid int `json:"shopid,string"`
}

func NewSocketDisconnectedEvent(shopid int) *SocketDisconnectedEvent {
	return &SocketDisconnectedEvent{
		Shopid: shopid,
	}
}

type SocketSyncEvent struct {
	Shopid int `json:"shopid,string"`
}

func NewSocketSyncEvent(shopid int) *SocketSyncEvent {
	return &SocketSyncEvent{
		Shopid: shopid,
	}
}
