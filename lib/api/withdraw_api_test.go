package api_test

import (
	"encoding/base64"
	"log"
	"strconv"
	"testing"

	"github.com/pdcgo/tokopedia_lib/lib/api"
	"github.com/pdcgo/tokopedia_lib/scenario"
	"github.com/stretchr/testify/assert"
)

func TestRandomAccountsAuthorization(t *testing.T) {
	l := 8
	res := api.RandomAccountsAuthorization(l)

	assert.Equal(t, len(res), l)
	log.Println(res)
}

func TestEncryptPIN(t *testing.T) {
	key, err := base64.StdEncoding.DecodeString("LS0tLS1CRUdJTiBQVUJMSUMgS0VZLS0tLS0KTUlJQklqQU5CZ2txaGtpRzl3MEJBUUVGQUFPQ0FROEFNSUlCQ2dLQ0FRRUF3S1JNeFFmZGtWT25lQnUvWUExeAowTUJER1hJcEZpZU1LaDlFcnA3RWc4Ny9RL3F4TlZYVk9rVUJ4WGF5SWd0K0lpbXc0L3ZObE52T2M3Uit2M1BaClRnK1h0WVNGM2NLdjJxK1pYYkVQVDNJdzEzS3ZjVjdHc0x1dDhtVXZKcEZ3WFNjUjZXY2lraFBVQ3h6UlcrZzEKSVI4Q0l6VlkvaHE2ekVvS3NRdkpGeFlwNmpxbWs3enB2cFNjZmc0MzJpNVlTSXpaK1Z0Yi9hQ3BDbmE0bU9DcQpoU0VVZGp0VVVTQUNaTDVRa0JyT1dXRSt3czRnY0ZjcFFDU0x6bFpIUXdvK3kzTUpIVmpaMlhPaGFuOFFLQ2pzCi9FVTc3UGIwQjVlcnpFaGRvWnFUbU96d05aTU10OW00REUrTDZsangxdnp1MnZEOVk4aTE4Q2l1dndQU1FyM3EKUVFJREFRQUIKLS0tLS1FTkQgUFVCTElDIEtFWS0tLS0t")
	assert.Nil(t, err)

	pinEncrypt, err := api.EncryptPIN("778899", string(key))
	assert.Nil(t, err)
	log.Println(pinEncrypt)
}

func TestWithdraw(t *testing.T) {
	sellerApi, saveSession := scenario.GetTokopediaApiClient()
	defer saveSession()

	hasil, err := sellerApi.UserDataQuery()
	assert.NotEmpty(t, hasil)
	assert.Nil(t, err)

	msisdn := hasil.Data.User.Phone

	banks, err := sellerApi.BankListQuery(false)
	assert.NotEmpty(t, banks)
	assert.Nil(t, err)

	bank := banks.Data.GetBankListWDV2.Data.GetDefaultBank()

	balance, err := sellerApi.WithDrawBalance()
	assert.Nil(t, err)

	var generateKey *api.GenerateKey

	var otpValidate *api.OtpValidate

	t.Run("test mode list query withdraw", func(t *testing.T) {
		hasil, err := sellerApi.OTPModeListQuery(msisdn, strconv.Itoa(bank.BankAccountID))
		assert.NotEmpty(t, hasil)
		assert.Nil(t, err)
		assert.True(t, hasil.Data.OTPModeList.Success)
	})

	t.Run("test otp request withdraw", func(t *testing.T) {
		hasil, err := sellerApi.WithdrawOtpRequest(msisdn)
		assert.NotEmpty(t, hasil)
		assert.Nil(t, err)
	})

	t.Run("test withdraw generate key", func(t *testing.T) {
		hasil, err := sellerApi.WindrawnGenerateKey()
		assert.NotEmpty(t, hasil)
		assert.Nil(t, err)

		generateKey = hasil.Data.GenerateKey
	})

	t.Run("test pinv2check", func(t *testing.T) {
		hasil, err := sellerApi.PinV2Check(msisdn)
		assert.NotEmpty(t, hasil)
		assert.Nil(t, err)
	})

	t.Run("test otp validate withdraw", func(t *testing.T) {
		payload, err := api.NewOtpValidateVariable(msisdn, strconv.Itoa(bank.BankAccountID), "778899", generateKey)
		assert.Nil(t, err)

		hasil, err := sellerApi.WithdrawOtpValidate(payload)
		assert.Nil(t, err)
		assert.NotNil(t, hasil)

		otpValidate = hasil.Data.OTPValidate
	})

	t.Run("test withdraw saldo mutation", func(t *testing.T) {
		if otpValidate.ValidateToken == "" {
			otpValidate.ValidateToken = "d6b35dc5792a4d72bf6e12eec2c58ca3"
		}
		payload := api.NewWithdrawVariable(hasil.Data.User, bank, otpValidate, strconv.Itoa(balance.Data.MidasGetAllDepositAmount.BuyerAll))

		hasil, err := sellerApi.WithdrawSaldoMutation(payload)
		assert.Nil(t, err)
		assert.NotNil(t, hasil)
	})
}
