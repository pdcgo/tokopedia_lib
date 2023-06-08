package api

type Category struct {
	ID       int         `json:"id"`
	Name     string      `json:"name"`
	URL      string      `json:"url"`
	Children []*Category `json:"children"`
	Typename string      `json:"__typename"`
}

type CategoryAllListLite struct {
	Categories []*Category `json:"categories"`
	Typename   string      `json:"__typename"`
}

type CategoryAllListLiteData struct {
	CategoryAllListLite *CategoryAllListLite `json:"categoryAllListLite"`
}
type CategoryAllListLiteRes struct {
	Data CategoryAllListLiteData `json:"data"`
}

func (api *TokopediaApi) CategoryAllListLite() (*CategoryAllListLiteRes, error) {

	query := GraphqlPayload{
		OperationName: "categoryAllListLite",
		Query: `
		query categoryAllListLite {
			categoryAllListLite(filter: "seller") {
			  categories {
				id
				name
				url
				children {
				  id
				  name
				  url
				  children {
					id
					name
					url
					__typename
				  }
				  __typename
				}
				__typename
			  }
			  __typename
			}
		  }`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil CategoryAllListLiteRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
