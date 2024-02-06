package api

import (
	crand "crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"log"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type WithdrawVariable struct {
	IsJoinRP      bool   `json:"isJoinRP"`
	Action        string `json:"action"`
	Type          int    `json:"type"`
	Token         string `json:"token"`
	ValidateToken string `json:"validateToken"`
	DeviceType    string `json:"deviceType"`
	UserID        string `json:"userId"`
	Email         string `json:"email"`
	Amount        string `json:"amount"`
	MasterEmail   string `json:"masterEmail"`
	MasterID      string `json:"masterID"`
	AccountID     string `json:"accountID"`
	AccountName   string `json:"accountName"`
	AccountNumber string `json:"accountNumber"`
	BankID        string `json:"bankId"`
	BankName      string `json:"bankName"`
	Password      string `json:"password"`
	Lang          string `json:"lang"`
	IsSeller      bool   `json:"isSeller"`
	Program       string `json:"program"`
	IsAdmin       bool   `json:"isAdmin"`
}

func NewWithdrawVariable(user *User, bank *GetBankWDV2, otpValidate *OtpValidate, amount string) *WithdrawVariable {
	variable := &WithdrawVariable{
		Action:        "1",
		Type:          1,
		ValidateToken: otpValidate.ValidateToken,
		DeviceType:    "desktop",
		UserID:        user.ID,
		Email:         user.Email,
		Amount:        amount,
		AccountID:     strconv.Itoa(bank.BankAccountID),
		AccountName:   bank.AccountName,
		AccountNumber: bank.AccountNo,
		BankID:        strconv.Itoa(bank.BankID),
		BankName:      bank.BankName,
		Lang:          "id",
		IsSeller:      true,
	}

	return variable
}

type JoinPromptMessageResponse struct {
	Title      string `json:"title"`
	Desc       string `json:"desc"`
	ActionText string `json:"actionText"`
	ActionLink string `json:"actionLink"`
	IsSuccess  bool   `json:"isSuccess"`
	StatusCode int    `json:"statusCode"`
	Typename   string `json:"__typename"`
}

type RichieSubmitWithdrawal struct {
	ProcessTime               int                        `json:"process_time"`
	Message                   []interface{}              `json:"message"`
	Status                    string                     `json:"status"`
	MessageError              string                     `json:"message_error"`
	AccountName               string                     `json:"accountName"`
	AccountNumber             string                     `json:"accountNumber"`
	BankName                  string                     `json:"bankName"`
	Amount                    int                        `json:"amount"`
	AdminFee                  int                        `json:"adminFee"`
	WithdrawalNote            string                     `json:"withdrawalNote"`
	ErrorCode                 string                     `json:"errorCode"`
	JoinPromptMessageResponse *JoinPromptMessageResponse `json:"joinPromptMessageResponse"`
	Title                     string                     `json:"title"`
	Description               string                     `json:"description"`
	Header                    string                     `json:"header"`
	CtaLink                   string                     `json:"ctaLink"`
	CtaWording                string                     `json:"ctaWording"`
	Image                     string                     `json:"image"`
	Typename                  string                     `json:"__typename"`
}

type WithdrawSaldoMutationResp struct {
	Data struct {
		RichieSubmitWithdrawal *RichieSubmitWithdrawal `json:"richieSubmitWithdrawal"`
	} `json:"data"`
}

func (api *TokopediaApi) WithdrawSaldoMutation(variable *WithdrawVariable) (*WithdrawSaldoMutationResp, error) {
	query := GraphqlPayload{
		OperationName: "withdrawSaldoMutation",
		Variables:     variable,
		Query:         "mutation withdrawSaldoMutation($isSeller: Boolean, $amount: String!, $userId: String!, $email: String!, $action: String!, $type: Int!, $deviceType: String!, $token: String!, $masterEmail: String!, $masterID: String!, $accountID: String!, $accountName: String!, $accountNumber: String!, $bankId: String!, $bankName: String!, $lang: String!, $validateToken: String!, $password: String, $program: String, $isAdmin: Boolean, $isJoinRP: Boolean = false) {\n  richieSubmitWithdrawal(input: {isSeller: $isSeller, amount: $amount, userId: $userId, email: $email, action: $action, type: $type, deviceType: $deviceType, token: $token, masterEmail: $masterEmail, masterID: $masterID, accountID: $accountID, accountName: $accountName, accountNumber: $accountNumber, bankId: $bankId, bankName: $bankName, lang: $lang, validateToken: $validateToken, password: $password, program: $program, isAdmin: $isAdmin, isJoinRP: $isJoinRP}) {\n    process_time\n    message\n    status\n    message_error\n    accountName\n    accountNumber\n    bankName\n    amount\n    adminFee\n    withdrawalNote\n    errorCode\n    joinPromptMessageResponse {\n      title\n      desc\n      actionText\n      actionLink\n      isSuccess\n      statusCode\n      __typename\n    }\n    title\n    description\n    header\n    ctaLink\n    ctaWording\n    image\n    __typename\n  }\n}\n",
	}

	req := api.NewGraphqlReq(&query)

	ts := int(time.Now().UnixMilli())
	req.Header.Set("x-auth-signature", "") // after testing this can be blank
	req.Header.Set("x-auth-hash", "")      // after testing this can be blank
	req.Header.Set("x-auth-timestamp", strconv.Itoa(ts))

	var hasil WithdrawSaldoMutationResp
	err := api.SendRequest(req, &hasil)
	if err != nil {
		return nil, err
	}

	if hasil.Data.RichieSubmitWithdrawal.Status != "success" {
		return nil, errors.New(hasil.Data.RichieSubmitWithdrawal.MessageError)
	}

	return &hasil, err

}

