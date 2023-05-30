package query

const (
	ShopCoreInfo = `query ShopInfoCore($id: Int!, $domain: String) {
		  shopInfoByID(input: {shopIDs: [$id], fields: ["active_product", "allow_manage_all", "assets", "core", "closed_info", "create_info", "favorite", "location", "status", "is_open", "other-goldos", "shipment", "shopstats", "shop-snippet", "other-shiploc", "shopHomeType", "branch-link", "goapotik", "fs_type"], domain: $domain, source: "shoppage"}) {
		    result {
		      shopCore {
		        description
		        domain
		        shopID
		        name
		        tagLine
		        defaultSort
		        __typename
		      }
	      createInfo {
		        openSince
		        __typename
		      }
	      favoriteData {
		        totalFavorite
		        alreadyFavorited
		        __typename
		      }
	      activeProduct
	      shopAssets {
		        avatar
		        cover
		        __typename
		      }
	      location
	      isAllowManage
	      branchLinkDomain
	      isOpen
	      shipmentInfo {
		        isAvailable
		        image
		        name
		        product {
		          isAvailable
		          productName
		          uiHidden
		          __typename
		        }
	        __typename
	      }
	      shippingLoc {
		        districtName
		        cityName
		        __typename
		      }
	      shopStats {
		        productSold
		        totalTxSuccess
		        totalShowcase
		        __typename
		      }
	      statusInfo {
		        shopStatus
		        statusMessage
		        statusTitle
		        tickerType
		        __typename
		      }
	      closedInfo {
		        closedNote
		        until
		        reason
		        detail {
		          status
		          __typename
		        }
	        __typename
	      }
	      bbInfo {
		        bbName
		        bbDesc
		        bbNameEN
		        bbDescEN
		        __typename
		      }
	      goldOS {
		        isGold
		        isGoldBadge
		        isOfficial
		        badge
		        shopTier
		        __typename
		      }
	      shopSnippetURL
	      customSEO {
		        title
		        description
		        bottomContent
		        __typename
		      }
	      isQA
	      isGoApotik
	      partnerInfo {
		        fsType
		        __typename
		      }
	      epharmacyInfo {
		        siaNumber
		        sipaNumber
		        apj
		        __typename
		      }
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
