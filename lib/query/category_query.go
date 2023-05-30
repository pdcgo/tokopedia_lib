package query

const (
	CategoryAllListLite = `query categoryAllListLite {
		categoryAllListLite(filter: "seller") {
		  categories {
			id
			name
			url
			children {
			  id
			  name
			  url
			  children {
				id
				name
				url
				__typename
			  }
		  __typename
		}
		__typename
	  }
	  __typename
	}
  }`
	JarvisRecommendation = `query jarvisRecommendation($productName: String) {
		getJarvisRecommendation(product_name: $productName) {
		  categories {
			id
			name
			__typename
		  }
	  __typename
	}
  }`
)
