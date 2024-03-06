/* eslint-disable @typescript-eslint/no-explicit-any*/

import axios from "axios";
import { useState } from "react";

export class Uint extends Number {}
export class Int extends Number {}
export class Float32 extends Number {}
export class Float64 extends Number {}

export type MaybeUndefined<T> = T | undefined;
export type MaybeNull<T> = T | null;

export interface GrabShopeeQueryCli {
	base: string
}

export interface GrabShopeeResponse {
	use_csv: boolean
	empty_csv: boolean
}

export interface GrabTokopediaQueryCli {
	base: string
}

export interface GrabTokopediaResponse {
	deprecated: boolean
	use_csv: boolean
	empty_csv: boolean
}

export interface GrabJakmallQueryCli {
	base: string
}

export interface GrabJakmallResponse {
	use_csv: boolean
	empty_csv: boolean
}

export interface DataSpinQuery {
	name: string
}

export interface DataSpinItemResponse {
	name: string
	data: Array<string>
}

export interface BaseResponse {
	errcode: number
	message: string
	status: string
}

export interface DataSpinDeleteQuery {
	name: string
}

export interface SettingSpinData {
	smin: number
	smax: number
	merek_ins_t: boolean
	ignore_first_word: boolean
	title: string
	desc: string
}

export interface SettingSpinTitlePool {
	name: string
	data: string
}

export interface SettingSpinResponse {
	data: SettingSpinData | undefined
	errcode: number
	titlePool: Array<SettingSpinTitlePool>
}

export interface SettingSpinConfigUpdatePayload {
	name: string
	titlePool: Array<SettingSpinTitlePool>
}

export interface ProductMatchStageQuery {
	is_public: boolean
	kota: string
	marketplace: string
	namespace: string
	pmax: number
	pmin: number
	use_empty_city: boolean
}

export interface ProductNamespaceAgg {
	count: number
	price_min: number
	price_max: number
	name: string
}

export interface ProductPriceRangeAggQuery {
	is_public: boolean
	kota: string
	marketplace: string
	namespace: string
	pmax: number
	pmin: number
	use_empty_city: boolean
	range_price: number
}

export interface ProductPriceRangeAgg {
	_id: Array<number>
	count: number
}

export interface ProductCategoryAgg {
	_id: number
	price_min: number
	price_max: number
	count: number
	name: Array<string>
}

export interface ProductCityAgg {
	_id: string
	price_min: number
	price_max: number
	count: number
}

export interface ShopeeCategoryOld {
	parent_category: number
	catid: number
	parent_display_name: string
	display_name: string
	status: string
}

export interface ProductNamespaceRenamePayload {
	marketplace: string
	namespace: string
	update_namespace: string
}

export interface ListMarkupResponse {
	errcode: number
	message: string
	status: string
	data: Array<string>
}

export interface MarkupData {
	mark: string
	type: string
	range: any
	up: Array<number>
}

export interface Markup {
	data: Array<MarkupData>
	fix_mark: number
	name: string
}

export interface MarkupGetQuery {
	name: string
}

export interface MarkupGetResponse {
	data: Array<MarkupData>
	fix_mark: number
	name: string
	errorcode: number
}

export interface SearchFilterDynamicShipping {
	positionid: number
	name: string
	channelids: Array<string>
	display_name: string
	item_tag_ids: Array<number>
}

export interface RegionSetting {
	enable_size_chart: boolean
	low_stock_value: number
	dimension_mandatory: boolean
}

export interface CategoryTreeItem {
	id: number
	name: string
	display_name: string
	parent_id: number
	has_active_children: boolean
	has_children: boolean
	region_setting: RegionSetting | undefined
	is_prohibit: boolean
	children: Array<CategoryTreeItem | undefined>
}

export interface ShopeeCategory {
	id: number
	name: string
	display_name: string
	parent_id: number
	has_active_children: boolean
	has_children: boolean
	region_setting: RegionSetting | undefined
	is_prohibit: boolean
	children: Array<CategoryTreeItem | undefined>
	chain_name: Array<string>
	chain_ids: Array<number>
}

export interface CategoryMain {
	catid: number
	display_name: string
	name: string
	image: string
	is_adult: any
	parent_category: number
	sort_weight: number
	block_buyer_platform: any
}

export interface CategorySubSub {
	catid: number
	display_name: string
	image: string
	block_buyer_platform: any
}

export interface CategorySub {
	catid: number
	display_name: string
	name: string
	image: string
	is_adult: any
	parent_category: number
	sort_weight: number
	block_buyer_platform: any
	sub_sub: Array<CategorySubSub>
}

export interface CategoryItem {
	main: CategoryMain
	sub: Array<CategorySub>
}

export interface ManifestResponse {
	category: Array<ShopeeCategory>
	public_category_repo: Array<CategoryItem>
}

export interface PredictWeightPayload {
	itemid: number
	shopid: number
}

export interface PredictWeightResponse {
	harga: number
	jarak: number
}

export interface PredictWeightSaveQuery {
	predict: number
}

export interface PredictWeightLoadResponse {
	predict_weight: number
}

export interface SettingCropResponse {
	name: string
	data: null
	errcode: number
}

export interface SettingCrop {
	name: string
	data: null
}

export interface SettingIntResponse {
	name: string
	data: number
}

export interface SettingUploadIntervalData {
	uptmax: number
	uptmin: number
}

export interface SettingUploadIntervalResponse {
	name: string
	data: SettingUploadIntervalData
}

export interface SettingStrResponse {
	name: string
	data: string
}

export interface SettingGrabData {
	penjualan: number
	prosentase: number
	stock: number
	tokped_point: Array<number>
}

export interface SettingGrabResponse {
	name: string
	data: SettingGrabData
}

export interface SettingLastActiveData {
	active: boolean
	days: number
}

export interface SettingLastActiveResponse {
	name: string
	data: SettingLastActiveData
}

export interface GrabShopeeQuery {
	by: string
	locations: Array<string>
	official_mall: boolean
	price_max: number
	price_min: number
	rating_filter: number
	shopee24: boolean
	shopee_verified: boolean
}

export interface SettingGrabFilterShopeeResponse {
	name: string
	data: GrabShopeeQuery
}

export interface GrabShopeeProductCreated {
	active: boolean
	max: number
	min: number
}

export interface GrabShopeeShipping {
	channelids: Array<string>
	display_name: string
	name: string
	item_tag_ids: Array<number>
	positionid: number
}

export interface SettingGrabFilterShopeeExtraResponse {
	product_created: GrabShopeeProductCreated
	shippings: Array<GrabShopeeShipping>
}

export interface CrawlerConfig {
	username: string
	password: string
	email: string
	email_password: string
}

export interface GrabTokopediaQuery {
	pmin: number
	pmax: number
	ob: string
	rt: string
	condition: string
	fcity: Array<string>
	goldmerchant: boolean
	official: boolean
	shipping: Array<string>
	preorder: boolean
}

export interface SettingGrabFilterTokopedia {
	name: string
	data: GrabTokopediaQuery
}

export interface SettingGrabFilterTokopediaResponse {
	errcode: number
	data: SettingGrabFilterTokopedia
}

export interface Cloudinary {
	active: boolean
	url: string
}

export interface SettingCloudinaryResponse {
	name: string
	data: Cloudinary
}

export interface SettingBoolResponse {
	name: string
	data: boolean
}

export interface RandomAttribute {
	active: boolean
	force_tidakada: boolean
}

export interface SettingRandomAttributeResponse {
	name: string
	data: RandomAttribute
}

export interface SettingTextFilterResponse {
	data: Array<string>
	errcode: number
	name: string
}

export interface SettingTextFilterPayload {
	data: Array<string>
	name: string
}

export interface GrabBasicBlacklistData {
	data: Array<string>
	filename: string
}

export interface GrabBasicBlacklist {
	active: boolean
	shopee: GrabBasicBlacklistData
}

export interface SettingDefaultResponse {
	_id: string
	name: string
	use_price_discount: boolean
	blacklist_username: GrabBasicBlacklist
}

export interface SettingFilterWordQuery {
	name: string
	makedefault: boolean
}

export interface ReportQuery {
	output: string
}

export interface Akun {
	_id: string
	user: string
	name: string
	pass: string
	password: string
	active: boolean
	limit_upload: number
	count_upload: number
	namespace: string
	water: string
	shopee_categ: Array<string>
	tokped_categ: Array<string>
	type: string
	markup: string
	hastag: string
	last_up: number
	polatitle: string
	is_running_upload: boolean
}

export interface AkunFormPayload {
	data: Akun
}

export interface AkunDeletePayload {
	data: Array<string>
}

export interface AkunListQuery {
	start: number
	limit: number
	active: string
	last: number
	search: string
	sort: string
	reverse: number
}

export interface AkunListResponse {
	akuns: Array<Akun>
	count: number
	errcode: number
}

export interface UploadShipping {
	channelid: number
	enabled: boolean
}

export interface UploadConfig {
	shipping: Array<UploadShipping>
	use_custom_shipping: boolean
}

export interface UploadProductShippingLimits {
	item_min_weight: number
}

export interface UploadProductShipping {
	channel_id: number
	name: string
	display_name: string
	name_key: string
	flag: string
	service_type: string
	default_price: string
	min_default_price: string
	max_default_price: string
	limits: UploadProductShippingLimits
	sizes: Array<any>
	cover_shipping_fee: boolean
	enabled: boolean
	item_flag: string
	price: string
	size: number
	size_id: number
	save_into_item: boolean
	volumetric_factor: number
	conflicting_enabled_channels: any
	is_mask_channel: boolean
	parent_channel_id: number
	item_exclusive_channels: Array<number>
	related_enabled_channels: Array<any>
	related_disabled_channels: Array<any>
}

export interface ConfigAccountPayload {
	data: string
}

export interface BotConfigCaptcha {
	user: string
	pass: string
}

export interface BotConfigDatabase {
	DB_URI: string
	DB_NAME: string
}

export interface BotConfigLisensi {
	email: string
	pwd: string
}

export interface BotConfig {
	captcha: BotConfigCaptcha
	database: BotConfigDatabase
	lisensi: BotConfigLisensi
	suplier_csv: boolean
	session_storage_source: string
	session_storage_local_location: string
	thread_count: number
	tor_directory: string
	tor_print_ip: boolean
	tor_proxy: boolean
}

export interface DeleteConfig {
	akun: string
	awaltanggal: string
	blokir: boolean
	delete: number
	diarsipkan: boolean
	diperiksa: boolean
	sold: number
	tanggal: string
	view: number
}

export interface DeleteProductHarga {
	min: number
	max: number
}

export interface DeleteProduct {
	fil_category: boolean
	fil_harga: boolean
	fil_keyword: boolean
	category: Array<Array<number>>
	harga: DeleteProductHarga
	keyword: string
}

export interface ConfigDeleteExtraResponse {
	name: string
	data: DeleteProduct
}

export interface ShopeeCateg {
	parent_category: number
	catid: number
	parent_display_name: string
	display_name: string
	is_collection: number
}

export interface GrabTasker {
	_id: string
	toko_username: string
	mode: string
	marketplace: string
	product_url: string
	namespace: string
	tokped_categ: Array<string>
	use_filter: boolean
	keyword: string
	shopee_categ: ShopeeCateg
	jakmall_categs: Array<string>
}

export interface DumpCategoryQuery {
	mp: string
}

export interface CategoryTokopediaMapQuery {
	action: string
}

export interface CategoryTokopediaMapItem {
	shopee_categ: Array<string>
}

export interface HastagQuery {
	name: string
}

export interface HastagDataResponse {
	name: string
	data: Array<string>
}

export interface HastagUpdatePayload {
	name: string
	data: Array<string>
}

export interface HastagDeleteQuery {
	name: string
}

export interface HastagLimitResponse {
	name: string
	data: Array<number>
}

export interface HastagLimitData {
	min: number
	max: number
}

export interface ExampleSpinProductsQuery {
	harga: string
	hastag: string
	polatitle: string
}

export interface MarkupDebug {
	fix_harga: number
	harga_asli: number
	harga_up: number
	markup_item_index: number
	up_percent: number
	up_price: number
	up_fix: number
}

export interface SpinProductExample {
	name: string
	price: number
	price_untung: number
	desc: string
	markup_debug: MarkupDebug
}

export interface ExampleSpinTitleQuery {
	title: string
}

export interface ExampleSpinTitleResponse {
	text: string
}

export interface CheckbotQueryCli {
	base: string
	cek: string
	out: string
}

