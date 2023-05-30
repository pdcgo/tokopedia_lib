package query

const (
	PdpGetDataP2 = `query PDPGetDataP2($productID: String!, $pdpSession: String!, $deviceID: String, $userLocation: pdpUserLocation, $affiliate: pdpAffiliate) {
		  pdpGetData(productID: $productID, pdpSession: $pdpSession, deviceID: $deviceID, userLocation: $userLocation, affiliate: $affiliate) {
		    error {      Code
		      Message
		      DevMessage
		      __typename
		    }
	    callsError {
		      shopInfo {
		        Code
		        Message
		        __typename
		      }
	      cartRedirection {
		        Code
		        Message
		        __typename
		      }
	      nearestWarehouse {
		        Code
		        Message
		        __typename
		      }
	      __typename
	    }
	    productView
	    wishlistCount
	    shopFinishRate {
		      finishRate
		      __typename
		    }
	    shopInfo {
		      shopTier
		      badgeURL
		      closedInfo {
		        closedNote
		        reason
		        detail {
		          openDate
		          __typename
		        }
	        __typename
	      }
	      isOpen
	      favoriteData {
		        totalFavorite
		        alreadyFavorited
		        __typename
		      }
	      activeProduct
	      createInfo {
		        epochShopCreated
		        __typename
		      }
	      shopAssets {
		        avatar
		        __typename
		      }
	      shopCore {
		        domain
		        shopID
		        name
		        shopScore
		        url
		        ownerID
		        __typename
		      }
	      shopLastActive
	      location
	      statusInfo {
		        statusMessage
		        shopStatus
		        isIdle
		        __typename
		      }
	      isAllowManage
	      isOwner
	      ownerInfo {
		        id
		        __typename
		      }
	      isCOD
	      shopType
	      tickerData {
		        title
		        message
		        color
		        link
		        action
		        actionLink
		        tickerType
		        actionBottomSheet {
		          title
		          message
		          reason
		          buttonText
		          buttonLink
		          __typename
		        }
	        __typename
	      }
	      partnerLabel
	      __typename
	    }
	    merchantVoucher {
		      vouchers {
		        voucher_id
		        voucher_name
		        voucher_type {
		          voucher_type
		          identifier
		          __typename
		        }
	        voucher_code
	        amount {
		          amount
		          amount_type
		          amount_formatted
		          __typename
		        }
	        minimum_spend
	        valid_thru
	        tnc
	        banner {
		          desktop_url
		          mobile_url
		          __typename
		        }
	        status {
		          status
		          identifier
		          __typename
		        }
	        in_use_expiry
	        __typename
	      }
	      __typename
	    }
	    nearestWarehouse {
		      product_id
		      stock
		      stock_wording
		      price
		      warehouse_info {
		        warehouse_id
		        is_fulfillment
		        district_id
		        postal_code
		        geolocation
		        __typename
		      }
	      __typename
	    }
	    installmentRecommendation {
		      data {
		        term
		        mdr_value
		        mdr_type
		        interest_rate
		        minimum_amount
		        maximum_amount
		        monthly_price
		        os_monthly_price
		        partner_code
		        partner_name
		        partner_icon
		        subtitle
		        __typename
		      }
	      __typename
	    }
	    productWishlistQuery {
		      value
		      __typename
		    }
	    cartRedirection {
		      status
		      error_message
		      data {
		        product_id
		        config_name
		        hide_floating_button
		        available_buttons {
		          text
		          color
		          cart_type
		          onboarding_message
		          show_recommendation
		          __typename
		        }
	        unavailable_buttons
	        __typename
	      }
	      __typename
	    }
	    shopTopChatSpeed {
		      messageResponseTime
		      __typename
		    }
	    shopRatingsQuery {
		      ratingScore
		      __typename
		    }
	    shopPackSpeed {
		      speedFmt
		      hour
		      __typename
		    }
	    ratesEstimate {
		      warehouseID
		      products
		      data {
		        destination
		        title
		        subtitle
		        chipsLabel
		        courierLabel
		        eTAText
		        cheapestShippingPrice
		        fulfillmentData {
		          icon
		          prefix
		          description
		          __typename
		        }
	        errors {
		          code: Code
		          message: Message
		          devMessage: DevMessage
		          __typename
		        }
	        __typename
	      }
	      bottomsheet {
		        title
		        iconURL
		        subtitle
		        buttonCopy
		        __typename
		      }
	      __typename
	    }
	    restrictionInfo {
		      message
		      restrictionData {
		        productID
		        isEligible
		        action {
		          actionType
		          title
		          description
		          attributeName
		          badgeURL
		          buttonText
		          buttonLink
		          __typename
		        }
	        __typename
	      }
	      __typename
	    }
	    ticker {
		      tickerInfo {
		        productIDs
		        tickerData {
		          title
		          message
		          color
		          link
		          action
		          actionLink
		          tickerType
		          actionBottomSheet {
		            title
		            message
		            reason
		            buttonText
		            buttonLink
		            __typename
		          }
	          __typename
	        }
	        __typename
	      }
	      __typename
	    }
	    navBar {
		      name
		      items {
		        componentName
		        title
		        __typename
		      }
	      __typename
	    }
	    bebasOngkir {
		      products {
		        productID
		        boType
		        __typename
		      }
	      __typename
	    }
	    __typename
	  }
	}
	`
)
