package api

type CloseShopSchedule struct {
	Success  bool   `json:"success"`
	Message  string `json:"message"`
	Typename string `json:"__typename"`
}

type CloseShopScheduleData struct {
	CloseShopSchedule *CloseShopSchedule `json:"closeShopSchedule"`
}

type CloseShopScheduleRes struct {
	Data *CloseShopScheduleData `json:"data"`
}

type CloseShopScheduleInput struct {
	Action     int    `json:"action"`
	CloseNote  string `json:"closeNote"`
	CloseStart int    `json:"closeStart,string,omitempty"`
	CloseEnd   int    `json:"closeEnd,string,omitempty"`
}

type CloseShopScheduleVar struct {
	Input *CloseShopScheduleInput `json:"input"`
}

func (api *TokopediaApi) CloseShopSchedule(input *CloseShopScheduleInput) (*CloseShopScheduleRes, error) {

	query := GraphqlPayload{
		OperationName: "CloseShopSchedule",
		Variables: CloseShopScheduleVar{
			Input: input,
		},
		Query: `
		mutation CloseShopSchedule($input: CloseShop!) {
			closeShopSchedule(input: $input) {
			  success
			  message
			  __typename
			}
		  }
		`,
	}

	var hasil CloseShopScheduleRes
	req := api.NewGraphqlReq(&query)
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
