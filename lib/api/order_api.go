package api

import (
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApi) OrderList(payload *query.OrderListQuery) (*model.OrderListResp, error) {

	payload.Input.ShopID = int(api.AuthenticatedData.UserShopInfo.Info.ShopID)

	gqlQuery := GraphqlPayload{
		OperationName: "OrderList",
		Variables:     payload,
		Query: `fragment AddonDetail on SOMAddon {
			order_id
			addon_id: id
			reference_id
			level
			name
			price
			price_str
			subtotal_price
			subtotal_price_str
			quantity
			type
			image_url
			metadata {
			  add_on_note {
				is_custom_note
				from
				to
				notes
				short_notes
				__typename
			  }
			  __typename
			}
			create_time
			__typename
		  }
		  
		  fragment AddonSummary on SOMAddonSummary {
			addons {
			  ...AddonDetail
			  __typename
			}
			total
			total_price
			total_price_str
			total_quantity
			__typename
		  }
		  
		  fragment ProductDetail on OrderProduct {
			productId: product_id
			snapshotUrl: snapshot_url
			productName: product_name
			orderNote: order_note
			productPrice: product_price
			productQty: product_qty
			subTotal: sub_total
			totalWeight: total_weight
			addonSummary: addon_summary {
			  ...AddonSummary
			  __typename
			}
			returnable
			sku
			picture
			__typename
		  }
		  
		  fragment BmgmProduct on BMGMSOMList {
			id
			bmgm_tier_name
			tier_discount_amount
			tier_discount_amount_formatted
			price_before_benefit
			price_before_benefit_formatted
			price_after_benefit
			price_after_benefit_formatted
			order_detail {
			  ...ProductDetail
			  __typename
			}
			__typename
		  }
		  
		  fragment BundleProduct on BundleDataSOMList {
			bundleId: bundle_id
			bundleVariantId: bundle_variant_id
			bundleName: bundle_name
			bundlePrice: bundle_price
			bundleSubtotalPrice: bundle_subtotal_price
			bundleItems: order_detail {
			  ...ProductDetail
			  __typename
			}
			__typename
		  }
		  
		  fragment OrderDetails on SOMDetails {
			totalProducts: total_products
			addonLabel: addon_label
			iconAddon: addon_icon
			iconBundle: bundle_icon
			iconBMGM: bmgm_icon
			productsBundle: bundles {
			  ...BundleProduct
			  __typename
			}
			productsSingle: non_bundles {
			  ...ProductDetail
			  __typename
			}
			productsBMGM: bmgms {
			  ...BmgmProduct
			  __typename
			}
			__typename
		  }
		  
		  fragment AddonOrderLevel on SOMAddonInfo {
			label
			icon_url
			orderLevel: order_level {
			  ...AddonSummary
			  __typename
			}
			__typename
		  }
		  
		  query OrderList($input: OrderListArgs!) {
			orderList(input: $input) {
			  firstOrderId: first_order_id
			  lastOrderId: last_order_id
			  firstDate: first_date
			  lastDate: last_date
			  total_data_per_batch
			  emptyStateMetadata: empty_state {
				title
				subtitle
				image_url
				__typename
			  }
			  user_id
			  shop_id
			  user_token
			  is_warehouse_admin
			  isShowBulkAction: is_show_bulk_action
			  isShowDownloadReport: is_show_download_report_button
			  available_courier {
				agency {
				  value: shipping_id
				  text: shipping_name
				  __typename
				}
				product {
				  value: shipping_product_id
				  text: product_name
				  __typename
				}
				__typename
			  }
			  paging {
				show_back_button
				show_next_button
				pages_show_value_list
				pages_real_value_list
				current_batch_page
				current_page
				next_changer_value
				prev_changer_value
				__typename
			  }
			  list {
				orderDetails: details {
				  ...OrderDetails
				  __typename
				}
				pofData: pof_data {
				  isPof: is_pof
				  pofStatus: pof_status
				  __typename
				}
				isMitra: is_mitra
				ticker {
				  text
				  action_text
				  action_key
				  type
				  cta_action_type
				  cta_action_value
				  cta_text
				  __typename
				}
				addonInfo: addon_info {
				  ...AddonOrderLevel
				  __typename
				}
				seller_notes_text
				is_flagged_order
				fulfill_by
				kero_code
				courier_type
				courier_product_id
				id: order_id
				is_additional_cost
				has_booking_info
				booking_info_url
				courier_product_name
				courier_info
				cancel_request
				cancel_request_origin_note
				cancel_request_note
				cancel_request_time
				dropship_name
				dropship_phone
				insurance_type
				status
				order_status_id
				order_resi
				order_resi_url
				origin_address
				origin_district
				origin_geo
				origin_postal_code
				is_purchase_protection
				order_total_price
				labels: order_label {
				  text: flag_name
				  color: flag_color
				  backgroundColor: flag_background
				  __typename
				}
				preorder_process_time_days_left
				buyer_name
				order_date
				buyer_id
				courier_id
				cashback
				courier_name
				isChecked: is_checked
				is_topads
				is_broadcast_chat
				is_shipping_printed
				is_replacement_taken
				trade_in_fee
				deadline_color
				deadline_text
				deadline_time_left
				deadline_style
				destination_street
				destination_district
				destination_city
				destination_province
				destination_postal_code
				destination_phone
				destination_receiver_name
				free_return_order
				courier_ref
				is_free_shipping
				is_penalty_reject
				penalty_reject_wording
				is_show_print_label
				is_tokocabang
				expired_finish_notif_label {
				  is_expired_finish_label
				  unix_expired_label_time
				  unix_finish_order_time
				  finish_order_time
				  expired_label_time
				  __typename
				}
				online_booking {
				  is_hide_input_awb
				  is_remove_input_awb
				  is_show_info
				  info_text
				  __typename
				}
				order_product {
				  productId: product_id
				  snapshotUrl: snapshot_url
				  productName: product_name
				  orderNote: order_note
				  productPrice: product_price
				  productQty: product_qty
				  subTotal: sub_total
				  totalWeight: total_weight
				  addonSummary: addon_summary {
					...AddonSummary
					__typename
				  }
				  returnable
				  sku
				  picture
				  __typename
				}
				logistic_info {
				  all {
					id
					priority
					description
					info_text_short
					info_text_long
					__typename
				  }
				  __typename
				}
				buttonList: button {
				  bulk_url
				  key
				  display_name
				  url
				  title
				  color
				  type
				  popup {
					template {
					  code
					  params
					  __typename
					}
					__typename
				  }
				  __typename
				}
				warehouse_name
				cancelRequestStatus: cancel_request_status
				isShowSellerNotes: is_show_seller_notes
				isShowChatButton: is_show_chat_button
				plus_data {
				  description
				  edu_url
				  logo_url
				  __typename
				}
				shipment_logo
				has_reso_status
				tx_id
				group_type
				tag {
				  isAffiliate: is_affiliate
				  __typename
				}
				__typename
			  }
			  __typename
			}
		  }
		`,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil *model.OrderListResp
	err := api.SendRequest(req, &hasil)
	return hasil, err
}

func (api *TokopediaApi) OrderIncomeDetail(orderid int) (*model.SOMIncomeDetailRes, error) {
	payload := query.OrderIncomeDetailQuery{
		Input: &query.OrderIncomeDetailInput{
			OrderID: orderid,
		},
	}

	gqlQuery := GraphqlPayload{
		OperationName: "GetSOMIncomeDetail",
		Variables:     payload,
		Query: `fragment DetailAttribute on SOMIncomeDetailAttribute {
			tooltip {
			  title
			  value
			  __typename
			}
			data {
			  __typename
			  ... on SOMIncomeDetailAttributeIconData {
				icon_url
				icon_url_dark
				__typename
			  }
			  ... on SOMIncomeDetailAttributeLabelData {
				level
				label
				__typename
			  }
			}
			__typename
		  }
		  
		  query GetSOMIncomeDetail($input: SOMIncomeDetailRequest!) {
			get_som_income_detail(input: $input) {
			  title
			  sections {
				key
				label
				sub_label
				value
				value_raw
				attributes {
				  ...DetailAttribute
				  __typename
				}
				components {
				  key
				  label
				  sub_label
				  value
				  value_raw
				  attributes {
					...DetailAttribute
					__typename
				  }
				  type
				  __typename
				}
				__typename
			  }
			  summary {
				key
				label
				sub_label
				value
				value_raw
				attributes {
				  ...DetailAttribute
				  __typename
				}
				state
				note
				__typename
			  }
			  __typename
			}
		  }
		  `,
	}

	req := api.NewGraphqlReq(&gqlQuery)

	var hasil *model.SOMIncomeDetailRes
	err := api.SendRequest(req, &hasil)
	return hasil, err
}
