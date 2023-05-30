package api

import "strconv"

type NotificationCounterRes struct {
	Data struct {
		Notifications struct {
			Resolution   int `json:"resolution"`
			ResolutionAs struct {
				Seller   int    `json:"seller"`
				Typename string `json:"__typename"`
			} `json:"resolutionAs"`
			NotifcenterTrxUnread struct {
				NotifUnreadSellerInt int    `json:"notif_unread_seller_int"`
				Typename             string `json:"__typename"`
			} `json:"notifcenter_trx_unread"`
			NotifcenterUnread struct {
				NotifUnreadSellerInt int    `json:"notif_unread_seller_int"`
				Typename             string `json:"__typename"`
			} `json:"notifcenter_unread"`
			SellerOrderStatus struct {
				NewOrder            int    `json:"newOrder"`
				ReadyToShip         int    `json:"readyToShip"`
				Shipped             int    `json:"shipped"`
				ArriveAtDestination int    `json:"arriveAtDestination"`
				Typename            string `json:"__typename"`
			} `json:"sellerOrderStatus"`
			Inbox struct {
				TalkSeller int    `json:"talk_seller"`
				Ticket     int    `json:"ticket"`
				Review     int    `json:"review"`
				Reputation int    `json:"reputation"`
				Typename   string `json:"__typename"`
			} `json:"inbox"`
			Chat struct {
				UnreadsSeller int    `json:"unreadsSeller"`
				Typename      string `json:"__typename"`
			} `json:"chat"`
			Typename string `json:"__typename"`
		} `json:"notifications"`
	} `json:"data"`
}

type NotificationCounterVar struct {
	ShopID string `json:"shopId"`
}

func (api *TokopediaApi) NotificationCounter() (*NotificationCounterRes, error) {
	shopid := strconv.Itoa(int(api.AuthenticatedData.UserShopInfo.Info.ShopID))
	query := GraphqlPayload{
		OperationName: "NotificationCounterQuery",
		Variables: NotificationCounterVar{
			ShopID: shopid,
		},
		Query: `
		query NotificationCounterQuery($shopId: String) {
			notifications(input: {shop_id: $shopId}) {
			  resolution
			  resolutionAs {
				seller
				__typename
			  }
			  notifcenter_trx_unread {
				notif_unread_seller_int
				__typename
			  }
			  notifcenter_unread {
				notif_unread_seller_int
				__typename
			  }
			  sellerOrderStatus {
				newOrder
				readyToShip
				shipped
				arriveAtDestination
				__typename
			  }
			  inbox {
				talk_seller
				ticket
				review
				reputation
				__typename
			  }
			  chat {
				unreadsSeller
				__typename
			  }
			  __typename
			}
		  }`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil *NotificationCounterRes
	err := api.SendRequest(req, &hasil)

	return hasil, err
}
