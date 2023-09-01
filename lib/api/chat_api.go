package api

type ChatListItemAttributesContact struct {
	Name      string `json:"name"`
	Role      string `json:"role"`
	Thumbnail string `json:"thumbnail"`
	Typename  string `json:"__typename"`
}

type ChatListItemAttributes struct {
	Contact          ChatListItemAttributesContact `json:"contact"`
	LastReplyMessage string                        `json:"lastReplyMessage"`
	LastReplyTime    string                        `json:"lastReplyTime"`
	ReadStatus       int                           `json:"readStatus"`
	Unreads          int                           `json:"unreads"`
	PinStatus        int                           `json:"pinStatus"`
	IsReplyByTopbot  bool                          `json:"isReplyByTopbot"`
	Typename         string                        `json:"__typename"`
}

type ChatListItem struct {
	Key        string                 `json:"key"`
	MessageID  int64                  `json:"messageId"`
	Attributes ChatListItemAttributes `json:"attributes"`
	Typename   string                 `json:"__typename"`
}

type ChatListData struct {
	ChatList struct {
		List            []ChatListItem `json:"list"`
		HasNext         bool           `json:"hasNext"`
		ShowTimeMachine int            `json:"showTimeMachine"`
		Typename        string         `json:"__typename"`
	} `json:"chatList"`
}

type ChatListRes struct {
	Data ChatListData `json:"data"`
}

type ChatListVar struct {
	Page     int    `json:"page"`
	PerPage  int    `json:"perPage"`
	Order    string `json:"order"`
	Platform string `json:"platform"`
	Filter   string `json:"filter"`
	Tab      string `json:"tab"`
}

func (api *TokopediaApi) GetChatList(variables ChatListVar) (*ChatListRes, error) {

	query := GraphqlPayload{
		OperationName: "ChatListQuery",
		Variables:     variables,
		Query: `
		query ChatListQuery($tab: String, $filter: String, $page: Int = 1, $perPage: Int = 10, $order: String = "desc", $platform: String = "mobile") {
			chatList: chatListMessage(tab: $tab, filter: $filter, page: $page, perPage: $perPage, order: $order, platform: $platform) {
			  list {
				key: messageKey
				messageId: msgID
				attributes {
				  contact {
					name
					role
					thumbnail
					__typename
				  }
				  lastReplyMessage
				  lastReplyTime: lastReplyTimeStr
				  readStatus
				  unreads
				  pinStatus
				  isReplyByTopbot
				  __typename
				}
				__typename
			  }
			  hasNext
			  showTimeMachine
			  __typename
			}
		  }`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil ChatListRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
