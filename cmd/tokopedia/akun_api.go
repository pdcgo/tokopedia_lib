package main

import (
	"net/http"

	"github.com/pdcgo/v2_gots_sdk"
)

type Response struct {
	Msg string `json:"msg"`
	Err string `json:"error"`
}

type AkunListQuery struct {
	Page    int    `json:"page"`
	PerPage int    `json:"per_page"`
	Search  string `json:"Search"`
}

type Pagination struct {
	Page    int `json:"page"`
	PerPage int `json:"per_page"`
	Count   int `json:"count"`
}

type AkunItem struct {
	AkunUploadStatus
	Username   string `json:"username"`
	Password   string `json:"password"`
	Secret     string `json:"secret"`
	Markup     string `json:"markup"`
	Spin       string `json:"spin"`
	Collection string `json:"collection"`
}

type AkunListResponse struct {
	Response
	Data []*AkunItem `json:"data"`

	Pagination Pagination `json:"pagination"`
}

type BulkItem struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Secret   string `json:"secret"`
}

type BulkPayload struct {
	Data []*BulkItem `json:"data"`
}

type AkunUpdatePayload struct {
	Data []*AkunItem `json:"data"`
}

type AkunDeletePayload struct {
	Usernames []string `json:"usernames"`
}

type AkunResetPayload struct {
	Usernames []string `json:"usernames"`
}

func RegisterAkunApi(g *v2_gots_sdk.SdkGroup) {
	akun := g.Group("akun")

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodGet,
		RelativePath: "list",
		Query:        AkunListQuery{},
		Response:     AkunListResponse{},
	})

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "bulk_add",
		Payload:      BulkPayload{},
		Response:     Response{},
	})

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "update",
		Payload:      AkunUpdatePayload{},
		Response:     Response{},
	})

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "delete",
		Payload:      AkunDeletePayload{},
		Response:     Response{},
	})

	akun.Register(&v2_gots_sdk.Api{
		Method:       http.MethodPost,
		RelativePath: "reset",
		Response:     Response{},
		Payload:      AkunResetPayload{},
	})

}
