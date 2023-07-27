package chat

import (
	"fmt"
	"math"
	"math/rand"
	"strings"
	"time"
)

type ParentReply struct {
	AttachmentId    int64  `json:"attachment_id"`
	AttachmentType  int64  `json:"attachment_type"`
	SenderId        int64  `json:"sender_id"`
	Name            string `json:"string"`
	ReplyTime       int64  `json:"reply_time"`
	ReplyTimeString string `json:"reply_time_string"`
	MainText        string `json:"main_text"`
	Source          string `json:"inbox"`
}

type FreeOngkir struct {
	IsActive bool   `json:"is_active"`
	ImageUrl string `json:"image_url"`
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

type ProductProfile struct {
	ImageUrl                  string        `json:"image_url"`
	Name                      string        `json:"name"`
	ParentId                  int64         `json:"parent_id"`
	Price                     string        `json:"price"`
	PriceInt                  int64         `json:"price_int"`
	Url                       string        `json:"url"`
	AndroidUrl                string        `json:"android_url"`
	IosUrl                    string        `json:"ios_url"`
	IsUpcomingCampaignProduct bool          `json:"is_upcoming_campaign_product"`
	PlayStoreProductData      interface{}   `json:"playstore_product_data"`
	ShopId                    int64         `json:"shop_id"`
	Status                    int64         `json:"status"`
	MinOrder                  int64         `json:"mmin_order"`
	Categoryid                int64         `json:"category_id"`
	RemainingStock            int64         `json:"remaining_stock"`
	CategoryBreadCrumb        string        `json:"category_breadcrumb"`
	ListImageUrl              []string      `json:"list_image_url"`
	Variant                   []interface{} `json:"variant"`
	Wishlist                  bool          `json:"wishlist"`
	FreeOngkir                FreeOngkir    `json:"free_ongkir"`
	Rating                    interface{}   `json:"rating"`
	IsFulfillment             bool          `json:"is_fulfillment"`
	IconTokoCabang            string        `json:"icon_tokocabang"`
	DescTokoCabang            string        `json:"desc_tokocabang"`
	StockInfo                 interface{}   `json:"stock_info"`
	LocationStock             LocationStock `json:"location_stock"`
	Type                      string        `json:"type"`
	DefaultChild              int64         `json:"defaultChild"`
	Id                        int64         `json:"id"`
	DropPercentage            string        `json:"dropPercentage"`
	PriceBefore               string        `json:"priceBefore"`
	Text                      string        `json:"text"`
}

type Payload struct {
	GroupId   string `json:"group_id"`
	StickerId string `json:"sticker_id"`
	ImageUrl  string `json:"image_url"`
	Intention string `json:"intention"`
}

type ExtrasProduct struct {
	Url       string `json:"url"`
	ProductId string `json:"product_id"`
}

type Extras struct {
	ExtrasProduct []ExtrasProduct `json:"extras_product"`
}

type FallbackAttachment struct {
	Html    string `json:"html"`
	Message string `json:"message"`
}

type Attributes struct {
	ProductId      int64          `json:"product_id,omitempty"`
	ProductProfile ProductProfile `json:"product_profile,omitempty"`
	Source         string         `json:"source,omitempty"`
	StickerProfile *Payload       `json:"sticker_profile,omitempty"`
}

type Attachment struct {
	Attributes         Attributes         `json:"attributes"`
	FallbackAttachment FallbackAttachment `json:"fallback_attachment"`
	Id                 int64              `json:"id"`
	Type               int64              `json:"type"`
}

type MessageResp struct {
	CensoredReply     string `json:"censored_reply,omitempty"`
	OriginalReply     string `json:"original_reply,omitempty"`
	Timestamp         string `json:"timestamp,omitempty"`
	TimestampFmt      string `json:"timestamp_fmt,omitempty"`
	TimestampUnix     uint   `json:"timestamp_unix,omitempty"`
	TimestampUnixNano uint   `json:"timestamp_unix_nano,omitempty"`
}

type DataResp struct {
	MsgId             uint           `json:"msg_id"`
	From              string         `json:"from,omitempty"`
	FromUid           string         `json:"from_uid,omitempty"`
	FromUserName      string         `json:"from_user_name,omitempty"`
	FromRole          string         `json:"from_role,omitempty"`
	ReminderTicker    any            `json:"reminder_ticker,omitempty"`
	IsOpposite        bool           `json:"is_opposite,omitempty"`
	ToUid             uint           `json:"to_uid,omitempty"`
	ClientConnectTime string         `json:"client_connect_time,omitempty"`
	Thumbnail         string         `json:"thumbnail,omitempty"`
	Message           *MessageResp   `json:"message,omitempty"`
	StartTime         string         `json:"start_time,omitempty"`
	ShowRating        bool           `json:"show_rating,omitempty"`
	LocalId           string         `json:"local_id,omitempty"`
	Source            string         `json:"source,omitempty"`
	ParentReply       *ParentReply   `json:"parent_reply,omitempty"`
	AttachmentId      int64          `json:"attachment_id,omitempty"`
	Attachment        *Attachment    `json:"attachment,omitempty"`
	Extras            *ExtrasProduct `json:"extras_product,omitempty"`
}

type DataSend struct {
	MessageId      uint            `json:"message_id"`
	Message        string          `json:"message,omitempty"`
	StartTime      string          `json:"start_time,omitempty"`
	ParentReply    *ParentReply    `json:"parent_reply,omitempty"`
	From           string          `json:"from,omitempty"`
	FromUserName   string          `json:"from_user_name,omitempty"`
	LocalId        string          `json:"local_id,omitempty"`
	Source         string          `json:"source,omitempty"`
	AttachmentType int64           `json:"attachment_type,omitempty"`
	ProductProfile *ProductProfile `json:"product_profile,omitempty"`
	Payload        *Payload        `json:"payload,omitempty"`
	Extras         *Extras         `json:"extras,omitempty"`
	FilePath       string          `json:"file_path,omitempty"`
}
type MessageId struct {
	MsgId uint `json:"msg_id"`
}

type BaseSocketType struct {
	Code uint        `json:"code"`
	Data interface{} `json:"data"`
}

func StartTime() string {
	dateTime := time.Now().Format("2006-01-02 15:04:05")

	startTime := strings.Replace(dateTime, " ", "T", 1) + "Z"
	return startTime
}

func ReplyTime() int64 {
	now := time.Now().UnixNano()
	return now
}

func ReplyTimeStr() string {
	now := time.Now().UnixNano()
	return fmt.Sprintf("%d", now)
}

func LocalId() string {
	localId := fmt.Sprintf("-%f", math.Round(1e6*rand.Float64()))
	return localId
}
