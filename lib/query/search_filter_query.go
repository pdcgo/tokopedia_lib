package query

const (
	FilterSortProductQuery = `query FilterSortProductQuery($params: String!) {
		  filter_sort_product(params: $params) {
		    data {
		      filter {
		        title
		        template_name
		        search {
		          searchable
		          placeholder
		          __typename
		        }
	        options {
		          name
		          Description
		          key
		          icon
		          value
		          inputType
		          totalData
		          valMax
		          valMin
		          hexColor
		          child {
		            key
		            value
		            name
		            icon
		            inputType
		            totalData
		            child {
		              key
		              value
		              name
		              icon
		              inputType
		              totalData
		              child {
		                key
		                value
		                name
		                icon
		                inputType
		                totalData
		                __typename
		              }
	              __typename
	            }
	            isPopular
	            __typename
	          }
	          isPopular
	          isNew
	          __typename
	        }
	        __typename
	      }
	      sort {
		        name
		        key
		        value
		        inputType
		        applyFilter
		        __typename
		      }
	      __typename
	    }
	    __typename
	  }
	}
	`
)
