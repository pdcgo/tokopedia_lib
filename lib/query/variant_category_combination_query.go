package query

const (
	VariantCategoryCombination = `
	query variantCategoryCombination($categoryID: Int!, $type: String!, $productVariants: String, $allVariants: String) {
	  getVariantCategoryCombination(categoryID: $categoryID, type: $type, productVariants: $productVariants, allVariants: $allVariants) {
		header {
		  processTime
		  messages
		  reason
		  errorCode
		  __typename
		}
		data {
		  categoryID
		  variantIDCombinations
		  allVariants {
			VariantID
			Name
			__typename
		  }
		  variantDetails {
			VariantID
			HasUnit
			Identifier
			Name
			Status
			Units {
			  VariantUnitID
			  Status
			  UnitName
			  UnitShortName
			  UnitValues {
				VariantUnitValueID
				Status
				Value
				EquivalentValueID
				EnglishValue
				Hex
				__typename
			  }
			  __typename
			}
			__typename
		  }
		  __typename
		}
		__typename
	  }
	}`
)
