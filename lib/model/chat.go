package model

import "encoding/json"

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

type ChatRepliesBlock struct {
	IsPromoBlocked bool   `json:"isPromoBlocked"`
	IsBlocked      bool   `json:"isBlocked"`
	BlockedUntil   string `json:"blockedUntil"`
	Typename       string `json:"__typename"`
}

type ChatRepliesContactStatus struct {
	Timestamp int    `json:"timestamp"`
	IsOnline  bool   `json:"isOnline"`
	Typename  string `json:"__typename"`
}

type ChatRepliesContact struct {
	UserID       int                      `json:"userId"`
	ShopID       int                      `json:"shopId"`
	Name         string                   `json:"name"`
	Role         string                   `json:"role"`
	Interlocutor bool                     `json:"interlocutor"`
	Badge        string                   `json:"badge"`
	IsGold       bool                     `json:"isGold"`
	Domain       string                   `json:"domain"`
	Thumbnail    string                   `json:"thumbnail"`
	Tag          string                   `json:"tag"`
	Status       ChatRepliesContactStatus `json:"status"`
	Typename     string                   `json:"__typename"`
}

type ChatMessageAttachment struct {
	ID         int      `json:"id"`
	Type       int      `json:"type"`
	Fallback   Fallback `json:"fallback"`
	Attributes string   `json:"attributes"`
	Typename   string   `json:"__typename"`
}

type ChatMessage struct {
	AttachmentID    int                   `json:"attachmentId"`
	Attachment      ChatMessageAttachment `json:"attachment"`
	ParentReply     any                   `json:"parentReply"`
	BlastID         int                   `json:"blastId"`
	Source          string                `json:"source"`
	IsOpposite      bool                  `json:"isOpposite"`
	IsRead          bool                  `json:"isRead"`
	Message         string                `json:"message"`
	MessageID       int64                 `json:"messageId"`
	OldMessageID    int                   `json:"oldMessageId"`
	OldMessageTitle string                `json:"oldMessageTitle"`
	ReplyID         string                `json:"replyId"`
	ReplyTime       string                `json:"replyTime"`
	Role            string                `json:"role"`
	SenderID        int                   `json:"senderId"`
	SenderName      string                `json:"senderName"`
	Status          int                   `json:"status"`
	FraudStatus     int                   `json:"fraudStatus"`
	AllowDelete     bool                  `json:"allowDelete"`
	Typename        string                `json:"__typename"`
}

type ChatRepliesItemChat struct {
	Time     string        `json:"time"`
	Messages []ChatMessage `json:"messages"`
	Typename string        `json:"__typename"`
}

type ChatRepliesItem struct {
	Date     string                `json:"date"`
	Chats    []ChatRepliesItemChat `json:"chats"`
	Typename string                `json:"__typename"`
}

type ChatReplies struct {
	Block           ChatRepliesBlock     `json:"block"`
	Contacts        []ChatRepliesContact `json:"contacts"`
	TextareaReply   int                  `json:"textareaReply"`
	List            []ChatRepliesItem    `json:"list"`
	HasNext         bool                 `json:"hasNext"`
	HasNextAfter    bool                 `json:"hasNextAfter"`
	ShowTimeMachine int                  `json:"showTimeMachine"`
	MinReplyTime    string               `json:"minReplyTime"`
	MaxReplyTime    string               `json:"maxReplyTime"`
	AttachmentIDs   string               `json:"attachmentIDs"`
	Typename        string               `json:"__typename"`
}

type ChatRepliesData struct {
	ChatReplies ChatReplies `json:"chatReplies"`
}

type ChatAttachmentProductVariant struct {
	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Option     struct {
		ID    int    `json:"id"`
		Hex   string `json:"hex"`
		Value string `json:"value"`
	} `json:"option"`
	Position int    `json:"position"`
	UnitName string `json:"unit_name"`
}

type FreeOngkir struct {
	IsActive bool   `json:"is_active"`
	ImageURL string `json:"image_url"`
}

type LocationStock struct {
	Addressid            int64  `json:"address_id"`
	DistrictId           int64  `json:"district_id"`
	DistrictName         string `json:"district_name"`
	DistrictIcon         string `json:"district_icon"`
	DistrictNameFullText string `json:"district_name_full_text"`
	PostalCode           string `json:"postal_code"`
	Latlon               string `json:"latlon"`
}

