package query

const (
	ProductRatingandTopics = "query productRatingAndTopics($productID: String!) {\n  productrevGetProductRatingAndTopics(productID: $productID) {\n    productID\n    rating {\n      positivePercentageFmt\n      ratingScore\n      totalRating\n      totalRatingWithImage\n      totalRatingTextAndImage\n      detail {\n        rate\n        totalReviews\n        formattedTotalReviews\n        percentageFloat\n        __typename\n      }\n      __typename\n    }\n    topics {\n      rating\n      ratingFmt\n      formatted\n      key\n      reviewCount\n      reviewCountFmt\n      show\n      __typename\n    }\n    availableFilters {\n      withAttachment\n      rating\n      topics\n      helpfulness\n      __typename\n    }\n    __typename\n  }\n}\n"
)