export interface CheckOrderQueryCli {
	base: string
	akun_limit: number
	queue_size: number
	filepath: string
	output: string
}

export interface DeleteProductQueryCli {
	base: string
	report: string
}

export interface UpdaterAttributeCli {
	base: string
}

export interface ExportSupplierQuery {
	namespace: string
}

export interface ExportUrlQuery {
	namespace: string
}

export interface WeightPredictionItem {
	price_min: number
	price_max: number
	ratio: number
}

export interface WeightPrediction {
	data: Array<WeightPredictionItem | undefined>
}

export interface AttributeQuery {
	cat_id: number
}

export interface Lang {
	language: string
	value: string
}

export interface AttrChild {
	multiLang: Array<Lang>
	displayName: string
	valueId: number
	valueType: number
}

export interface AttributeInfo {
	attributeUnitList: Array<string>
	dateFormatType: number
	formatType: number
	inputType: number
	inputValidationType: number
	isNcc: boolean
	maxValueCount: number
}

export interface AttributeTree {
	attributeId: number
	displayName: string
	mandatory: boolean
	children: Array<AttrChild | undefined>
	attributeInfo: AttributeInfo | undefined
}

export interface AttributeResponse {
	exist: boolean
	attributes: Array<AttributeTree | undefined>
}

export interface QlobotShopeeImportCSVQuery {
	namespace: string
	per_item: number
}

export interface BulkItem {
	username: string
	password: string
	secret: string
}

export interface BulkPayload {
	data: Array<BulkItem | undefined>
}

export interface Response {
	msg: string
	error: string
}

export interface AkunItem {
	limit_upload: number
	count_upload: number
	active_upload: boolean
	lastup: number
	in_upload: boolean
	last_error: string
	username: string
	password: string
	secret: string
	markup: string
	spin: string
	collection: string
	hastag: string
	title_pattern: string
}

export interface AkunUpdatePayload {
	data: Array<AkunItem | undefined>
}

export interface ShopeeMapItem {
	shopee_id: number
	tokopedia_id: number
}

export interface GetMapQuery {
	collection: string
}

export interface ShopeeTopedMapResponse {
	data: Array<ShopeeMapItem | undefined>
}

export interface AutoSuggestStatus {
	status: string
}

export interface ShopeeMapperConfig {
	use_mapper: boolean
}

export interface TokopediaMapQuery {
	namespace: string
}

export interface TokopediaMapItem {
	shopee_id: number
	tokopedia_id: number
	product_count: number
	shopee_category_name: Array<string>
	tokopedia_category_name: Array<string>
}

export interface Category {
	id: number
	name: string
	url: string
	children: Array<Category | undefined>
	__typename: string
}

export interface CategoryAllListLite {
	categories: Array<Category | undefined>
	__typename: string
}

export interface CategoryAllListLiteData {
	categoryAllListLite: CategoryAllListLite | undefined
}

export interface CategoryAllListLiteRes {
	data: CategoryAllListLiteData
}

export interface UpdateTopedCategoryPayload {
	username: string
	password: string
	secret: string
}

export interface AkunDeleteItem {
	username: string
	password: string
	secret: string
}

export interface SoldConfig {
	min: number
	max: number
}

export interface ViewConfig {
	min: number
	max: number
}

export interface PriceConfig {
	min: number
	max: number
}

export interface TokopediaDeleteConfig {
	limit_concurent: number
	limit_product: number
	title: Array<string>
	product_status: string
	category_id: string
	start_time: number
	end_time: number
	akuns: Array<AkunDeleteItem | undefined>
	sold_filter: SoldConfig | undefined
	view_filter: ViewConfig | undefined
	price_filter: PriceConfig | undefined
}

export interface DeleteSettingRes {
	data: TokopediaDeleteConfig | undefined
}

export interface DumpCategoryResponse {
	errcode: number
	message: string
	status: string
}

export interface CheckVerifQueryCli {
	base: string
	fname: string
}

export interface RunCheckVerifPayload {
	fname: string
}

export interface DeleterQueryCli {
	base: string
}

export interface AutoSubmit {
	base_ktp: string
	filename: string
}

export interface Fcity {
	Description: string
	__typename: string
	child: Array<any>
	hexColor: string
	icon: string
	inputType: string
	isNew: boolean
	isPopular: boolean
	key: string
	name: string
	totalData: string
	valMax: string
	valMin: string
	value: string
}

export interface PubShippingItem {
	name: string
	Description: string
	key: string
	icon: string
	value: string
	inputType: string
	totalData: string
	valMax: string
	valMin: string
	hexColor: string
	child: Array<any>
	isPopular: boolean
	isNew: boolean
	__typename: string
}

export interface ListMapEtalaseQuery {
	namespace: string
}

export interface ShopeeEtalaseMapItem {
	shopee_id: number
	tokopedia_id: number
	product_count: number
	category_no_mapping: boolean
}

export interface ListMapEtalaseRes {
	data: Array<ShopeeEtalaseMapItem | undefined>
}

export interface DeleteEtalaseQuery {
	name: string
}

export interface EtalaseMapItem {
	etalase_name: string
	category_id: number
}

export interface EtalasePayload {
	etalase: string
	cat_ids: Array<number>
}

export interface TokopediaAttributeQuery {
	cat_id: number
}

export interface DataValue {
	id: number
	name: string
	selected: boolean
	__typename: string
}

export interface AnnotationData {
	variant: string
	sortOrder: number
	values: Array<DataValue>
	__typename: string
}

export interface TokopediaAttributeResponse {
	exist: boolean
	attributes: Array<AnnotationData | undefined>
}

export interface CheckOrderConfig {
	useDateRange: boolean
	startDate: string
	endDate: string
	useStatus: boolean
	statusKeys: Array<string>
}

export interface TokopediaCheckOrderQueryCli {
	base: string
	fname: string
	output: string
}

export interface CheckOrderAkunItem {
	password: string
	secret: string
	username: string
}

export interface CategoryNavigation {
	name: string
	url: string
	icon: string
	color: string
	children: Array<CategoryNavigation | undefined>
}

export interface JkmlCategoryListResponse {
	msg: string
	error: string
	data: Array<CategoryNavigation | undefined>
}

export interface JkmlWebResponse {
	msg: string
	error: string
}

export interface SearchCityOption {
	label: string
	value: string
}

export interface SearchSortOption {
	label: string
	value: string
}

export interface JkmlSearchDataResponse {
	msg: string
	error: string
	cities: Array<SearchCityOption | undefined>
	delivery_types: Array<string>
	sorts: Array<SearchSortOption | undefined>
}

export interface GrabSearchFilter {
	category: string
	price_min: number
	price_max: number
	untung_paling_besar: number
	in_stock: number
	bulk_price: number
	delivery_types: Array<string>
	cities: Array<string>
	sort: string
	rating: number
}

export interface JkmlSearchFilterResponse {
	msg: string
	error: string
	data: GrabSearchFilter | undefined
}

export interface JkmlCategoryMapListQuery {
	type: string
	namespace: string
}

export interface MapCateg {
	id: number
	name: string
}

export interface JkmlCategoryMap {
	type: string
	name: string
	categs: Array<MapCateg | undefined>
	mapper_id: number
	mapper_name: string
	mapper_categs: Array<MapCateg | undefined>
	count: number
}

export interface JkmlCategoryMapListResponse {
	msg: string
	error: string
	data: Array<JkmlCategoryMap | undefined>
}

export interface CategoryMapper {
	type: string
	name: string
	categs: Array<MapCateg | undefined>
	mapper_id: number
	mapper_name: string
	mapper_categs: Array<MapCateg | undefined>
}

export interface ManualShopeeUploadQueryCli {
	base: string
	reset: boolean
	one_to_multi: boolean
	limit: number
}

export interface ManualTokopediaUploadQueryCli {
	base: string
	use_mapper: boolean
	reset: boolean
	one_to_multi: boolean
	limit: number
}

export interface ShopeeUploadQueryCli {
	base: string
}

export interface TopedShopeeUploadQueryCli {
	base: string
	use_mapper: boolean
}

export interface QlobotShopeeUploadQueryCli {
	base: string
}

export interface JakmallShopeeUploadQueryCli {
	base: string
	use_mapper: boolean
}

export interface JakmallTokopediaUploadQueryCli {
	base: string
	use_mapper: boolean
}

export interface ShopeeTopedUploadQueryCli {
	base: string
}

export interface TopedTopedUploadQueryCli {
	base: string
}

export interface RenameCollectionPayload {
	marketplace: string
	namespace: string
	update_namespace: string
}

export interface BaseWebResponse {
	errcode: number
	message: string
	status: string
}

export interface TokopediaToShopeeAutoSuggestQuery {
	namespace: string
}

export interface TempAkunRes {
	data: string
	message: string
	error: number
}

export interface InfoRes {
	lisensi: string
	version: string
}

export interface CacheSizeQuery {
	reset: boolean
}

export interface SizeSum {
	size: number
	size_kb: number
	size_mb: number
	size_gb: number
}

export interface CacheSizeRes {
	processing: boolean
	cache_size: SizeSum | undefined
	webdriver_size: SizeSum | undefined
}

export interface SourceAttributeQuery {
	product_id: number
	attribute_type: string
}

export interface TokopediaAttribute {
	categories: Array<number>
	attributes: Array<string>
}

export interface AttributeResTokopediaAttribute {
	err_msg: string
	data: Array<TokopediaAttribute | undefined>
}

export interface CreateAttributePayloadTokopediaAttribute {
	product_id: number
	attribute_type: string
	data: TokopediaAttribute | undefined
}

export interface CreateAttributeResTokopediaAttribute {
	err_msg: string
	data: TokopediaAttribute | undefined
}

export interface CustomValue {
	raw_value: string
	unit: string
}

export interface AttributeValue {
	raw_value: string
	unit: string
	value_id: number
}

export interface AttributePayload {
	attribute_id: number
	attribute_value_id: number
	custom_value: CustomValue | undefined
	attribute_values: Array<AttributeValue>
}

export interface ShopeeAttribute {
	categories: Array<number>
	attributes: Array<AttributePayload | undefined>
}

export interface AttributeResShopeeAttribute {
	err_msg: string
	data: Array<ShopeeAttribute | undefined>
}

export interface CreateAttributePayloadShopeeAttribute {
	product_id: number
	attribute_type: string
	data: ShopeeAttribute | undefined
}

export interface CreateAttributeResShopeeAttribute {
	err_msg: string
	data: ShopeeAttribute | undefined
}

export interface Collection {
	id: number
	name: string
	created_at: string
	updated_at: string
}

export interface CollectionCreateRes {
	err_msg: string
	id: number
	name: string
	created_at: string
	updated_at: string
}

export interface PaginationQuery {
	page: number
	limit: number
}

export interface CollectionItem {
	id: number
	name: string
	created_at: string
	updated_at: string
	count: number
}

export interface CollectionResList {
	err_msg: string
	data: Array<CollectionItem | undefined>
	page: number
	limit: number
	count: number
}

export interface CollectionItemPayload {
	col_id: number
}

export interface ColDeletePayload {
	collection_ids: Array<number>
}

export interface ProductListQuery {
	page: number
	limit: number
	coll_id: number
	status: string
}

export interface AttributeProduct {
	id: number
	product_id: number
	attribute_type: string
	data: string
}

export interface UseOnceText {
	id: number
	field_config_id: number
	text: string
}

export interface FieldConfig {
	id: number
	product_id: number
	field_type: string
	use_spin: boolean
	spin_text: string
	use_once_text: boolean
	once_text: Array<UseOnceText | undefined>
}

export interface VariantOption {
	id: number
	product_id: number
	name: string
	option: Array<string>
}

export interface VariantImage {
	id: number
	product_id: number
	option_name: string
	image_preview: string
	image_collection_path: string
	count_image: number
}

export interface Variant {
	id: number
	product_id: number
	names: Array<string>
	values: Array<string>
	price: number
	stock: number
}

export interface WatermarkConfig {
	id: number
	product_id: number
	text: string
	font_location: string
	pos_x: string
	pos_y: string
	use_watermark: boolean
}

export interface FrameConfig {
	id: number
	product_id: number
	use_frame: boolean
	frame_location: string
}

export interface ProductMap {
	id: number
	map_id: string
	product_id: number
	url_product: string
}