type OtpRequestVariable struct {
	Msisdn   string `json:"msisdn"`
	OtpType  string `json:"otpType"`
	Mode     string `json:"mode"`
	OtpDigit int    `json:"otpDigit"`
}

func NewOtpRequestPINVariable(msisdn string) *OtpRequestVariable {
	return &OtpRequestVariable{
		Msisdn:   msisdn,
		OtpType:  "120",
		Mode:     "PIN",
		OtpDigit: 6,
	}
}

type OtpRequest struct {
	Success            bool          `json:"success"`
	Message            string        `json:"message"`
	ErrorMessage       string        `json:"errorMessage"`
	SseSessionID       string        `json:"sse_session_id"`
	ListDeviceReceiver []interface{} `json:"list_device_receiver"`
	ErrorCode          string        `json:"error_code"`
	MessageTitle       string        `json:"message_title"`
	MessageSubTitle    string        `json:"message_sub_title"`
	MessageImgLink     string        `json:"message_img_link"`
	Typename           string        `json:"__typename"`
}

type OtpRequestResp struct {
	Data struct {
		OTPRequest *OtpRequest `json:"OTPRequest"`
	} `json:"data"`
}

func (api *TokopediaApi) withdrawOtpRequest(variable *OtpRequestVariable) (*OtpRequestResp, error) {
	query := GraphqlPayload{
		OperationName: "OTPRequest",
		Variables:     variable,
		Query:         "query OTPRequest($otpType: String!, $mode: String, $msisdn: String, $email: String, $otpDigit: Int, $ValidateToken: String, $UserIDEnc: String, $UserIDSigned: Int, $Signature: String, $MsisdnEnc: String, $EmailEnc: String) {\n  OTPRequest(otpType: $otpType, mode: $mode, msisdn: $msisdn, email: $email, otpDigit: $otpDigit, ValidateToken: $ValidateToken, UserIDEnc: $UserIDEnc, UserIDSigned: $UserIDSigned, Signature: $Signature, MsisdnEnc: $MsisdnEnc, EmailEnc: $EmailEnc) {\n    success\n    message\n    errorMessage\n    sse_session_id\n    list_device_receiver\n    error_code\n    message_title\n    message_sub_title\n    message_img_link\n    __typename\n  }\n}\n",
	}

	req := api.NewGraphqlReq(&query)

	var hasil OtpRequestResp
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

func (api *TokopediaApi) WithdrawOtpRequest(msisdn string) (*OtpRequestResp, error) {
	payload := NewOtpRequestPINVariable(msisdn)

	return api.withdrawOtpRequest(payload)
}

type OtpValidateVariable struct {
	Msisdn     string `json:"msisdn"`
	BankAccID  string `json:"BankAccID"`
	UsePINHash bool   `json:"UsePINHash"`
	PIN        string `json:"PIN"`
	PINHash    string `json:"PINHash"`
	Code       string `json:"code"`
	OtpType    string `json:"otpType"`
	Mode       string `json:"mode"`
}

func NewOtpValidateVariable(msisdn, bankAccountId, pin string, generateKey *GenerateKey) *OtpValidateVariable {
	variable := &OtpValidateVariable{
		Msisdn:     msisdn,
		BankAccID:  bankAccountId,
		UsePINHash: true,
		PIN:        pin,
		PINHash:    generateKey.H,
		Mode:       "PIN",
		OtpType:    "120",
	}
	return variable
}

type OtpValidate struct {
	Success       bool          `json:"success"`
	Message       string        `json:"message"`
	ErrorMessage  string        `json:"errorMessage"`
	ValidateToken string        `json:"validateToken"`
	CookieList    []interface{} `json:"cookieList"`
	Typename      string        `json:"__typename"`
}

type OtpValidateResp struct {
	Data struct {
		OTPValidate *OtpValidate `json:"OTPValidate"`
	} `json:"data"`
}

func (api *TokopediaApi) WithdrawOtpValidate(payload *OtpValidateVariable) (*OtpValidateResp, error) {
	query := GraphqlPayload{
		OperationName: "OTPValidate",
		Variables:     payload,
		Query:         "query OTPValidate($msisdn: String, $code: String!, $otpType: String, $fpData: String, $getSL: String, $email: String, $mode: String, $ValidateToken: String, $UserIDEnc: String, $UserID: Int, $signature: String, $UsePINHash: Boolean, $PIN: String, $PINHash: String, $BankAccID: String) {\n  OTPValidate(code: $code, otpType: $otpType, msisdn: $msisdn, fpData: $fpData, getSL: $getSL, email: $email, mode: $mode, ValidateToken: $ValidateToken, UserIDEnc: $UserIDEnc, UserID: $UserID, signature: $signature, UsePINHash: $UsePINHash, PIN: $PIN, PINHash: $PINHash, BankAccID: $BankAccID) {\n    success\n    message\n    errorMessage\n    validateToken\n    cookieList {\n      key\n      value\n      expire\n      __typename\n    }\n    __typename\n  }\n}\n",
	}

	req := api.NewGraphqlReq(&query)

	var hasil OtpValidateResp
	err := api.SendRequest(req, &hasil)

	if err != nil {
		return nil, err
	}

	if !hasil.Data.OTPValidate.Success {
		return nil, errors.New(hasil.Data.OTPValidate.ErrorMessage)
	}

	return &hasil, nil
}

type WithdrawGenerateKeyVariable struct {
	Module string `json:"module"`
}

type GenerateKey struct {
	Key             string `json:"key"`
	ServerTimestamp int    `json:"server_timestamp"`
	H               string `json:"h"`
	Typename        string `json:"__typename"`
}

func (g *GenerateKey) GetRsaPublicKey() (string, error) {
	key, err := base64.StdEncoding.DecodeString(g.Key)
	return string(key), err
}

type WindrawnGenerateKeyResp struct {
	Data struct {
		GenerateKey *GenerateKey `json:"generate_key"`
	} `json:"data"`
}

func (w *WindrawnGenerateKeyResp) GetRSAPublicKeyContent() (string, error) {
	publicKey, err := w.Data.GenerateKey.GetRsaPublicKey()
	if err != nil {
		return "", err
	}

	resSplits := strings.Split(publicKey, "\n")
	content := strings.Join(resSplits[1:len(resSplits)-1], "")

	return content, nil
}

func (api *TokopediaApi) WindrawnGenerateKey() (*WindrawnGenerateKeyResp, error) {
	payload := &WithdrawGenerateKeyVariable{
		Module: "pinv2",
	}
	query := GraphqlPayload{
		OperationName: "GenerateKey",
		Variables:     payload,
		Query:         "query GenerateKey($module: String!) {\n  generate_key(module: $module) {\n    key\n    server_timestamp\n    h\n    __typename\n  }\n}\n",
	}

	req := api.NewGraphqlReq(&query)

	req.Header.Set("Accounts-Authorization", RandomAccountsAuthorization(8))

	var hasil WindrawnGenerateKeyResp
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}

type PinV2CheckVariable struct {
	ID   string `json:"id"`
	Type string `json:"type"`
}

type PinV2CheckResp struct {
	Data struct {
		PinV2Check struct {
			Uh           bool   `json:"uh"`
			ErrorMessage string `json:"error_message"`
			Typename     string `json:"__typename"`
		} `json:"pinV2Check"`
	} `json:"data"`
}

func (api *TokopediaApi) pinV2Check(payload *PinV2CheckVariable) (*PinV2CheckResp, error) {
	query := GraphqlPayload{
		OperationName: "pinV2Check",
		Variables:     payload,
		Query:         "query pinV2Check($id: String, $type: String) {\n  pinV2Check(id: $id, type: $type) {\n    uh\n    error_message\n    __typename\n  }\n}\n",
	}

	req := api.NewGraphqlReq(&query)

	req.Header.Set("Accounts-Authorization", RandomAccountsAuthorization(8))

	var hasil PinV2CheckResp
	err := api.SendRequest(req, &hasil)
	if err != nil {
		return nil, err
	}

	if hasil.Data.PinV2Check.ErrorMessage != "" {
		return nil, errors.New(hasil.Data.PinV2Check.ErrorMessage)
	}

	return &hasil, nil
}

func (api *TokopediaApi) PinV2Check(phone string) (*PinV2CheckResp, error) {
	payload := &PinV2CheckVariable{
		ID:   phone,
		Type: "phone",
	}

	return api.pinV2Check(payload)
}

type OTPModeListQueryVariable struct {
	OtpType   string `json:"otpType"`
	Msisdn    string `json:"msisdn"`
	Email     string `json:"email"`
	BankAccID string `json:"BankAccID"`
}

func NewOTPModelListQueryVariable(msisdn, bankAccId string) *OTPModeListQueryVariable {
	return &OTPModeListQueryVariable{
		OtpType:   "120",
		Msisdn:    msisdn,
		BankAccID: bankAccId,
	}
}

type ModelList struct {
	OtpListText          string `json:"otpListText"`
	ModeText             string `json:"modeText"`
	AfterOtpListText     string `json:"afterOtpListText"`
	AfterOtpListTextHTML string `json:"afterOtpListTextHtml"`
	OtpListImgURL        string `json:"otpListImgUrl"`
	Typename             string `json:"__typename"`
}

type OtpModelList struct {
	Success             bool         `json:"success"`
	Message             string       `json:"message"`
	AutoReadLite        bool         `json:"autoReadLite"`
	TickerTrouble       string       `json:"tickerTrouble"`
	EnableTicker        bool         `json:"enableTicker"`
	DefaultBehaviorMode int          `json:"defaultBehaviorMode"`
	ModeLists           []*ModelList `json:"modeLists"`
	OtpDigit            int          `json:"otpDigit"`
	Typename            string       `json:"__typename"`
}

type OTPModeListQueryResp struct {
	Data struct {
		OTPModeList *OtpModelList `json:"OTPModeList"`
	} `json:"data"`
}

func (api *TokopediaApi) otpModeListQuery(payload *OTPModeListQueryVariable) (*OTPModeListQueryResp, error) {
	query := GraphqlPayload{
		OperationName: "OTPModeListQuery",
		Variables:     payload,
		Query:         "query OTPModeListQuery($otpType: String!, $msisdn: String, $email: String, $ValidateToken: String, $UserIDEnc: String, $userID: String, $Signature: String, $MsisdnEnc: String, $EmailEnc: String, $BankAccID: String) {\n  OTPModeList(otpType: $otpType, msisdn: $msisdn, email: $email, ValidateToken: $ValidateToken, UserIDEnc: $UserIDEnc, userID: $userID, Signature: $Signature, MsisdnEnc: $MsisdnEnc, EmailEnc: $EmailEnc, BankAccID: $BankAccID) {\n    success\n    message\n    autoReadLite\n    tickerTrouble\n    enableTicker\n    defaultBehaviorMode\n    modeLists {\n      otpListText\n      modeText\n      afterOtpListText\n      afterOtpListTextHtml\n      otpListImgUrl\n      __typename\n    }\n    otpDigit\n    __typename\n  }\n}\n",
	}

	req := api.NewGraphqlReq(&query)

	var hasil OTPModeListQueryResp
	err := api.SendRequest(req, &hasil)

	if err != nil {
		return nil, err
	}

	if !hasil.Data.OTPModeList.Success {
		return nil, errors.New(hasil.Data.OTPModeList.Message)
	}

	return &hasil, nil
}

func (api *TokopediaApi) OTPModeListQuery(msisdn, bankAccId string) (*OTPModeListQueryResp, error) {
	payload := NewOTPModelListQueryVariable(msisdn, bankAccId)

	return api.otpModeListQuery(payload)
}

func RandomAccountsAuthorization(len int) string {
	t := ""
	n := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	for r := 0; r < len; r++ {
		rand := rand.Float64() * float64(len)
		floor := int(math.Floor(rand))

		t += string(n[floor])
	}

	return t
}

func GetPublicKey(pub string) (*rsa.PublicKey, error) {
	block, _ := pem.Decode([]byte(pub))
	b := block.Bytes
	ifc, err := x509.ParsePKIXPublicKey(b)
	if err != nil {
		return nil, err
	}
	switch pub := ifc.(type) {
	case *rsa.PublicKey:
		return pub, nil
	default:
		break
	}
	return nil, errors.New("public key type incorrect")
}

// TODO: This is still wrong.
//
// The result if used in OTPValidate is wrong.
func EncryptPIN(msg string, key string) (string, error) {
	pinSalt := "b9f14c8ed04a41c7a5361b648a088b69"
	saltedPin := fmt.Sprintf("%s%s", msg, pinSalt)

	log.Println(saltedPin)

	pub, err := GetPublicKey(key)
	if err != nil {
		return "", err
	}

	hash := sha256.New()
	ciphertext, err := rsa.EncryptOAEP(hash, crand.Reader, pub, []byte(saltedPin), nil)
	if err != nil {
		return "", err
	}

	encryptedPin := base64.StdEncoding.EncodeToString(ciphertext)

	return encryptedPin, nil
}
