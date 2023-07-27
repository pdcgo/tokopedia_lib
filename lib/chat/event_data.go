package chat

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"
)

type EventType int

const (
	ReadUserChatEvent = 301
	ChatEvent         = 103
)

type EventCode struct {
	Code EventType `json:"code"`
}

type EmitEventSocket struct {
	*EventCode
	Data interface{} `json:"data"`
}

func (event *EmitEventSocket) UnmarshalJSON(data []byte) error {
	type Alias EmitEventSocket

	code := EventCode{}

	err := json.Unmarshal(data, &code)
	if err != nil {
		return err
	}

	parse := func(event *EmitEventSocket) error {
		aux := &struct {
			*Alias
		}{
			Alias: (*Alias)(event),
		}
		return json.Unmarshal(data, &aux)
	}

	event.EventCode = &code
	switch code.Code {
	case ReadUserChatEvent:
		event.Data = &ReaduserChat{}
		err = parse(event)

	case ChatEvent:
		event.Data = &SendChat{}
		err = parse(event)

	default:
		code := strconv.Itoa(int(code.Code))
		err = errors.New("parsing socket json event error " + code)
	}

	return err
}

type RcvEventSocket struct {
	*EventCode
	Data interface{} `json:"data"`
}

func (event *RcvEventSocket) UnmarshalJSON(data []byte) error {
	type Alias RcvEventSocket

	code := EventCode{}

	err := json.Unmarshal(data, &code)
	if err != nil {
		return err
	}

	parse := func(event *RcvEventSocket) error {
		aux := &struct {
			*Alias
		}{
			Alias: (*Alias)(event),
		}
		return json.Unmarshal(data, &aux)
	}

	event.EventCode = &code
	switch code.Code {
	case ReadUserChatEvent:
		event.Data = &ReaduserChat{}
		err = parse(event)

	case ChatEvent:
		event.Data = &RcvChat{}
		err = parse(event)

	default:
		code := strconv.Itoa(int(code.Code))
		err = errors.New("parsing socket json event error " + code)
	}

	return err
}

type SendChat struct {
	MessageID    int64     `json:"message_id"`
	Message      string    `json:"message"`
	StartTime    time.Time `json:"start_time"`
	ParentReply  any       `json:"parent_reply"`
	From         string    `json:"from"`
	FromUserName string    `json:"from_user_name"`
	Source       string    `json:"source"`
	LocalID      string    `json:"local_id"`
}

type ReaduserChat struct {
	MsgID int64 `json:"msg_id"`
}

type ReminderTicker struct {
	Enable       bool   `json:"enable"`
	EnableClose  bool   `json:"enable_close"`
	FeatureID    int    `json:"feature_id"`
	LocalID      string `json:"local_id"`
	MainText     string `json:"main_text"`
	RegexMessage string `json:"regex_message"`
	SubText      string `json:"sub_text"`
	TickerType   string `json:"ticker_type"`
	URL          string `json:"url"`
	URLLabel     string `json:"url_label"`
}
type Message struct {
	CensoredReply     string    `json:"censored_reply"`
	OriginalReply     string    `json:"original_reply"`
	Timestamp         time.Time `json:"timestamp"`
	TimestampFmt      string    `json:"timestamp_fmt"`
	TimestampUnix     int64     `json:"timestamp_unix"`
	TimestampUnixNano int64     `json:"timestamp_unix_nano"`
}
type RcvChat struct {
	MsgID             int64          `json:"msg_id"`
	From              string         `json:"from"`
	FromUID           int            `json:"from_uid"`
	FromUserName      string         `json:"from_user_name"`
	FromRole          string         `json:"from_role"`
	Thumbnail         string         `json:"thumbnail"`
	ReminderTicker    ReminderTicker `json:"reminder_ticker"`
	IsOpposite        bool           `json:"is_opposite"`
	ToUID             int            `json:"to_uid"`
	Message           Message        `json:"message"`
	StartTime         time.Time      `json:"start_time"`
	ShowRating        bool           `json:"show_rating"`
	LocalID           string         `json:"local_id"`
	ClientConnectTime time.Time      `json:"client_connect_time"`
	Source            string         `json:"source"`
}
