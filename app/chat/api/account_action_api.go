package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/report"
)

func (api *AccountApi) reconnect(ctx *gin.Context) {

	shopid, err := strconv.Atoi(ctx.Param("shopid"))
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.chatGroup.Reconnect(shopid)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}

func (api *AccountApi) togglePin(ctx *gin.Context) {

	shopid, err := strconv.Atoi(ctx.Param("shopid"))
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.accountService.TogglePinned(shopid)
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}

func (api *AccountApi) withdraw(ctx *gin.Context) {

	shopid, err := strconv.Atoi(ctx.Param("shopid"))
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.accountRepo.WithAccount(api.initConfig.ActiveGroup, shopid, func(account *model.Account) error {
		report := report.NewWitdrawReport(fmt.Sprintf("withdraw_%s_report.csv", account.GetUsername()))
		return api.accountService.Withdraw(account, account.AccountData.Pin, report)
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}

func (api *AccountApi) autoWithdraw(ctx *gin.Context) {

	accounts, err := api.accountRepo.List(&repo.ListAccountFilter{
		GroupName: api.initConfig.ActiveGroup,
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	report := report.NewWitdrawReport("withdraw_report.csv")
	for _, account := range accounts {
		api.accountService.Withdraw(account, account.AccountData.Pin, report)
	}
}

type Setpinpayload struct {
	Pin string `json:"pin"`
}

func (api *AccountApi) setPin(ctx *gin.Context) {

	query := BaseQuery{}
	err := ctx.BindQuery(&query)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	payload := Setpinpayload{}
	err = ctx.BindJSON(&payload)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.accountService.SetPin(query.Shopid, payload.Pin)
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}
