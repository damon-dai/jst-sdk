/**
 * @Author: Damon
 * @Date: 2025/9/4 14:19
 * @Description: 聚水潭电商erp
 **/

package jushuitan

import (
	"fmt"
	"github.com/damon-dai/jst-sdk/intf"
	"github.com/damon-dai/jst-sdk/util"
	"time"
)

// 实现接口
type jstErp struct {
	conf intf.StoreConfig
}

func NewJstErp(conf intf.StoreConfig) *jstErp {
	jstClient := &jstErp{
		conf: conf,
	}

	return jstClient
}

func (j *jstErp) getApiUrl(sandbox bool) string {
	apiUrl := "https://openapi.jushuitan.com"
	if sandbox {
		apiUrl = "https://dev-api.jushuitan.com"
	}

	return apiUrl
}

func (j *jstErp) getPublicRequestParams() map[string]string {
	return map[string]string{
		"app_key":      j.conf.AppKey,
		"access_token": "b7e3b1e24e174593af8ca5c397e53dad",
		"timestamp":    fmt.Sprintf("%d", time.Now().Unix()),
		"charset":      "utf-8",
		"version":      "2",
	}
}

// 获取接口请求参数
func (j *jstErp) getParams(bizParam string) map[string]string {
	param := j.getPublicRequestParams()
	if len(bizParam) == 0 {
		bizParam = "{}"
	}
	param["biz"] = bizParam
	param["sign"] = util.GenerateSign(j.conf.AppSecret, param)

	return param
}

// ShopsQuery 店铺查询
func (j *jstErp) ShopsQuery(bizParam string) string {
	response, err := util.HttpPostForm(fmt.Sprintf("%s/open/shops/query", j.getApiUrl(j.conf.Sandbox)), j.getParams(bizParam), 2*time.Second)
	if err != nil {
		return fmt.Sprintf(`{"code": 10000, "msg": %s}`, err.Error())
	}

	return string(response)
}

// UploadItemSku 普通商品资料上传
func (j *jstErp) UploadItemSku(bizParam string) string {
	return "{}"
}
