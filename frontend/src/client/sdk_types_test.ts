/* eslint-disable @typescript-eslint/no-explicit-any */


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
export interface DataSpinItemResponse {
    name: string;
    data: string[];
}
export interface DataSpinQuery {
    name: string;
}

export interface BaseResponse {
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
export interface ListMarkupResponse {
    errcode: number;
    message?: string;
    status?: string;
    data: string[];
}

export interface MarkupData {
    mark: string;
    type: string;
    range: any;
    up: number[];
}
export interface Markup {
    data: MarkupData[];
    fix_mark: number;
    name: string;
}
export interface MarkupGetQuery {
    name: string;
}
export interface MarkupGetResponse {
    data: MarkupData[];
    fix_mark: number;
    name: string;
    errorcode: number;
}




export interface ProductNamespaceAgg {
    count: number;
    price_min: number;
    price_max: number;
    name: string;
}
export interface ProductPriceRangeAgg {
    _id: number[];
    count: number;
}
export interface ProductCategoryAgg {
    _id: number;
    price_min: number;
    price_max: number;
    count: number;
    name: string[];
}
export interface ProductCityAgg {
    _id: string;
    price_min: number;
    price_max: number;
    count: number;
}
export interface ShopeeCategoryOld {
    parent_category: number;
    catid: number;
    parent_display_name: string;
    display_name: string;
    status: string;
}
export interface ProductMatchStageQuery {
    is_public: boolean;
    kota: string;
    marketplace: string;
    namespace: string;
    pmax: number;
    pmin: number;
}

export interface ProductPriceRangeAggQuery {
    is_public: boolean;
    kota: string;
    marketplace: string;
    namespace: string;
    pmax: number;
    pmin: number;
    range_price: number;
}













export interface PredictWeightResponse {
    harga: number;
    jarak: number;
}
export interface PredictWeightPayload {
    itemid: number;
    shopid: number;
}
export interface PredictWeightSaveQuery {
    predict: number;
}

export interface PredictWeightLoadResponse {
    predict_weight: number;
}
export interface SearchFilterDynamicShipping {
    positionid: number;
    name: string;
    channelids: string[];
    display_name: string;
    item_tag_ids: number[];
}

export interface CategorySubSub {
    catid: number;
    display_name: string;
    image: string;
    block_buyer_platform: any;
}
export interface CategorySub {
    catid: number;
    display_name: string;
    name: string;
    image: string;
    is_adult: any;
    parent_category: number;
    sort_weight: number;
    block_buyer_platform: any;
    sub_sub: CategorySubSub[];
}
export interface CategoryMain {
    catid: number;
    display_name: string;
    name: string;
    image: string;
    is_adult: any;
    parent_category: number;
    sort_weight: number;
    block_buyer_platform: any;
}
export interface CategoryItem {
    main: CategoryMain;
    sub: CategorySub[];
}
export interface ShopeeCategoryRegionSetting {
    enable_size_chart: boolean;
    low_stock_value: number;
    dimension_mandatory: boolean;
}
export interface ShopeeCategory {
    id: number;
    name: string;
    display_name: string;
    parent_id: number;
    has_active_children: boolean;
    has_children: boolean;
    children?: any[];
    region_setting: ShopeeCategoryRegionSetting;
    is_prohibit: boolean;
    chain_name: string[];
    chain_ids: number[];
}
export interface ManifestResponse {
    category: ShopeeCategory[];
    public_category_repo: CategoryItem[];
}


export interface ShopeeMapItem {
    shopee_id: number;
    tokopedia_id: number;
}
export interface GetMapQuery {
    collection: string;
}
export interface ShopeeTopedMapResponse {
    data: ShopeeMapItem[];
}

export interface AutoSuggestStatus {
    status: string;
}
export interface ShopeeMapperConfig {
    use_mapper: boolean;
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
export interface ViewConfig {
    min: number;
    max: number;
}
export interface SoldConfig {
    min: number;
    max: number;
}
export interface AkunDeleteItem {
    username: string;
    password: string;
    secret: string;
}
export interface DeleteConfig {
    limit_concurent: number;
    limit_product: number;
    title: string[];
    product_status: string;
    category_id: string;
    start_time: number;
    end_time: number;
    akuns: AkunDeleteItem[];
    sold_filter?: SoldConfig;
    view_filter?: ViewConfig;
}
export interface DeleteSettingRes {
    data?: DeleteConfig;
}



export interface DriverAccount {
    username: string;
    password: string;
    secret: string;
}
export interface RunCheckbotPayload {
    fname: string;
    Akuns: DriverAccount[];
}
export interface VerifDriverAccount {
    username: string;
    password: string;
    secret: string;
    Pesan: string;
    Status: string;
}
export interface RunCheckVerifPayload {
    fname: string;
    Akuns: VerifDriverAccount[];
}
export interface AutoSubmit {
    base_ktp: string;
    filename: string;
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

