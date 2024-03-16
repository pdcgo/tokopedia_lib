package service

import (
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/chat/group"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/app/chat/report"
	"github.com/pdcgo/tokopedia_lib/app/withdraw"
	"github.com/pdcgo/tokopedia_lib/lib/api"
)

type AccountService struct {
	sync.Mutex
	accountRepo *repo.AccountRepo
	driverGroup *group.DriverGroup
}

func NewAccountService(accountRepo *repo.AccountRepo, driverGroup *group.DriverGroup) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
		driverGroup: driverGroup,
	}
}

type Account struct {
	OtpPassword string `json:"otp_password"`
	Password    string `json:"password"`
	Username    string `json:"username"`
}

func (s *AccountService) createAccount(account Account, api *api.TokopediaApi) (*model.AccountData, error) {
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

func (s *AccountService) AddAccount(account Account, groupName string) error {

	lock := s.TryLock()
	if !lock {
		return errors.New("masih dalam pemrosesan akun lain")
	}
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

		err = s.accountRepo.AddAccountData(groupName, accountData)
		return err
	})
}

func (s *AccountService) SyncAccount(shopid int, notifHash string, notif *api.NotificationCounterRes) (err error) {

	notifData := notif.Data.Notifications
	err = s.accountRepo.UpdateAccount(shopid, func(account *model.Account) {
		account.UnreadChat = notifData.Chat.UnreadsSeller
		account.NewOrder = notifData.SellerOrderStatus.NewOrder
		account.Diskusi = notifData.Inbox.TalkSeller
		account.Online = true
		account.NotifHash = notifHash
	})
	return err
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
