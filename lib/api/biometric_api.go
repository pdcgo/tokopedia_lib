package api

type PublicKeyRp struct {
	Name     string `json:"name"`
	Icon     string `json:"icon"`
	ID       string `json:"id"`
	Typename string `json:"__typename"`
}

type PublicKeyUser struct {
	Name        string `json:"name"`
	Icon        string `json:"icon"`
	DisplayName string `json:"displayName"`
	ID          string `json:"id"`
	Typename    string `json:"__typename"`
}

type PublicKeyParam struct {
	Type     string `json:"type"`
	Alg      int    `json:"alg"`
	Typename string `json:"__typename"`
}

type AuthenticatorSelection struct {
	AuthenticatorAttachment string `json:"authenticatorAttachment"`
	RequireResidentKey      bool   `json:"requireResidentKey"`
	UserVerification        string `json:"userVerification"`
	Typename                string `json:"__typename"`
}

type BiometricPublicKey struct {
	Challenge              string                  `json:"challenge"`
	Rp                     *PublicKeyRp            `json:"rp"`
	User                   *PublicKeyUser          `json:"user"`
	PubKeyCredParams       []*PublicKeyParam       `json:"pubKeyCredParams"`
	AuthenticatorSelection *AuthenticatorSelection `json:"authenticatorSelection"`
	Timeout                int                     `json:"timeout"`
	Typename               string                  `json:"__typename"`
}

type OTPBiometricBeginRegister struct {
	IsSuccess    bool                `json:"isSuccess"`
	ErrorMessage string              `json:"errorMessage"`
	PublicKey    *BiometricPublicKey `json:"publicKey"`
	Typename     string              `json:"__typename"`
}

type BiometricBeginRegister struct {
	OTPBiometricBeginRegister *OTPBiometricBeginRegister `json:"OTPBiometricBeginRegister"`
}

type BiometricBeginRegisterResp struct {
	Data *BiometricBeginRegister `json:"data"`
}

func (api *TokopediaApi) BiometricBeginRegister() (*BiometricBeginRegisterResp, error) {
	query := GraphqlPayload{
		OperationName: "biometricBeginRegister",
		Variables:     struct{}{},
		Query:         "query biometricBeginRegister {\n  OTPBiometricBeginRegister {\n    isSuccess\n    errorMessage\n    publicKey {\n      challenge\n      rp {\n        name\n        icon\n        id\n        __typename\n      }\n      user {\n        name\n        icon\n        displayName\n        id\n        __typename\n      }\n      pubKeyCredParams {\n        type\n        alg\n        __typename\n      }\n      authenticatorSelection {\n        authenticatorAttachment\n        requireResidentKey\n        userVerification\n        __typename\n      }\n      timeout\n      __typename\n    }\n    __typename\n  }\n}\n",
	}

	req := api.NewGraphqlReq(&query)

	var hasil BiometricBeginRegisterResp
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