export interface ManualProduct {
	id: number
	as_draft: boolean
	image_preview: string
	image_collection_path: string
	count_image: number
	title: string
	price: number
	desc: string
	use_variant: boolean | undefined
	use_markup: string
	weight: number
	stock: number
	attribute: Array<AttributeProduct | undefined>
	field_spin: Array<FieldConfig | undefined>
	variant_option: Array<VariantOption | undefined>
	variant_image: Array<VariantImage | undefined>
	variant: Array<Variant | undefined>
	collection: Array<Collection | undefined>
	watermark: WatermarkConfig | undefined
	frame_config: FrameConfig | undefined
	last_error: string
	map: Array<ProductMap | undefined>
}

export interface ProductListRes {
	err_msg: string
	data: Array<ManualProduct | undefined>
	page: number
	limit: number
	count: number
}

export interface ProductCreatePayload {
	coll_id: number
}

export interface ApiResponse {
	err_msg: string
}

export interface DeleteProductPayload {
	ids: Array<number>
}

export interface ItemQuery {
	product_id: number
}

export interface ItemRes {
	err_msg: string
	data: ManualProduct | undefined
}

export interface NewProductQuery {
	coll_id: number
}

export interface NewProductResponse {
	err_msg: string
	data: ManualProduct | undefined
}

export interface CreatePayload {
	product_id: number
}

export interface BasicUpdatePayload {
	product_id: number
	image_preview: string
	image_collection_path: string
	count_image: number
	title: string
	price: number
	desc: string
	use_variant: boolean | undefined
	use_markup: string
	weight: number
	stock: number
}

export interface UpdateVariationPayload {
	product_id: number
	variant_option: Array<VariantOption | undefined>
	variant_image: Array<VariantImage | undefined>
	variant: Array<Variant | undefined>
}

export interface UpdateFieldConfigPayload {
	product_id: number
	field_spin: Array<FieldConfig | undefined>
}

export interface SpinConfigQuery {
	field_config_id: number
}

export interface ListTextRes {
	err_msg: string
	count: number
	data: Array<UseOnceText | undefined>
}

export interface AddTextPayload {
	field_config_id: number
	data: Array<string>
}

export interface CreateFieldConfigPayload {
	product_id: number
	field_type: string
}

export interface CreateFieldConfigRes {
	err_msg: string
	data: FieldConfig | undefined
}

export interface DeleteFieldConfigPayload {
	id: number
}

export interface PathCheckPayload {
	path: string
}

export interface PathCheckResponse {
	err_msg: string
	is_abs: boolean
	Absolute: string
	exist: boolean
}



