package api

type AccountRes struct {
	Data *AccountData `json:"data"`
}
type AccountUser struct {
	ID             string `json:"id"`
	IsLoggedIn     bool   `json:"isLoggedIn"`
	Name           string `json:"name"`
	ProfilePicture string `json:"profilePicture"`
	Completion     int    `json:"completion"`
	PhoneVerified  bool   `json:"phoneVerified"`
	Typename       string `json:"__typename"`
}
type AccountInfo struct {
	ShopID     int    `json:"shopId,string"`
	ShopName   string `json:"shopName"`
	ShopDomain string `json:"shopDomain"`
	ShopAvatar string `json:"shopAvatar"`
	IsOfficial string `json:"isOfficial"`
	Typename   string `json:"__typename"`
}
type AccountOwner struct {
	IsPowerMerchant bool   `json:"isPowerMerchant"`
	PmStatus        string `json:"pmStatus"`
	Typename        string `json:"__typename"`
}
type AccountUserShopInfo struct {
	Info     *AccountInfo  `json:"info"`
	Owner    *AccountOwner `json:"owner"`
	Typename string        `json:"__typename"`
}
type AccountWallet struct {
	OvoCash   string `json:"ovoCash"`
	OvoPoints string `json:"ovoPoints"`
	Linked    bool   `json:"linked"`
	Typename  string `json:"__typename"`
}
type AccountWalletPending struct {
	PendingBalance string `json:"pendingBalance"`
	Typename       string `json:"__typename"`
}
type AccountBalance struct {
	BalanceStr string `json:"balanceStr"`
	Typename   string `json:"__typename"`
}
type AccountTier struct {
	NameDesc    string `json:"nameDesc"`
	EggImageURL string `json:"eggImageURL"`
	Typename    string `json:"__typename"`
}
type Points struct {
	Reward   int    `json:"reward"`
	Typename string `json:"__typename"`
}
type Status struct {
	Tier     AccountTier `json:"tier"`
	Points   Points      `json:"points"`
	Typename string      `json:"__typename"`
}
type AccountTokopoints struct {
	Status   *Status `json:"status"`
	Typename string  `json:"__typename"`
}
type AccountCta struct {
	Text     string `json:"text"`
	URL      string `json:"url"`
	Typename string `json:"__typename"`
}
type AccountShortcutList struct {
	ID          int         `json:"id"`
	Cta         *AccountCta `json:"cta"`
	Description string      `json:"description"`
	Typename    string      `json:"__typename"`
}
type AccountShortcutGroupList struct {
	ShortcutList []*AccountShortcutList `json:"shortcutList"`
	Typename     string                 `json:"__typename"`
}
type AccountTokopointsShortcutList struct {
	ShortcutGroupList []*AccountShortcutGroupList `json:"shortcutGroupList"`
	Typename          string                      `json:"__typename"`
}
type AccountTokopointsSumCoupon struct {
	SumCouponStr string `json:"sumCouponStr"`
	Typename     string `json:"__typename"`
}
type AccountData struct {
	User                   *AccountUser                   `json:"user"`
	UserShopInfo           *AccountUserShopInfo           `json:"userShopInfo"`
	Wallet                 *AccountWallet                 `json:"wallet"`
	WalletPending          *AccountWalletPending          `json:"walletPending"`
	Balance                *AccountBalance                `json:"balance"`
	Tokopoints             *AccountTokopoints             `json:"tokopoints"`
	TokopointsShortcutList *AccountTokopointsShortcutList `json:"tokopointsShortcutList"`
	TokopointsSumCoupon    *AccountTokopointsSumCoupon    `json:"tokopointsSumCoupon"`
}

func (api *TokopediaApi) AccountInfo() (*AccountRes, error) {
	query := GraphqlPayload{
		OperationName: "Account",
		Query: `query Account {
			user {
			  id
			  isLoggedIn
			  name
			  profilePicture
			  completion
			  phoneVerified: phone_verified
			  __typename
			}
			userShopInfo {
			  info {
				shopId: shop_id
				shopName: shop_name
				shopDomain: shop_domain
				shopAvatar: shop_avatar
				isOfficial: shop_is_official
				__typename
			  }
			  owner {
				isPowerMerchant: is_gold_merchant
				pmStatus: pm_status
				__typename
			  }
			  __typename
			}
			wallet {
			  ovoCash: cash_balance
			  ovoPoints: point_balance
			  linked
			  __typename
			}
			walletPending: goalPendingBalance {
			  pendingBalance: point_balance_text
			  __typename
			}
			balance: saldo {
			  balanceStr: deposit_fmt
			  __typename
			}
			tokopoints {
			  status {
				tier {
				  nameDesc
				  eggImageURL
				  __typename
				}
				points {
				  reward
				  __typename
				}
				__typename
			  }
			  __typename
			}
			tokopointsShortcutList(groupCodes: ["account_page_widget"]) {
			  shortcutGroupList {
				shortcutList {
				  id
				  cta {
					text
					url
					__typename
				  }
				  description
				  __typename
				}
				__typename
			  }
			  __typename
			}
			tokopointsSumCoupon {
			  sumCouponStr
			  __typename
			}
		  }`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil AccountRes
	err := api.SendRequest(req, &hasil)
	if err != nil {
		return nil, err
	}

	return &hasil, nil

}

type User struct {
	ID              string `json:"id"`
	IsLoggedIn      bool   `json:"isLoggedIn"`
	ProfilePicture  string `json:"profilePicture"`
	Name            string `json:"name"`
	Email           string `json:"email"`
	PhoneVerified   bool   `json:"phone_verified"`
	Gender          string `json:"gender"`
	Bday            string `json:"bday"`
	Completion      int    `json:"completion"`
	CreatedPassword bool   `json:"createdPassword"`
	RegisterDate    string `json:"registerDate"`
	Phone           string `json:"phone"`
	PhoneMasked     string `json:"phoneMasked"`
	Age             string `json:"age"`
	Typename        string `json:"__typename"`
}

type UserDataQueryResp struct {
	Data struct {
		User *User `json:"user"`
		Shop struct {
			Domain   string `json:"domain"`
			Typename string `json:"__typename"`
		} `json:"shop"`
	} `json:"data"`
}

func (api *TokopediaApi) UserDataQuery() (*UserDataQueryResp, error) {
	query := GraphqlPayload{
		OperationName: "user_data_query",
		Variables:     struct{}{},
		Query:         "query user_data_query {\n  user {\n    id\n    isLoggedIn\n    profilePicture\n    name\n    email\n    phone_verified\n    gender\n    bday\n    completion\n    createdPassword\n    registerDate\n    phone\n    phoneMasked\n    age\n    __typename\n  }\n  shop {\n    domain\n    __typename\n  }\n}\n",
	}

	req := api.NewGraphqlReq(&query)

	var hasil UserDataQueryResp
	err := api.SendRequest(req, &hasil)
	if err != nil {
		return nil, err
	}

	return &hasil, nil
}
