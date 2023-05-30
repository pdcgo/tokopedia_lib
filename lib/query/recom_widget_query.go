package query

const (
	RecomWidget = `query RecomWidget($userID: Int!, $pageName: String!, $xSource: String!, $xDevice: String!, $productIDs: String, $LayoutPageType: String, $ref: String, $categoryIDs: String, $queryParam: String, $pageNumber: Int!) {
		  productRecommendationWidget(userID: $userID, pageName: $pageName, pageNumber: $pageNumber, xSource: $xSource, xDevice: $xDevice, productIDs: $productIDs, LayoutPageType: $LayoutPageType, ref: $ref, categoryIDs: $categoryIDs, queryParam: $queryParam) {
		    data {
		      tID
		      source
		      title
		      foreignTitle
		      seeMoreUrlLink
		      layoutType
		      pageName
		      widgetUrl
		      pagination {
		        hasNext
		        __typename
		      }
	      recommendation {
		        productSlashedPrice: slashedPrice
		        slashedPriceInt
		        productDiscountPercentage: discountPercentage
		        productReviewCount: countReview
		        isWishlist: isWishlist
		        productImageUrl: imageUrl
		        isTopads
		        clickUrl
		        trackerImageUrl
		        productUrl: url
		        productRating: rating
		        productPrice: price
		        priceInt
		        id
		        productName: name
		        categoryBreadcrumbs
		        recommendationType
		        stock
		        departmentID: departmentId
		        shop {
		          id
		          name
		          location
		          city
		          url
		          __typename
		        }
	        productLabels: labels {
		          title
		          color
		          __typename
		        }
	        labelGroup: labelgroup {
		          type
		          title
		          position
		          url
		          __typename
		        }
	        wholesalePrice {
		          price
		          quantityMax
		          quantityMin
		          priceString
		          __typename
		        }
	        badges {
		          title
		          imageUrl
		          __typename
		        }
	        __typename
	      }
	      __typename
	    }
	    meta {
		      recommendation
		      size
		      failSize
		      processTime
		      experimentVersion
		      __typename
		    }
	    __typename
	  }
	}`
)
