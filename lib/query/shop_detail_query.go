package query

const (
	ShopCoreInfo       = "query ShopInfoCore($id: Int!, $domain: String) {\n  shopInfoByID(input: {shopIDs: [$id], fields: [\"active_product\", \"allow_manage_all\", \"assets\", \"core\", \"closed_info\", \"create_info\", \"favorite\", \"location\", \"status\", \"is_open\", \"other-goldos\", \"shipment\", \"shopstats\", \"shop-snippet\", \"other-shiploc\", \"shopHomeType\", \"branch-link\", \"goapotik\", \"fs_type\"], domain: $domain, source: \"shoppage\"}) {\n    result {\n      shopCore {\n        description\n        domain\n        shopID\n        name\n        tagLine\n        defaultSort\n        __typename\n      }\n      createInfo {\n        openSince\n        __typename\n      }\n      favoriteData {\n        totalFavorite\n        alreadyFavorited\n        __typename\n      }\n      activeProduct\n      shopAssets {\n        avatar\n        cover\n        __typename\n      }\n      location\n      isAllowManage\n      branchLinkDomain\n      isOpen\n      shipmentInfo {\n        isAvailable\n        image\n        name\n        product {\n          isAvailable\n          productName\n          uiHidden\n          __typename\n        }\n        __typename\n      }\n      shippingLoc {\n        districtName\n        cityName\n        __typename\n      }\n      shopStats {\n        productSold\n        totalTxSuccess\n        totalShowcase\n        __typename\n      }\n      statusInfo {\n        shopStatus\n        statusMessage\n        statusTitle\n        tickerType\n        __typename\n      }\n      closedInfo {\n        closedNote\n        until\n        reason\n        detail {\n          status\n          __typename\n        }\n        __typename\n      }\n      bbInfo {\n        bbName\n        bbDesc\n        bbNameEN\n        bbDescEN\n        __typename\n      }\n      goldOS {\n        isGold\n        isGoldBadge\n        isOfficial\n        badge\n        shopTier\n        __typename\n      }\n      shopSnippetURL\n      customSEO {\n        title\n        description\n        bottomContent\n        __typename\n      }\n      isQA\n      isGoApotik\n      partnerInfo {\n        fsType\n        __typename\n      }\n      epharmacyInfo {\n        siaNumber\n        sipaNumber\n        apj\n        __typename\n      }\n      __typename\n    }\n    error {\n      message\n      __typename\n    }\n    __typename\n  }\n}\n"
	ShopStatisticQuery = `query ShopStatisticQuery($shopID: Int!, $shopIDStr: String!) {
		shopSatisfaction: ShopSatisfactionQuery(shopId: $shopID) {
		  recentOneMonth {
			bad
			good
			neutral
			__typename
		  }
	  __typename
	}
	shopRating: productrevGetShopRating(shopID: $shopIDStr) {
		  detail {
			formattedTotalReviews
			rate
			percentage
			percentageFloat
			totalReviews
			__typename
		  }
	  totalRating
	  ratingScore
	  __typename
	}
	shopReputation: reputation_shops(shop_ids: [$shopID]) {
		  badge
		  score
		  score_map
		  __typename
		}
  }`
	GetShopOperationalHourStatus = `query GetShopOperationalHourStatus($shopID: String!) {
		getShopOperationalHourStatus(shopID: $shopID, type: 1) {
		  timestamp
		  statusActive
		  startTime
		  endTime
		  timestampUTCUnix
		  tickerTitle
		  tickerMessage
		  __typename
		}
  }`
	ShopPageGetLayoutV2 = `fragment DisplayWidgetFrag on DisplayWidget {
		imageUrl
		desktopImageUrl
		videoUrl
		appLink
		webLink
		linkType
		linkID
		__typename
	  }
  
  fragment ProductWidgetFrag on ProductWidget {
		productID
		name
		imageUrl
		productUrl
		displayPrice
		originalPrice
		discountPercentage
		isShowFreeOngkir
		freeOngkirPromoIcon
		recommendationType
		isSoldOut
		rating
		totalReview
		isPO
		cashback
		labelGroups {
		  position
		  type
		  title
		  url
		  __typename
		}
	__typename
  }
  
  fragment PromoWidgetFrag on PromoWidget {
		voucherID
		imageUrl
		name
		voucherType {
		  voucherType
		  identifier
		  __typename
		}
	voucherCode
	amount {
		  amountType
		  amount
		  amountFormatted
		  __typename
		}
	minimumSpend
	minimumSpendFormatted
	owner {
		  ownerID
		  identifier
		  __typename
		}
	validThru
	tnc
	inUseExpiry
	status {
		  status
		  identifier
		  __typename
		}
	__typename
  }
  
  fragment EtalaseWidgetFrag on EtalaseWidget {
		imageUrl
		desktopImageUrl
		webLink
		Name
		linkID
		__typename
	  }
  
  fragment CampaignWidgetFrag on CampaignWidget {
		campaignID
		name
		description
		startDate
		endDate
		statusCampaign
		timeDescription
		timeCounter
		totalNotify
		totalNotifyWording
		totalProduct
		totalProductWording
		backgroundGradientColor {
		  firstColor
		  secondColor
		  __typename
		}
	dynamicRule {
		  dynamicRoleData {
			ruleName
			ruleID
			isActive
			ruleAdditionalData
			__typename
		  }
	  descriptionHeader
	  __typename
	}
	banners {
		  imageID
		  imageURL
		  bannerType
		  __typename
		}
	products {
		  id
		  name
		  url
		  urlApps
		  urlMobile
		  imageURL
		  price
		  countSold
		  stock
		  status
		  discountedPrice
		  discountPercentage
		  position
		  rating
		  stockWording {
			title
			__typename
		  }
	  hideGimmick
	  labelGroups {
			position
			type
			title
			url
			__typename
		  }
	  __typename
	}
	__typename
  }
  
  query ShopPageGetLayoutV2($widgetRequest: [ShopPageWidgetRequest!]!, $shopID: String!, $districtID: String, $cityID: String, $latitude: String, $longitude: String) {
		shopPageGetLayout: shopPageGetLayoutV2(widgetRequest: $widgetRequest, shopID: $shopID, districtID: $districtID, cityID: $cityID, latitude: $latitude, longitude: $longitude) {
		  widgets {
			widgetID
			widgetMasterID
			layoutOrder
			name
			type
			header {
			  title
			  subtitle
			  ctaText
			  ctaLink
			  cover
			  ratio
			  isATC
			  isActive
			  etalaseID
			  isShowEtalaseName
			  __typename
			}
		data {
			  ...DisplayWidgetFrag
			  ...ProductWidgetFrag
			  ...PromoWidgetFrag
			  ...EtalaseWidgetFrag
			  ...CampaignWidgetFrag
			  __typename
			}
		__typename
	  }
	  __typename
	}
  }`
	VoucherListQuery = `query VoucherListQuery($shopID: Int!) {
	getPublicMerchantVoucherList(shop_id: $shopID) {
	  vouchers {
		amount {
		  amount
		  amountFormatted: amount_formatted
		  amountType: amount_type
		  __typename
		}
	inUseExpiry: in_use_expiry
	minimumSpend: minimum_spend
	minimumSpendFormatted: minimum_spend_formatted
	owner {
		  identifier
		  __typename
		}
	status {
		  identifier
		  status
		  __typename
		}
	validThru: valid_thru
	voucherID: voucher_id
	name: voucher_name
	voucherCode: voucher_code
	voucherType: voucher_type {
		  identifier
		  voucherType: voucher_type
		  __typename
		}
	banner {
		  desktopUrl: desktop_url
		  __typename
		}
	tnc
	__typename
  }
  __typename
}
}`
	ShopNote = `query ShopNote($sid: String!, $id: String) {
	shopNotesByShopID(shopID: $sid, id: $id) {
	  result {
		title
		content
		link: url
		updateTime
		__typename
	  }
  error {
		message
		__typename
	  }
  __typename
}
}`
)
