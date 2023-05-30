package query

const (
	TopadsProductQuery = `query TopadsProductQuery($adParams: String) {
		  displayAdsV3(displayParams: $adParams) {
		    data {
		      clickTrackUrl: product_click_url
		      product_wishlist_url
		      product {
		        id
		        name
		        wishlist
		        image {
		          imageUrl: s_ecs
		          trackerImageUrl: s_url
		          __typename
		        }
	        url: uri
	        relative_uri
	        price: price_format
	        wholeSalePrice: wholesale_price {
		          quantityMin: quantity_min_format
		          quantityMax: quantity_max_format
		          price: price_format
		          __typename
		        }
	        count_talk_format
	        countReviewFormat: count_review_format
	        category {
		          id
		          __typename
		        }
	        categoryBreadcrumb: category_breadcrumb
	        preorder: product_preorder
	        product_wholesale
	        free_return
	        isNewProduct: product_new_label
	        cashback: product_cashback_rate
	        rating: product_rating
	        ratingAverage: product_rating_format
	        top_label
	        bottomLabel: bottom_label
	        labelGroups: label_group {
		          position
		          type
		          title
		          url
		          __typename
		        }
	        campaign {
		          discountPercentage: discount_percentage
		          originalPrice: original_price
		          __typename
		        }
	        customvideo_url
	        __typename
	      }
	      shop {
		        shopId: id
		        name
		        domain
		        city
		        tagline
		        uri
		        isOfficial: shop_is_official
		        isPowerBadge: gold_shop
		        badges {
		          title
		          imageURL: image_url
		          show
		          __typename
		        }
	        __typename
	      }
	      tag
	      __typename
	    }
	    header {
		      meta {
		        ab_test
		        templating
		        __typename
		      }
	      __typename
	    }
	    __typename
	  }
	}
	`
)