export const clients = {
	GetLauncherV1RunGrabShopee: {
		url: "launcher/v1/run_grab_shopee" as const,
		method: "GET" as const,
		query: {
				base: ``
			} as GrabShopeeQueryCli ,
		body: {},
		response: {
			use_csv: false,
			empty_csv: false
		}
	},
	GetLauncherV1RunGrabTokopedia: {
		url: "launcher/v1/run_grab_tokopedia" as const,
		method: "GET" as const,
		query: {
				base: ``
			} as GrabTokopediaQueryCli ,
		body: {},
		response: {
			deprecated: false,
			use_csv: false,
			empty_csv: false
		}
	},
	GetLauncherV1RunGrabJakmall: {
		url: "launcher/v1/run_grab_jakmall" as const,
		method: "GET" as const,
		query: {
				base: ``
			} as GrabJakmallQueryCli ,
		body: {},
		response: {
			use_csv: false,
			empty_csv: false
		}
	},
	GetLegacyApiDataspin: {
		url: "legacy/api/dataspin" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: [
			``
		] as Array<string>
	},
	GetLegacyApiDataspinGet: {
		url: "legacy/api/dataspin/get" as const,
		method: "GET" as const,
		query: {
			name: ``
		},
		body: {},
		response: [
			{
					name: ``,
					data: [
					``
					] as Array<string>
				} as DataSpinItemResponse | undefined
		] as Array<DataSpinItemResponse | undefined>
	},
	PostLegacyApiDataspin: {
		url: "legacy/api/dataspin" as const,
		method: "POST" as const,
		query: undefined,
		body: {
					name: ``,
					data: [
					``
					] as Array<string>
				},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	DeleteLegacyApiDataspin: {
		url: "legacy/api/dataspin" as const,
		method: "DELETE" as const,
		query: {
			name: ``
		},
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiSettingSpin: {
		url: "legacy/api/settingSpin" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			data: {
				smin: 0,
				smax: 0,
				merek_ins_t: false,
				ignore_first_word: false,
				title: ``,
				desc: ``
			} as SettingSpinData | undefined,
			errcode: 0,
			titlePool: [
			{
				name: ``,
				data: ``
			}
			] as Array<SettingSpinTitlePool>
		}
	},
	PostLegacyApiSettingSpin: {
		url: "legacy/api/settingSpin" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			smin: 0,
			smax: 0,
			merek_ins_t: false,
			ignore_first_word: false,
			title: ``,
			desc: ``
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	PostLegacyApiConfigSettingSpin: {
		url: "legacy/api/config/settingSpin" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			titlePool: [
			{
				name: ``,
				data: ``
			}
			] as Array<SettingSpinTitlePool>
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyV1ProductNamespaceAll: {
		url: "legacy/v1/product/namespace_all" as const,
		method: "GET" as const,
		query: {
			is_public: false,
			kota: ``,
			marketplace: ``,
			namespace: ``,
			pmax: 0,
			pmin: 0,
			use_empty_city: false
		},
		body: {},
		response: [
			{
				count: 0,
				price_min: 0,
				price_max: 0,
				name: ``
			}
		] as Array<ProductNamespaceAgg>
	},
	GetLegacyV1ProductPriceRange: {
		url: "legacy/v1/product/price_range" as const,
		method: "GET" as const,
		query: {
			is_public: false,
			kota: ``,
			marketplace: ``,
			namespace: ``,
			pmax: 0,
			pmin: 0,
			use_empty_city: false,
			range_price: 0
		},
		body: {},
		response: [
			{
				_id: [
				0
				] as Array<number>,
				count: 0
			}
		] as Array<ProductPriceRangeAgg>
	},
	GetLegacyV1ProductCategory: {
		url: "legacy/v1/product/category" as const,
		method: "GET" as const,
		query: {
			is_public: false,
			kota: ``,
			marketplace: ``,
			namespace: ``,
			pmax: 0,
			pmin: 0,
			use_empty_city: false
		},
		body: {},
		response: [
			{
				_id: 0,
				price_min: 0,
				price_max: 0,
				count: 0,
				name: [
				``
				] as Array<string>
			}
		] as Array<ProductCategoryAgg>
	},
	GetLegacyV1ProductKota: {
		url: "legacy/v1/product/kota" as const,
		method: "GET" as const,
		query: {
			is_public: false,
			kota: ``,
			marketplace: ``,
			namespace: ``,
			pmax: 0,
			pmin: 0,
			use_empty_city: false
		},
		body: {},
		response: [
			{
				_id: ``,
				price_min: 0,
				price_max: 0,
				count: 0
			}
		] as Array<ProductCityAgg>
	},
	GetLegacyV1ProductDelete: {
		url: "legacy/v1/product/delete" as const,
		method: "GET" as const,
		query: {
			is_public: false,
			kota: ``,
			marketplace: ``,
			namespace: ``,
			pmax: 0,
			pmin: 0,
			use_empty_city: false
		},
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	PostLegacyApiDeleteItem: {
		url: "legacy/api/deleteItem" as const,
		method: "POST" as const,
		query: {
			is_public: false,
			kota: ``,
			marketplace: ``,
			namespace: ``,
			pmax: 0,
			pmin: 0,
			use_empty_city: false
		},
		body: [
			0
		] as Array<number>,
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	PostLegacyV1ProductCategstatToCsv: {
		url: "legacy/v1/product/categstat_to_csv" as const,
		method: "POST" as const,
		query: undefined,
		body: [
			{
				parent_category: 0,
				catid: 0,
				parent_display_name: ``,
				display_name: ``,
				status: ``
			}
		] as Array<ShopeeCategoryOld>,
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyV1ProductResync: {
		url: "legacy/v1/product/resync" as const,
		method: "GET" as const,
		query: {
			is_public: false,
			kota: ``,
			marketplace: ``,
			namespace: ``,
			pmax: 0,
			pmin: 0,
			use_empty_city: false
		},
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	PostLegacyV1ProductRenameNamespace: {
		url: "legacy/v1/product/rename_namespace" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			marketplace: ``,
			namespace: ``,
			update_namespace: ``
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiListMarkup: {
		url: "legacy/api/listMarkup" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``,
			data: [
			``
			] as Array<string>
		}
	},
	PostLegacyApiAddMarkup: {
		url: "legacy/api/addMarkup" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			data: [
			{
				mark: ``,
				type: ``,
				range: {},
				up: [
				0
				] as Array<number>
			}
			] as Array<MarkupData>,
			fix_mark: 0,
			name: ``
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiMarkup: {
		url: "legacy/api/markup" as const,
		method: "GET" as const,
		query: {
			name: ``
		},
		body: {},
		response: {
			data: [
			{
				mark: ``,
				type: ``,
				range: {},
				up: [
				0
				] as Array<number>
			}
			] as Array<MarkupData>,
			fix_mark: 0,
			name: ``,
			errorcode: 0
		}
	},
	PostLegacyApiMarkup: {
		url: "legacy/api/markup" as const,
		method: "POST" as const,
		query: {
			name: ``
		},
		body: {
			data: [
			{
				mark: ``,
				type: ``,
				range: {},
				up: [
				0
				] as Array<number>
			}
			] as Array<MarkupData>,
			fix_mark: 0,
			name: ``
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	PostLegacyApiDeleteMarkup: {
		url: "legacy/api/deleteMarkup" as const,
		method: "POST" as const,
		query: undefined,
		body: [
			``
		] as Array<string>,
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiShopeeShipping: {
		url: "legacy/api/shopee_shipping" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: [
			{
				positionid: 0,
				name: ``,
				channelids: [
				``
				] as Array<string>,
				display_name: ``,
				item_tag_ids: [
				0
				] as Array<number>
			}
		] as Array<SearchFilterDynamicShipping>
	},
	GetLegacyShopeeManifest: {
		url: "legacy/shopee/manifest" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			category: [
			{
				id: 0,
				name: ``,
				display_name: ``,
				parent_id: 0,
				has_active_children: false,
				has_children: false,
				region_setting: {
					enable_size_chart: false,
					low_stock_value: 0,
					dimension_mandatory: false
				} as RegionSetting | undefined,
				is_prohibit: false,
				children: [
				{
						id: 0,
						name: ``,
						display_name: ``,
						parent_id: 0,
						has_active_children: false,
						has_children: false,
						region_setting: {
							enable_size_chart: false,
							low_stock_value: 0,
							dimension_mandatory: false
						} as RegionSetting | undefined,
						is_prohibit: false,
						children: [
						] as Array<CategoryTreeItem | undefined>
					} as CategoryTreeItem | undefined
				] as Array<CategoryTreeItem | undefined>,
				chain_name: [
				``
				] as Array<string>,
				chain_ids: [
				0
				] as Array<number>
			}
			] as Array<ShopeeCategory>,
			public_category_repo: [
			{
				main: {
				catid: 0,
				display_name: ``,
				name: ``,
				image: ``,
				is_adult: {},
				parent_category: 0,
				sort_weight: 0,
				block_buyer_platform: {}
			},
				sub: [
				{
					catid: 0,
					display_name: ``,
					name: ``,
					image: ``,
					is_adult: {},
					parent_category: 0,
					sort_weight: 0,
					block_buyer_platform: {},
					sub_sub: [
					{
						catid: 0,
						display_name: ``,
						image: ``,
						block_buyer_platform: {}
					}
					] as Array<CategorySubSub>
				}
				] as Array<CategorySub>
			}
			] as Array<CategoryItem>
		}
	},
	PostLegacyV4ShopeeWeightPredict: {
		url: "legacy/v4/shopee/weight/predict" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			itemid: 0,
			shopid: 0
		},
		response: {
			harga: 0,
			jarak: 0.00
		}
	},
	GetLegacyV3PredictweightSave: {
		url: "legacy/v3/predictweight/save" as const,
		method: "GET" as const,
		query: {
			predict: 0.00
		},
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyV3PredictweightLoad: {
		url: "legacy/v3/predictweight/load" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			predict_weight: 0.00
		}
	},
	GetLegacyApiSettingGetCropSetting: {
		url: "legacy/api/setting/get/cropSetting" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: null,
			errcode: 0
		}
	},
	PostLegacyApiSettingUpdateCropSetting: {
		url: "legacy/api/setting/update/cropSetting" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: null
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigConcurentRequest: {
		url: "legacy/api/config/concurentRequest" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: 0
		}
	},
	PostLegacyApiConfigConcurentRequest: {
		url: "legacy/api/config/concurentRequest" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: 0
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiUpInterval: {
		url: "legacy/api/upInterval" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: {
			uptmax: 0,
			uptmin: 0
		}
		}
	},
	PostLegacyApiUpInterval: {
		url: "legacy/api/upInterval" as const,
		method: "POST" as const,
		query: undefined,
		body: {
		uptmax: 0,
		uptmin: 0
	},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigUpThread: {
		url: "legacy/api/config/upThread" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: 0
		}
	},
	PostLegacyApiConfigUpThread: {
		url: "legacy/api/config/upThread" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: 0
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigUpMode: {
		url: "legacy/api/config/upMode" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: ``
		}
	},
	PostLegacyApiConfigUpMode: {
		url: "legacy/api/config/upMode" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: ``
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigGrabFilter: {
		url: "legacy/api/config/grabFilter" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: {
			penjualan: 0,
			prosentase: 0,
			stock: 0,
			tokped_point: [
			0
			] as Array<number>
		}
		}
	},
	PostLegacyApiConfigGrabFilter: {
		url: "legacy/api/config/grabFilter" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: {
			penjualan: 0,
			prosentase: 0,
			stock: 0,
			tokped_point: [
			0
			] as Array<number>
		}
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigLastLogin: {
		url: "legacy/api/config/last_login" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: {
			active: false,
			days: 0
		}
		}
	},
	PostLegacyApiConfigLastLogin: {
		url: "legacy/api/config/last_login" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: {
			active: false,
			days: 0
		}
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigLastReview: {
		url: "legacy/api/config/lastReview" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: {
			active: false,
			days: 0
		}
		}
	},
	PostLegacyApiConfigLastReview: {
		url: "legacy/api/config/lastReview" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: {
			active: false,
			days: 0
		}
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiLimitGrab: {
		url: "legacy/api/limitGrab" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: 0
		}
	},
	PostLegacyApiLimitGrab: {
		url: "legacy/api/limitGrab" as const,
		method: "POST" as const,
		query: undefined,
		body: 0,
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigShopeeGrabSetting: {
		url: "legacy/api/config/shopeeGrabSetting" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: {
			by: ``,
			locations: [
			``
			] as Array<string>,
			official_mall: false,
			price_max: 0,
			price_min: 0,
			rating_filter: 0,
			shopee24: false,
			shopee_verified: false
		}
		}
	},
	PostLegacyApiConfigShopeeGrabSetting: {
		url: "legacy/api/config/shopeeGrabSetting" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: {
			by: ``,
			locations: [
			``
			] as Array<string>,
			official_mall: false,
			price_max: 0,
			price_min: 0,
			rating_filter: 0,
			shopee24: false,
			shopee_verified: false
		}
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyShopeeFilterGrabber: {
		url: "legacy/shopee/filter_grabber" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			product_created: {
			active: false,
			max: 0.00,
			min: 0.00
		},
			shippings: [
			{
				channelids: [
				``
				] as Array<string>,
				display_name: ``,
				name: ``,
				item_tag_ids: [
				0
				] as Array<number>,
				positionid: 0
			}
			] as Array<GrabShopeeShipping>
		}
	},
	PutLegacyShopeeFilterGrabber: {
		url: "legacy/shopee/filter_grabber" as const,
		method: "PUT" as const,
		query: undefined,
		body: {
			product_created: {
			active: false,
			max: 0.00,
			min: 0.00
		},
			shippings: [
			{
				channelids: [
				``
				] as Array<string>,
				display_name: ``,
				name: ``,
				item_tag_ids: [
				0
				] as Array<number>,
				positionid: 0
			}
			] as Array<GrabShopeeShipping>
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyShopeeConfigCrawler: {
		url: "legacy/shopee/config_crawler" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			username: ``,
			password: ``,
			email: ``,
			email_password: ``
		}
	},
	PutLegacyShopeeConfigCrawler: {
		url: "legacy/shopee/config_crawler" as const,
		method: "PUT" as const,
		query: undefined,
		body: {
			username: ``,
			password: ``,
			email: ``,
			email_password: ``
		},
		response: {
			username: ``,
			password: ``,
			email: ``,
			email_password: ``
		}
	},
	GetLegacyApiSettingGrab: {
		url: "legacy/api/settingGrab" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			errcode: 0,
			data: {
			name: ``,
			data: {
			pmin: 0,
			pmax: 0,
			ob: ``,
			rt: ``,
			condition: ``,
			fcity: [
			``
			] as Array<string>,
			goldmerchant: false,
			official: false,
			shipping: [
			``
			] as Array<string>,
			preorder: false
		}
		}
		}
	},
	PostLegacyApiSettingGrab: {
		url: "legacy/api/settingGrab" as const,
		method: "POST" as const,
		query: undefined,
		body: {
		pmin: 0,
		pmax: 0,
		ob: ``,
		rt: ``,
		condition: ``,
		fcity: [
		``
		] as Array<string>,
		goldmerchant: false,
		official: false,
		shipping: [
		``
		] as Array<string>,
		preorder: false
	},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigCloudinary: {
		url: "legacy/api/config/cloudinary" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: {
			active: false,
			url: ``
		}
		}
	},
	PostLegacyApiConfigCloudinary: {
		url: "legacy/api/config/cloudinary" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: {
			active: false,
			url: ``
		}
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigGetRandom: {
		url: "legacy/api/config/get_random" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: false
		}
	},
	PostLegacyApiConfigGetRandom: {
		url: "legacy/api/config/get_random" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: false
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigRndAttribute: {
		url: "legacy/api/config/rnd_attribute" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: {
			active: false,
			force_tidakada: false
		}
		}
	},
	PostLegacyApiConfigRndAttribute: {
		url: "legacy/api/config/rnd_attribute" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: {
			active: false,
			force_tidakada: false
		}
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigUserSameResource: {
		url: "legacy/api/config/userSameResource" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: false
		}
	},
	PostLegacyApiConfigUserSameResource: {
		url: "legacy/api/config/userSameResource" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: false
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiSettingGetBlacklistTitle: {
		url: "legacy/api/setting/get/blacklistTitle" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			data: [
			``
			] as Array<string>,
			errcode: 0,
			name: ``
		}
	},
	PostLegacyApiSettingUpdateBlacklistTitle: {
		url: "legacy/api/setting/update/blacklistTitle" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			data: [
			``
			] as Array<string>,
			name: ``
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyV3SettingDefault: {
		url: "legacy/v3/setting/default" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			_id: ``,
			name: ``,
			use_price_discount: false,
			blacklist_username: {
			active: false,
			shopee: {
			data: [
			``
			] as Array<string>,
			filename: ``
		}
		}
		}
	},
	PostLegacyV3SettingDefault: {
		url: "legacy/v3/setting/default" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			_id: ``,
			name: ``,
			use_price_discount: false,
			blacklist_username: {
			active: false,
			shopee: {
			data: [
			``
			] as Array<string>,
			filename: ``
		}
		}
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigFilterTitle: {
		url: "legacy/api/config/filterTitle" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			data: [
			``
			] as Array<string>,
			errcode: 0,
			name: ``
		}
	},
	PostLegacyApiConfigFilterTitle: {
		url: "legacy/api/config/filterTitle" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			data: [
			``
			] as Array<string>,
			name: ``
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiFilterWord: {
		url: "legacy/api/filterWord" as const,
		method: "GET" as const,
		query: {
			name: ``,
			makedefault: false
		},
		body: {},
		response: {
			data: [
			``
			] as Array<string>,
			errcode: 0,
			name: ``
		}
	},
	PostLegacyApiFilterWord: {
		url: "legacy/api/filterWord" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			data: [
			``
			] as Array<string>,
			name: ``
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	PostLegacyApiFilterWordMakedefault: {
		url: "legacy/api/filterWord/makedefault" as const,
		method: "POST" as const,
		query: undefined,
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiBackupAkun: {
		url: "legacy/api/backupAkun" as const,
		method: "GET" as const,
		query: {
			output: ``
		},
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	PostLegacyApiUser: {
		url: "legacy/api/user" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			data: {
			_id: ``,
			user: ``,
			name: ``,
			pass: ``,
			password: ``,
			active: false,
			limit_upload: 0,
			count_upload: 0,
			namespace: ``,
			water: ``,
			shopee_categ: [
			``
			] as Array<string>,
			tokped_categ: [
			``
			] as Array<string>,
			type: ``,
			markup: ``,
			hastag: ``,
			last_up: 0,
			polatitle: ``,
			is_running_upload: false
		}
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	DeleteLegacyApiUser: {
		url: "legacy/api/user" as const,
		method: "DELETE" as const,
		query: undefined,
		body: {
			data: [
			``
			] as Array<string>
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiAkuns: {
		url: "legacy/api/akuns" as const,
		method: "GET" as const,
		query: {
			start: 0,
			limit: 0,
			active: ``,
			last: 0,
			search: ``,
			sort: ``,
			reverse: 0
		},
		body: {},
		response: {
			akuns: [
			{
			_id: ``,
			user: ``,
			name: ``,
			pass: ``,
			password: ``,
			active: false,
			limit_upload: 0,
			count_upload: 0,
			namespace: ``,
			water: ``,
			shopee_categ: [
			``
			] as Array<string>,
			tokped_categ: [
			``
			] as Array<string>,
			type: ``,
			markup: ``,
			hastag: ``,
			last_up: 0,
			polatitle: ``,
			is_running_upload: false
		}
			] as Array<Akun>,
			count: 0,
			errcode: 0
		}
	},
	GetLegacyV1AkunReset: {
		url: "legacy/v1/akun/reset" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	PutLegacyApiUser: {
		url: "legacy/api/user" as const,
		method: "PUT" as const,
		query: undefined,
		body: {
			data: {
			_id: ``,
			user: ``,
			name: ``,
			pass: ``,
			password: ``,
			active: false,
			limit_upload: 0,
			count_upload: 0,
			namespace: ``,
			water: ``,
			shopee_categ: [
			``
			] as Array<string>,
			tokped_categ: [
			``
			] as Array<string>,
			type: ``,
			markup: ``,
			hastag: ``,
			last_up: 0,
			polatitle: ``,
			is_running_upload: false
		}
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyShopeeConfigUpload: {
		url: "legacy/shopee/config_upload" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			shipping: [
			{
				channelid: 0,
				enabled: false
			}
			] as Array<UploadShipping>,
			use_custom_shipping: false
		}
	},
	PutLegacyShopeeConfigUpload: {
		url: "legacy/shopee/config_upload" as const,
		method: "PUT" as const,
		query: undefined,
		body: {
			shipping: [
			{
				channelid: 0,
				enabled: false
			}
			] as Array<UploadShipping>,
			use_custom_shipping: false
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyV4ShopeeUploadConfigProductShippingList: {
		url: "legacy/v4/shopee/upload_config/product_shipping_list" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: [
			{
				channel_id: 0,
				name: ``,
				display_name: ``,
				name_key: ``,
				flag: ``,
				service_type: ``,
				default_price: ``,
				min_default_price: ``,
				max_default_price: ``,
				limits: {
				item_min_weight: 0
			},
				sizes: [
				{}
				] as Array<any>,
				cover_shipping_fee: false,
				enabled: false,
				item_flag: ``,
				price: ``,
				size: 0,
				size_id: 0,
				save_into_item: false,
				volumetric_factor: 0,
				conflicting_enabled_channels: {},
				is_mask_channel: false,
				parent_channel_id: 0,
				item_exclusive_channels: [
				0
				] as Array<number>,
				related_enabled_channels: [
				{}
				] as Array<any>,
				related_disabled_channels: [
				{}
				] as Array<any>
			}
		] as Array<UploadProductShipping>
	},
	PostLegacyApiTool: {
		url: "legacy/api/tool" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			data: ``
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyV3ConfigurationGet: {
		url: "legacy/v3/configuration/get" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			captcha: {
			user: ``,
			pass: ``
		},
			database: {
			DB_URI: ``,
			DB_NAME: ``
		},
			lisensi: {
			email: ``,
			pwd: ``
		},
			suplier_csv: false,
			session_storage_source: ``,
			session_storage_local_location: ``,
			thread_count: 0,
			tor_directory: ``,
			tor_print_ip: false,
			tor_proxy: false
		}
	},
	GetLegacyApiConfigDelete: {
		url: "legacy/api/config_delete" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			akun: ``,
			awaltanggal: `2021-12-01T07:00:00+07:00`,
			blokir: false,
			delete: 0,
			diarsipkan: false,
			diperiksa: false,
			sold: 0,
			tanggal: `2021-12-01T07:00:00+07:00`,
			view: 0
		}
	},
	PostLegacyApiConfigDelete: {
		url: "legacy/api/config_delete" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			akun: ``,
			awaltanggal: `2021-12-01T07:00:00+07:00`,
			blokir: false,
			delete: 0,
			diarsipkan: false,
			diperiksa: false,
			sold: 0,
			tanggal: `2021-12-01T07:00:00+07:00`,
			view: 0
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigDeleteProduct: {
		url: "legacy/api/config/deleteProduct" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: {
			fil_category: false,
			fil_harga: false,
			fil_keyword: false,
			category: [
			[
				0
			] as Array<number>
			] as Array<Array<number>>,
			harga: {
			min: 0,
			max: 0
		},
			keyword: ``
		}
		}
	},
	PostLegacyApiConfigDeleteProduct: {
		url: "legacy/api/config/deleteProduct" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: {
			fil_category: false,
			fil_harga: false,
			fil_keyword: false,
			category: [
			[
				0
			] as Array<number>
			] as Array<Array<number>>,
			harga: {
			min: 0,
			max: 0
		},
			keyword: ``
		}
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyV1TaskerAll: {
		url: "legacy/v1/tasker/all" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: [
			{
				_id: ``,
				toko_username: ``,
				mode: ``,
				marketplace: ``,
				product_url: ``,
				namespace: ``,
				tokped_categ: [
				``
				] as Array<string>,
				use_filter: false,
				keyword: ``,
				shopee_categ: {
				parent_category: 0,
				catid: 0,
				parent_display_name: ``,
				display_name: ``,
				is_collection: 0
			},
				jakmall_categs: [
				``
				] as Array<string>
			}
		] as Array<GrabTasker>
	},
	PostLegacyV1TaskerSave: {
		url: "legacy/v1/tasker/save" as const,
		method: "POST" as const,
		query: undefined,
		body: [
			{
				_id: ``,
				toko_username: ``,
				mode: ``,
				marketplace: ``,
				product_url: ``,
				namespace: ``,
				tokped_categ: [
				``
				] as Array<string>,
				use_filter: false,
				keyword: ``,
				shopee_categ: {
				parent_category: 0,
				catid: 0,
				parent_display_name: ``,
				display_name: ``,
				is_collection: 0
			},
				jakmall_categs: [
				``
				] as Array<string>
			}
		] as Array<GrabTasker>,
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	DeleteLegacyV1Taskerid: {
		url: "legacy/v1/tasker/:id" as const,
		method: "DELETE" as const,
		query: undefined,
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiCategoryDumpCsv: {
		url: "legacy/api/category/dump_csv" as const,
		method: "GET" as const,
		query: {
			mp: ``
		},
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyV4ShopeeSearchFilterDumpFacetCategory: {
		url: "legacy/v4/shopee/search_filter/dump_facet_category" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {} as any
	},
	GetLegacyV4ShopeeSearchFilterDumpFacetCategoryLink: {
		url: "legacy/v4/shopee/search_filter/dump_facet_category_link" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyV1CategoryMapperTokopedia: {
		url: "legacy/v1/category_mapper/tokopedia" as const,
		method: "GET" as const,
		query: {
			action: ``
		},
		body: {},
		response: {}
	},
	GetLegacyApiHastag: {
		url: "legacy/api/hastag" as const,
		method: "GET" as const,
		query: {
			name: ``
		},
		body: {},
		response: {
			name: ``,
			data: [
			``
			] as Array<string>
		}
	},
	GetLegacyApiHastagList: {
		url: "legacy/api/hastag/list" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: [
			``
		] as Array<string>
	},
	PostLegacyApiHastag: {
		url: "legacy/api/hastag" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			name: ``,
			data: [
			``
			] as Array<string>
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	DeleteLegacyApiHastag: {
		url: "legacy/api/hastag" as const,
		method: "DELETE" as const,
		query: {
			name: ``
		},
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyApiConfigHastagLimit: {
		url: "legacy/api/config/hastagLimit" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			name: ``,
			data: [
			0
			] as Array<number>
		}
	},
	PostLegacyApiConfigHastagLimit: {
		url: "legacy/api/config/hastagLimit" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			min: 0,
			max: 0
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetLegacyV1ExamplespinProducts: {
		url: "legacy/v1/examplespin/products" as const,
		method: "GET" as const,
		query: {
			harga: ``,
			hastag: ``,
			polatitle: ``
		},
		body: {},
		response: [
			{
				name: ``,
				price: 0,
				price_untung: 0,
				desc: ``,
				markup_debug: {
				fix_harga: 0,
				harga_asli: 0,
				harga_up: 0,
				markup_item_index: 0,
				up_percent: 0.00,
				up_price: 0,
				up_fix: 0
			}
			}
		] as Array<SpinProductExample>
	},
	GetLegacyV1ExamplespinTitle: {
		url: "legacy/v1/examplespin/title" as const,
		method: "GET" as const,
		query: {
			title: ``
		},
		body: {},
		response: {
			text: ``
		}
	},
	GetShopeeV5RunCheckbot: {
		url: "shopee/v5/run_checkbot" as const,
		method: "GET" as const,
		query: {
				base: ``,
				cek: ``,
				out: ``
			} as CheckbotQueryCli ,
		body: {},
		response: {} as any
	},
	GetShopeeV5RunCheckOrder: {
		url: "shopee/v5/run_check_order" as const,
		method: "GET" as const,
		query: {
				base: ``,
				akun_limit: 0,
				queue_size: 0,
				filepath: ``,
				output: ``
			} as CheckOrderQueryCli ,
		body: {},
		response: {} as any
	},
	PostShopeeV5RunDeleteProduct: {
		url: "shopee/v5/run_delete_product" as const,
		method: "POST" as const,
		query: {
				base: ``,
				report: ``
			} as DeleteProductQueryCli ,
		body: {},
		response: {} as any
	},
	GetShopeeV5UpdaterAttribute: {
		url: "shopee/v5/updater_attribute" as const,
		method: "GET" as const,
		query: {
				base: ``
			} as UpdaterAttributeCli ,
		body: {},
		response: {} as any
	},
	PutShopeeV5ProductExportSupplier: {
		url: "shopee/v5/product/export_supplier" as const,
		method: "PUT" as const,
		query: {
				namespace: ``
			} as ExportSupplierQuery ,
		body: {},
		response: {} as any
	},
	PutShopeeV5ProductExportUrl: {
		url: "shopee/v5/product/export_url" as const,
		method: "PUT" as const,
		query: {
				namespace: ``
			} as ExportUrlQuery ,
		body: {},
		response: {} as any
	},
	GetShopeeV5FilterGrabLocation: {
		url: "shopee/v5/filter/grab_location" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: [
			``
		] as Array<string>
	},
	GetShopeeV5ConfigWeightRatio: {
		url: "shopee/v5/config/weight_ratio" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
				data: [
				{
						price_min: 0,
						price_max: 0,
						ratio: 0.00
					} as WeightPredictionItem | undefined
				] as Array<WeightPredictionItem | undefined>
			} as WeightPrediction 
	},
	PutShopeeV5ConfigWeightRatio: {
		url: "shopee/v5/config/weight_ratio" as const,
		method: "PUT" as const,
		query: undefined,
		body: {
				data: [
				{
						price_min: 0,
						price_max: 0,
						ratio: 0.00
					} as WeightPredictionItem | undefined
				] as Array<WeightPredictionItem | undefined>
			} as WeightPrediction ,
		response: {
				data: [
				{
						price_min: 0,
						price_max: 0,
						ratio: 0.00
					} as WeightPredictionItem | undefined
				] as Array<WeightPredictionItem | undefined>
			} as WeightPrediction 
	},
	GetShopeeV5Attribute: {
		url: "shopee/v5/attribute" as const,
		method: "GET" as const,
		query: {
				cat_id: 0
			} as AttributeQuery ,
		body: {},
		response: {
				exist: false,
				attributes: [
				{
						attributeId: 0,
						displayName: ``,
						mandatory: false,
						children: [
						{
								multiLang: [
								{
									language: ``,
									value: ``
								}
								] as Array<Lang>,
								displayName: ``,
								valueId: 0,
								valueType: 0
							} as AttrChild | undefined
						] as Array<AttrChild | undefined>,
						attributeInfo: {
							attributeUnitList: [
							``
							] as Array<string>,
							dateFormatType: 0,
							formatType: 0,
							inputType: 0,
							inputValidationType: 0,
							isNcc: false,
							maxValueCount: 0
						} as AttributeInfo | undefined
					} as AttributeTree | undefined
				] as Array<AttributeTree | undefined>
			} as AttributeResponse 
	},
	PostShopeeV5QlobotShopeeImportCsv: {
		url: "shopee/v5/qlobot/shopee_import_csv" as const,
		method: "POST" as const,
		query: {
			namespace: ``,
			per_item: 0
		},
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetTokopediaAkunList: {
		url: "tokopedia/akun/list" as const,
		method: "GET" as const,
		query: {
			start: 0,
			limit: 0,
			active: ``,
			last: 0,
			search: ``,
			sort: ``,
			reverse: 0
		},
		body: {},
		response: {
			akuns: [
			{
			_id: ``,
			user: ``,
			name: ``,
			pass: ``,
			password: ``,
			active: false,
			limit_upload: 0,
			count_upload: 0,
			namespace: ``,
			water: ``,
			shopee_categ: [
			``
			] as Array<string>,
			tokped_categ: [
			``
			] as Array<string>,
			type: ``,
			markup: ``,
			hastag: ``,
			last_up: 0,
			polatitle: ``,
			is_running_upload: false
		}
			] as Array<Akun>,
			count: 0,
			errcode: 0
		}
	},
	PostTokopediaAkunBulkAdd: {
		url: "tokopedia/akun/bulk_add" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			data: [
			{
					username: ``,
					password: ``,
					secret: ``
				} as BulkItem | undefined
			] as Array<BulkItem | undefined>
		},
		response: {
			msg: ``,
			error: ``
		}
	},
	PostTokopediaAkunUpdate: {
		url: "tokopedia/akun/update" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			data: [
			{
					limit_upload: 0,
					count_upload: 0,
					active_upload: false,
					lastup: 0,
					in_upload: false,
					last_error: ``,
					username: ``,
					password: ``,
					secret: ``,
					markup: ``,
					spin: ``,
					collection: ``,
					hastag: ``,
					title_pattern: ``
				} as AkunItem | undefined
			] as Array<AkunItem | undefined>
		},
		response: {
			msg: ``,
			error: ``
		}
	},
	PostTokopediaAkunDelete: {
		url: "tokopedia/akun/delete" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			data: [
			``
			] as Array<string>
		},
		response: {
			msg: ``,
			error: ``
		}
	},
	PutTokopediaAkunResetAllCount: {
		url: "tokopedia/akun/reset_all_count" as const,
		method: "PUT" as const,
		query: undefined,
		body: {},
		response: {
			msg: ``,
			error: ``
		}
	},
	PutTokopediaMapperMap: {
		url: "tokopedia/mapper/map" as const,
		method: "PUT" as const,
		query: undefined,
		body: [
			{
				shopee_id: 0,
				tokopedia_id: 0
			}
		] as Array<ShopeeMapItem>,
		response: {
			msg: ``,
			error: ``
		}
	},
	GetTokopediaMapperMap: {
		url: "tokopedia/mapper/map" as const,
		method: "GET" as const,
		query: {
			collection: ``
		},
		body: {},
		response: {
			data: [
			{
					shopee_id: 0,
					tokopedia_id: 0
				} as ShopeeMapItem | undefined
			] as Array<ShopeeMapItem | undefined>
		}
	},
	PutTokopediaMapperAutosuggest: {
		url: "tokopedia/mapper/autosuggest" as const,
		method: "PUT" as const,
		query: {
			collection: ``
		},
		body: {},
		response: {} as any
	},
	GetTokopediaMapperAutosuggest: {
		url: "tokopedia/mapper/autosuggest" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			status: ``
		}
	},
	GetTokopediaMapperSetting: {
		url: "tokopedia/mapper/setting" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			use_mapper: false
		}
	},
	PutTokopediaMapperSetting: {
		url: "tokopedia/mapper/setting" as const,
		method: "PUT" as const,
		query: undefined,
		body: {
			use_mapper: false
		} as ShopeeMapperConfig ,
		response: {
			use_mapper: false
		}
	},
	GetTokopediaMapperCategory: {
		url: "tokopedia/mapper/category" as const,
		method: "GET" as const,
		query: {
				namespace: ``
			} as TokopediaMapQuery ,
		body: {},
		response: [
			{
				shopee_id: 0,
				tokopedia_id: 0,
				product_count: 0,
				shopee_category_name: [
				``
				] as Array<string>,
				tokopedia_category_name: [
				``
				] as Array<string>
			}
		] as Array<TokopediaMapItem>
	},
	GetTokopediaCategoryList: {
		url: "tokopedia/category/list" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			data: {
			categoryAllListLite: {
				categories: [
				{
						id: 0,
						name: ``,
						url: ``,
						children: [
						] as Array<Category | undefined>,
						__typename: ``
					} as Category | undefined
				] as Array<Category | undefined>,
				__typename: ``
			} as CategoryAllListLite | undefined
		}
		}
	},
	PutTokopediaCategoryUpdateCategory: {
		url: "tokopedia/category/update_category" as const,
		method: "PUT" as const,
		query: undefined,
		body: {
			username: ``,
			password: ``,
			secret: ``
		},
		response: {
			msg: ``,
			error: ``
		}
	},
	GetTokopediaDeleterSetting: {
		url: "tokopedia/deleter/setting" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			data: {
				limit_concurent: 0,
				limit_product: 0,
				title: [
				``
				] as Array<string>,
				product_status: ``,
				category_id: ``,
				start_time: 0,
				end_time: 0,
				akuns: [
				{
						username: ``,
						password: ``,
						secret: ``
					} as AkunDeleteItem | undefined
				] as Array<AkunDeleteItem | undefined>,
				sold_filter: {
					min: 0,
					max: 0
				} as SoldConfig | undefined,
				view_filter: {
					min: 0,
					max: 0
				} as ViewConfig | undefined,
				price_filter: {
					min: 0,
					max: 0
				} as PriceConfig | undefined
			} as TokopediaDeleteConfig | undefined
		}
	},
	PutTokopediaDeleterSetting: {
		url: "tokopedia/deleter/setting" as const,
		method: "PUT" as const,
		query: undefined,
		body: {
			limit_concurent: 0,
			limit_product: 0,
			title: [
			``
			] as Array<string>,
			product_status: ``,
			category_id: ``,
			start_time: 0,
			end_time: 0,
			akuns: [
			{
					username: ``,
					password: ``,
					secret: ``
				} as AkunDeleteItem | undefined
			] as Array<AkunDeleteItem | undefined>,
			sold_filter: {
				min: 0,
				max: 0
			} as SoldConfig | undefined,
			view_filter: {
				min: 0,
				max: 0
			} as ViewConfig | undefined,
			price_filter: {
				min: 0,
				max: 0
			} as PriceConfig | undefined
		},
		response: {
			data: {
				limit_concurent: 0,
				limit_product: 0,
				title: [
				``
				] as Array<string>,
				product_status: ``,
				category_id: ``,
				start_time: 0,
				end_time: 0,
				akuns: [
				{
						username: ``,
						password: ``,
						secret: ``
					} as AkunDeleteItem | undefined
				] as Array<AkunDeleteItem | undefined>,
				sold_filter: {
					min: 0,
					max: 0
				} as SoldConfig | undefined,
				view_filter: {
					min: 0,
					max: 0
				} as ViewConfig | undefined,
				price_filter: {
					min: 0,
					max: 0
				} as PriceConfig | undefined
			} as TokopediaDeleteConfig | undefined
		}
	},
	GetTokopediaDumpCategoryDump: {
		url: "tokopedia/dump/category_dump" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	PutTokopediaCekbotRun: {
		url: "tokopedia/cekbot/run" as const,
		method: "PUT" as const,
		query: {
				base: ``,
				cek: ``,
				out: ``
			} as CheckbotQueryCli ,
		body: {},
		response: {} as any
	},
	PutTokopediaCheckVerifRun: {
		url: "tokopedia/check_verif/run" as const,
		method: "PUT" as const,
		query: {
				base: ``,
				fname: ``
			} as CheckVerifQueryCli ,
		body: {
			fname: ``
		},
		response: {} as any
	},
	PutTokopediaDeleterRunDelete: {
		url: "tokopedia/deleter/run_delete" as const,
		method: "PUT" as const,
		query: {
				base: ``
			} as DeleterQueryCli ,
		body: {},
		response: {} as any
	},
	PostTokopediaAutosubmitRun: {
		url: "tokopedia/autosubmit/run" as const,
		method: "POST" as const,
		query: undefined,
		body: {
				base_ktp: ``,
				filename: ``
			} as AutoSubmit ,
		response: {} as any
	},
	GetTokopediaFilterFcity: {
		url: "tokopedia/filter/fcity" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: [
			{
					Description: ``,
					__typename: ``,
					child: [
					{}
					] as Array<any>,
					hexColor: ``,
					icon: ``,
					inputType: ``,
					isNew: false,
					isPopular: false,
					key: ``,
					name: ``,
					totalData: ``,
					valMax: ``,
					valMin: ``,
					value: ``
				} as Fcity | undefined
		] as Array<Fcity | undefined>
	},
	GetTokopediaFilterCategories: {
		url: "tokopedia/filter/categories" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: [
			{
					Description: ``,
					__typename: ``,
					child: [
					{}
					] as Array<any>,
					hexColor: ``,
					icon: ``,
					inputType: ``,
					isNew: false,
					isPopular: false,
					key: ``,
					name: ``,
					totalData: ``,
					valMax: ``,
					valMin: ``,
					value: ``
				} as Fcity | undefined
		] as Array<Fcity | undefined>
	},
	GetTokopediaFilterShipping: {
		url: "tokopedia/filter/shipping" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: [
			{
					name: ``,
					Description: ``,
					key: ``,
					icon: ``,
					value: ``,
					inputType: ``,
					totalData: ``,
					valMax: ``,
					valMin: ``,
					hexColor: ``,
					child: [
					{}
					] as Array<any>,
					isPopular: false,
					isNew: false,
					__typename: ``
				} as PubShippingItem | undefined
		] as Array<PubShippingItem | undefined>
	},
	GetTokopediaEtalaseMapList: {
		url: "tokopedia/etalase_map/list" as const,
		method: "GET" as const,
		query: {
			namespace: ``
		},
		body: {},
		response: {
				data: [
				{
						shopee_id: 0,
						tokopedia_id: 0,
						product_count: 0,
						category_no_mapping: false
					} as ShopeeEtalaseMapItem | undefined
				] as Array<ShopeeEtalaseMapItem | undefined>
			} as ListMapEtalaseRes 
	},
	DeleteTokopediaEtalaseMapDelete: {
		url: "tokopedia/etalase_map/delete" as const,
		method: "DELETE" as const,
		query: {
				name: ``
			} as DeleteEtalaseQuery ,
		body: {},
		response: {} as any
	},
	PutTokopediaEtalaseMapUpdate: {
		url: "tokopedia/etalase_map/update" as const,
		method: "PUT" as const,
		query: undefined,
		body: [
			{
					etalase_name: ``,
					category_id: 0
				} as EtalaseMapItem | undefined
		] as Array<EtalaseMapItem | undefined>,
		response: {
			msg: ``,
			error: ``
		} as Response 
	},
	GetTokopediaEtalaseMapListEtalase: {
		url: "tokopedia/etalase_map/list_etalase" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: [
			{
					etalase: ``,
					cat_ids: [
					0
					] as Array<number>
				} as EtalasePayload | undefined
		] as Array<EtalasePayload | undefined>
	},
	GetTokopediaAttributeUpdaterAttribute: {
		url: "tokopedia/attribute/updater_attribute" as const,
		method: "GET" as const,
		query: {
				base: ``
			} as UpdaterAttributeCli ,
		body: {},
		response: {} as any
	},
	GetTokopediaAttributeGetAttribute: {
		url: "tokopedia/attribute/get_attribute" as const,
		method: "GET" as const,
		query: {
			cat_id: 0
		},
		body: {},
		response: {
			exist: false,
			attributes: [
			{
					variant: ``,
					sortOrder: 0,
					values: [
					{
						id: 0,
						name: ``,
						selected: false,
						__typename: ``
					}
					] as Array<DataValue>,
					__typename: ``
				} as AnnotationData | undefined
			] as Array<AnnotationData | undefined>
		}
	},
	PutTokopediaCekorderSaveConfig: {
		url: "tokopedia/cekorder/save_config" as const,
		method: "PUT" as const,
		query: undefined,
		body: {
				useDateRange: false,
				startDate: ``,
				endDate: ``,
				useStatus: false,
				statusKeys: [
				``
				] as Array<string>
			} as CheckOrderConfig ,
		response: {} as any
	},
	PutTokopediaCekorderRun: {
		url: "tokopedia/cekorder/run" as const,
		method: "PUT" as const,
		query: {
				base: ``,
				fname: ``,
				output: ``
			} as TokopediaCheckOrderQueryCli ,
		body: [
			{
					password: ``,
					secret: ``,
					username: ``
				} as CheckOrderAkunItem | undefined
		] as Array<CheckOrderAkunItem | undefined>,
		response: {} as any
	},
	GetJakmallCategoryList: {
		url: "jakmall/category/list" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			msg: ``,
			error: ``,
			data: [
			{
					name: ``,
					url: ``,
					icon: ``,
					color: ``,
					children: [
					] as Array<CategoryNavigation | undefined>
				} as CategoryNavigation | undefined
			] as Array<CategoryNavigation | undefined>
		}
	},
	GetJakmallCategoryDumpCsv: {
		url: "jakmall/category/dump_csv" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			msg: ``,
			error: ``
		}
	},
	GetJakmallSearchFilterData: {
		url: "jakmall/search_filter/data" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			msg: ``,
			error: ``,
			cities: [
			{
					label: ``,
					value: ``
				} as SearchCityOption | undefined
			] as Array<SearchCityOption | undefined>,
			delivery_types: [
			``
			] as Array<string>,
			sorts: [
			{
					label: ``,
					value: ``
				} as SearchSortOption | undefined
			] as Array<SearchSortOption | undefined>
		}
	},
	GetJakmallSearchFilter: {
		url: "jakmall/search_filter" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			msg: ``,
			error: ``,
			data: {
				category: ``,
				price_min: 0,
				price_max: 0,
				untung_paling_besar: 0,
				in_stock: 0,
				bulk_price: 0,
				delivery_types: [
				``
				] as Array<string>,
				cities: [
				``
				] as Array<string>,
				sort: ``,
				rating: 0.00
			} as GrabSearchFilter | undefined
		}
	},
	PostJakmallSearchFilter: {
		url: "jakmall/search_filter" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			category: ``,
			price_min: 0,
			price_max: 0,
			untung_paling_besar: 0,
			in_stock: 0,
			bulk_price: 0,
			delivery_types: [
			``
			] as Array<string>,
			cities: [
			``
			] as Array<string>,
			sort: ``,
			rating: 0.00
		},
		response: {
			msg: ``,
			error: ``
		}
	},
	GetJakmallCategoryMapperList: {
		url: "jakmall/category_mapper/list" as const,
		method: "GET" as const,
		query: {
			type: ``,
			namespace: ``
		},
		body: {},
		response: {
			msg: ``,
			error: ``,
			data: [
			{
					type: ``,
					name: ``,
					categs: [
					{
							id: 0,
							name: ``
						} as MapCateg | undefined
					] as Array<MapCateg | undefined>,
					mapper_id: 0,
					mapper_name: ``,
					mapper_categs: [
					{
							id: 0,
							name: ``
						} as MapCateg | undefined
					] as Array<MapCateg | undefined>,
					count: 0
				} as JkmlCategoryMap | undefined
			] as Array<JkmlCategoryMap | undefined>
		}
	},
	PostJakmallCategoryMapperSave: {
		url: "jakmall/category_mapper/save" as const,
		method: "POST" as const,
		query: undefined,
		body: [
			{
				type: ``,
				name: ``,
				categs: [
				{
						id: 0,
						name: ``
					} as MapCateg | undefined
				] as Array<MapCateg | undefined>,
				mapper_id: 0,
				mapper_name: ``,
				mapper_categs: [
				{
						id: 0,
						name: ``
					} as MapCateg | undefined
				] as Array<MapCateg | undefined>
			}
		] as Array<CategoryMapper>,
		response: {
			msg: ``,
			error: ``
		}
	},
	PutJakmallCategoryMapperAutosuggest: {
		url: "jakmall/category_mapper/autosuggest" as const,
		method: "PUT" as const,
		query: {
			type: ``,
			namespace: ``
		},
		body: {},
		response: {
			msg: ``,
			error: ``
		}
	},
	GetUploadV6ManualToShopee: {
		url: "upload/v6/manual_to_shopee" as const,
		method: "GET" as const,
		query: {
				base: ``,
				reset: false,
				one_to_multi: false,
				limit: 0
			} as ManualShopeeUploadQueryCli ,
		body: {},
		response: {} as any
	},
	GetUploadV6ManualToTokopedia: {
		url: "upload/v6/manual_to_tokopedia" as const,
		method: "GET" as const,
		query: {
				base: ``,
				use_mapper: false,
				reset: false,
				one_to_multi: false,
				limit: 0
			} as ManualTokopediaUploadQueryCli ,
		body: {},
		response: {} as any
	},
	GetUploadV6ShopeeToShopee: {
		url: "upload/v6/shopee_to_shopee" as const,
		method: "GET" as const,
		query: {
				base: ``
			} as ShopeeUploadQueryCli ,
		body: {},
		response: {} as any
	},
	GetUploadV6TokopediaToShopee: {
		url: "upload/v6/tokopedia_to_shopee" as const,
		method: "GET" as const,
		query: {
				base: ``,
				use_mapper: false
			} as TopedShopeeUploadQueryCli ,
		body: {},
		response: {} as any
	},
	GetUploadV6QlobotToShopee: {
		url: "upload/v6/qlobot_to_shopee" as const,
		method: "GET" as const,
		query: {
			base: ``
		},
		body: {},
		response: {
			msg: ``,
			error: ``
		}
	},
	GetUploadV6JakmallToShopee: {
		url: "upload/v6/jakmall_to_shopee" as const,
		method: "GET" as const,
		query: {
			base: ``,
			use_mapper: false
		},
		body: {},
		response: {
			msg: ``,
			error: ``
		}
	},
	GetUploadV6JakmallToTokopedia: {
		url: "upload/v6/jakmall_to_tokopedia" as const,
		method: "GET" as const,
		query: {
			base: ``,
			use_mapper: false
		},
		body: {},
		response: {
			msg: ``,
			error: ``
		}
	},
	GetTokopediaUploadShopee: {
		url: "tokopedia/upload/shopee" as const,
		method: "GET" as const,
		query: {
			base: ``
		},
		body: {},
		response: {
			msg: ``,
			error: ``
		}
	},
	GetTokopediaUploadTokopedia: {
		url: "tokopedia/upload/tokopedia" as const,
		method: "GET" as const,
		query: {
			base: ``
		},
		body: {},
		response: {
			msg: ``,
			error: ``
		}
	},
	PostV1ProductRenameNamespace: {
		url: "v1/product/rename_namespace" as const,
		method: "POST" as const,
		query: undefined,
		body: {
			marketplace: ``,
			namespace: ``,
			update_namespace: ``
		},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetV1CategoryMapperTokopediaToShopeeAutosuggest: {
		url: "v1/category/mapper/tokopedia_to_shopee_autosuggest" as const,
		method: "GET" as const,
		query: {
			namespace: ``
		},
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetV1AkunTempAkun: {
		url: "v1/akun/temp_akun" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			data: ``,
			message: ``,
			error: 0
		}
	},
	GetV1MainInfo: {
		url: "v1/main/info" as const,
		method: "GET" as const,
		query: undefined,
		body: {},
		response: {
			lisensi: ``,
			version: ``
		}
	},
	GetV1MainCacheSize: {
		url: "v1/main/cache_size" as const,
		method: "GET" as const,
		query: {
			reset: false
		},
		body: {},
		response: {
			processing: false,
			cache_size: {
				size: 0.00,
				size_kb: 0.00,
				size_mb: 0.00,
				size_gb: 0.00
			} as SizeSum | undefined,
			webdriver_size: {
				size: 0.00,
				size_kb: 0.00,
				size_mb: 0.00,
				size_gb: 0.00
			} as SizeSum | undefined
		}
	},
	DeleteV1MainClearCache: {
		url: "v1/main/clear_cache" as const,
		method: "DELETE" as const,
		query: undefined,
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	DeleteV1MainClearCacheWebdriver: {
		url: "v1/main/clear_cache_webdriver" as const,
		method: "DELETE" as const,
		query: undefined,
		body: {},
		response: {
			errcode: 0,
			message: ``,
			status: ``
		}
	},
	GetPdcsourceAttrToped: {
		url: "pdcsource/attr_toped" as const,
		method: "GET" as const,
		query: {
				product_id: 0,
				attribute_type: ``
			} as SourceAttributeQuery ,
		body: {},
		response: {
				err_msg: ``,
				data: [
				{
						categories: [
						0
						] as Array<number>,
						attributes: [
						``
						] as Array<string>
					} as TokopediaAttribute | undefined
				] as Array<TokopediaAttribute | undefined>
			} as AttributeResTokopediaAttribute 
	},
	PutPdcsourceAttrToped: {
		url: "pdcsource/attr_toped" as const,
		method: "PUT" as const,
		query: undefined,
		body: {
				product_id: 0,
				attribute_type: ``,
				data: {
						categories: [
						0
						] as Array<number>,
						attributes: [
						``
						] as Array<string>
					} as TokopediaAttribute | undefined
			} as CreateAttributePayloadTokopediaAttribute ,
		response: {
				err_msg: ``,
				data: {
						categories: [
						0
						] as Array<number>,
						attributes: [
						``
						] as Array<string>
					} as TokopediaAttribute | undefined
			} as CreateAttributeResTokopediaAttribute 
	},
	GetPdcsourceAttShopee: {
		url: "pdcsource/att_shopee" as const,
		method: "GET" as const,
		query: {
				product_id: 0,
				attribute_type: ``
			} as SourceAttributeQuery ,
		body: {},
		response: {
				err_msg: ``,
				data: [
				{
						categories: [
						0
						] as Array<number>,
						attributes: [
						{
								attribute_id: 0,
								attribute_value_id: 0,
								custom_value: {
									raw_value: ``,
									unit: ``
								} as CustomValue | undefined,
								attribute_values: [
								{
									raw_value: ``,
									unit: ``,
									value_id: 0
								}
								] as Array<AttributeValue>
							} as AttributePayload | undefined
						] as Array<AttributePayload | undefined>
					} as ShopeeAttribute | undefined
				] as Array<ShopeeAttribute | undefined>
			} as AttributeResShopeeAttribute 
	},
	PutPdcsourceAttShopee: {
		url: "pdcsource/att_shopee" as const,
		method: "PUT" as const,
		query: undefined,
		body: {
				product_id: 0,
				attribute_type: ``,
				data: {
						categories: [
						0
						] as Array<number>,
						attributes: [
						{
								attribute_id: 0,
								attribute_value_id: 0,
								custom_value: {
									raw_value: ``,
									unit: ``
								} as CustomValue | undefined,
								attribute_values: [
								{
									raw_value: ``,
									unit: ``,
									value_id: 0
								}
								] as Array<AttributeValue>
							} as AttributePayload | undefined
						] as Array<AttributePayload | undefined>
					} as ShopeeAttribute | undefined
			} as CreateAttributePayloadShopeeAttribute ,
		response: {
				err_msg: ``,
				data: {
						categories: [
						0
						] as Array<number>,
						attributes: [
						{
								attribute_id: 0,
								attribute_value_id: 0,
								custom_value: {
									raw_value: ``,
									unit: ``
								} as CustomValue | undefined,
								attribute_values: [
								{
									raw_value: ``,
									unit: ``,
									value_id: 0
								}
								] as Array<AttributeValue>
							} as AttributePayload | undefined
						] as Array<AttributePayload | undefined>
					} as ShopeeAttribute | undefined
			} as CreateAttributeResShopeeAttribute 
	},
	PostPdcsourceCollectionCreate: {
		url: "pdcsource/collection/create" as const,
		method: "POST" as const,
		query: undefined,
		body: {
				id: 0,
				name: ``,
				created_at: `2021-12-01T07:00:00+07:00`,
				updated_at: `2021-12-01T07:00:00+07:00`
			} as Collection ,
		response: {
				err_msg: ``,
				id: 0,
				name: ``,
				created_at: `2021-12-01T07:00:00+07:00`,
				updated_at: `2021-12-01T07:00:00+07:00`
			} as CollectionCreateRes 
	},
	GetPdcsourceCollectionList: {
		url: "pdcsource/collection/list" as const,
		method: "GET" as const,
		query: {
				page: 0,
				limit: 0
			} as PaginationQuery ,
		body: {},
		response: {
				err_msg: ``,
				data: [
				{
						id: 0,
						name: ``,
						created_at: `2021-12-01T07:00:00+07:00`,
						updated_at: `2021-12-01T07:00:00+07:00`,
						count: 0
					} as CollectionItem | undefined
				] as Array<CollectionItem | undefined>,
				page: 0,
				limit: 0,
				count: 0
			} as CollectionResList 
	},
	GetPdcsourceCollectionItem: {
		url: "pdcsource/collection/item" as const,
		method: "GET" as const,
		query: {
				col_id: 0
			} as CollectionItemPayload ,
		body: {},
		response: {
				id: 0,
				name: ``,
				created_at: `2021-12-01T07:00:00+07:00`,
				updated_at: `2021-12-01T07:00:00+07:00`
			} as Collection 
	},
	DeletePdcsourceCollection: {
		url: "pdcsource/collection" as const,
		method: "DELETE" as const,
		query: undefined,
		body: {
				collection_ids: [
				0
				] as Array<number>
			} as ColDeletePayload ,
		response: {} as any
	},
	GetPdcsourceProductList: {
		url: "pdcsource/product/list" as const,
		method: "GET" as const,
		query: {
				page: 0,
				limit: 0,
				coll_id: 0,
				status: ``
			} as ProductListQuery ,
		body: {},
		response: {
				err_msg: ``,
				data: [
				{
						id: 0,
						as_draft: false,
						image_preview: ``,
						image_collection_path: ``,
						count_image: 0,
						title: ``,
						price: 0,
						desc: ``,
						use_variant: false as boolean | undefined,
						use_markup: ``,
						weight: 0,
						stock: 0,
						attribute: [
						{
								id: 0,
								product_id: 0,
								attribute_type: ``,
								data: ``
							} as AttributeProduct | undefined
						] as Array<AttributeProduct | undefined>,
						field_spin: [
						{
								id: 0,
								product_id: 0,
								field_type: ``,
								use_spin: false,
								spin_text: ``,
								use_once_text: false,
								once_text: [
								{
										id: 0,
										field_config_id: 0,
										text: ``
									} as UseOnceText | undefined
								] as Array<UseOnceText | undefined>
							} as FieldConfig | undefined
						] as Array<FieldConfig | undefined>,
						variant_option: [
						{
								id: 0,
								product_id: 0,
								name: ``,
								option: [
								``
								] as Array<string>
							} as VariantOption | undefined
						] as Array<VariantOption | undefined>,
						variant_image: [
						{
								id: 0,
								product_id: 0,
								option_name: ``,
								image_preview: ``,
								image_collection_path: ``,
								count_image: 0
							} as VariantImage | undefined
						] as Array<VariantImage | undefined>,
						variant: [
						{
								id: 0,
								product_id: 0,
								names: [
								``
								] as Array<string>,
								values: [
								``
								] as Array<string>,
								price: 0,
								stock: 0
							} as Variant | undefined
						] as Array<Variant | undefined>,
						collection: [
						{
								id: 0,
								name: ``,
								created_at: `2021-12-01T07:00:00+07:00`,
								updated_at: `2021-12-01T07:00:00+07:00`
							} as Collection | undefined
						] as Array<Collection | undefined>,
						watermark: {
							id: 0,
							product_id: 0,
							text: ``,
							font_location: ``,
							pos_x: ``,
							pos_y: ``,
							use_watermark: false
						} as WatermarkConfig | undefined,
						frame_config: {
							id: 0,
							product_id: 0,
							use_frame: false,
							frame_location: ``
						} as FrameConfig | undefined,
						last_error: ``,
						map: [
						{
								id: 0,
								map_id: ``,
								product_id: 0,
								url_product: ``
							} as ProductMap | undefined
						] as Array<ProductMap | undefined>
					} as ManualProduct | undefined
				] as Array<ManualProduct | undefined>,
				page: 0,
				limit: 0,
				count: 0
			} as ProductListRes 
	},
	PostPdcsourceProductCreate: {
		url: "pdcsource/product/create" as const,
		method: "POST" as const,
		query: undefined,
		body: {
				coll_id: 0
			} as ProductCreatePayload ,
		response: {
				err_msg: ``
			} as ApiResponse 
	},
	DeletePdcsourceProduct: {
		url: "pdcsource/product" as const,
		method: "DELETE" as const,
		query: undefined,
		body: {
				ids: [
				0
				] as Array<number>
			} as DeleteProductPayload ,
		response: {
				err_msg: ``
			} as ApiResponse 
	},
	GetPdcsourceProductItem: {
		url: "pdcsource/product/item" as const,
		method: "GET" as const,
		query: {
				product_id: 0
			} as ItemQuery ,
		body: {},
		response: {
				err_msg: ``,
				data: {
						id: 0,
						as_draft: false,
						image_preview: ``,
						image_collection_path: ``,
						count_image: 0,
						title: ``,
						price: 0,
						desc: ``,
						use_variant: false as boolean | undefined,
						use_markup: ``,
						weight: 0,
						stock: 0,
						attribute: [
						{
								id: 0,
								product_id: 0,
								attribute_type: ``,
								data: ``
							} as AttributeProduct | undefined
						] as Array<AttributeProduct | undefined>,
						field_spin: [
						{
								id: 0,
								product_id: 0,
								field_type: ``,
								use_spin: false,
								spin_text: ``,
								use_once_text: false,
								once_text: [
								{
										id: 0,
										field_config_id: 0,
										text: ``
									} as UseOnceText | undefined
								] as Array<UseOnceText | undefined>
							} as FieldConfig | undefined
						] as Array<FieldConfig | undefined>,
						variant_option: [
						{
								id: 0,
								product_id: 0,
								name: ``,
								option: [
								``
								] as Array<string>
							} as VariantOption | undefined
						] as Array<VariantOption | undefined>,
						variant_image: [
						{
								id: 0,
								product_id: 0,
								option_name: ``,
								image_preview: ``,
								image_collection_path: ``,
								count_image: 0
							} as VariantImage | undefined
						] as Array<VariantImage | undefined>,
						variant: [
						{
								id: 0,
								product_id: 0,
								names: [
								``
								] as Array<string>,
								values: [
								``
								] as Array<string>,
								price: 0,
								stock: 0
							} as Variant | undefined
						] as Array<Variant | undefined>,
						collection: [
						{
								id: 0,
								name: ``,
								created_at: `2021-12-01T07:00:00+07:00`,
								updated_at: `2021-12-01T07:00:00+07:00`
							} as Collection | undefined
						] as Array<Collection | undefined>,
						watermark: {
							id: 0,
							product_id: 0,
							text: ``,
							font_location: ``,
							pos_x: ``,
							pos_y: ``,
							use_watermark: false
						} as WatermarkConfig | undefined,
						frame_config: {
							id: 0,
							product_id: 0,
							use_frame: false,
							frame_location: ``
						} as FrameConfig | undefined,
						last_error: ``,
						map: [
						{
								id: 0,
								map_id: ``,
								product_id: 0,
								url_product: ``
							} as ProductMap | undefined
						] as Array<ProductMap | undefined>
					} as ManualProduct | undefined
			} as ItemRes 
	},
	GetPdcsourceEditNew: {
		url: "pdcsource/edit/new" as const,
		method: "GET" as const,
		query: {
				coll_id: 0
			} as NewProductQuery ,
		body: {},
		response: {
				err_msg: ``,
				data: {
						id: 0,
						as_draft: false,
						image_preview: ``,
						image_collection_path: ``,
						count_image: 0,
						title: ``,
						price: 0,
						desc: ``,
						use_variant: false as boolean | undefined,
						use_markup: ``,
						weight: 0,
						stock: 0,
						attribute: [
						{
								id: 0,
								product_id: 0,
								attribute_type: ``,
								data: ``
							} as AttributeProduct | undefined
						] as Array<AttributeProduct | undefined>,
						field_spin: [
						{
								id: 0,
								product_id: 0,
								field_type: ``,
								use_spin: false,
								spin_text: ``,
								use_once_text: false,
								once_text: [
								{
										id: 0,
										field_config_id: 0,
										text: ``
									} as UseOnceText | undefined
								] as Array<UseOnceText | undefined>
							} as FieldConfig | undefined
						] as Array<FieldConfig | undefined>,
						variant_option: [
						{
								id: 0,
								product_id: 0,
								name: ``,
								option: [
								``
								] as Array<string>
							} as VariantOption | undefined
						] as Array<VariantOption | undefined>,
						variant_image: [
						{
								id: 0,
								product_id: 0,
								option_name: ``,
								image_preview: ``,
								image_collection_path: ``,
								count_image: 0
							} as VariantImage | undefined
						] as Array<VariantImage | undefined>,
						variant: [
						{
								id: 0,
								product_id: 0,
								names: [
								``
								] as Array<string>,
								values: [
								``
								] as Array<string>,
								price: 0,
								stock: 0
							} as Variant | undefined
						] as Array<Variant | undefined>,
						collection: [
						{
								id: 0,
								name: ``,
								created_at: `2021-12-01T07:00:00+07:00`,
								updated_at: `2021-12-01T07:00:00+07:00`
							} as Collection | undefined
						] as Array<Collection | undefined>,
						watermark: {
							id: 0,
							product_id: 0,
							text: ``,
							font_location: ``,
							pos_x: ``,
							pos_y: ``,
							use_watermark: false
						} as WatermarkConfig | undefined,
						frame_config: {
							id: 0,
							product_id: 0,
							use_frame: false,
							frame_location: ``
						} as FrameConfig | undefined,
						last_error: ``,
						map: [
						{
								id: 0,
								map_id: ``,
								product_id: 0,
								url_product: ``
							} as ProductMap | undefined
						] as Array<ProductMap | undefined>
					} as ManualProduct | undefined
			} as NewProductResponse 
	},
	PutPdcsourceEditPublish: {
		url: "pdcsource/edit/publish" as const,
		method: "PUT" as const,
		query: undefined,
		body: {
				product_id: 0
			} as CreatePayload ,
		response: {
				err_msg: ``
			} as ApiResponse 
	},
	PostPdcsourceEditSetBasic: {
		url: "pdcsource/edit/set_basic" as const,
		method: "POST" as const,
		query: undefined,
		body: {
				product_id: 0,
				image_preview: ``,
				image_collection_path: ``,
				count_image: 0,
				title: ``,
				price: 0,
				desc: ``,
				use_variant: false as boolean | undefined,
				use_markup: ``,
				weight: 0,
				stock: 0
			} as BasicUpdatePayload ,
		response: {
				err_msg: ``
			} as ApiResponse 
	},
	PostPdcsourceEditVariationUpdate: {
		url: "pdcsource/edit/variation_update" as const,
		method: "POST" as const,
		query: undefined,
		body: {
				product_id: 0,
				variant_option: [
				{
						id: 0,
						product_id: 0,
						name: ``,
						option: [
						``
						] as Array<string>
					} as VariantOption | undefined
				] as Array<VariantOption | undefined>,
				variant_image: [
				{
						id: 0,
						product_id: 0,
						option_name: ``,
						image_preview: ``,
						image_collection_path: ``,
						count_image: 0
					} as VariantImage | undefined
				] as Array<VariantImage | undefined>,
				variant: [
				{
						id: 0,
						product_id: 0,
						names: [
						``
						] as Array<string>,
						values: [
						``
						] as Array<string>,
						price: 0,
						stock: 0
					} as Variant | undefined
				] as Array<Variant | undefined>
			} as UpdateVariationPayload ,
		response: {
				err_msg: ``
			} as ApiResponse 
	},
	PostPdcsourceEditFieldConfig: {
		url: "pdcsource/edit/field_config" as const,
		method: "POST" as const,
		query: undefined,
		body: {
				product_id: 0,
				field_spin: [
				{
						id: 0,
						product_id: 0,
						field_type: ``,
						use_spin: false,
						spin_text: ``,
						use_once_text: false,
						once_text: [
						{
								id: 0,
								field_config_id: 0,
								text: ``
							} as UseOnceText | undefined
						] as Array<UseOnceText | undefined>
					} as FieldConfig | undefined
				] as Array<FieldConfig | undefined>
			} as UpdateFieldConfigPayload ,
		response: {
				err_msg: ``
			} as ApiResponse 
	},
	GetPdcsourceSpinListText: {
		url: "pdcsource/spin/list_text" as const,
		method: "GET" as const,
		query: {
				field_config_id: 0
			} as SpinConfigQuery ,
		body: {},
		response: {
				err_msg: ``,
				count: 0,
				data: [
				{
						id: 0,
						field_config_id: 0,
						text: ``
					} as UseOnceText | undefined
				] as Array<UseOnceText | undefined>
			} as ListTextRes 
	},
	PostPdcsourceSpinAddText: {
		url: "pdcsource/spin/add_text" as const,
		method: "POST" as const,
		query: undefined,
		body: {
				field_config_id: 0,
				data: [
				``
				] as Array<string>
			} as AddTextPayload ,
		response: {
				err_msg: ``
			} as ApiResponse 
	},
	PostPdcsourceSpinFieldConfig: {
		url: "pdcsource/spin/field_config" as const,
		method: "POST" as const,
		query: undefined,
		body: {
				product_id: 0,
				field_type: ``
			} as CreateFieldConfigPayload ,
		response: {
				err_msg: ``,
				data: {
						id: 0,
						product_id: 0,
						field_type: ``,
						use_spin: false,
						spin_text: ``,
						use_once_text: false,
						once_text: [
						{
								id: 0,
								field_config_id: 0,
								text: ``
							} as UseOnceText | undefined
						] as Array<UseOnceText | undefined>
					} as FieldConfig | undefined
			} as CreateFieldConfigRes 
	},
	DeletePdcsourceSpinFieldConfig: {
		url: "pdcsource/spin/field_config" as const,
		method: "DELETE" as const,
		query: undefined,
		body: {
				id: 0
			} as DeleteFieldConfigPayload ,
		response: {
				err_msg: ``
			} as ApiResponse 
	},
	PostPdcsourceFsCheck: {
		url: "pdcsource/fs/check" as const,
		method: "POST" as const,
		query: undefined,
		body: {
				path: ``
			} as PathCheckPayload ,
		response: {
				err_msg: ``,
				is_abs: false,
				Absolute: ``,
				exist: false
			} as PathCheckResponse 
	}
}

export type Fn<T> = (a: T) => void;

export type SendOptions<Data, Query, Err = Error> = {
    onSuccess?: Fn<Data>;
    onError?: Fn<Err>;
    query?: Query;
};

export type ClientReturn<Data, Query, Err = Error> = {
    pending: boolean;
    data: MaybeNull<Data>;
    error: MaybeNull<Err>;
    send: Fn<SendOptions<Data, Query, Err>>;
};

export type Clients = typeof clients;
export type Target = keyof Clients;

export function useQuery<
    K extends Target,
    R extends Clients[K]["response"],
    Q extends Clients[K]["query"]
>(action: K, options?: SendOptions<R, Q>): ClientReturn<R, Q> {
    const uri = clients[action].url;
    const method = clients[action].method;
    const queryOptions = options;

    const [pending, setPending] = useState(false);
    const [data, setData] = useState<MaybeNull<R>>(null);
    const [error, setError] = useState<MaybeNull<Error>>(null);

    async function send(options: SendOptions<R, Q> | undefined = queryOptions) {
        setPending(true);

        const query = options?.query;

        try {
            const { data } = await axios({
                method,
                url: uri,
                ...(query
                    ? {
                          params: query
                      }
                    : {})
            });

            options?.onSuccess?.(data);
            setData(data);
            setError(null);
        } catch (e) {
            options?.onError?.(e as any);
            setError(e as any);
            setData(null);
        } finally {
            setPending(false);
        }
    }

    return {
        data,
        error,
        pending,
        send
    };
}
