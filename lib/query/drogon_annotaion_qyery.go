package query

const (
	GetDrogonAnnotation = `query getDrogonAnnotation($categoryID: String!, $productID: String, $vendorName: String, $excludeSensitive: String) {
	  drogonAnnotationCategoryV2(categoryID: $categoryID, productID: $productID, vendorName: $vendorName, excludeSensitive: $excludeSensitive) {
		header {
		  processTime
		  messages
		  reason
		  errorCode
		  __typename
		}
		productID
		data {
		  variant
		  sortOrder
		  values {
			id
			name
			selected
			__typename
		  }
		  __typename
		}
		__typename
	  }
	}`
)
