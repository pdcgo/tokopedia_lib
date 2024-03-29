package query

const (
	SearchProductQuery = `query SearchProductQuery($params: String, $adParams: String) {
          CategoryProducts: searchProduct(params: $params) {
            count
            data: products {
              id
              url
              imageUrl: image_url
              imageUrlLarge: image_url_700
              catId: category_id
              gaKey: ga_key
              countReview: count_review
              discountPercentage: discount_percentage
              preorder: is_preorder
              name
              price
              priceInt: price_int
              original_price
              rating
              wishlist
              labels {
                title
                color
                __typename
              }
          badges {
                imageUrl: image_url
                show
                __typename
              }
          shop {
                id
                url
                name
                goldmerchant: is_power_badge
                official: is_official
                reputation
                clover
                location
                __typename
              }
          labelGroups: label_groups {
                position
                title
                type
                __typename
              }
          __typename
        }
        __typename
      }
      displayAdsV3(displayParams: $adParams) {
            data {
              id
              ad_ref_key
              redirect
              sticker_id
              sticker_image
              productWishListUrl: product_wishlist_url
              clickTrackUrl: product_click_url
              shop_click_url
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
            campaign {
                  original_price
                  discountPercentage: discount_percentage
                  __typename
                }
            wholeSalePrice: wholesale_price {
                  quantityMin: quantity_min_format
                  quantityMax: quantity_max_format
                  price: price_format
                  __typename
                }
            count_talk_format
            countReview: count_review_format
            category {
                  id
                  __typename
                }
            preorder: product_preorder
            product_wholesale
            free_return
            isNewProduct: product_new_label
            cashback: product_cashback_rate
            rating: product_rating
            top_label
            bottomLabel: bottom_label
            __typename
          }
          shop {
                image_product {
                  image_url
                  __typename
                }
            id
            name
            domain
            location
            city
            tagline
            goldmerchant: gold_shop
            gold_shop_badge
            official: shop_is_official
            lucky_shop
            uri
            owner_id
            is_owner
            badges {
                  title
                  image_url
                  show
                  __typename
                }
            __typename
          }
          applinks
          __typename
        }
        template {
              isAd: is_ad
              __typename
            }
        __typename
      }
    }
    `
)
