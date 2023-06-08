package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	mongolib "github.com/pdcgo/go_v2_shopeelib/lib/mongo"
	"github.com/pdcgo/v2_gots_sdk"
	"gorm.io/gorm"
)

type ShopeeTopedMapApi struct {
	db       *gorm.DB
	prodRepo *mongolib.ProductRepo
}

type MapItem struct {
	ShopeeID    int64 `gorm:"primaryKey;autoIncrement:false"  json:"shopee_id"`
	TokopediaID int   `gorm:"primaryKey;autoIncrement:false" json:"tokopedia_id"`
}

func (mapi *ShopeeTopedMapApi) UpdateMap(c *gin.Context) {
	var payload []*MapItem
	c.BindJSON(&payload)

	hasil := Response{
		Msg: "success",
	}

	err := mapi.db.CreateInBatches(payload, 100).Error
	if err != nil {
		hasil.Msg = err.Error()
		hasil.Err = "error"
		c.JSON(http.StatusInternalServerError, &hasil)
		return
	}

	c.JSON(http.StatusOK, &hasil)

}

type GetMapQuery struct {
	Collection string `json:"collection" form:"collection"`
}

func (mapi *ShopeeTopedMapApi) GetMap(c *gin.Context) {
	var query GetMapQuery
	c.BindQuery(&query)

	aggre, err := mapi.prodRepo.CategoryAgg("shopee", query.Collection, false)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Msg: err.Error(),
			Err: "error",
		})
		return
	}

	ids := make([]int64, len(aggre))

	for index, agg := range aggre {
		ids[index] = agg.ID
	}

	hasil := []*MapItem{}

	err = mapi.db.Model(&MapItem{}).Where("shopee_id IN ?", ids).Find(&hasil).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Msg: err.Error(),
			Err: "error",
		})
		return
	}

	c.JSON(http.StatusOK, &ShopeeTopedMapResponse{
		Data: hasil,
	})
}

func (mapi *ShopeeTopedMapApi) AutoSuggest(c *gin.Context) {
	panic("not Implemented")
}

type ShopeeTopedMapResponse struct {
	Data []*MapItem `json:"data"`
}

func RegisterShopeeTopedMap(grp *v2_gots_sdk.SdkGroup, db *gorm.DB, prodrepo *mongolib.ProductRepo) *ShopeeTopedMapApi {
	// untuk migrasi data
	db.AutoMigrate(MapItem{})

	mapapi := ShopeeTopedMapApi{
		db:       db,
		prodRepo: prodrepo,
	}

	grp = grp.Group("mapper")

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: "map",
		Payload:      []MapItem{},
		Response:     Response{},
	}, mapapi.UpdateMap)

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "map",
		Query:        GetMapQuery{},
		Response:     ShopeeTopedMapResponse{},
	}, mapapi.GetMap)

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: "autosuggest",
	}, mapapi.AutoSuggest)

	return &mapapi
}
