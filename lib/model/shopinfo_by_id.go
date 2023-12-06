package model

import "strings"

type SIBIDFavoriteData struct {
	TotalFavorite int    `json:"totalFavorite"`
	Typename      string `json:"__typename"`
}

type SIBIDGoldOS struct {
	IsGold           int    `json:"isGold"`
	IsOfficial       int    `json:"isOfficial"`
	Badge            string `json:"badge"`
	ShopTier         int    `json:"shopTier"`
	ShopTierWording  string `json:"shopTierWording"`
	ShopGrade        int    `json:"shopGrade"`
	ShopGradeWording string `json:"shopGradeWording"`
	Typename         string `json:"__typename"`
}

type SIBIDShopAssets struct {
	Avatar       string `json:"avatar"`
	Cover        string `json:"cover"`
	DefaultCover []struct {
		ID       string `json:"id"`
		Path     string `json:"path"`
		Typename string `json:"__typename"`
	} `json:"defaultCover"`
	Typename string `json:"__typename"`
}

type SIBIDShopCore struct {
	Name        string `json:"name"`
	ShopID      string `json:"shopID"`
	Domain      string `json:"domain"`
	Description string `json:"description"`
	TagLine     string `json:"tagLine"`
	Typename    string `json:"__typename"`
}

type SIBIDClosedInfoDetail struct {
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	OpenDate  string `json:"openDate"`
	Status    int    `json:"status"`
	Typename  string `json:"__typename"`
}

type SIBIDClosedInfo struct {
	ClosedNote string                `json:"closedNote"`
	Until      string                `json:"until"`
	Detail     SIBIDClosedInfoDetail `json:"detail"`
	Typename   string                `json:"__typename"`
}

type SIBIDStatusInfo struct {
	ShopStatus int    `json:"shopStatus"`
	StatusName string `json:"statusName"`
	Typename   string `json:"__typename"`
}

type SIBIDOs struct {
	IsOfficial int    `json:"isOfficial"`
	Expired    string `json:"expired"`
	Typename   string `json:"__typename"`
}

type SIBIDResult struct {
	FavoriteData  SIBIDFavoriteData `json:"favoriteData"`
	GoldOS        SIBIDGoldOS       `json:"goldOS"`
	Location      string            `json:"location"`
	ShopAssets    SIBIDShopAssets   `json:"shopAssets"`
	IsAllowManage int               `json:"isAllowManage"`
	IsOwner       int               `json:"isOwner"`
	ShopCore      SIBIDShopCore     `json:"shopCore"`
	ShopHomeType  string            `json:"shopHomeType"`
	ClosedInfo    SIBIDClosedInfo   `json:"closedInfo"`
	StatusInfo    SIBIDStatusInfo   `json:"statusInfo"`
	Os            SIBIDOs           `json:"os"`
	Typename      string            `json:"__typename"`
}

type ShopInfoByID struct {
	Result   []SIBIDResult `json:"result"`
	Typename string        `json:"__typename"`
}

type SIBIDErrorExtension struct {
	DeveloperMessage string `json:"developerMessage"`
	Timestamp        string `json:"timestamp"`
}

type SIBIDError struct {
	Message    string              `json:"message"`
	Path       []string            `json:"path"`
	Extensions SIBIDErrorExtension `json:"extensions"`
}

type ShopInfoIDData struct {
	ShopInfoByID `json:"shopInfoByID"`
}

type ShopInfoByIDError []SIBIDError

func (errs ShopInfoByIDError) IsNotAuthorized() bool {
	for _, err := range errs {
		if strings.Contains(err.Message, "not authorized") {
			return true
		}
	}

	return false
}

func (errs ShopInfoByIDError) Error() string {
	for _, err := range errs {
		return err.Message
	}
	return ""
}

type ShopInfoByIDRes struct {
	Errors ShopInfoByIDError `json:"errors"`
	Data   ShopInfoIDData    `json:"data"`
}
