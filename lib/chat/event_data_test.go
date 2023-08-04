package chat_test

import (
	"encoding/json"
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/chat"
	"github.com/stretchr/testify/assert"
)

func TestParsingEmitModel(t *testing.T) {
	t.Run("test chat read data model", func(t *testing.T) {
		data := `{"code":301,"data":{"msg_id":2483170385}}`

		event := chat.EmitEventSocket{}
		err := json.Unmarshal([]byte(data), &event)
		assert.Nil(t, err)
		assert.IsType(t, &chat.ReaduserChat{}, event.Data)
	})

	t.Run("test send model data", func(t *testing.T) {
		data := `{"code":103,"data":{"message_id":2483170385,"message":"asdasd","start_time":"2023-07-26T05:54:44.138Z","parent_reply":null,"from":"Zahrulsch","from_user_name":"Zahrulsch","source":"inbox","local_id":"-981438"}}`

		event := chat.EmitEventSocket{}
		err := json.Unmarshal([]byte(data), &event)
		assert.Nil(t, err)
		assert.IsType(t, &chat.SendChat{}, event.Data)
	})

}

func TestRcvModel(t *testing.T) {
	t.Run("test on chat", func(t *testing.T) {
		data := `{"code":103,"data":{"msg_id":2483170385,"from":"sch_mart","from_uid":80205851,"from_user_name":"Zahrulsch","from_role":"Shop Owner","thumbnail":"https://images.tokopedia.net/img/cache/215-square/shops-1/2020/5/5/7125740/7125740_aa0db54e-b543-4942-9584-d7a4a96e7d55.png","reminder_ticker":{"enable":false,"enable_close":false,"feature_id":0,"local_id":"","main_text":"","regex_message":"","sub_text":"","ticker_type":"","url":"","url_label":""},"is_opposite":false,"to_uid":80205851,"message":{"censored_reply":"test data","original_reply":"test data","timestamp":"2023-07-26T13:31:10.926526764+07:00","timestamp_fmt":"26 July 2023, 13:31 WIB","timestamp_unix":1690353070926,"timestamp_unix_nano":1690353070926527000},"start_time":"2023-07-26T06:31:11.084Z","show_rating":false,"local_id":"-291357","client_connect_time":"0001-01-01T00:00:00Z","source":"inbox"}}`

		event := chat.RcvEventSocket{}
		err := json.Unmarshal([]byte(data), &event)
		assert.Nil(t, err)
		assert.IsType(t, &chat.RcvChat{}, event.Data)

	})
}
