package socket

type MessageResp struct {
	CensoredReply     string `json:"censored_reply,omitempty"`
	OriginalReply     string `json:"original_reply,omitempty"`
	Timestamp         string `json:"timestamp,omitempty"`
	TimestampFmt      string `json:"timestamp_fmt,omitempty"`
	TimestampUnix     uint   `json:"timestamp_unix,omitempty"`
	TimestampUnixNano uint   `json:"timestamp_unix_nano,omitempty"`
}

type DataResp struct {
	MsgId             uint         `json:"msg_id"`
	From              string       `json:"from,omitempty"`
	FromUid           string       `json:"from_uid,omitempty"`
	FromUserName      string       `json:"from_user_name,omitempty"`
	FromRole          string       `json:"from_role,omitempty"`
	ReminderTicker    any          `json:"reminder_ticker,omitempty"`
	IsOpposite        bool         `json:"is_opposite,omitempty"`
	ToUid             uint         `json:"to_uid,omitempty"`
	ClientConnectTime string       `json:"client_connect_time,omitempty"`
	Thumbnail         string       `json:"thumbnail,omitempty"`
	Message           *MessageResp `json:"message,omitempty"`
	StartTime         string       `json:"start_time,omitempty"`
	ShowRating        bool         `json:"show_rating,omitempty"`
	LocalId           string       `json:"local_id,omitempty"`
	Source            string       `json:"source,omitempty"`
}

type DataSend struct {
	MessageId    uint   `json:"message_id"`
	Message      string `json:"message,omitempty"`
	StartTime    string `json:"start_time,omitempty"`
	ParentReply  any    `json:"parent_reply,omitempty"`
	From         string `json:"from,omitempty"`
	FromUserName string `json:"from_user_name,omitempty"`
	LocalId      string `json:"local_id,omitempty"`
	Source       string `json:"source,omitempty"`
}

type MessageId struct {
	MsgId uint `json:"msg_id"`
}

type BaseSocketType struct {
	Code uint        `json:"code"`
	Data interface{} `json:"data"`
}
