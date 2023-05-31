package query

const (
	ChatGetBundleSticker = `query chatGetBundleSticker($id: String!, $maxId: String) {
		  chatBundleSticker(groupUUID: $id, limit: 8, maxUUID: $maxId) {
		    list {
		      imageUrl
		      intention
		      stickerUUID
		      groupUUID
		      __typename
		    }
	    hasNext
	    maxUUID
	    __typename
	  }
	}`
)
