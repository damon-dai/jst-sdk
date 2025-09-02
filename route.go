/**
 * @Author: Damon
 * @Date: 2025/9/2 15:57
 * @Description: 聚水潭接口路由统一定义
 **/

package jst_sdk

/**
 * 基础API路由
 */
const (
	// GET_INIT_TOKEN 获取初始化token
	GET_INIT_TOKEN = "openWeb/auth/getInitToken"

	// QUERY_SHOPS 店铺查询
	QUERY_SHOPS = "open/shops/query"

	// QUERY_LOGISTICSCOMPANY 物流公司查询
	QUERY_LOGISTICSCOMPANY = "open/logisticscompany/query"

	// QUERY_PARTNER_WMS 仓库查询
	QUERY_PARTNER_WMS = "open/wms/partner/query"

	// GET_COMPANY_USERS 商家用户信息
	GET_COMPANY_USERS = "/open/webapi/userapi/company/getcompanyusers"
)

/**
 * 商品API路由
 */
const (
	// UPLOAD_ITEMSKU 普通商品资料上传
	UPLOAD_ITEMSKU = "/open/jushuitan/itemsku/upload"

	// UPLOAD_SKUMAP 店铺商品资料上传
	UPLOAD_SKUMAP = "/open/jushuitan/skumap/upload"

	// UPLOAD_COMBINESKU_ITEM 组合装商品上传
	UPLOAD_COMBINESKU_ITEM = "/open/jushuitan/item/combinesku/upload"

	// QUERY_SKUMAP 店铺商品资料查询
	QUERY_SKUMAP = "/open/skumap/query"

	// QUERY_COMBINE_SKU 组合装商品查询
	QUERY_COMBINE_SKU = "/open/combine/sku/query"

	// QUERY_CATEGORY 商品类目查询
	QUERY_CATEGORY = "/open/category/query"

	// QUERY_MALL_ITEM 普通商品查询（按款查询）
	QUERY_MALL_ITEM = "/open/mall/item/query"

	// QUERY_SKU 普通商品资料查询（按sku查询）
	QUERY_SKU = "/open/sku/query"
)

/**
 * 库存API
 */
const (
	// QUERY_INVENTORY 库存查询
	QUERY_INVENTORY = "/open/inventory/query"

	// QUERY_INVENTORY_COUNT 库存盘点查询
	QUERY_INVENTORY_COUNT = "/open/inventory/count/query"

	// UPLOAD_INVENTORY_V2 新建盘点单
	UPLOAD_INVENTORY_V2 = "/open/jushuitan/inventoryv2/upload"
)

/**
 * 订单API
 */
const (
	// UPLOAD_ORDERS 订单上传(商家自有商城、跨境线下)
	UPLOAD_ORDERS = "/open/jushuitan/orders/upload"

	// CANCEL_ORDERBYOID 订单取消-按内部单号取消
	CANCEL_ORDERBYOID = "/open/jushuitan/orderbyoid/cancel"

	// QUERY_ORDERS_SINGLE 订单查询
	QUERY_ORDERS_SINGLE = "/open/orders/single/query"

	// UPLOAD_ORDER_SENT 订单发货
	UPLOAD_ORDER_SENT = "/open/order/sent/upload"
)

/**
 * 物流API
 */
const (
	// UPLOAD_EXPRESS_REGISTER 批量快递登记
	UPLOAD_EXPRESS_REGISTER = "/open/express/register/upload"

	// UPLOAD_ORDERS_WEIGHT_SEND 称重并发货
	UPLOAD_ORDERS_WEIGHT_SEND = "/open/orders/weight/send/upload"

	// QUERY_LOGISTIC 发货信息查询
	QUERY_LOGISTIC = "/open/logistic/query"
)

/**
 * 采购API
 */
const (
	// QUERY_PURCHASE 采购单查询
	QUERY_PURCHASE = "/open/purchase/query"

	// UPLOAD_PURCHASE 采购单上传
	UPLOAD_PURCHASE = "/open/jushuitan/purchase/upload"

	// UPLOAD_SUPPLIER 供应商上传
	UPLOAD_SUPPLIER = "/open/supplier/upload"

	// QUERY_SUPPLIER 供应商查询
	QUERY_SUPPLIER = "/open/supplier/query"

	// QUERY_MANUFACTURE 加工单查询
	QUERY_MANUFACTURE = "/open/jushuitan/manufacture/query"
)

/**
 * 入库API
 */
const (
	// UPLOAD_PURCHASEIN 新建采购入库单
	UPLOAD_PURCHASEIN = "/open/jushuitan/purchasein/upload"

	// QUERY_PURCHASEIN 采购入库查询
	QUERY_PURCHASEIN = "/open/purchasein/query"

	// UPLOAD_PURCHASEIN_RECEIVED 入库单确认
	UPLOAD_PURCHASEIN_RECEIVED = "/open/purchasein/received/upload"

	// UPLOAD_PURCHASEIN_CONFIRM 采购按箱入库
	UPLOAD_PURCHASEIN_CONFIRM = "/open/purchasein/confirm/upload"
)

/**
 * 出库API
 */
const (
	// QUERY_ORDER_OUT_SIMPLE 销售出库查询
	QUERY_ORDER_OUT_SIMPLE = "/open/orders/out/simple/query"

	// UPLOAD_ORDERS_WMS_SENT 出库发货
	UPLOAD_ORDERS_WMS_SENT = "/open/orders/wms/sent/upload"

	// QUERY_PURCHASEOUT 采购退货查询
	QUERY_PURCHASEOUT = "/open/purchaseout/query"

	// CANCEL_PURCHASEOUT 采购退货取消
	CANCEL_PURCHASEOUT = "/open/jushuitan/purchaseout/cancel"
)

/**
 * 售后API
 */
const (
	// QUERY_REFUND_SINGLE 退货退款查询
	QUERY_REFUND_SINGLE = "/open/refund/single/query"

	// QUERY_AFTERSALE_RECEIVED 实际收货查询
	QUERY_AFTERSALE_RECEIVED = "/open/aftersale/received/query"

	// UPLOAD_AFTERSALE 售后上传
	UPLOAD_AFTERSALE = "/open/aftersale/upload"

	// CONFIRM_AFTERSALEAPI 售后单确认
	CONFIRM_AFTERSALEAPI = "/open/webapi/aftersaleapi/open/confirm"

	// PAYQUERYASMODIFIED 退款单查询
	PAYQUERYASMODIFIED = "/open/webapi/aftersaleapi/pay/payqueryasmodified"
)

/**
 *其他出入库API
 */
const (
	// QUERY_OTHER_INOUT 其它出入库查询
	QUERY_OTHER_INOUT = "/open/other/inout/query"

	// UPLOAD_OTHERINOUT 新建其它出入库（新）
	UPLOAD_OTHERINOUT = "/open/jushuitan/otherinout/upload"
)
