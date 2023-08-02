package query

const (
	HeaderMainData = `query headerMainData {
	  dynamicHomeIcon {
	    categoryGroup {
	      id
	      title
	      desc
	      categoryRows {
	        id
	        name
	        url
	        imageUrl
	        type
	        categoryId
	        __typename
	      }
      __typename
    }
    __typename
  }
  categoryAllListLite {
	    categories {
	      id
	      name
	      url
	      iconImageUrl
	      isCrawlable
	      children {
	        id
	        name
	        url
	        isCrawlable
	        children {
	          id
	          name
	          url
	          isCrawlable
	          __typename
	        }
        __typename
      }
      __typename
    }
    __typename
  }
}
`
)
