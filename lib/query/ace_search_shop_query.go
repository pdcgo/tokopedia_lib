package query

const (
	AceSearchShop = `query AceSearchShop($params: String!) {
		  aceSearchShop(params: $params) {
		    totalData: total_shop
		    shops {
		      id: shop_id
		      name: shop_name
		      domain: shop_domain
		      ownerId: shop_is_owner
		      city: shop_location
		      shopStatus: shop_status
		      tagLine: shop_tag_line
		      desc: shop_description
		      reputationScore: reputation_score
		      totalFave: shop_total_favorite
		      isPowerBadge: shop_gold_shop
		      isPMPro: is_pm_pro
		      isOfficial: is_official
		      url: shop_url
		      imageURL: shop_image
		      reputationImageURL: reputation_image_uri
		      shopLucky: shop_lucky
		      products {
		        id
		        name
		        url
		        price
		        productImg: image_url
		        priceText: price_format
		        __typename
		      }
	      GAKey: ga_key
	      favorited
	      voucher {
		        freeShipping: free_shipping
		        cashback {
		          cashbackValue: cashback_value
		          isPercentage: is_percentage
		          __typename
		        }
	        __typename
	      }
	      __typename
	    }
	    suggestion {
		      currentKeyword
		      suggestionText: text
		      suggestionTextQuery: query
		      __typename
		    }
	    header {
		      keywordProcess: keyword_process
		      responseCode: response_code
		      totalData: total_data
		      __typename
		    }
	    __typename
	  }
	}
	`
)
