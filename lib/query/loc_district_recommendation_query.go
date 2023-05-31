package query

const (
	LocDisctricRecommendation = `
	query keroDistrictRecommendation($query: String, $page: String) {
	  kero_district_recommendation(query: $query, page: $page) {
		next_available
		district {
		  districtId: district_id
		  district_name
		  cityId: city_id
		  city_name
		  provinceId: province_id
		  province_name
		  zip_code
		  __typename
		}
		__typename
	  }
	}`
)
