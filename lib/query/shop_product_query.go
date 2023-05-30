package query

const (
	ShopProducts = `query ShopProducts($sid: String!, $page: Int, $perPage: Int, $keyword: String, $etalaseId: String, $sort: Int, $user_districtId: String, $user_cityId: String, $user_lat: String, $user_long: String) {
		GetShopProduct(shopID: $sid, filter: {page: $page, perPage: $perPage, fkeyword: $keyword, fmenu: $etalaseId, sort: $sort, user_districtId: $user_districtId, user_cityId: $user_cityId, user_lat: $user_lat, user_long: $user_long}) {
		  status
		  errors
		  links {
			prev
			next
			__typename
		  }
	  data {
			name
			product_url
			product_id
			price {
			  text_idr
			  __typename
			}
		primary_image {
			  original
			  thumbnail
			  resize300
			  __typename
			}
		flags {
			  isSold
			  isPreorder
			  isWholesale
			  isWishlist
			  __typename
			}
		campaign {
			  discounted_percentage
			  original_price_fmt
			  start_date
			  end_date
			  __typename
			}
		label {
			  color_hex
			  content
			  __typename
			}
		label_groups {
			  position
			  title
			  type
			  url
			  __typename
			}
		badge {
			  title
			  image_url
			  __typename
			}
		stats {
			  reviewCount
			  rating
			  averageRating
			  __typename
			}
		category {
			  id
			  __typename
			}
		__typename
	  }
	  __typename
	}
  }`
)
