package api

import (
	"github.com/pdcgo/tokopedia_lib/lib/model"
	"github.com/pdcgo/tokopedia_lib/lib/query"
)

func (api *TokopediaApi) GetInvoiceV3(invoice string) (*model.InvoiceRes, error) {
	variable := query.NewGetInvoiceVariable(invoice)
	query := GraphqlPayload{
		OperationName: "GetInvoiceV3",
		Variables:     variable,
		Query: `fragment InvoiceV3AddonSummary on InvoiceV3AddonSummary {
			addons {
			  order_id
			  addons_id: id
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
			total
			total_price
			total_price_str
			total_quantity
			__typename
		  }
		  
		  query GetInvoiceV3($input: GetInvoiceV3Request!) {
			get_invoice_v3(input: $input) {
			  invoices {
				order_data {
				  invoice_date
				  invoice_ref_num
				  status
				  update_time
				  shop_name
				  customer_name
				  is_toko_cabang
				  is_pre_order
				  pre_order_day
				  additional_status_label
				  have_product_bundle
				  payment_summary_invoice {
					label
					value
					key
					__typename
				  }
				  shipping_data {
					shipping_address
					shipping_courier
					shipping_weight
					receiver_name
					receiver_phone
					is_bebas_ongkir
					bebas_ongkir_img_url
					is_shipping_insurance
					shipping_insurance_label
					__typename
				  }
				  dropship_info {
					name
					phone
					__typename
				  }
				  tokopedia_merchandise {
					name
					address
					npwp
					wording_ppn
					is_stamp_duty
					wording_stamp_duty
					__typename
				  }
				  addon_info {
					all_level {
					  ...InvoiceV3AddonSummary
					  __typename
					}
					order_level {
					  ...InvoiceV3AddonSummary
					  __typename
					}
					__typename
				  }
				  features {
					partial_fulfillment {
					  footer_info
					  order_unfulfilled_details {
						order_dtl_id
						product_id
						product_name
						product_price
						product_price_total
						product_qty
						product_sku_id
						product_weight
						product_notes
						product_cashback_label
						product_snapshot_url
						is_ppp
						__typename
					  }
					  partial_order_fulfillment_summary {
						label
						value
						key
						__typename
					  }
					  __typename
					}
					__typename
				  }
				  details {
					bundles {
					  bundle_id
					  bundle_variant_id
					  bundle_name
					  bundle_price
					  bundle_quantity
					  bundle_subtotal_price
					  order_detail {
						product_id
						product_name
						product_price
						product_price_total
						product_qty
						product_sku_id
						product_weight
						product_notes
						product_cashback_label
						product_snapshot_url
						is_ppp
						bundle_id
						bundle_variant_id
						addon_summary {
						  addons {
							order_id
							addons_id: id
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
						  total
						  total_price
						  total_price_str
						  total_quantity
						  __typename
						}
						__typename
					  }
					  __typename
					}
					non_bundles {
					  product_id
					  product_name
					  product_price
					  product_price_total
					  product_qty
					  product_sku_id
					  product_weight
					  product_notes
					  product_cashback_label
					  product_snapshot_url
					  is_ppp
					  addon_summary {
						...InvoiceV3AddonSummary
						__typename
					  }
					  __typename
					}
					bundle_icon
					total_products
					addon_icon
					addon_label
					bmgm_icon
					bmgms {
					  bmgm_id: id
					  bmgm_tier_name
					  label
					  tier_discount_amount
					  tier_discount_amount_formatted
					  price_before_benefit
					  price_before_benefit_formatted
					  price_after_benefit
					  price_after_benefit_formatted
					  order_detail {
						product_id
						product_name
						product_price
						product_price_total
						product_qty
						product_sku_id
						product_weight
						product_notes
						product_cashback_label
						product_snapshot_url
						is_ppp
						addon_summary {
						  addons {
							order_id
							addons_id: id
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
						  total
						  total_price
						  total_price_str
						  total_quantity
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
				payment_data {
				  payment_method {
					name
					img_url
					amount
					installment_detail
					__typename
				  }
				  payment_details {
					label
					value
					key
					__typename
				  }
				  order_count
				  orders {
					invoice_ref_num
					shop_name
					amount
					__typename
				  }
				  voucher_details {
					voucher_name
					voucher_code
					cashback_details {
					  amount
					  amount_detail
					  __typename
					}
					is_merchant_coupon
					is_tokopedia_coupon
					discount_details {
					  amount
					  type
					  __typename
					}
					__typename
				  }
				  voucher_disclaimer
				  voucher_disclaimer_display_type
				  __typename
				}
				__typename
			  }
			  __typename
			}
		  }
		`,
	}

	req := api.NewGraphqlReq(&query)

	var hasil model.InvoiceRes
	err := api.SendRequest(req, &hasil)

	return &hasil, err
}
