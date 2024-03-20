package api

import (
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/report"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
)

func (api *AccountApi) reconnect(ctx *gin.Context) {

	shopid, err := strconv.Atoi(ctx.Param("shopid"))
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	err = api.accountService.Reconnect(shopid)
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

	err = api.accountService.UpdateAccount(shopid, func(account *model.Account) error {
		account.AccountData.Pinned = !account.AccountData.Pinned
		return nil
	})
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

	err = api.accountService.WithAccount(api.initConfig.ActiveGroup, shopid, func(account *model.Account) error {
		username := account.GetUsername()
		report := report.NewWitdrawReport(fmt.Sprintf("withdraw_%s_report.csv", username))
		return api.accountService.Withdraw(username, account.AccountData.Pin, report)
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}

func (api *AccountApi) autoWithdraw(ctx *gin.Context) {

	accounts, err := api.accountService.List(&repo.ListAccountFilter{
		GroupName: api.initConfig.ActiveGroup,
	})
	if err != nil {
		ctx.JSON(api.BaseResponseInternalServerError(err))
		return
	}

	report := report.NewWitdrawReport("withdraw_report.csv")
	for _, account := range accounts {
		username := account.GetUsername()
		event := sio_event.WithdrawEvent{
			Name:    username,
			Type:    "success",
			Message: "success",
		}

		err = api.accountService.Withdraw(username, account.AccountData.Pin, report)
		if err != nil {
			event.Type = "error"
			event.Message = err.Error()
		}

		api.sio.BroadcastToNamespace("", "withdraw", &sio_event.AccountWithdrawEvent{
			Shopid: account.AccountData.ShopID,
			Event:  &event,
		})
	}
}

type SetpinQuery struct {
	Shopid int `json:"shopid" form:"shopid"`
}

type Setpinpayload struct {
	Pin string `json:"pin"`
}

func (api *AccountApi) setPin(ctx *gin.Context) {

	query := SetpinQuery{}
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

	err = api.accountService.UpdateAccount(query.Shopid, func(account *model.Account) error {
		account.AccountData.Pin = payload.Pin
		return nil
	})
	if err != nil {
		ctx.JSON(api.BaseResponseBadRequest(err))
		return
	}

	ctx.JSON(api.BaseResponseSuccess())
}
