package service

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/pdcgo/common_conf/common_concept"
	"github.com/pdcgo/common_conf/pdc_common"
	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/chat/config"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/report"
	"github.com/pdcgo/tokopedia_lib/app/chat/sio_event"
	"github.com/pdcgo/tokopedia_lib/app/withdraw"
	"github.com/pdcgo/tokopedia_lib/lib/api"
	apimodel "github.com/pdcgo/tokopedia_lib/lib/model"
)

type AccountService struct {
	sync.Mutex
	*group.ChatGroup
	*repo.AccountRepo
	initConfig    *config.InitConfig
	event         *common_concept.CoreEvent
	driverGroup   *group.DriverGroup
	browserCancel context.CancelFunc
}

func NewAccountService(
	initConfig *config.InitConfig,
	event *common_concept.CoreEvent,
	accountRepo *repo.AccountRepo,
	driverGroup *group.DriverGroup,
	chatGroup *group.ChatGroup,
) *AccountService {

	accountService := AccountService{
		initConfig:  initConfig,
		event:       event,
		ChatGroup:   chatGroup,
		AccountRepo: accountRepo,
		driverGroup: driverGroup,
	}

	go accountService.handleEvent()
	return &accountService
}

type AccountPayload struct {
	OtpPassword string `json:"otp_password"`
	Password    string `json:"password"`
	Username    string `json:"username"`
}

func (s *AccountService) createAccount(account AccountPayload, api *api.TokopediaApi) (*model.AccountData, error) {
	auth, err := api.IsAutheticated()
	if err != nil {
		return nil, err
	}

	user := auth.Data.User
	if strings.EqualFold(account.Username, user.Name) &&
		strings.EqualFold(account.Username, user.Email) {

		msg := fmt.Sprintf(
			"username %s tidak sama dengan %s / %s",
			account.Username, user.Name, user.Email,
		)
		return nil, errors.New(msg)
	}

	info, err := api.AccountInfo()
	if err != nil {
		return nil, err
	}

	notif, err := api.NotificationCounter()
	if err != nil {
		return nil, err
	}

	shopInfo := info.Data.UserShopInfo.Info
	notifications := notif.Data.Notifications
	accountData := model.AccountData{
		Username:    account.Username,
		Password:    account.Password,
		OtpPassword: account.OtpPassword,
		ShopID:      shopInfo.ShopID,
		Account: model.Account{
			ID:         shopInfo.ShopID,
			ProfileUrl: shopInfo.ShopAvatar,
			ShopName:   shopInfo.ShopName,
			ShopDomain: shopInfo.ShopDomain,
			UnreadChat: notifications.NotifcenterTrxUnread.NotifUnreadSellerInt,
			NewOrder:   notifications.SellerOrderStatus.NewOrder,
		},
	}

	return &accountData, nil
}

func (s *AccountService) AddAccount(account AccountPayload, groupName string) error {

	s.Lock()
	defer s.Unlock()

	err := s.driverGroup.AddDriver(account.Username, account.Password, account.OtpPassword)
	if err != nil {
		return err
	}

	return s.driverGroup.WithDriverApi(account.Username, func(driver *tokopedia_lib.DriverAccount, api *api.TokopediaApi) error {

		accountData, err := s.createAccount(account, api)
		if err != nil {
			return err
		}

		err = s.AddAccountData(groupName, accountData)
		return err
	})
}

func (s *AccountService) SyncAccount(shopid int, notifHash string, notif *api.NotificationCounterRes) (err error) {

	s.Lock()
	defer s.Unlock()

	notifData := notif.Data.Notifications
	err = s.UpdateAccount(shopid, func(account *model.Account) error {
		account.UnreadChat = notifData.Chat.UnreadsSeller
		account.NewOrder = notifData.SellerOrderStatus.NewOrder
		account.Diskusi = notifData.Inbox.TalkSeller
		account.Online = true
		account.NotifHash = notifHash
		return nil
	})
	return err
}

func (s *AccountService) OpenBrowser(username string) {

	s.Lock()
	defer s.Unlock()

	if s.browserCancel != nil {
		s.browserCancel()
	}

	cancel, err := s.driverGroup.OpenDriver(username)
	s.browserCancel = cancel
	if err != nil {
		pdc_common.ReportError(err)
	}
}

var ErrPinKosong = errors.New("pin kosong")
var WdLock sync.Mutex

func (s *AccountService) Withdraw(username string, pin string, report *report.WitdrawReport) (err error) {

	WdLock.Lock()
	defer func() {
		time.Sleep(time.Second)
		WdLock.Unlock()
	}()

	item := &withdraw.WithdrawReport{
		Jumlah:    "Rp0",
		SisaSaldo: "Rp0",
	}
	if report != nil {
		report.Add(item)
		defer report.Save()
	}

	return s.driverGroup.WithDriverApi(username, func(driver *tokopedia_lib.DriverAccount, api *api.TokopediaApi) error {

		items, err := withdraw.GetUnwithdrawTransaction(api)
		if err != nil {
			return err
		}
		if report != nil {
			report.Add(items...)
		}

		if pin == "" {
			item.Keterangan = ErrPinKosong.Error()
			return ErrPinKosong
		}

		return driver.Run(false, func(dctx *tokopedia_lib.DriverContext) error {
			wd := withdraw.NewWithdraw(api)
			err := wd.Run(dctx, pin, item)
			if err != nil {
				item.Keterangan = err.Error()
			}
			return err
		})
	})
}

func (s *AccountService) GetLocations(username string) ([]apimodel.ShopLocationLegacy, error) {

	var locations []apimodel.ShopLocationLegacy
	err := s.driverGroup.WithDriverApi(username, func(driver *tokopedia_lib.DriverAccount, tapi *api.TokopediaApi) error {

		shopid := int(tapi.AuthenticatedData.UserShopInfo.Info.ShopID)
		locationAll, err := tapi.GetShopLocationAll(shopid)
		if err != nil {
			return err
		}

		locations = locationAll.Data.ShopLocGetAllLocations.Data.Warehouses.GetLocations()
		return nil
	})

	return locations, err
}

func (s *AccountService) updateActive(shopid int) error {
	return s.WithAccount(s.initConfig.ActiveGroup, shopid, func(account *model.Account) error {
		log.Printf("[ %s ] set active", account.AccountData.Username)
		return s.driverGroup.WithDriverApi(account.AccountData.Username, func(driver *tokopedia_lib.DriverAccount, tapi *api.TokopediaApi) error {
			_, err := tapi.SetShopActive()
			return err
		})
	})
}

func (s *AccountService) updateSaldo(shopid int) error {
	return s.WithAccount(s.initConfig.ActiveGroup, shopid, func(account *model.Account) error {
		log.Printf("[ %s ] getting saldo", account.AccountData.Username)
		return s.driverGroup.WithDriverApi(account.AccountData.Username, func(driver *tokopedia_lib.DriverAccount, api *api.TokopediaApi) error {
			balance, err := api.GetBalance()
			if err != nil {
				return err
			}
			return s.UpdateAccount(shopid, func(account *model.Account) error {
				account.Saldo = balance.Data.Balance.SellerAll
				return nil
			})
		})
	})
}

func (s *AccountService) handleEvent() {

	for event := range s.event.GetEvent() {
		switch ev := event.(type) {

		case *sio_event.AccountActiveEvent:
			err := s.updateActive(ev.Shopid)
			if err != nil {
				pdc_common.ReportError(err)
			}

		case *sio_event.SocketConnectEvent:
			err := s.updateSaldo(ev.Shopid)
			if err != nil {
				pdc_common.ReportError(err)
			}
		}
	}
}
