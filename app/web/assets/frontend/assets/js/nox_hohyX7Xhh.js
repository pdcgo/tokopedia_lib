import{c as k}from"./nox_kgoAykAyg.js";import{r}from"./nox_Xoteyeg77.js";const n={GetLauncherV1RunGrabShopee:{url:"launcher/v1/run_grab_shopee",method:"GET",query:{base:""},body:{},response:{use_csv:!1,empty_csv:!1}},GetLauncherV1RunGrabTokopedia:{url:"launcher/v1/run_grab_tokopedia",method:"GET",query:{base:""},body:{},response:{deprecated:!1,use_csv:!1,empty_csv:!1}},GetLauncherV1RunGrabJakmall:{url:"launcher/v1/run_grab_jakmall",method:"GET",query:{base:""},body:{},response:{use_csv:!1,empty_csv:!1}},GetLegacyApiDataspin:{url:"legacy/api/dataspin",method:"GET",query:{name:""},body:{},response:[{name:"",data:[""]}]},PostLegacyApiDataspin:{url:"legacy/api/dataspin",method:"POST",query:void 0,body:{name:"",data:[""]},response:{errcode:0,message:"",status:""}},DeleteLegacyApiDataspin:{url:"legacy/api/dataspin",method:"DELETE",query:{name:""},body:{},response:{errcode:0,message:"",status:""}},GetLegacyApiSettingSpin:{url:"legacy/api/settingSpin",method:"GET",query:void 0,body:{},response:{data:{smin:0,smax:0,merek_ins_t:!1,title:"",desc:""},errcode:0,titlePool:[{name:"",data:""}]}},PostLegacyApiSettingSpin:{url:"legacy/api/settingSpin",method:"POST",query:void 0,body:{smin:0,smax:0,merek_ins_t:!1,title:"",desc:""},response:{errcode:0,message:"",status:""}},PostLegacyApiConfigSettingSpin:{url:"legacy/api/config/settingSpin",method:"POST",query:void 0,body:{name:"",titlePool:[{name:"",data:""}]},response:{errcode:0,message:"",status:""}},GetLegacyV1ProductNamespaceAll:{url:"legacy/v1/product/namespace_all",method:"GET",query:{is_public:!1,kota:"",marketplace:"",namespace:"",pmax:0,pmin:0,use_empty_city:!1},body:{},response:[{count:0,price_min:0,price_max:0,name:""}]},GetLegacyV1ProductPriceRange:{url:"legacy/v1/product/price_range",method:"GET",query:{is_public:!1,kota:"",marketplace:"",namespace:"",pmax:0,pmin:0,use_empty_city:!1,range_price:0},body:{},response:[{_id:[0],count:0}]},GetLegacyV1ProductCategory:{url:"legacy/v1/product/category",method:"GET",query:{is_public:!1,kota:"",marketplace:"",namespace:"",pmax:0,pmin:0,use_empty_city:!1},body:{},response:[{_id:0,price_min:0,price_max:0,count:0,name:[""]}]},GetLegacyV1ProductKota:{url:"legacy/v1/product/kota",method:"GET",query:{is_public:!1,kota:"",marketplace:"",namespace:"",pmax:0,pmin:0,use_empty_city:!1},body:{},response:[{_id:"",price_min:0,price_max:0,count:0}]},GetLegacyV1ProductDelete:{url:"legacy/v1/product/delete",method:"GET",query:{is_public:!1,kota:"",marketplace:"",namespace:"",pmax:0,pmin:0,use_empty_city:!1},body:{},response:{errcode:0,message:"",status:""}},PostLegacyApiDeleteItem:{url:"legacy/api/deleteItem",method:"POST",query:{is_public:!1,kota:"",marketplace:"",namespace:"",pmax:0,pmin:0,use_empty_city:!1},body:[0],response:{errcode:0,message:"",status:""}},PostLegacyV1ProductCategstatToCsv:{url:"legacy/v1/product/categstat_to_csv",method:"POST",query:void 0,body:[{parent_category:0,catid:0,parent_display_name:"",display_name:"",status:""}],response:{errcode:0,message:"",status:""}},GetLegacyV1ProductResync:{url:"legacy/v1/product/resync",method:"GET",query:{is_public:!1,kota:"",marketplace:"",namespace:"",pmax:0,pmin:0,use_empty_city:!1},body:{},response:{errcode:0,message:"",status:""}},PostLegacyV1ProductRenameNamespace:{url:"legacy/v1/product/rename_namespace",method:"POST",query:void 0,body:{marketplace:"",namespace:"",update_namespace:""},response:{errcode:0,message:"",status:""}},GetLegacyApiListMarkup:{url:"legacy/api/listMarkup",method:"GET",query:void 0,body:{},response:{errcode:0,message:"",status:"",data:[""]}},PostLegacyApiAddMarkup:{url:"legacy/api/addMarkup",method:"POST",query:void 0,body:{data:[{mark:"",type:"",range:{},up:null}],fix_mark:0,name:""},response:{errcode:0,message:"",status:""}},GetLegacyApiMarkup:{url:"legacy/api/markup",method:"GET",query:{name:""},body:{},response:{data:[{mark:"",type:"",range:{},up:null}],fix_mark:0,name:"",errorcode:0}},PostLegacyApiMarkup:{url:"legacy/api/markup",method:"POST",query:{name:""},body:{data:[{mark:"",type:"",range:{},up:null}],fix_mark:0,name:""},response:{errcode:0,message:"",status:""}},PostLegacyApiDeleteMarkup:{url:"legacy/api/deleteMarkup",method:"POST",query:void 0,body:[""],response:{errcode:0,message:"",status:""}},GetLegacyApiShopeeShipping:{url:"legacy/api/shopee_shipping",method:"GET",query:void 0,body:{},response:[{positionid:0,name:"",channelids:[""],display_name:"",item_tag_ids:[0]}]},GetLegacyShopeeManifest:{url:"legacy/shopee/manifest",method:"GET",query:void 0,body:{},response:{category:[{id:0,name:"",display_name:"",parent_id:0,has_active_children:!1,has_children:!1,region_setting:{enable_size_chart:!1,low_stock_value:0,dimension_mandatory:!1},is_prohibit:!1,children:[{id:0,name:"",display_name:"",parent_id:0,has_active_children:!1,has_children:!1,region_setting:{enable_size_chart:!1,low_stock_value:0,dimension_mandatory:!1},is_prohibit:!1,children:[]}],chain_name:[""],chain_ids:[0]}],public_category_repo:[{main:{catid:0,display_name:"",name:"",image:"",is_adult:{},parent_category:0,sort_weight:0,block_buyer_platform:{}},sub:[{catid:0,display_name:"",name:"",image:"",is_adult:{},parent_category:0,sort_weight:0,block_buyer_platform:{},sub_sub:[{catid:0,display_name:"",image:"",block_buyer_platform:{}}]}]}]}},PostLegacyV4ShopeeWeightPredict:{url:"legacy/v4/shopee/weight/predict",method:"POST",query:void 0,body:{itemid:0,shopid:0},response:{harga:0,jarak:0}},GetLegacyV3PredictweightSave:{url:"legacy/v3/predictweight/save",method:"GET",query:{predict:0},body:{},response:{errcode:0,message:"",status:""}},GetLegacyV3PredictweightLoad:{url:"legacy/v3/predictweight/load",method:"GET",query:void 0,body:{},response:{predict_weight:0}},GetLegacyApiSettingGetCropSetting:{url:"legacy/api/setting/get/cropSetting",method:"GET",query:void 0,body:{},response:{name:"",data:null,errcode:0}},PostLegacyApiSettingUpdateCropSetting:{url:"legacy/api/setting/update/cropSetting",method:"POST",query:void 0,body:{name:"",data:null},response:{errcode:0,message:"",status:""}},GetLegacyApiConfigConcurentRequest:{url:"legacy/api/config/concurentRequest",method:"GET",query:void 0,body:{},response:{name:"",data:0}},PostLegacyApiConfigConcurentRequest:{url:"legacy/api/config/concurentRequest",method:"POST",query:void 0,body:{name:"",data:0},response:{errcode:0,message:"",status:""}},GetLegacyApiUpInterval:{url:"legacy/api/upInterval",method:"GET",query:void 0,body:{},response:{name:"",data:{uptmax:0,uptmin:0}}},PostLegacyApiUpInterval:{url:"legacy/api/upInterval",method:"POST",query:void 0,body:{uptmax:0,uptmin:0},response:{errcode:0,message:"",status:""}},GetLegacyApiConfigUpThread:{url:"legacy/api/config/upThread",method:"GET",query:void 0,body:{},response:{name:"",data:0}},PostLegacyApiConfigUpThread:{url:"legacy/api/config/upThread",method:"POST",query:void 0,body:{name:"",data:0},response:{errcode:0,message:"",status:""}},GetLegacyApiConfigUpMode:{url:"legacy/api/config/upMode",method:"GET",query:void 0,body:{},response:{name:"",data:""}},PostLegacyApiConfigUpMode:{url:"legacy/api/config/upMode",method:"POST",query:void 0,body:{name:"",data:""},response:{errcode:0,message:"",status:""}},GetLegacyApiConfigGrabFilter:{url:"legacy/api/config/grabFilter",method:"GET",query:void 0,body:{},response:{name:"",data:{penjualan:0,prosentase:0,stock:0,tokped_point:[0]}}},PostLegacyApiConfigGrabFilter:{url:"legacy/api/config/grabFilter",method:"POST",query:void 0,body:{name:"",data:{penjualan:0,prosentase:0,stock:0,tokped_point:[0]}},response:{errcode:0,message:"",status:""}},GetLegacyApiConfigLastLogin:{url:"legacy/api/config/last_login",method:"GET",query:void 0,body:{},response:{name:"",data:{active:!1,days:0}}},PostLegacyApiConfigLastLogin:{url:"legacy/api/config/last_login",method:"POST",query:void 0,body:{name:"",data:{active:!1,days:0}},response:{errcode:0,message:"",status:""}},GetLegacyApiConfigLastReview:{url:"legacy/api/config/lastReview",method:"GET",query:void 0,body:{},response:{name:"",data:{active:!1,days:0}}},PostLegacyApiConfigLastReview:{url:"legacy/api/config/lastReview",method:"POST",query:void 0,body:{name:"",data:{active:!1,days:0}},response:{errcode:0,message:"",status:""}},GetLegacyApiLimitGrab:{url:"legacy/api/limitGrab",method:"GET",query:void 0,body:{},response:{name:"",data:0}},PostLegacyApiLimitGrab:{url:"legacy/api/limitGrab",method:"POST",query:void 0,body:0,response:{errcode:0,message:"",status:""}},GetLegacyApiConfigShopeeGrabSetting:{url:"legacy/api/config/shopeeGrabSetting",method:"GET",query:void 0,body:{},response:{name:"",data:{by:"",locations:[""],official_mall:!1,price_max:0,price_min:0,rating_filter:0,shopee24:!1,shopee_verified:!1}}},PostLegacyApiConfigShopeeGrabSetting:{url:"legacy/api/config/shopeeGrabSetting",method:"POST",query:void 0,body:{name:"",data:{by:"",locations:[""],official_mall:!1,price_max:0,price_min:0,rating_filter:0,shopee24:!1,shopee_verified:!1}},response:{errcode:0,message:"",status:""}},GetLegacyShopeeFilterGrabber:{url:"legacy/shopee/filter_grabber",method:"GET",query:void 0,body:{},response:{product_created:{active:!1,max:0,min:0},shippings:[{channelids:[""],display_name:"",name:"",item_tag_ids:[0],positionid:0}]}},PutLegacyShopeeFilterGrabber:{url:"legacy/shopee/filter_grabber",method:"PUT",query:void 0,body:{product_created:{active:!1,max:0,min:0},shippings:[{channelids:[""],display_name:"",name:"",item_tag_ids:[0],positionid:0}]},response:{errcode:0,message:"",status:""}},GetLegacyShopeeConfigCrawler:{url:"legacy/shopee/config_crawler",method:"GET",query:void 0,body:{},response:{username:"",password:"",email:"",email_password:""}},PutLegacyShopeeConfigCrawler:{url:"legacy/shopee/config_crawler",method:"PUT",query:void 0,body:{username:"",password:"",email:"",email_password:""},response:{username:"",password:"",email:"",email_password:""}},GetLegacyApiSettingGrab:{url:"legacy/api/settingGrab",method:"GET",query:void 0,body:{},response:{errcode:0,data:{name:"",data:{pmin:0,pmax:0,ob:"",rt:"",condition:"",fcity:[""],goldmerchant:!1,official:!1,shipping:[""],preorder:!1}}}},PostLegacyApiSettingGrab:{url:"legacy/api/settingGrab",method:"POST",query:void 0,body:{pmin:0,pmax:0,ob:"",rt:"",condition:"",fcity:[""],goldmerchant:!1,official:!1,shipping:[""],preorder:!1},response:{errcode:0,message:"",status:""}},GetLegacyApiConfigCloudinary:{url:"legacy/api/config/cloudinary",method:"GET",query:void 0,body:{},response:{name:"",data:{active:!1,url:""}}},PostLegacyApiConfigCloudinary:{url:"legacy/api/config/cloudinary",method:"POST",query:void 0,body:{name:"",data:{active:!1,url:""}},response:{errcode:0,message:"",status:""}},GetLegacyApiConfigGetRandom:{url:"legacy/api/config/get_random",method:"GET",query:void 0,body:{},response:{name:"",data:!1}},PostLegacyApiConfigGetRandom:{url:"legacy/api/config/get_random",method:"POST",query:void 0,body:{name:"",data:!1},response:{errcode:0,message:"",status:""}},GetLegacyApiConfigRndAttribute:{url:"legacy/api/config/rnd_attribute",method:"GET",query:void 0,body:{},response:{name:"",data:{active:!1,force_tidakada:!1}}},PostLegacyApiConfigRndAttribute:{url:"legacy/api/config/rnd_attribute",method:"POST",query:void 0,body:{name:"",data:{active:!1,force_tidakada:!1}},response:{errcode:0,message:"",status:""}},GetLegacyApiConfigUserSameResource:{url:"legacy/api/config/userSameResource",method:"GET",query:void 0,body:{},response:{name:"",data:!1}},PostLegacyApiConfigUserSameResource:{url:"legacy/api/config/userSameResource",method:"POST",query:void 0,body:{name:"",data:!1},response:{errcode:0,message:"",status:""}},GetLegacyApiSettingGetBlacklistTitle:{url:"legacy/api/setting/get/blacklistTitle",method:"GET",query:void 0,body:{},response:{data:[""],errcode:0,name:""}},PostLegacyApiSettingUpdateBlacklistTitle:{url:"legacy/api/setting/update/blacklistTitle",method:"POST",query:void 0,body:{data:[""],name:""},response:{errcode:0,message:"",status:""}},GetLegacyV3SettingDefault:{url:"legacy/v3/setting/default",method:"GET",query:void 0,body:{},response:{_id:"",name:"",use_price_discount:!1,blacklist_username:{active:!1,shopee:{data:[""],filename:""}}}},PostLegacyV3SettingDefault:{url:"legacy/v3/setting/default",method:"POST",query:void 0,body:{_id:"",name:"",use_price_discount:!1,blacklist_username:{active:!1,shopee:{data:[""],filename:""}}},response:{errcode:0,message:"",status:""}},GetLegacyApiConfigFilterTitle:{url:"legacy/api/config/filterTitle",method:"GET",query:void 0,body:{},response:{data:[""],errcode:0,name:""}},PostLegacyApiConfigFilterTitle:{url:"legacy/api/config/filterTitle",method:"POST",query:void 0,body:{data:[""],name:""},response:{errcode:0,message:"",status:""}},GetLegacyApiFilterWord:{url:"legacy/api/filterWord",method:"GET",query:{name:"",makedefault:!1},body:{},response:{data:[""],errcode:0,name:""}},PostLegacyApiFilterWord:{url:"legacy/api/filterWord",method:"POST",query:void 0,body:{data:[""],name:""},response:{errcode:0,message:"",status:""}},PostLegacyApiFilterWordMakedefault:{url:"legacy/api/filterWord/makedefault",method:"POST",query:void 0,body:{},response:{errcode:0,message:"",status:""}},GetLegacyApiBackupAkun:{url:"legacy/api/backupAkun",method:"GET",query:void 0,body:{},response:{errcode:0,message:"",status:""}},PostLegacyApiUser:{url:"legacy/api/user",method:"POST",query:void 0,body:{data:{_id:"",user:"",name:"",pass:"",password:"",active:!1,limit_upload:0,count_upload:0,namespace:"",water:"",shopee_categ:[""],tokped_categ:[""],type:"",markup:"",hastag:"",last_up:0,polatitle:"",is_running_upload:!1}},response:{errcode:0,message:"",status:""}},DeleteLegacyApiUser:{url:"legacy/api/user",method:"DELETE",query:void 0,body:{data:[""]},response:{errcode:0,message:"",status:""}},GetLegacyApiAkuns:{url:"legacy/api/akuns",method:"GET",query:{start:0,limit:0,active:"",last:0,search:"",sort:"",reverse:0},body:{},response:{akuns:[{_id:"",user:"",name:"",pass:"",password:"",active:!1,limit_upload:0,count_upload:0,namespace:"",water:"",shopee_categ:[""],tokped_categ:[""],type:"",markup:"",hastag:"",last_up:0,polatitle:"",is_running_upload:!1}],count:0,errcode:0}},GetLegacyV1AkunReset:{url:"legacy/v1/akun/reset",method:"GET",query:void 0,body:{},response:{errcode:0,message:"",status:""}},PutLegacyApiUser:{url:"legacy/api/user",method:"PUT",query:void 0,body:{data:{_id:"",user:"",name:"",pass:"",password:"",active:!1,limit_upload:0,count_upload:0,namespace:"",water:"",shopee_categ:[""],tokped_categ:[""],type:"",markup:"",hastag:"",last_up:0,polatitle:"",is_running_upload:!1}},response:{errcode:0,message:"",status:""}},GetLegacyShopeeConfigUpload:{url:"legacy/shopee/config_upload",method:"GET",query:void 0,body:{},response:{shipping:[{channelid:0,enabled:!1}],use_custom_shipping:!1}},PutLegacyShopeeConfigUpload:{url:"legacy/shopee/config_upload",method:"PUT",query:void 0,body:{shipping:[{channelid:0,enabled:!1}],use_custom_shipping:!1},response:{errcode:0,message:"",status:""}},GetLegacyV4ShopeeUploadConfigProductShippingList:{url:"legacy/v4/shopee/upload_config/product_shipping_list",method:"GET",query:void 0,body:{},response:[{channel_id:0,name:"",display_name:"",name_key:"",flag:"",service_type:"",default_price:"",min_default_price:"",max_default_price:"",limits:{item_min_weight:0},sizes:[{}],cover_shipping_fee:!1,enabled:!1,item_flag:"",price:"",size:0,size_id:0,save_into_item:!1,volumetric_factor:0,conflicting_enabled_channels:{},is_mask_channel:!1,parent_channel_id:0,item_exclusive_channels:[0],related_enabled_channels:[{}],related_disabled_channels:[{}]}]},PostLegacyApiTool:{url:"legacy/api/tool",method:"POST",query:void 0,body:{data:""},response:{errcode:0,message:"",status:""}},GetLegacyV3ConfigurationGet:{url:"legacy/v3/configuration/get",method:"GET",query:void 0,body:{},response:{captcha:{user:"",pass:""},database:{DB_URI:"",DB_NAME:""},lisensi:{email:"",pwd:""},suplier_csv:!1,session_storage_source:"",session_storage_local_location:"",thread_count:0,tor_directory:"",tor_print_ip:!1,tor_proxy:!1}},GetLegacyApiConfigDelete:{url:"legacy/api/config_delete",method:"GET",query:void 0,body:{},response:{akun:"",awaltanggal:"2021-12-01T07:00:00+07:00",blokir:!1,delete:0,diarsipkan:!1,diperiksa:!1,sold:0,tanggal:"2021-12-01T07:00:00+07:00",view:0}},PostLegacyApiConfigDelete:{url:"legacy/api/config_delete",method:"POST",query:void 0,body:{akun:"",awaltanggal:"2021-12-01T07:00:00+07:00",blokir:!1,delete:0,diarsipkan:!1,diperiksa:!1,sold:0,tanggal:"2021-12-01T07:00:00+07:00",view:0},response:{errcode:0,message:"",status:""}},GetLegacyApiConfigDeleteProduct:{url:"legacy/api/config/deleteProduct",method:"GET",query:void 0,body:{},response:{name:"",data:{fil_category:!1,fil_harga:!1,fil_keyword:!1,category:[[0]],harga:{min:0,max:0},keyword:""}}},PostLegacyApiConfigDeleteProduct:{url:"legacy/api/config/deleteProduct",method:"POST",query:void 0,body:{name:"",data:{fil_category:!1,fil_harga:!1,fil_keyword:!1,category:[[0]],harga:{min:0,max:0},keyword:""}},response:{errcode:0,message:"",status:""}},GetLegacyV1TaskerAll:{url:"legacy/v1/tasker/all",method:"GET",query:void 0,body:{},response:[{_id:"",toko_username:"",mode:"",marketplace:"",product_url:"",namespace:"",tokped_categ:[""],use_filter:!1,keyword:"",shopee_categ:{parent_category:0,catid:0,parent_display_name:"",display_name:"",is_collection:0},jakmall_categs:[""]}]},PostLegacyV1TaskerSave:{url:"legacy/v1/tasker/save",method:"POST",query:void 0,body:[{_id:"",toko_username:"",mode:"",marketplace:"",product_url:"",namespace:"",tokped_categ:[""],use_filter:!1,keyword:"",shopee_categ:{parent_category:0,catid:0,parent_display_name:"",display_name:"",is_collection:0},jakmall_categs:[""]}],response:{errcode:0,message:"",status:""}},DeleteLegacyV1Taskerid:{url:"legacy/v1/tasker/:id",method:"DELETE",query:void 0,body:{},response:{errcode:0,message:"",status:""}},GetLegacyApiCategoryDumpCsv:{url:"legacy/api/category/dump_csv",method:"GET",query:{mp:""},body:{},response:{errcode:0,message:"",status:""}},GetLegacyV4ShopeeSearchFilterDumpFacetCategory:{url:"legacy/v4/shopee/search_filter/dump_facet_category",method:"GET",query:void 0,body:{},response:{}},GetLegacyV4ShopeeSearchFilterDumpFacetCategoryLink:{url:"legacy/v4/shopee/search_filter/dump_facet_category_link",method:"GET",query:void 0,body:{},response:{errcode:0,message:"",status:""}},GetLegacyV1CategoryMapperTokopedia:{url:"legacy/v1/category_mapper/tokopedia",method:"GET",query:{action:""},body:{},response:{}},GetLegacyApiHastag:{url:"legacy/api/hastag",method:"GET",query:void 0,body:{},response:{name:"",data:[""]}},PostLegacyApiHastag:{url:"legacy/api/hastag",method:"POST",query:void 0,body:{name:"",data:[""]},response:{errcode:0,message:"",status:""}},DeleteLegacyApiHastag:{url:"legacy/api/hastag",method:"DELETE",query:{name:""},body:{},response:{errcode:0,message:"",status:""}},GetLegacyApiConfigHastagLimit:{url:"legacy/api/config/hastagLimit",method:"GET",query:void 0,body:{},response:{name:"",data:null}},GetLegacyV1ExamplespinProducts:{url:"legacy/v1/examplespin/products",method:"GET",query:{harga:"",hastag:"",polatitle:""},body:{},response:[{name:"",price:0,price_untung:0,desc:"",markup_debug:{fix_harga:0,harga_asli:0,harga_up:0,markup_item_index:0,up_percent:0,up_price:0,up_fix:0}}]},GetLegacyV1ExamplespinTitle:{url:"legacy/v1/examplespin/title",method:"GET",query:{title:""},body:{},response:{text:""}},GetShopeeV5RunCheckbot:{url:"shopee/v5/run_checkbot",method:"GET",query:{base:"",cek:"",out:""},body:{},response:{}},GetShopeeV5RunCheckOrder:{url:"shopee/v5/run_check_order",method:"GET",query:{base:"",akun_limit:0,queue_size:0,filepath:"",output:""},body:{},response:{}},PostShopeeV5RunDeleteProduct:{url:"shopee/v5/run_delete_product",method:"POST",query:{base:"",report:""},body:{},response:{}},GetShopeeV5UpdaterAttribute:{url:"shopee/v5/updater_attribute",method:"GET",query:{base:""},body:{},response:{}},PutShopeeV5ProductExportSupplier:{url:"shopee/v5/product/export_supplier",method:"PUT",query:{namespace:""},body:{},response:{}},PutShopeeV5ProductExportUrl:{url:"shopee/v5/product/export_url",method:"PUT",query:{namespace:""},body:{},response:{}},GetShopeeV5FilterGrabLocation:{url:"shopee/v5/filter/grab_location",method:"GET",query:void 0,body:{},response:[""]},GetShopeeV5ConfigWeightRatio:{url:"shopee/v5/config/weight_ratio",method:"GET",query:void 0,body:{},response:{data:[{price_min:0,price_max:0,ratio:0}]}},PutShopeeV5ConfigWeightRatio:{url:"shopee/v5/config/weight_ratio",method:"PUT",query:void 0,body:{data:[{price_min:0,price_max:0,ratio:0}]},response:{data:[{price_min:0,price_max:0,ratio:0}]}},GetShopeeV5Attribute:{url:"shopee/v5/attribute",method:"GET",query:{cat_id:0},body:{},response:{exist:!1,attributes:[{attributeId:0,displayName:"",mandatory:!1,children:[{multiLang:[{language:"",value:""}],displayName:"",valueId:0,valueType:0}],attributeInfo:{attributeUnitList:[""],dateFormatType:0,formatType:0,inputType:0,inputValidationType:0,isNcc:!1,maxValueCount:0}}]}},PostShopeeV5QlobotShopeeImportCsv:{url:"shopee/v5/qlobot/shopee_import_csv",method:"POST",query:{namespace:"",per_item:0},body:{},response:{errcode:0,message:"",status:""}},GetTokopediaAkunList:{url:"tokopedia/akun/list",method:"GET",query:{start:0,limit:0,active:"",last:0,search:"",sort:"",reverse:0},body:{},response:{akuns:[{_id:"",user:"",name:"",pass:"",password:"",active:!1,limit_upload:0,count_upload:0,namespace:"",water:"",shopee_categ:[""],tokped_categ:[""],type:"",markup:"",hastag:"",last_up:0,polatitle:"",is_running_upload:!1}],count:0,errcode:0}},PostTokopediaAkunBulkAdd:{url:"tokopedia/akun/bulk_add",method:"POST",query:void 0,body:{data:[{username:"",password:"",secret:""}]},response:{msg:"",error:""}},PostTokopediaAkunUpdate:{url:"tokopedia/akun/update",method:"POST",query:void 0,body:{data:[{limit_upload:0,count_upload:0,active_upload:!1,lastup:0,in_upload:!1,last_error:"",username:"",password:"",secret:"",markup:"",spin:"",collection:"",hastag:"",title_pattern:""}]},response:{msg:"",error:""}},PostTokopediaAkunDelete:{url:"tokopedia/akun/delete",method:"POST",query:void 0,body:{data:[""]},response:{msg:"",error:""}},PutTokopediaAkunResetAllCount:{url:"tokopedia/akun/reset_all_count",method:"PUT",query:void 0,body:{},response:{msg:"",error:""}},PutTokopediaMapperMap:{url:"tokopedia/mapper/map",method:"PUT",query:void 0,body:[{shopee_id:0,tokopedia_id:0}],response:{msg:"",error:""}},GetTokopediaMapperMap:{url:"tokopedia/mapper/map",method:"GET",query:{collection:""},body:{},response:{data:[{shopee_id:0,tokopedia_id:0}]}},PutTokopediaMapperAutosuggest:{url:"tokopedia/mapper/autosuggest",method:"PUT",query:{collection:""},body:{},response:{}},GetTokopediaMapperAutosuggest:{url:"tokopedia/mapper/autosuggest",method:"GET",query:void 0,body:{},response:{status:""}},GetTokopediaMapperSetting:{url:"tokopedia/mapper/setting",method:"GET",query:void 0,body:{},response:{use_mapper:!1}},PutTokopediaMapperSetting:{url:"tokopedia/mapper/setting",method:"PUT",query:void 0,body:{use_mapper:!1},response:{use_mapper:!1}},GetTokopediaMapperCategory:{url:"tokopedia/mapper/category",method:"GET",query:{namespace:""},body:{},response:[{shopee_id:0,tokopedia_id:0,product_count:0,shopee_category_name:[""],tokopedia_category_name:[""]}]},GetTokopediaCategoryList:{url:"tokopedia/category/list",method:"GET",query:void 0,body:{},response:{data:{categoryAllListLite:{categories:[{id:0,name:"",url:"",children:[],__typename:""}],__typename:""}}}},PutTokopediaCategoryUpdateCategory:{url:"tokopedia/category/update_category",method:"PUT",query:void 0,body:{username:"",password:"",secret:""},response:{msg:"",error:""}},GetTokopediaDeleterSetting:{url:"tokopedia/deleter/setting",method:"GET",query:void 0,body:{},response:{data:{akun:"",awaltanggal:"2021-12-01T07:00:00+07:00",blokir:!1,delete:0,diarsipkan:!1,diperiksa:!1,sold:0,tanggal:"2021-12-01T07:00:00+07:00",view:0}}},PutTokopediaDeleterSetting:{url:"tokopedia/deleter/setting",method:"PUT",query:void 0,body:{akun:"",awaltanggal:"2021-12-01T07:00:00+07:00",blokir:!1,delete:0,diarsipkan:!1,diperiksa:!1,sold:0,tanggal:"2021-12-01T07:00:00+07:00",view:0},response:{data:{akun:"",awaltanggal:"2021-12-01T07:00:00+07:00",blokir:!1,delete:0,diarsipkan:!1,diperiksa:!1,sold:0,tanggal:"2021-12-01T07:00:00+07:00",view:0}}},GetTokopediaDumpCategoryDump:{url:"tokopedia/dump/category_dump",method:"GET",query:void 0,body:{},response:{errcode:0,message:"",status:""}},PutTokopediaCekbotRun:{url:"tokopedia/cekbot/run",method:"PUT",query:{base:"",cek:"",out:""},body:{},response:{}},PutTokopediaCheckVerifRun:{url:"tokopedia/check_verif/run",method:"PUT",query:{base:"",fname:""},body:{fname:""},response:{}},PutTokopediaDeleterRunDelete:{url:"tokopedia/deleter/run_delete",method:"PUT",query:{base:""},body:{},response:{}},PostTokopediaAutosubmitRun:{url:"tokopedia/autosubmit/run",method:"POST",query:void 0,body:{base_ktp:"",filename:""},response:{}},GetTokopediaFilterFcity:{url:"tokopedia/filter/fcity",method:"GET",query:void 0,body:{},response:[{Description:"",__typename:"",child:[{}],hexColor:"",icon:"",inputType:"",isNew:!1,isPopular:!1,key:"",name:"",totalData:"",valMax:"",valMin:"",value:""}]},GetTokopediaFilterCategories:{url:"tokopedia/filter/categories",method:"GET",query:void 0,body:{},response:[{Description:"",__typename:"",child:[{}],hexColor:"",icon:"",inputType:"",isNew:!1,isPopular:!1,key:"",name:"",totalData:"",valMax:"",valMin:"",value:""}]},GetTokopediaFilterShipping:{url:"tokopedia/filter/shipping",method:"GET",query:void 0,body:{},response:[{name:"",Description:"",key:"",icon:"",value:"",inputType:"",totalData:"",valMax:"",valMin:"",hexColor:"",child:[{}],isPopular:!1,isNew:!1,__typename:""}]},GetTokopediaEtalaseMapList:{url:"tokopedia/etalase_map/list",method:"GET",query:{namespace:""},body:{},response:{data:[{shopee_id:0,tokopedia_id:0,product_count:0,category_no_mapping:!1}]}},DeleteTokopediaEtalaseMapDelete:{url:"tokopedia/etalase_map/delete",method:"DELETE",query:{name:""},body:{},response:{}},PutTokopediaEtalaseMapUpdate:{url:"tokopedia/etalase_map/update",method:"PUT",query:void 0,body:[{etalase_name:"",category_id:0}],response:{msg:"",error:""}},GetTokopediaEtalaseMapListEtalase:{url:"tokopedia/etalase_map/list_etalase",method:"GET",query:void 0,body:{},response:[{etalase:"",cat_ids:[0]}]},GetTokopediaAttributeUpdaterAttribute:{url:"tokopedia/attribute/updater_attribute",method:"GET",query:{base:""},body:{},response:{}},GetTokopediaAttributeGetAttribute:{url:"tokopedia/attribute/get_attribute",method:"GET",query:{cat_id:0},body:{},response:{exist:!1,attributes:[{variant:"",sortOrder:0,values:[{id:0,name:"",selected:!1,__typename:""}],__typename:""}]}},PutTokopediaCekorderSaveConfig:{url:"tokopedia/cekorder/save_config",method:"PUT",query:void 0,body:{useDateRange:!1,startDate:"",endDate:"",useStatus:!1,statusKeys:[""]},response:{}},PutTokopediaCekorderRun:{url:"tokopedia/cekorder/run",method:"PUT",query:{base:"",fname:"",output:""},body:[{password:"",secret:"",username:""}],response:{}},GetJakmallCategoryList:{url:"jakmall/category/list",method:"GET",query:void 0,body:{},response:{msg:"",error:"",data:[{name:"",url:"",icon:"",color:"",children:[]}]}},GetJakmallCategoryDumpCsv:{url:"jakmall/category/dump_csv",method:"GET",query:void 0,body:{},response:{msg:"",error:""}},GetJakmallSearchFilterData:{url:"jakmall/search_filter/data",method:"GET",query:void 0,body:{},response:{msg:"",error:"",cities:[{label:"",value:""}],delivery_types:[""],sorts:[{label:"",value:""}]}},GetJakmallSearchFilter:{url:"jakmall/search_filter",method:"GET",query:void 0,body:{},response:{msg:"",error:"",data:{category:"",price_min:0,price_max:0,untung_paling_besar:0,in_stock:0,bulk_price:0,delivery_types:[""],cities:[""],sort:"",rating:0}}},PostJakmallSearchFilter:{url:"jakmall/search_filter",method:"POST",query:void 0,body:{category:"",price_min:0,price_max:0,untung_paling_besar:0,in_stock:0,bulk_price:0,delivery_types:[""],cities:[""],sort:"",rating:0},response:{msg:"",error:""}},GetJakmallCategoryMapperList:{url:"jakmall/category_mapper/list",method:"GET",query:{type:"",namespace:""},body:{},response:{msg:"",error:"",data:[{type:"",name:"",categs:[{id:0,name:""}],mapper_id:0,mapper_name:"",mapper_categs:[{id:0,name:""}],count:0}]}},PostJakmallCategoryMapperSave:{url:"jakmall/category_mapper/save",method:"POST",query:void 0,body:[{type:"",name:"",categs:[{id:0,name:""}],mapper_id:0,mapper_name:"",mapper_categs:[{id:0,name:""}]}],response:{msg:"",error:""}},PutJakmallCategoryMapperAutosuggest:{url:"jakmall/category_mapper/autosuggest",method:"PUT",query:{type:"",namespace:""},body:{},response:{msg:"",error:""}},GetUploadV6ManualToShopee:{url:"upload/v6/manual_to_shopee",method:"GET",query:{base:"",reset:!1,one_to_multi:!1,limit:0},body:{},response:{}},GetUploadV6ManualToTokopedia:{url:"upload/v6/manual_to_tokopedia",method:"GET",query:{base:"",use_mapper:!1,reset:!1,one_to_multi:!1,limit:0},body:{},response:{}},GetUploadV6ShopeeToShopee:{url:"upload/v6/shopee_to_shopee",method:"GET",query:{base:""},body:{},response:{}},GetUploadV6TokopediaToShopee:{url:"upload/v6/tokopedia_to_shopee",method:"GET",query:{base:"",use_mapper:!1},body:{},response:{}},GetUploadV6QlobotToShopee:{url:"upload/v6/qlobot_to_shopee",method:"GET",query:{base:""},body:{},response:{msg:"",error:""}},GetUploadV6JakmallToShopee:{url:"upload/v6/jakmall_to_shopee",method:"GET",query:{base:""},body:{},response:{msg:"",error:""}},GetUploadV6JakmallToTokopedia:{url:"upload/v6/jakmall_to_tokopedia",method:"GET",query:{base:"",use_mapper:!1},body:{},response:{msg:"",error:""}},GetTokopediaUploadShopee:{url:"tokopedia/upload/shopee",method:"GET",query:{base:""},body:{},response:{msg:"",error:""}},GetTokopediaUploadTokopedia:{url:"tokopedia/upload/tokopedia",method:"GET",query:{base:""},body:{},response:{msg:"",error:""}},PostV1ProductRenameNamespace:{url:"v1/product/rename_namespace",method:"POST",query:void 0,body:{marketplace:"",namespace:"",update_namespace:""},response:{errcode:0,message:"",status:""}},GetV1CategoryMapperTokopediaToShopeeAutosuggest:{url:"v1/category/mapper/tokopedia_to_shopee_autosuggest",method:"GET",query:{namespace:""},body:{},response:{errcode:0,message:"",status:""}},GetV1AkunTempAkun:{url:"v1/akun/temp_akun",method:"GET",query:void 0,body:{},response:{data:"",message:"",error:0}},GetV1MainInfo:{url:"v1/main/info",method:"GET",query:void 0,body:{},response:{lisensi:"",version:""}},GetV1MainCacheSize:{url:"v1/main/cache_size",method:"GET",query:{reset:!1},body:{},response:{processing:!1,cache_size:{size:0,size_kb:0,size_mb:0,size_gb:0},webdriver_size:{size:0,size_kb:0,size_mb:0,size_gb:0}}},DeleteV1MainClearCache:{url:"v1/main/clear_cache",method:"DELETE",query:void 0,body:{},response:{errcode:0,message:"",status:""}},DeleteV1MainClearCacheWebdriver:{url:"v1/main/clear_cache_webdriver",method:"DELETE",query:void 0,body:{},response:{errcode:0,message:"",status:""}},GetPdcsourceAttrToped:{url:"pdcsource/attr_toped",method:"GET",query:{product_id:0,attribute_type:""},body:{},response:{err_msg:"",data:[{categories:[0],attributes:[""]}]}},PutPdcsourceAttrToped:{url:"pdcsource/attr_toped",method:"PUT",query:void 0,body:{product_id:0,attribute_type:"",data:{categories:[0],attributes:[""]}},response:{err_msg:"",data:{categories:[0],attributes:[""]}}},GetPdcsourceAttShopee:{url:"pdcsource/att_shopee",method:"GET",query:{product_id:0,attribute_type:""},body:{},response:{err_msg:"",data:[{categories:[0],attributes:[{attribute_id:0,attribute_value_id:0,custom_value:{raw_value:"",unit:""},attribute_values:[{raw_value:"",unit:"",value_id:0}]}]}]}},PutPdcsourceAttShopee:{url:"pdcsource/att_shopee",method:"PUT",query:void 0,body:{product_id:0,attribute_type:"",data:{categories:[0],attributes:[{attribute_id:0,attribute_value_id:0,custom_value:{raw_value:"",unit:""},attribute_values:[{raw_value:"",unit:"",value_id:0}]}]}},response:{err_msg:"",data:{categories:[0],attributes:[{attribute_id:0,attribute_value_id:0,custom_value:{raw_value:"",unit:""},attribute_values:[{raw_value:"",unit:"",value_id:0}]}]}}},PostPdcsourceCollectionCreate:{url:"pdcsource/collection/create",method:"POST",query:void 0,body:{id:0,name:"",created_at:"2021-12-01T07:00:00+07:00",updated_at:"2021-12-01T07:00:00+07:00"},response:{err_msg:"",id:0,name:"",created_at:"2021-12-01T07:00:00+07:00",updated_at:"2021-12-01T07:00:00+07:00"}},GetPdcsourceCollectionList:{url:"pdcsource/collection/list",method:"GET",query:{page:0,limit:0},body:{},response:{err_msg:"",data:[{id:0,name:"",created_at:"2021-12-01T07:00:00+07:00",updated_at:"2021-12-01T07:00:00+07:00",count:0}],page:0,limit:0,count:0}},GetPdcsourceCollectionItem:{url:"pdcsource/collection/item",method:"GET",query:{col_id:0},body:{},response:{id:0,name:"",created_at:"2021-12-01T07:00:00+07:00",updated_at:"2021-12-01T07:00:00+07:00"}},DeletePdcsourceCollection:{url:"pdcsource/collection",method:"DELETE",query:void 0,body:{collection_ids:[0]},response:{}},GetPdcsourceProductList:{url:"pdcsource/product/list",method:"GET",query:{page:0,limit:0,coll_id:0,status:""},body:{},response:{err_msg:"",data:[{id:0,as_draft:!1,image_preview:"",image_collection_path:"",count_image:0,title:"",price:0,desc:"",use_variant:!1,use_markup:"",weight:0,stock:0,attribute:[{id:0,product_id:0,attribute_type:"",data:""}],field_spin:[{id:0,product_id:0,field_type:"",use_spin:!1,spin_text:"",use_once_text:!1,once_text:[{id:0,field_config_id:0,text:""}]}],variant_option:[{id:0,product_id:0,name:"",option:[""]}],variant_image:[{id:0,product_id:0,option_name:"",image_preview:"",image_collection_path:"",count_image:0}],variant:[{id:0,product_id:0,names:[""],values:[""],price:0,stock:0}],collection:[{id:0,name:"",created_at:"2021-12-01T07:00:00+07:00",updated_at:"2021-12-01T07:00:00+07:00"}],watermark:{id:0,product_id:0,text:"",font_location:"",pos_x:"",pos_y:"",use_watermark:!1},frame_config:{id:0,product_id:0,use_frame:!1,frame_location:""},last_error:"",map:[{id:0,map_id:"",product_id:0,url_product:""}]}],page:0,limit:0,count:0}},PostPdcsourceProductCreate:{url:"pdcsource/product/create",method:"POST",query:void 0,body:{coll_id:0},response:{err_msg:""}},DeletePdcsourceProduct:{url:"pdcsource/product",method:"DELETE",query:void 0,body:{ids:[0]},response:{err_msg:""}},GetPdcsourceProductItem:{url:"pdcsource/product/item",method:"GET",query:{product_id:0},body:{},response:{err_msg:"",data:{id:0,as_draft:!1,image_preview:"",image_collection_path:"",count_image:0,title:"",price:0,desc:"",use_variant:!1,use_markup:"",weight:0,stock:0,attribute:[{id:0,product_id:0,attribute_type:"",data:""}],field_spin:[{id:0,product_id:0,field_type:"",use_spin:!1,spin_text:"",use_once_text:!1,once_text:[{id:0,field_config_id:0,text:""}]}],variant_option:[{id:0,product_id:0,name:"",option:[""]}],variant_image:[{id:0,product_id:0,option_name:"",image_preview:"",image_collection_path:"",count_image:0}],variant:[{id:0,product_id:0,names:[""],values:[""],price:0,stock:0}],collection:[{id:0,name:"",created_at:"2021-12-01T07:00:00+07:00",updated_at:"2021-12-01T07:00:00+07:00"}],watermark:{id:0,product_id:0,text:"",font_location:"",pos_x:"",pos_y:"",use_watermark:!1},frame_config:{id:0,product_id:0,use_frame:!1,frame_location:""},last_error:"",map:[{id:0,map_id:"",product_id:0,url_product:""}]}}},GetPdcsourceEditNew:{url:"pdcsource/edit/new",method:"GET",query:{coll_id:0},body:{},response:{err_msg:"",data:{id:0,as_draft:!1,image_preview:"",image_collection_path:"",count_image:0,title:"",price:0,desc:"",use_variant:!1,use_markup:"",weight:0,stock:0,attribute:[{id:0,product_id:0,attribute_type:"",data:""}],field_spin:[{id:0,product_id:0,field_type:"",use_spin:!1,spin_text:"",use_once_text:!1,once_text:[{id:0,field_config_id:0,text:""}]}],variant_option:[{id:0,product_id:0,name:"",option:[""]}],variant_image:[{id:0,product_id:0,option_name:"",image_preview:"",image_collection_path:"",count_image:0}],variant:[{id:0,product_id:0,names:[""],values:[""],price:0,stock:0}],collection:[{id:0,name:"",created_at:"2021-12-01T07:00:00+07:00",updated_at:"2021-12-01T07:00:00+07:00"}],watermark:{id:0,product_id:0,text:"",font_location:"",pos_x:"",pos_y:"",use_watermark:!1},frame_config:{id:0,product_id:0,use_frame:!1,frame_location:""},last_error:"",map:[{id:0,map_id:"",product_id:0,url_product:""}]}}},PutPdcsourceEditPublish:{url:"pdcsource/edit/publish",method:"PUT",query:void 0,body:{product_id:0},response:{err_msg:""}},PostPdcsourceEditSetBasic:{url:"pdcsource/edit/set_basic",method:"POST",query:void 0,body:{product_id:0,image_preview:"",image_collection_path:"",count_image:0,title:"",price:0,desc:"",use_variant:!1,use_markup:"",weight:0,stock:0},response:{err_msg:""}},PostPdcsourceEditVariationUpdate:{url:"pdcsource/edit/variation_update",method:"POST",query:void 0,body:{product_id:0,variant_option:[{id:0,product_id:0,name:"",option:[""]}],variant_image:[{id:0,product_id:0,option_name:"",image_preview:"",image_collection_path:"",count_image:0}],variant:[{id:0,product_id:0,names:[""],values:[""],price:0,stock:0}]},response:{err_msg:""}},PostPdcsourceEditFieldConfig:{url:"pdcsource/edit/field_config",method:"POST",query:void 0,body:{product_id:0,field_spin:[{id:0,product_id:0,field_type:"",use_spin:!1,spin_text:"",use_once_text:!1,once_text:[{id:0,field_config_id:0,text:""}]}]},response:{err_msg:""}},GetPdcsourceSpinListText:{url:"pdcsource/spin/list_text",method:"GET",query:{field_config_id:0},body:{},response:{err_msg:"",count:0,data:[{id:0,field_config_id:0,text:""}]}},PostPdcsourceSpinAddText:{url:"pdcsource/spin/add_text",method:"POST",query:void 0,body:{field_config_id:0,data:[""]},response:{err_msg:""}},PostPdcsourceSpinFieldConfig:{url:"pdcsource/spin/field_config",method:"POST",query:void 0,body:{product_id:0,field_type:""},response:{err_msg:"",data:{id:0,product_id:0,field_type:"",use_spin:!1,spin_text:"",use_once_text:!1,once_text:[{id:0,field_config_id:0,text:""}]}}},DeletePdcsourceSpinFieldConfig:{url:"pdcsource/spin/field_config",method:"DELETE",query:void 0,body:{id:0},response:{err_msg:""}},PostPdcsourceFsCheck:{url:"pdcsource/fs/check",method:"POST",query:void 0,body:{path:""},response:{err_msg:"",is_abs:!1,Absolute:"",exist:!1}}};function q(s,m){const y=n[s].url,_=n[s].method,a=m,[g,i]=r.useState(!1),[h,d]=r.useState(null),[f,p]=r.useState(null);async function b(e=a){var l,u;i(!0);const c=e==null?void 0:e.query;try{const{data:t}=await k({method:_,url:y,...c?{params:c}:{}});(l=e==null?void 0:e.onSuccess)==null||l.call(e,t),d(t),p(null)}catch(t){(u=e==null?void 0:e.onError)==null||u.call(e,t),p(t),d(null)}finally{i(!1)}}return{data:h,error:f,pending:g,send:b}}function E(s,m){const y=n[s].url,_=n[s].method,a=m,[g,i]=r.useState(!1),[h,d]=r.useState(null),[f,p]=r.useState(null);async function b(e=a,c){var u,t,v,T;i(!0);const l=(a==null?void 0:a.query)||(e==null?void 0:e.query);try{const{data:o}=await k({method:_,url:y,data:c,...l?{params:l}:{}});(u=a==null?void 0:a.onSuccess)==null||u.call(a,o),(t=e==null?void 0:e.onSuccess)==null||t.call(e,o),d(o),p(null)}catch(o){(v=a==null?void 0:a.onError)==null||v.call(a,o),(T=e==null?void 0:e.onError)==null||T.call(e,o),p(o),d(null)}finally{i(!1)}}return{data:h,error:f,pending:g,mutate:b}}export{E as a,q as u};
