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
