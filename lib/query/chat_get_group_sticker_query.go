package query

const (
	ChatGetGroupSticker = `query chatGetGroupSticker($type: Int!) {
		  chatListGroupSticker(stickerType: $type) {
		    list {
		      groupUUID
		      lastUpdate
		      thumbnail
		      title
		      __typename
		    }
	    __typename
	  }
	}
	`
)
