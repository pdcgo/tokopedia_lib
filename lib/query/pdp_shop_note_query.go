package query

const (
	PdpShopNote = `query PDPShopNote($shopID: String!) {
		  shopNotesByShopID(shopID: $shopID) {
		    result {
		      id
		      title
		      content
		      url
		      updateTime
		      __typename
		    }
	    error {
		      message
		      __typename
		    }
	    __typename
	  }
	}`
)
