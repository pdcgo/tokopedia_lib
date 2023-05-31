package query

const (
	DiscussionDataByProductID = `query DiscussionDataByProductID($productID: String!, $shopID: String, $page: Int!, $limit: Int!, $sortBy: String, $category: String) {
		  discussionDataByProductID(productID: $productID, shopID: $shopID, page: $page, limit: $limit, sortBy: $sortBy, category: $category) {
		    shopID
		    shopURL
		    productID
		    hasNext
		    totalQuestion
		    question {
		      questionID
		      content
		      maskedContent
		      userName
		      userThumbnail
		      userID
		      createTime
		      createTimeFormatted
		      state {
		        isMasked
		        allowReply
		        isYours
		        isFollowed
		        allowFollow
		        allowReport
		        allowDelete
		        __typename
		      }
	      totalAnswer
	      answer {
		        answerID
		        content
		        maskedContent
		        userName
		        userThumbnail
		        userID
		        isSeller
		        createTime
		        createTimeFormatted
		        likeCount
		        state {
		          isMasked
		          isLiked
		          allowLike
		          isYours
		          allowReport
		          allowDelete
		          __typename
		        }
	        attachedProductCount
	        __typename
	      }
	      __typename
	    }
	    __typename
	  }
	}
	`
)