	PutTokopediaAkunResetAllCount: {
		method: "put"
		params: undefined
		payload: undefined
		response: Response
		path: "tokopedia/akun/reset_all_count"
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
		response: ListMarkupResponse
		path: "api/listMarkup"
	},

	PostApiAddMarkup: {
		method: "post"
		params: undefined
		payload: Markup
		response: BaseResponse
		path: "api/addMarkup"
	},

	GetApiMarkup: {
		method: "get"
		params: MarkupGetQuery
		payload: undefined
		response: MarkupGetResponse
		path: "api/markup"
	},

	PostApiMarkup: {
		method: "post"
		params: MarkupGetQuery
		payload: Markup
		response: BaseResponse
		path: "api/markup"
	},

	PostApiDeleteMarkup: {
		method: "post"
		params: undefined
		payload: string[]
		response: BaseResponse
		path: "api/deleteMarkup"
	},

	GetV1ProductNamespaceAll: {
		method: "get"
		params: ProductMatchStageQuery
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
		params: ProductMatchStageQuery
		payload: undefined
		response: ProductCategoryAgg[]
		path: "v1/product/category"
	},

	GetV1ProductKota: {
		method: "get"
		params: ProductMatchStageQuery
		payload: undefined
		response: ProductCityAgg[]
		path: "v1/product/kota"
	},

	GetV1ProductDelete: {
		method: "get"
		params: ProductMatchStageQuery
		payload: undefined
		response: BaseResponse
		path: "v1/product/delete"
	},

	PostApiDeleteItem: {
		method: "post"
		params: ProductMatchStageQuery
		payload: number[]
		response: BaseResponse
		path: "api/deleteItem"
	},

	PostV1ProductCategstatToCsv: {
		method: "post"
		params: undefined
		payload: ShopeeCategoryOld[]
		response: BaseResponse
		path: "v1/product/categstat_to_csv"
	},

	GetV1ProductResync: {
		method: "get"
		params: ProductMatchStageQuery
		payload: undefined
		response: BaseResponse
		path: "v1/product/resync"
	},

	PostV4ShopeeWeightPredict: {
		method: "post"
		params: undefined
		payload: PredictWeightPayload
		response: PredictWeightResponse
		path: "v4/shopee/weight/predict"
	},

	GetV3PredictweightSave: {
		method: "get"
		params: PredictWeightSaveQuery
		payload: undefined
		response: BaseResponse
		path: "v3/predictweight/save"
	},

	GetV3PredictweightLoad: {
		method: "get"
		params: undefined
		payload: undefined
		response: PredictWeightLoadResponse
		path: "v3/predictweight/load"
	},

	GetApiShopeeShipping: {
		method: "get"
		params: undefined
		payload: undefined
		response: SearchFilterDynamicShipping[]
		path: "api/shopee_shipping"
	},

	GetShopeeManifest: {
		method: "get"
		params: undefined
		payload: undefined
		response: ManifestResponse
		path: "shopee/manifest"
	},

	GetApiUpdateTokpedCategories: {
		method: "get"
		params: undefined
		payload: undefined
		response: BaseResponse
		path: "api/updateTokpedCategories"
	},

	PutTokopediaMapperMap: {
		method: "put"
		params: undefined
		payload: ShopeeMapItem[]
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
		params: GetMapQuery
		payload: undefined
		response: undefined
		path: "tokopedia/mapper/autosuggest"
	},

	GetTokopediaMapperAutosuggest: {
		method: "get"
		params: undefined
		payload: undefined
		response: AutoSuggestStatus
		path: "tokopedia/mapper/autosuggest"
	},

	GetTokopediaMapperSetting: {
		method: "get"
		params: undefined
		payload: undefined
		response: ShopeeMapperConfig
		path: "tokopedia/mapper/setting"
	},

	PutTokopediaMapperSetting: {
		method: "put"
		params: undefined
		payload: undefined
		response: ShopeeMapperConfig
		path: "tokopedia/mapper/setting"
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
	},

	GetTokopediaDeleterSetting: {
		method: "get"
		params: undefined
		payload: undefined
		response: DeleteSettingRes
		path: "tokopedia/deleter/setting"
	},

	PutTokopediaDeleterSetting: {
		method: "put"
		params: undefined
		payload: DeleteConfig
		response: DeleteSettingRes
		path: "tokopedia/deleter/setting"
	},

	PutTokopediaDeleterRunDelete: {
		method: "put"
		params: undefined
		payload: undefined
		response: Response
		path: "tokopedia/deleter/run_delete"
	},

	PutTokopediaCekbotRun: {
		method: "put"
		params: undefined
		payload: RunCheckbotPayload
		response: undefined
		path: "tokopedia/cekbot/run"
	},

	PutTokopediaCheckVerifRun: {
		method: "put"
		params: undefined
		payload: RunCheckVerifPayload
		response: undefined
		path: "tokopedia/check_verif/run"
	},

	PostTokopediaAutosubmitRun: {
		method: "post"
		params: undefined
		payload: AutoSubmit
		response: undefined
		path: "tokopedia/autosubmit/run"
	}
}
