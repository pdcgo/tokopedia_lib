
export interface AkunListQuery {
	offset: number;
	limit: number;
	search: string;
}
export interface Pagination {
	offset: number;
	limit: number;
	count: number;
}
export interface AkunItem {
	limit_upload: number;
	count_upload: number;
	active_upload: boolean;
	lastup: number;
	in_upload: boolean;
	last_error: string;
	username: string;
	password: string;
	secret: string;
	markup: string;
	spin: string;
	collection: string;
	hastag: string;
	title_pattern: string;
}
export interface AkunListResponse {
	msg: string;
	error: string;
	data: AkunItem[];
	pagination: Pagination;
}
export interface Response {
	msg: string;
	error: string;
}
export interface BulkItem {
	username: string;
	password: string;
	secret: string;
}
export interface BulkPayload {
	data: BulkItem[];
}

export interface AkunUpdatePayload {
	data: AkunItem[];
}

export interface AkunDeletePayload {
	usernames: string[];
}


export interface UploadAppStatus {
	status: string;
	akun_count: number;
	count_upload: number;
	limit_upload: number;
}
export interface DataSpinQuery {
	name: string;
}
export interface DataSpinItemResponse {
	name: string;
	data: string[];
}
export interface BaseResponse {
	data?: any;
	errcode: number;
	message?: string;
	status?: string;
}
export interface SettingSpinData {
	smin: number;
	smax: number;
	merek_ins_t: boolean;
	title: string;
	desc: string;
}
export interface DataSpinDeleteQuery {
	name: string;
}

export interface SettingSpinTitlePool {
	name: string;
	data: string;
}
export interface SettingSpinResponse {
	data?: SettingSpinData;
	errcode: number;
	titlePool: SettingSpinTitlePool[];
}



export interface SettingSpinConfigUpdatePayload {
	name: string;
	titlePool: SettingSpinTitlePool[];
}
export interface ProductAggQuery {
	marketplace: string;
	namespace: string;
}
export interface ProductNamespaceAgg {
	count: number;
	price_min: number;
	price_max: number;
	name: string;
}
export interface ProductPriceRangeAggQuery {
	marketplace: string;
	namespace: string;
	range_price: number;
}
export interface ProductPriceRangeAgg {
	_id: number[];
	count: number;
}
export interface ProductCategoryAggQuery {
	marketplace: string;
	namespace: string;
	is_public: boolean;
}
export interface ProductCategoryAgg {
	_id: number;
	price_min: number;
	price_max: number;
	count: number;
	name: any[];
}

export interface ProductCityAgg {
	_id: string;
	price_min: number;
	price_max: number;
	count: number;
}

export interface MapItem {
	shopee_id: number;
	tokopedia_id: number;
}
export interface GetMapQuery {
	collection: string;
}
export interface ShopeeTopedMapResponse {
	data: MapItem[];
}
export interface Category {
	id: number;
	name: string;
	url: string;
	children: Category[];
	__typename: string;
}
export interface CategoryAllListLite {
	categories: Category[];
	__typename: string;
}
export interface CategoryAllListLiteData {
	categoryAllListLite?: CategoryAllListLite;
}
export interface CategoryAllListLiteRes {
	data: CategoryAllListLiteData;
}

export interface UpdateTopedCategoryPayload {
	username: string;
	password: string;
	secret: string;
}
export type SdkConfig = {

	GetTokopediaAkunList: {
		method: "get"
		params: AkunListQuery
		payload: undefined
		response: AkunListResponse
		path: "tokopedia/akun/list"
	},

	PostTokopediaAkunBulkAdd: {
		method: "post"
		params: undefined
		payload: BulkPayload
		response: Response
		path: "tokopedia/akun/bulk_add"
	},

	PostTokopediaAkunUpdate: {
		method: "post"
		params: undefined
		payload: AkunUpdatePayload
		response: Response
		path: "tokopedia/akun/update"
	},

	PostTokopediaAkunDelete: {
		method: "post"
		params: undefined
		payload: AkunDeletePayload
		response: Response
		path: "tokopedia/akun/delete"
	},

	GetTokopediaUploadStart: {
		method: "get"
		params: undefined
		payload: undefined
		response: Response
		path: "tokopedia/upload/start"
	},

	GetTokopediaUploadStop: {
		method: "get"
		params: undefined
		payload: undefined
		response: Response
		path: "tokopedia/upload/stop"
	},

	GetTokopediaUploadStatus: {
		method: "get"
		params: undefined
		payload: undefined
		response: UploadAppStatus
		path: "tokopedia/upload/status"
	},

	GetApiDataspin: {
		method: "get"
		params: DataSpinQuery
		payload: undefined
		response: DataSpinItemResponse[]
		path: "api/dataspin"
	},

	PostApiDataspin: {
		method: "post"
		params: undefined
		payload: SettingSpinData
		response: BaseResponse
		path: "api/dataspin"
	},

	DeleteApiDataspin: {
		method: "delete"
		params: DataSpinDeleteQuery
		payload: undefined
		response: BaseResponse
		path: "api/dataspin"
	},

	GetApiSettingSpin: {
		method: "get"
		params: undefined
		payload: undefined
		response: SettingSpinResponse
		path: "api/settingSpin"
	},

	PostApiSettingSpin: {
		method: "post"
		params: undefined
		payload: SettingSpinData
		response: BaseResponse
		path: "api/settingSpin"
	},

	PostApiConfigSettingSpin: {
		method: "post"
		params: undefined
		payload: SettingSpinConfigUpdatePayload
		response: BaseResponse
		path: "api/config/settingSpin"
	},

	GetApiListMarkup: {
		method: "get"
		params: undefined
		payload: undefined
		response: { data: string[] }
		path: "api/listMarkup"
	},

	GetV1ProductNamespaceAll: {
		method: "get"
		params: ProductAggQuery
		payload: undefined
		response: ProductNamespaceAgg[]
		path: "v1/product/namespace_all"
	},

	GetV1ProductPriceRange: {
		method: "get"
		params: ProductPriceRangeAggQuery
		payload: undefined
		response: ProductPriceRangeAgg[]
		path: "v1/product/price_range"
	},

	GetV1ProductCategory: {
		method: "get"
		params: ProductCategoryAggQuery
		payload: undefined
		response: ProductCategoryAgg[]
		path: "v1/product/category"
	},

	GetV1ProductKota: {
		method: "get"
		params: ProductAggQuery
		payload: undefined
		response: ProductCityAgg[]
		path: "v1/product/kota"
	},

	PutTokopediaMapperMap: {
		method: "put"
		params: undefined
		payload: MapItem[]
		response: Response
		path: "tokopedia/mapper/map"
	},

	GetTokopediaMapperMap: {
		method: "get"
		params: GetMapQuery
		payload: undefined
		response: ShopeeTopedMapResponse
		path: "tokopedia/mapper/map"
	},

	PutTokopediaMapperAutosuggest: {
		method: "put"
		params: undefined
		payload: undefined
		response: undefined
		path: "tokopedia/mapper/autosuggest"
	},

	GetTokopediaCategoryList: {
		method: "get"
		params: undefined
		payload: undefined
		response: CategoryAllListLiteRes
		path: "tokopedia/category/list"
	},

	PutTokopediaCategoryUpdateCategory: {
		method: "put"
		params: undefined
		payload: UpdateTopedCategoryPayload
		response: Response
		path: "tokopedia/category/update_category"
	}
}
