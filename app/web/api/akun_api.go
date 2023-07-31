package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/lib/repo"
	"github.com/pdcgo/v2_gots_sdk"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Pagination struct {
	Offset int   `json:"offset"`
	Limit  int   `json:"limit"`
	Count  int64 `json:"count"`
}

type BulkItem struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

type AkunUpdatePayload struct {
	Data []*repo.AkunItem `json:"data"`
}

type AkunDeletePayload struct {
	Usernames []string `json:"usernames"`
}

type AkunResetPayload struct {
	Usernames []string `json:"usernames"`
}

type AkunApi struct {
	db   *gorm.DB
	repo *repo.AkunRepo
}

func NewAkunApi(db *gorm.DB, repo *repo.AkunRepo) *AkunApi {
	return &AkunApi{
		db:   db,
		repo: repo,
	}
}

type BulkPayload struct {
	Data []*BulkItem `json:"data"`
}

func (akapi *AkunApi) BulkAdd(ctx *gin.Context) {
	payload := BulkPayload{}
	hasil := Response{}

	err := ctx.BindJSON(&payload)
	if err != nil {
		hasil.Msg = "error"
		hasil.Err = err.Error()
		ctx.JSON(http.StatusInternalServerError, &hasil)
		return
	}
	akuns := make([]*repo.AkunItem, len(payload.Data))

	for ind, bitem := range payload.Data {
		akuns[ind] = &repo.AkunItem{
			Username: bitem.Username,
			Password: bitem.Password,
			Secret:   bitem.Secret,
		}
	}

	err = akapi.db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(akuns, 50).Error
	if err != nil {
		hasil.Msg = "error"
		hasil.Err = err.Error()
		ctx.JSON(http.StatusInternalServerError, &hasil)
		return
	}

	hasil.Msg = "success"
	ctx.JSON(http.StatusOK, &hasil)
}

type AkunListQuery struct {
	Offset int    `json:"offset" form:"offset"`
	Limit  int    `json:"limit" form:"limit"`
	Search string `json:"search" form:"search"`
}

type AkunListResponse struct {
	Response
	Data []*repo.AkunItem `json:"data"`

	Pagination Pagination `json:"pagination"`
}

func (api *AkunApi) List(ctx *gin.Context) {
	hasil := AkunListResponse{
		Data: []*repo.AkunItem{},
	}
	query := AkunListQuery{
		Offset: 0,
		Limit:  50,
	}
	ctx.BindQuery(&query)

	hasil.Pagination = Pagination{
		Offset: query.Offset,
		Limit:  query.Limit,
		Count:  0,
	}

	gentx := func() *gorm.DB {
		tx := api.db.Model(&repo.AkunItem{})

		if query.Search != "" {
			tx = tx.Where("username LIKE ?", "%"+query.Search+"%")
		}

		return tx
	}

	err := gentx().Limit(query.Limit).Offset(query.Offset).Find(&hasil.Data).Error
	if err != nil {
		hasil.Response.Err = err.Error()
		ctx.JSON(http.StatusInternalServerError, &hasil)
		return
	}

	err = gentx().Count(&hasil.Pagination.Count).Error
	if err != nil {
		hasil.Response.Err = err.Error()
		ctx.JSON(http.StatusInternalServerError, &hasil)
		return
	}

	ctx.JSON(http.StatusOK, &hasil)
}

func (akapi *AkunApi) Update(ctx *gin.Context) {
	hasil := Response{
		Msg: "success",
	}
	payload := AkunUpdatePayload{}
	err := ctx.BindJSON(&payload)
	if err != nil {
		hasil.Msg = "error"
		hasil.Err = err.Error()
		ctx.JSON(http.StatusInternalServerError, &hasil)
		return
	}
	for _, data := range payload.Data {
		err = akapi.db.Save(data).Error
		if err != nil {
			hasil.Msg = "error"
			hasil.Err = err.Error()
			ctx.JSON(http.StatusInternalServerError, &hasil)
			return
		}
	}

	ctx.JSON(http.StatusOK, &hasil)

}

func (akapi *AkunApi) Delete(ctx *gin.Context) {
	hasil := Response{
		Msg: "success",
	}
	payload := AkunDeletePayload{}
	err := ctx.BindJSON(&payload)
	if err != nil {
		hasil.Msg = "error"
		hasil.Err = err.Error()
		ctx.JSON(http.StatusInternalServerError, &hasil)
		return
	}

	for _, data := range payload.Usernames {
		err = akapi.db.Where("username = ?", data).Delete(&repo.AkunItem{}).Error
		if err != nil {
			hasil.Msg = "error"
			hasil.Err = err.Error()
			ctx.JSON(http.StatusInternalServerError, &hasil)
			return
		}
	}

	ctx.JSON(http.StatusOK, &hasil)
}

func (akapi *AkunApi) ResetAll(ctx *gin.Context) {
	hasil := Response{
		Msg: "success",
	}
	err := akapi.db.Session(&gorm.Session{AllowGlobalUpdate: true}).Model(&repo.AkunItem{}).Update("count_upload", 0).Error

	if err != nil {
		hasil.Msg = "error"
		hasil.Err = err.Error()
		ctx.JSON(http.StatusInternalServerError, &hasil)
		return
	}

	ctx.JSON(http.StatusOK, &hasil)
}

func RegisterAkunApi(g *v2_gots_sdk.SdkGroup, db *gorm.DB, repo *repo.AkunRepo) {

	akapi := NewAkunApi(db, repo)

	akun := g.Group("akun")

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "list",
		Query:        AkunListQuery{},
		Response:     AkunListResponse{},
	}, akapi.List)

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "bulk_add",
		Payload:      BulkPayload{},
		Response:     Response{},
	}, akapi.BulkAdd)

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "update",
		Payload:      AkunUpdatePayload{},
		Response:     Response{},
	}, akapi.Update)

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "delete",
		Payload:      AkunDeletePayload{},
		Response:     Response{},
	}, akapi.Delete)

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPut,
		RelativePath: "reset_all_count",
		Response:     Response{},
	}, akapi.ResetAll)

	// akun.Register(&v2_gots_sdk.Api{
	// 	Method:       http.MethodPost,
	// 	RelativePath: "reset",
	// 	Response:     Response{},
	// 	Payload:      AkunResetPayload{},
	// }, Reset)

}
