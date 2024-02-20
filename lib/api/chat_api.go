package api

import (
	"strconv"

	"github.com/pdcgo/tokopedia_lib/lib/model"
)

type ChatListRes struct {
	Data model.ChatListData `json:"data"`
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

type ChatRoomRes struct {
	Data model.ChatRepliesData `json:"data"`
}

type ChatRoomVar struct {
	PerPage    int    `json:"perPage"`
	MessageID  int64  `json:"messageId"`
	Keyword    string `json:"keyword"`
	IsTextOnly bool   `json:"isTextOnly"`
	Page       int    `json:"page"`
}

func (api *TokopediaApi) GetChatRoom(variables ChatRoomVar) (*ChatRoomRes, error) {

	query := GraphqlPayload{
		OperationName: "ChatRoomQuery",
		Variables:     variables,
		Query: `
		query ChatRoomQuery($messageId: Int!, $keyword: String, $page: Int, $perPage: Int = 10, $beforeReplyTime: String, $afterReplyTime: String, $isTextOnly: Boolean) {
			chatReplies(msgId: $messageId, keyword: $keyword, page: $page, perPage: $perPage, beforeReplyTime: $beforeReplyTime, afterReplyTime: $afterReplyTime, isTextOnly: $isTextOnly) {
			  block {
				isPromoBlocked
				isBlocked
				blockedUntil
				__typename
			  }
			  contacts {
				userId
				shopId
				name
				role
				interlocutor
				badge
				isGold
				domain
				thumbnail
				tag
				status {
				  timestamp
				  isOnline
				  __typename
				}
				__typename
			  }
			  textareaReply
			  list {
				date
				chats {
				  time
				  messages: replies {
					attachmentId: attachmentID
					attachment {
					  id
					  type
					  fallback {
						message
						html
						__typename
					  }
					  attributes
					  __typename
					}
					parentReply {
					  attachmentID
					  attachmentType
					  senderID
					  name
					  replyID
					  replyTimeUnixNano
					  fraudStatus
					  source
					  mainText
					  subText
					  imageURL
					  isExpired
					  __typename
					}
					blastId
					source
					isOpposite
					isRead
					message: msg
					messageId: msgId
					oldMessageId: oldMsgId
					oldMessageTitle: oldMsgTitle
					replyId
					replyTime
					role
					senderId
					senderName
					status
					fraudStatus
					allowDelete
					__typename
				  }
				  __typename
				}
				__typename
			  }
			  hasNext
			  hasNextAfter
			  showTimeMachine
			  minReplyTime
			  maxReplyTime
			  attachmentIDs
			  __typename
			}
		  }`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil ChatRoomRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

type ChatAttachmentRes struct {
	Data model.ChatAttachments `json:"data"`
}

type ChatAttachmentVar struct {
	MessageID     int64  `json:"messageId"`
	AttachmentIDs string `json:"attachmentIDs"`
}

func (api *TokopediaApi) GetChatAttachments(variables ChatAttachmentVar) (*ChatAttachmentRes, error) {

	query := GraphqlPayload{
		OperationName: "ChatAttachmentsQuery",
		Variables:     variables,
		Query: `
		query ChatAttachmentsQuery($messageId: Int!, $attachmentIDs: String, $postalCode: String, $addressId: Int, $districtId: Int, $latlon: String) {
			chatAttachments(msgId: $messageId, AttachmentIDs: $attachmentIDs, addressID: $addressId, districtID: $districtId, postalCode: $postalCode, latlon: $latlon) {
			  list {
				id
				type
				attributes
				fallback {
				  message
				  html
				  __typename
				}
				isActual
				__typename
			  }
			  __typename
			}
		  }
		`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil ChatAttachmentRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

type ChatPinRes struct {
	Data model.ChatPin `json:"data"`
}

type ChatPinVar struct {
	MessageID string `json:"msgId"`
}

func (api *TokopediaApi) ChatPin(msgId int64) (*ChatPinRes, error) {

	query := GraphqlPayload{
		OperationName: "ChatPin",
		Variables: ChatPinVar{
			MessageID: strconv.FormatInt(msgId, 10),
		},
		Query: `
		mutation ChatPin($msgId: String!) {
			chatPin(msgIDs: [$msgId]) {
			  success
			  __typename
			}
		  }
		`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil ChatPinRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

type ChatUnpinRes struct {
	Data model.ChatUnpin `json:"data"`
}

func (api *TokopediaApi) ChatUnpin(msgId int64) (*ChatUnpinRes, error) {

	query := GraphqlPayload{
		OperationName: "ChatUnpin",
		Variables: ChatPinVar{
			MessageID: strconv.FormatInt(msgId, 10),
		},
		Query: `
		mutation ChatUnpin($msgId: String!) {
			chatUnpin(msgIDs: [$msgId]) {
			  success
			  __typename
			}
		  }		  
		`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil ChatUnpinRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

type ChatSearchRes struct {
	Data model.ChatSearch `json:"data"`
}

type ChatSearchVar struct {
	Keyword  string `json:"keyword"`
	Size     int    `json:"size"`
	Status   int    `json:"status"`
	IsSeller int    `json:"isSeller"`
}

func (api *TokopediaApi) GetChatSearch(variables ChatSearchVar) (*ChatSearchRes, error) {

	query := GraphqlPayload{
		OperationName: "ChatSearchQuery",
		Variables:     variables,
		Query: `
		query ChatSearchQuery($keyword: String, $status: Int, $page: Int, $size: Int, $isSeller: Int) {
			chatSearch(keyword: $keyword, status: $status, page: $page, size: $size, isSeller: $isSeller) {
			  contact {
				data {
				  contact {
					id
					role
					attributes {
					  domain
					  name
					  shopStatus
					  tag
					  thumbnail
					  __typename
					}
					__typename
				  }
				  createBy
				  createTimeStr
				  lastMessage
				  msgId
				  oppositeId
				  oppositeType
				  replyId
				  roomId
				  __typename
				}
				__typename
			  }
			  __typename
			}
		  }					
		`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil ChatSearchRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

type ChatExistingRes struct {
	Data *model.ChatExistingData `json:"data"`
}

type ChatExistingVariables struct {
	Source   string `json:"source"`
	ToShopID int    `json:"toShopId"`
	ToUserID int    `json:"toUserId"`
}

func (api *TokopediaApi) GetChatExisting(shopId, userId int) (*ChatExistingRes, error) {

	variables := ChatExistingVariables{
		Source:   "pdp",
		ToShopID: shopId,
		ToUserID: userId,
	}

	query := GraphqlPayload{
		OperationName: "ChatExistingChatQuery",
		Variables:     variables,
		Query: `query ChatExistingChatQuery($toShopId: Int, $toUserId: Int, $source: String) {
			chatExistingChat(toShopId: $toShopId, toUserId: $toUserId, source: $source) {
			  messageId
			  contact {
				id
				role
				attributes {
				  name
				  domain
				  thumbnail
				  __typename
				}
				__typename
			  }
			  __typename
			}
		  }
		  `,
	}

	req := api.NewGraphqlReq(&query)

	var hasil ChatExistingRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
