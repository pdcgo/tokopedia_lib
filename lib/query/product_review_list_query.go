package query

const (
	ProductReviewList = `query productReviewList($productID: String!, $page: Int!, $limit: Int!, $sortBy: String, $filterBy: String) {
		  productrevGetProductReviewList(productID: $productID, page: $page, limit: $limit, sortBy: $sortBy, filterBy: $filterBy) {
		    productID
		    list {
		      id: feedbackID
		      variantName
		      message
		      productRating
		      reviewCreateTime
		      reviewCreateTimestamp
		      isReportable
		      isAnonymous
		      imageAttachments {
		        attachmentID
		        imageThumbnailUrl
		        imageUrl
		        __typename
		      }
	      videoAttachments {
		        attachmentID
		        videoUrl
		        __typename
		      }
	      reviewResponse {
		        message
		        createTime
		        __typename
		      }
	      user {
		        userID
		        fullName
		        image
		        url
		        __typename
		      }
	      likeDislike {
		        totalLike
		        likeStatus
		        __typename
		      }
	      stats {
		        key
		        formatted
		        count
		        __typename
		      }
	      badRatingReasonFmt
	      __typename
	    }
	    shop {
		      shopID
		      name
		      url
		      image
		      __typename
		    }
	    hasNext
	    totalReviews
	    __typename
	  }
	}`
)