type ChatAttachmentProduct struct {
	Name                      string                         `json:"name"`
	ParentID                  int64                          `json:"parent_id"`
	Price                     string                         `json:"price"`
	PriceInt                  int                            `json:"price_int"`
	ImageURL                  string                         `json:"image_url"`
	URL                       string                         `json:"url"`
	AndroidURL                string                         `json:"android_url"`
	IosURL                    string                         `json:"ios_url"`
	IsUpcomingCampaignProduct bool                           `json:"is_upcoming_campaign_product"`
	PlaystoreProductData      struct{}                       `json:"playstore_product_data"`
	PriceBefore               string                         `json:"price_before"`
	DropPercentage            string                         `json:"drop_percentage"`
	ShopID                    int                            `json:"shop_id"`
	Status                    int                            `json:"status"`
	MinOrder                  int                            `json:"min_order"`
	CategoryID                int                            `json:"category_id"`
	RemainingStock            int                            `json:"remaining_stock"`
	CategoryBreadcrumb        string                         `json:"category_breadcrumb"`
	ListImageURL              []string                       `json:"list_image_url"`
	Variant                   []ChatAttachmentProductVariant `json:"variant"`
	Wishlist                  bool                           `json:"wishlist"`
	FreeOngkir                FreeOngkir                     `json:"free_ongkir"`
	Rating                    struct{}                       `json:"rating"`
	IsFulfillment             bool                           `json:"is_fulfillment"`
	IconTokocabang            string                         `json:"icon_tokocabang"`
	DescTokocabang            string                         `json:"desc_tokocabang"`
	StockInfo                 struct{}                       `json:"stock_info"`
	IsVariant                 bool                           `json:"is_variant"`
	LocationStock             LocationStock                  `json:"location_stock"`
}

type ChatAttachmentSticker struct {
	GroupId   string `json:"group_id"`
	StickerId string `json:"sticker_id"`
	ImageUrl  string `json:"image_url"`
	Intention string `json:"intention"`
}

type ChatAttachmentAttributes struct {
	ProductID         int64                  `json:"product_id,omitempty"`
	ProductProfile    *ChatAttachmentProduct `json:"product_profile,omitempty"`
	Source            string                 `json:"source,omitempty"`
	StickerProfile    *ChatAttachmentSticker `json:"sticker_profile,omitempty"`
	ImageUrl          string                 `json:"image_url,omitempty"`
	ImageUrlThumbnail string                 `json:"image_url_thumbnail,omitempty"`
}

type ChatAttachmentItem struct {
	ID         int64                    `json:"id"`
	Type       int                      `json:"type"`
	Attributes ChatAttachmentAttributes `json:"attributes"`
	Fallback   Fallback                 `json:"fallback"`
	IsActual   bool                     `json:"isActual"`
	Typename   string                   `json:"__typename"`
}

func (c *ChatAttachmentItem) UnmarshalJSON(data []byte) error {
	type Alias ChatAttachmentItem
	aux := &struct {
		*Alias
		Attributes string `json:"attributes"`
	}{
		Alias: (*Alias)(c),
	}
	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	return json.Unmarshal([]byte(aux.Attributes), &c.Attributes)
}

type ChatAttachments struct {
	ChatAttachments struct {
		List     []ChatAttachmentItem `json:"list"`
		Typename string               `json:"__typename"`
	} `json:"chatAttachments"`
}

type ChatPin struct {
	ChatPin struct {
		Success  bool   `json:"success"`
		Typename string `json:"__typename"`
	} `json:"chatPin"`
}

type ChatUnpin struct {
	ChatUnpin struct {
		Success  bool   `json:"success"`
		Typename string `json:"__typename"`
	} `json:"chatUnpin"`
}

type ChatSearchContact struct {
	ID         int    `json:"id"`
	Role       string `json:"role"`
	Attributes struct {
		Domain     string `json:"domain"`
		Name       string `json:"name"`
		ShopStatus int    `json:"shopStatus"`
		Tag        string `json:"tag"`
		Thumbnail  string `json:"thumbnail"`
		Typename   string `json:"__typename"`
	} `json:"attributes"`
	Typename string `json:"__typename"`
}

type ChatSearchItem struct {
	Contact       ChatSearchContact `json:"contact"`
	CreateBy      int               `json:"createBy"`
	CreateTimeStr string            `json:"createTimeStr"`
	LastMessage   string            `json:"lastMessage"`
	MsgID         int64             `json:"msgId"`
	OppositeID    int               `json:"oppositeId"`
	OppositeType  int               `json:"oppositeType"`
	ReplyID       int               `json:"replyId"`
	RoomID        string            `json:"roomId"`
	Typename      string            `json:"__typename"`
}

type ChatSearch struct {
	ChatSearch struct {
		Contact struct {
			Data     []ChatSearchItem `json:"data"`
			Typename string           `json:"__typename"`
		} `json:"contact"`
		Typename string `json:"__typename"`
	} `json:"chatSearch"`
}

type ChatExistingAttributes struct {
	Name      string `json:"name"`
	Domain    string `json:"domain"`
	Thumbnail string `json:"thumbnail"`
	Typename  string `json:"__typename"`
}

type ChatExistingContact struct {
	ID         int                     `json:"id"`
	Role       string                  `json:"role"`
	Attributes *ChatExistingAttributes `json:"attributes"`
	Typename   string                  `json:"__typename"`
}

type ChatExistingChat struct {
	MessageID int64                `json:"messageId"`
	Contact   *ChatExistingContact `json:"contact"`
	Typename  string               `json:"__typename"`
}

type ChatExistingData struct {
	ChatExistingChat *ChatExistingChat `json:"chatExistingChat"`
}
