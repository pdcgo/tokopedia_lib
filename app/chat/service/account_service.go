package service

import (
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/pdcgo/tokopedia_lib"
	"github.com/pdcgo/tokopedia_lib/app/chat/model"
	"github.com/pdcgo/tokopedia_lib/app/chat/repo"
	"github.com/pdcgo/tokopedia_lib/lib/api"
)

type AccountService struct {
	sync.Mutex
	accountRepo *repo.AccountRepo
}

func NewAccountService(accountRepo *repo.AccountRepo) *AccountService {
	return &AccountService{
		accountRepo: accountRepo,
	}
}

type Account struct {
	OtpPassword string `json:"otp_password"`
	Password    string `json:"password"`
	Username    string `json:"username"`
}

func (s *AccountService) AddAccount(account Account, groupName string) error {

	lock := s.TryLock()
	if !lock {
		return errors.New("masih dalam pemrosesan akun lain")
	}
	defer s.Unlock()

	driver, err := tokopedia_lib.NewDriverAccount(
		account.Username,
		account.Password,
		account.OtpPassword,
	)
	if err != nil {
		return err
	}

	api, saveSession, err := driver.CreateApi()
	if err != nil {
		return err
	}
	defer saveSession()

	auth, err := api.IsAutheticated()
	if err != nil {
		return err
	}

	user := auth.Data.User
	if strings.EqualFold(account.Username, user.Name) &&
		strings.EqualFold(account.Username, user.Email) {

		msg := fmt.Sprintf(
			"username %s tidak sama dengan %s / %s",
			account.Username, user.Name, user.Email,
		)
		return errors.New(msg)
	}

	info, err := api.AccountInfo()
	if err != nil {
		return err
	}

	notif, err := api.NotificationCounter()
	if err != nil {
		return err
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

	err = s.accountRepo.AddAccountData(groupName, accountData)
	return err
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
