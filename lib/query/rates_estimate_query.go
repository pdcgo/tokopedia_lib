package query

const RatesEstimateQuery = `query ratesEstimateQuery($weight: Float!, $domain: String!, $origin: String, $productId: String, $destination: String, $POTime: Int, $isFulfillment: Boolean, $deviceType: String, $shopTier: Int, $bo_metadata: String, $free_shipping_flag: Int, $warehouse_id: String) {
	  ratesEstimateV3(input: {weight: $weight, domain: $domain, origin: $origin, product_id: $productId, destination: $destination, po_time: $POTime, type: $deviceType, is_fulfillment: $isFulfillment, shop_tier: $shopTier, bo_metadata: $bo_metadata, free_shipping_flag: $free_shipping_flag, warehouse_id: $warehouse_id}) {
	    data {
	      address {
	        city_name
	        province_name
	        district_name
	        country
	        postal_code
	        address
	        lat
	        long
	        phone
	        addr_name
	        address_1
	        receiver_name
	        __typename
	      }
      shop {
	        district_id
	        district_name
	        postal_code
	        origin
	        addr_street
	        latitude
	        longitude
	        province_id
	        city_id
	        city_name
	        __typename
	      }
      rates {
	        id
	        rates_id
	        type
	        services {
	          service_name
	          service_id
	          service_order
	          status
	          range_price {
	            min_price
	            max_price
	            __typename
	          }
          texts {
	            text_service_desc
	            text_service_notes
	            text_range_price
	            text_etd
	            text_price
	            __typename
	          }
          products {
	            shipper_name
	            shipper_id
	            shipper_product_id
	            shipper_product_name
	            shipper_weight
	            price {
	              price
	              formatted_price
	              __typename
	            }
            texts {
	              text_etd
	              text_range_price
	              text_eta_summarize
	              __typename
	            }
            cod {
	              is_cod_available
	              __typename
	            }
            eta {
	              text_eta
	              error_code
	              __typename
	            }
            features {
	              dynamic_price {
	                text_label
	                __typename
	              }
              __typename
            }
            __typename
          }
          service_based_shipment {
	            is_available
	            text_price
	            text_eta
	            __typename
	          }
          cod {
	            is_cod
	            cod_text
	            __typename
	          }
          order_priority {
	            is_now
	            __typename
	          }
          etd {
	            min_etd
	            max_etd
	            __typename
	          }
          range_price {
	            min_price
	            max_price
	            __typename
	          }
          __typename
        }
        __typename
      }
      texts {
	        text_min_price
	        text_destination
	        text_eta
	        __typename
	      }
      free_shipping {
	        flag
	        shipping_price
	        eta_text
	        error_code
	        icon_url
	        title
	        __typename
	      }
      tokocabang_from {
	        title
	        content
	        icon_url
	        __typename
	      }
      is_blackbox
      __typename
    }
    __typename
  }
}
`
