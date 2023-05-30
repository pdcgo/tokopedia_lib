package query

const (
	PdpGetReiewImageQuery = `query PDPGetReviewImageQuery($page: Int!, $productID: String!, $limit: Int!) {
		  productrevGetReviewImage(page: $page, productID: $productID, limit: $limit) {
		    list {
		      imageID
		      feedbackID
		      videoID
		      __typename
		    }
	    detail {
		      reviews: review {
		        reviewer: user {
		          userID
		          fullName
		          profilePicture: image
		          __typename
		        }
	        shopID
	        feedbackID
	        variantName
	        reviewText: review
	        rating
	        reviewTime: createTimestamp
	        likeCount: totalLike
	        badRatingReasonFmt
	        isLiked
	        isAnonymous
	        isReportable
	        __typename
	      }
	      images: image {
		        attachmentID
		        description
		        thumbnailURL
		        fullsizeURL
		        feedbackID
		        __typename
		      }
	      video {
		        attachmentID
		        url
		        feedbackID
		        __typename
		      }
	      mediaCount
	      __typename
	    }
	    hasNext
	    __typename
	  }
	}`
)
