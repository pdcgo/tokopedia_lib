package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/upload_app"
	"github.com/pdcgo/v2_gots_sdk"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Response struct {
	Msg string `json:"msg"`
	Err string `json:"error"`
}

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
	Data []*upload_app.AkunItem `json:"data"`
}

type AkunDeletePayload struct {
	Usernames []string `json:"usernames"`
}

type AkunResetPayload struct {
	Usernames []string `json:"usernames"`
}

type AkunApi struct {
	db   *gorm.DB
	repo *upload_app.AkunRepo
}

func NewAkunApi(db *gorm.DB, repo *upload_app.AkunRepo) *AkunApi {
	return &AkunApi{
		db:   db,
		repo: repo,
	}
}

type BulkPayload struct {
	Data []*BulkItem `json:"data"`
}

func (api *AkunApi) BulkAdd(ctx *gin.Context) {
	payload := BulkPayload{}
	hasil := Response{}

	err := ctx.BindJSON(&payload)
	if err != nil {
		hasil.Msg = "error"
		hasil.Err = err.Error()
		ctx.JSON(http.StatusInternalServerError, &hasil)
		return
	}

	err = api.db.Clauses(clause.OnConflict{DoNothing: true}).CreateInBatches(payload.Data, 50).Error
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
	Data []*upload_app.AkunItem `json:"data"`

	Pagination Pagination `json:"pagination"`
}

func (api *AkunApi) List(ctx *gin.Context) {
	hasil := AkunListResponse{
		Data: []*upload_app.AkunItem{},
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

	tx := api.db.Model(&upload_app.AkunItem{})

	if query.Search != "" {
		tx = api.db.Where("username LIKE ?", "%"+query.Search+"%")
	}
	err := tx.Find(&hasil.Data).Limit(query.Limit).Offset(query.Offset).Error
	if err != nil {
		hasil.Response.Err = err.Error()
		ctx.JSON(http.StatusInternalServerError, &hasil)
		return
	}

	err = tx.Count(&hasil.Pagination.Count).Error
	if err != nil {
		hasil.Response.Err = err.Error()
		ctx.JSON(http.StatusInternalServerError, &hasil)
		return
	}

	ctx.JSON(http.StatusOK, &hasil)
}

func (api *AkunApi) Update(ctx *gin.Context) {
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
		err = api.db.Save(data).Error
		if err != nil {
			hasil.Msg = "error"
			hasil.Err = err.Error()
			ctx.JSON(http.StatusInternalServerError, &hasil)
			return
		}
	}

	ctx.JSON(http.StatusOK, &hasil)

}

func (api *AkunApi) Delete(ctx *gin.Context) {
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
		err = api.db.Delete(&upload_app.AkunItem{}, data).Error
		if err != nil {
			hasil.Msg = "error"
			hasil.Err = err.Error()
			ctx.JSON(http.StatusInternalServerError, &hasil)
			return
		}
	}

	ctx.JSON(http.StatusOK, &hasil)
}

// func (api *AkunApi) Reset(ctx *gin.Context) {
// 	hasil := Response{
// 		Msg: "success",
// 	}
// 	payload := AkunResetPayload{}
// 	err := ctx.BindJSON(&payload)
// 	if err != nil {
// 		hasil.Msg = "error"
// 		hasil.Err = err.Error()
// 		ctx.JSON(http.StatusOK, &hasil)
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, &hasil)
// }

func RegisterAkunApi(g *v2_gots_sdk.SdkGroup, db *gorm.DB, repo *upload_app.AkunRepo) {

	api := NewAkunApi(db, repo)

	akun := g.Group("akun")

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "list",
		Query:        AkunListQuery{},
		Response:     AkunListResponse{},
	}, api.List)

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "bulk_add",
		Payload:      BulkPayload{},
		Response:     Response{},
	}, api.BulkAdd)

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "update",
		Payload:      AkunUpdatePayload{},
		Response:     Response{},
	}, api.Update)

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "delete",
		Payload:      AkunDeletePayload{},
		Response:     Response{},
	}, api.Delete)

	// akun.Register(&v2_gots_sdk.Api{
	// 	Method:       http.MethodPost,
	// 	RelativePath: "reset",
	// 	Response:     Response{},
	// 	Payload:      AkunResetPayload{},
	// }, api.Reset)

}
