# 聚水潭sdk

这里对接服务是「自有商城」应用

- 路由列表
```golang
/**
 * 基础API路由
 */
const (
	// GET_INIT_TOKEN 获取初始化token
	GET_INIT_TOKEN = "openWeb/auth/getInitToken"
	REFRESH_TOKEN  = "openWeb/auth/refreshToken"

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
```

| 参数        | 类型 | 名称    | 是否必填 |
| -----------| ------ | ------ |------ | 
| appKey   | string | 开发者应用Key                       | 必填 | 
| appSecret     | string | 开发者应用密钥                             | 必填 | 
| sandbox | bool | 沙箱环境：true、是          | 非必填 | 
| bizParam   | string | 业务请求参数，格式为jsonString                     | 必填 | 
| apiRoute   | string | api路由                    | 必填 | 

生成签名
```golang
// GenerateSign 生成签名
func GenerateSign(appSecret string, params map[string]string) string {
	if len(params) == 0 {
		return ""
	}

	// 获取所有键并排序
	keys := make([]string, 0, len(params))
	for key := range params {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	// 构建参数字符串
	resultStr := ""
	for _, key := range keys {
		if key == "sign" {
			continue
		}
		resultStr += key + params[key]
	}

	// 拼接应用密钥
	resultStr = appSecret + resultStr

	// 计算MD5并转换为16进制字符串
	hash := md5.Sum([]byte(resultStr))
	return hex.EncodeToString(hash[:])
}
```

1、店铺查询
```golang
    jstClientOptions := jst_sdk.JstClientOptions{}
    j := jst_sdk.NewJstClient(
        jstClientOptions.WithAppKey("b0b7d1db226d4216a3d58df9ffa2dde5"),
        jstClientOptions.WithAppSecret("99c4cef262f34ca882975a7064de0b87"),
        jstClientOptions.WithSandbox(true),
        jstClientOptions.WithBizParam("{\"shop_ids\": [10156218], \"page_index\": 1, \"page_size\": 2}"),
        jstClientOptions.WithApiRoute("QUERY_SHOPS"),
    )
    response := j.Execute()
```

2、订单查询
```golang
    jstClientOptions := jst_sdk.JstClientOptions{}
    j := jst_sdk.NewJstClient(
        jstClientOptions.WithAppKey("b0b7d1db226d4216a3d58df9ffa2dde5"),
        jstClientOptions.WithAppSecret("99c4cef262f34ca882975a7064de0b87"),
        jstClientOptions.WithSandbox(true),
        jstClientOptions.WithBizParam(`{"shop_id": 1, "is_offline_shop": false}`),
        jstClientOptions.WithApiRoute("QUERY_ORDERS_SINGLE"),
    )
    response := j.Execute()
```

### 附录
- 公共请求参数

  | 参数 | 类型 | 名称 | 示例值                            | 是否必填 |
  | ------ | ------ |--------------------------------| ------ | ------ | 
  | access_token | string | 商户授权access_token值 | 0ecde8631431a5ed6b3e7368afbabdabf5c                         | 必填 | 
  | app_key | string | 开发者应用app_key | 0ecde8631431a5ed6b3e7368afbabdadss                               | 必填 | 
  | timestamp | string | 当前请求的10位时间戳【单位是秒】 | 1577771730                     | 必填 | 
  | version | string | 接口版本，当前版本为【2】,目前只能传2，不能不传！ | 2                              | 必填 | 
  | charset | string | 交互数据的编码【utf-8】目前只能传utf-8，不能不传！ | utf-8                          | 必填 | 
  | sign | string | 请求的数字签名，是通过所有请求参数通过摘要生成的，保证请求参数没有被篡改。 | 0ecde8631431a5ed6b3e7368afbabdaoas | 必填 | 


- 公共响应参数

  | 参数 | 类型 | 描述 | 示例值  | 
    | ------ | ------ |------| ------ |
  | code | int | 响应状态码，code为0表示响应成功，其他状态码表示响应失败。具体系统错误码可以参考系统错误码。 | 0    | 
  | msg | string | 错误消息 | 鉴权失败 |  
  | data | object | 业务数据，为Json对象，如果请求错误，该字段可能会展示错误的详细信息。 | {}   | 
 

