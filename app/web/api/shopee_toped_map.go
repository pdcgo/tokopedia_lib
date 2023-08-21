package api

import (
	"errors"
	"log"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/go_v2_shopeelib/lib/mongorepo"
	"github.com/pdcgo/tokopedia_lib/app/config"
	"github.com/pdcgo/tokopedia_lib/lib/category_mapper"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/v2_gots_sdk"
	"gorm.io/gorm"
)

type SuggestStatusType string

const (
	SUGGEST_RUN  SuggestStatusType = "RUNNING"
	SUGGEST_STOP SuggestStatusType = "STOPPED"
)

type AutoSuggestStatus struct {
	sync.Mutex
	Status SuggestStatusType `json:"status"`
}

type ShopeeTopedMapApi struct {
	db            *gorm.DB
	prodRepo      *mongorepo.ProductRepo
	mapper        *category_mapper.Mapper
	smapper       *config.ShopeeMapper
	configRepo    *config.ConfigRepo
	SuggestStatus *AutoSuggestStatus
}

func (mapi *ShopeeTopedMapApi) UpdateMap(c *gin.Context) {
	var payload []*config.ShopeeMapItem
	c.BindJSON(&payload)

	hasil := Response{
		Msg: "success",
	}

	for _, mapitem := range payload {
		err := mapi.db.Save(mapitem).Error
		if err != nil {
			hasil.Msg = err.Error()
			hasil.Err = "error"
			c.JSON(http.StatusInternalServerError, &hasil)
			return
		}
	}

	c.JSON(http.StatusOK, &hasil)

}

type GetMapQuery struct {
	Collection string `json:"collection" form:"collection"`
}

func (mapi *ShopeeTopedMapApi) GetCollectionCategories(c *gin.Context) ([]mongorepo.ProductCategoryAgg, error) {
	var query GetMapQuery
	c.BindQuery(&query)

	aggQuery := mongorepo.ProductMatchStageQuery{
		Marketplace: mongorepo.MP_SHOPEE,
		Namespace:   query.Collection,
	}

	return mapi.prodRepo.CategoryAgg(aggQuery)
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
	data *mongorepo.ProductCategoryAgg
	db   *gorm.DB
}

func NewShopeeMapSuggestItem(db *gorm.DB, data *mongorepo.ProductCategoryAgg) *ShopeeMapSuggestItem {
	return &ShopeeMapSuggestItem{
		db:   db,
		data: data,
	}
}

func (item *ShopeeMapSuggestItem) SetTokopediaID(categid int) error {
	mapcateg := config.ShopeeMapItem{
		ShopeeID: item.data.ID,
	}

	item.db.First(&mapcateg, item.data.ID)

	if mapcateg.TokopediaID == 0 {
		mapcateg.TokopediaID = categid
		err := item.db.Save(mapcateg).Error
		if err != nil {
			return pdc_common.ReportError(err)
		}

		log.Println("mapping", item.data.Name)
	}

	return nil
}

func (item *ShopeeMapSuggestItem) GetName() string {
	data := item.data.Name[len(item.data.Name)-1]

	return data
}

func (mapi *ShopeeTopedMapApi) StatusAutoSuggest(c *gin.Context) {
	c.JSON(http.StatusOK, mapi.SuggestStatus)
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

	if mapi.SuggestStatus.TryLock() {
		log.Println("running auto suggest")
		mapi.SuggestStatus.Status = SUGGEST_RUN
		hasil := make([]category_mapper.ItemMap, len(aggre))
		for ind, agg := range aggre {
			item := agg
			hasil[ind] = NewShopeeMapSuggestItem(mapi.db, &item)
		}
		go func() {
			defer func() {
				mapi.SuggestStatus.Status = SUGGEST_STOP
				mapi.SuggestStatus.Unlock()
			}()
			mapi.mapper.RunMapper(hasil)
		}()

	}

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

type TokopediaMapQuery struct {
	Namespace string `json:"namespace" form:"namespace" schema:"namespace"`
}

type TokopediaMapItem struct {
	ShopeeID              int64    `json:"shopee_id"`
	TokopediaID           int      `json:"tokopedia_id"`
	Count                 int      `json:"product_count"`
	ShopeeCategoryName    []string `json:"shopee_category_name"`
	TokopediaCategoryName []string `json:"tokopedia_category_name"`
}

func (mapi *ShopeeTopedMapApi) TokopediaCollectionCategory(ctx *gin.Context) {
	query := TokopediaMapQuery{}
	ctx.BindQuery(&query)

	agg := repo.ProductAggregateIpml{
		Collection: mapi.prodRepo.Collection,
	}
	hasil := []*TokopediaMapItem{}
	err := agg.IterCategory(query.Namespace, func(tokopediaID, count int, name []string) error {
		item := TokopediaMapItem{
			TokopediaID:           tokopediaID,
			Count:                 count,
			TokopediaCategoryName: name,
		}

		mapitem, err := mapi.smapper.GetShopeeID(tokopediaID)
		if err == nil {
			item.ShopeeID = mapitem.ShopeeID

		} else {
			if !errors.Is(err, gorm.ErrRecordNotFound) {
				log.Println(err)
			}

		}

		hasil = append(hasil, &item)
		return nil
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, Response{
			Err: "error_map",
			Msg: err.Error(),
		})

		return
	}

	ctx.JSON(http.StatusOK, hasil)

}

type ShopeeTopedMapResponse struct {
	Data []*config.ShopeeMapItem `json:"data"`
}

func RegisterShopeeTopedMap(
	grp *v2_gots_sdk.SdkGroup,
	db *gorm.DB,
	prodrepo *mongorepo.ProductRepo,
	mapper *category_mapper.Mapper,
) *ShopeeTopedMapApi {
	// untuk migrasi data
	db.AutoMigrate(config.ShopeeMapItem{})

	mapapi := ShopeeTopedMapApi{
		db:         db,
		prodRepo:   prodrepo,
		mapper:     mapper,
		configRepo: config.NewConfigRepo(db),
		SuggestStatus: &AutoSuggestStatus{
			Status: SUGGEST_STOP,
		},
		smapper: config.NewShopeeMapper(db),
	}

	grp = grp.Group("mapper")

	// masih digunakan untuk toped shopee juga
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
		RelativePath: "autosuggest",
		Response:     AutoSuggestStatus{},
	}, mapapi.StatusAutoSuggest)

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "setting",
		Response:     config.ShopeeMapperConfig{},
	}, mapapi.GetConfig)

	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: "setting",
		Response:     config.ShopeeMapperConfig{},
		Payload:      &config.ShopeeMapperConfig{},
	}, mapapi.UpdateConfig)

	// TODO: kandidat yang baru
	grp.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "category",
		Response:     []TokopediaMapItem{},
		Query:        &TokopediaMapQuery{},
	}, mapapi.TokopediaCollectionCategory)

	return &mapapi
}
