package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/common_conf/pdc_common"
	mongolib "github.com/pdcgo/go_v2_shopeelib/lib/mongo"
	"github.com/pdcgo/tokopedia_lib/app/config"
	"github.com/pdcgo/tokopedia_lib/lib/category_mapper"
	"github.com/pdcgo/v2_gots_sdk"
	"gorm.io/gorm"
)

type ShopeeTopedMapApi struct {
	db         *gorm.DB
	prodRepo   *mongolib.ProductRepo
	mapper     *category_mapper.Mapper
	configRepo *config.ConfigRepo
}

func (mapi *ShopeeTopedMapApi) UpdateMap(c *gin.Context) {
	var payload []*config.ShopeeMapItem
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

func (mapi *ShopeeTopedMapApi) GetCollectionCategories(c *gin.Context) ([]mongolib.ProductCategoryAgg, error) {
	var query GetMapQuery
	c.BindQuery(&query)

	return mapi.prodRepo.CategoryAgg("shopee", query.Collection, false)
}

func (mapi *ShopeeTopedMapApi) GetMap(c *gin.Context) {
	var query GetMapQuery
	c.BindQuery(&query)

	aggre, err := mapi.GetCollectionCategories(c)
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

	hasil := []*config.ShopeeMapItem{}

	err = mapi.db.Model(&config.ShopeeMapItem{}).Where("shopee_id IN ?", ids).Find(&hasil).Error
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

type ShopeeMapSuggestItem struct {
	data *mongolib.ProductCategoryAgg
	db   *gorm.DB
}

func (item *ShopeeMapSuggestItem) SetTokopediaID(categid int) {
	mapcateg := config.ShopeeMapItem{
		ShopeeID:    item.data.ID,
		TokopediaID: categid,
	}

	err := item.db.Save(mapcateg).Error
	if err != nil {
		pdc_common.ReportError(err)
	}
}

func (item *ShopeeMapSuggestItem) GetName() string {
	data := item.data.Name[len(item.data.Name)-1]

	return data
}

func (mapi *ShopeeTopedMapApi) AutoSuggest(c *gin.Context) {
	aggre, err := mapi.GetCollectionCategories(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{
			Msg: err.Error(),
			Err: "error",
		})
		return
	}

	hasil := make([]category_mapper.ItemMap, len(aggre))
	for ind, agg := range aggre {
		hasil[ind] = &ShopeeMapSuggestItem{
			data: &agg,
			db:   mapi.db,
		}
	}

	mapi.mapper.RunMapper(hasil)
	c.JSON(http.StatusOK, &Response{})
}

func (mapi *ShopeeTopedMapApi) GetConfig(c *gin.Context) {
	data := config.ShopeeMapperConfig{}
	mapi.configRepo.GetConfig(&data)

	c.JSON(http.StatusOK, data)

}
func (mapi *ShopeeTopedMapApi) UpdateConfig(c *gin.Context) {
	data := config.ShopeeMapperConfig{}
	c.BindJSON(&data)

	mapi.configRepo.Save(&data)
	c.JSON(http.StatusOK, data)
}

type ShopeeTopedMapResponse struct {
	Data []*config.ShopeeMapItem `json:"data"`
}

func RegisterShopeeTopedMap(
	grp *v2_gots_sdk.SdkGroup,
	db *gorm.DB, prodrepo *mongolib.ProductRepo,
	mapper *category_mapper.Mapper,
) *ShopeeTopedMapApi {
	// untuk migrasi data
	db.AutoMigrate(config.ShopeeMapItem{})

	mapapi := ShopeeTopedMapApi{
		db:         db,
		prodRepo:   prodrepo,
		mapper:     mapper,
		configRepo: config.NewConfigRepo(db),
	}

	grp = grp.Group("mapper")

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: "map",
		Payload:      []config.ShopeeMapItem{},
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
		Query:        GetMapQuery{},
	}, mapapi.AutoSuggest)

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "setting",
		Response:     config.ShopeeMapperConfig{},
	}, mapapi.GetConfig)

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: "setting",
		Response:     config.ShopeeMapperConfig{},
	}, mapapi.UpdateConfig)

	return &mapapi
}
