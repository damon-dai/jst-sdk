/**
 * @Author: Damon
 * @Date: 2025/9/3 13:34
 * @Description:
 **/

package test

import (
	"fmt"
	"github.com/damon-dai/jst-sdk"
	"testing"
)

func TestJstClientOptions_getInitToken(t *testing.T) {
	type fields struct {
		appKey    string
		appSecret string
		sandbox   bool
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "getInitToken",
			fields: fields{
				appKey:    "b0b7d1db226d4216a3d58df9ffa2dde5",
				appSecret: "99c4cef262f34ca882975a7064de0b87",
				sandbox:   true,
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jstClientOptions := jst_sdk.JstClientOptions{}
			_ = jst_sdk.NewJstClient(
				jstClientOptions.WithAppKey(tt.fields.appKey),
				jstClientOptions.WithAppSecret(tt.fields.appSecret),
				jstClientOptions.WithSandbox(tt.fields.sandbox),
			)

			//got, err := j.getInitToken()
			//if (err != nil) != tt.wantErr {
			//	t.Errorf("getInitToken() error = %v, wantErr %v", err, tt.wantErr)
			//	return
			//}
			//if got != tt.want {
			//	t.Errorf("getInitToken() got = %v, want %v", got, tt.want)
			//}
		})
	}
}

func TestJstClientOptions_Execute(t *testing.T) {
	type fields struct {
		appKey    string
		appSecret string
		sandbox   bool
		bizParam  string
		apiRoute  string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{
			name: "查询门店列表",
			fields: fields{
				appKey:    "b0b7d1db226d4216a3d58df9ffa2dde5",
				appSecret: "99c4cef262f34ca882975a7064de0b87",
				sandbox:   true,
				bizParam:  "{\"shop_ids\": [10156218], \"page_index\": 1, \"page_size\": 2}",
				apiRoute:  "QUERY_SHOPS",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "订单查询",
			fields: fields{
				appKey:    "b0b7d1db226d4216a3d58df9ffa2dde5",
				appSecret: "99c4cef262f34ca882975a7064de0b87",
				sandbox:   true,
				bizParam:  `{"shop_id": 10156218}`,
				apiRoute:  "QUERY_ORDERS_SINGLE",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "物流公司查询",
			fields: fields{
				appKey:    "b0b7d1db226d4216a3d58df9ffa2dde5",
				appSecret: "99c4cef262f34ca882975a7064de0b87",
				sandbox:   true,
				bizParam:  "{\"page_index\": 1, \"page_size\": 10}",
				apiRoute:  "QUERY_LOGISTICSCOMPANY",
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "普通商品资料上传",
			fields: fields{
				appKey:    "b0b7d1db226d4216a3d58df9ffa2dde5",
				appSecret: "99c4cef262f34ca882975a7064de0b87",
				sandbox:   true,
				bizParam:  "{\"items\":[{\"i_id\":\"001C0812\",\"supplier_i_id\":\"sup_i_id12\",\"item_type\":\"成品\",\"sku_pic\":\"Sku_pic11\",\"remark\":\"123\",\"pic\":\"1231\",\"pic_big\":\"Pic_big1\",\"enabled\":-1,\"shelf_life\":10,\"batch_enabled\":true,\"s_price\":100,\"supplier_name\":\"供应商1\",\"brand\":\"NIKE12\",\"vc_name\":\"虚拟分类12\",\"other_price_5\":1.52,\"other_code\":\"\",\"h\":1,\"weight\":1.52,\"other_5\":\"2.51\",\"c_price\":0,\"sku_id\":\"001C0812\",\"other_4\":\"2.41\",\"other_3\":\"2.31\",\"l\":1,\"other_2\":\"2.21\",\"other_1\":\"2.11\",\"labels\":[],\"supplier_sku_id\":\"sup_sku_id12\",\"unit\":\"\",\"properties_value\":\"蓝色;XXL\",\"stock_disabled\":true,\"w\":3,\"is_series_number\":true,\"name\":\"兴动测试\",\"market_price\":300,\"c_name\":\"\",\"short_name\":\"短袖12\",\"other_price_3\":1.32,\"sku_code\":\"648125644612\",\"other_price_4\":1.42,\"other_price_1\":1.12,\"other_price_2\":1.22,\"hand_day\":180,\"rejectLifecycle\":10,\"lockupLifecycle\":20,\"adventLifecycle\":30,\"production_licence\":\"SC001\",\"CategoryPropertys\":{\"年份\":\"2021\",\"季节\":\"春\",\"波段\":\"春一波\",\"面料成分\":\"面料\",\"面料类别\":\"针织类\",\"执行标准\":\"2019\",\"安全技术类别\":\" B类\",\"计划上市日期\":\"2020-12-24\"}}]}",
				apiRoute:  "UPLOAD_ITEMSKU",
			},
			want:    "",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			jstClientOptions := jst_sdk.JstClientOptions{}
			j := jst_sdk.NewJstClient(
				jstClientOptions.WithAppKey(tt.fields.appKey),
				jstClientOptions.WithAppSecret(tt.fields.appSecret),
				jstClientOptions.WithSandbox(tt.fields.sandbox),
				jstClientOptions.WithBizParam(tt.fields.bizParam),
				jstClientOptions.WithApiRoute(tt.fields.apiRoute),
			)
			got := j.Execute()
			fmt.Println("got: ", got)
		})
	}
}
