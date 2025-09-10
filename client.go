/**
 * @Author: Damon
 * @Date: 2025/9/2 15:57
 * @Description: 客户端获取
 **/

package jst_sdk

import (
	"fmt"
	"github.com/damon-dai/jst-sdk/intf"
	"github.com/damon-dai/jst-sdk/util"
	"time"
)

type JstClientOptions struct {
	appKey    string // 开发者应用Key
	appSecret string // 开发者应用密钥
	sandbox   bool   // 沙箱环境：true、是
	bizParam  string // 业务请求参数，格式为jsonString
	apiRoute  string // api路由
	bizCode   string // 业务编码
	factory   *factory
}

type Option func(option *JstClientOptions)

func NewJstClient(opts ...Option) *JstClientOptions {
	options := defaultRequestOptions() // 默认的请求选项

	for _, optFunc := range opts {
		optFunc(options)
	}

	if len(options.appKey) == 0 || len(options.appSecret) == 0 {
		panic("初始化配置失败")
	}

	return &JstClientOptions{
		appKey:    options.appKey,
		appSecret: options.appSecret,
		sandbox:   options.sandbox,
		bizParam:  options.bizParam,
		apiRoute:  options.apiRoute,
		bizCode:   options.bizCode,
	}
}

// Execute 请求接口
func (j *JstClientOptions) Execute() string {
	route, ok := Routes[j.apiRoute]
	if !ok {
		return fmt.Sprintf(`{"code": 10000, "msg": [%s]路由参数错误}`, j.apiRoute)
	}

	response, err := util.HttpPostForm(fmt.Sprintf("%s/%s", j.getApiUrl(), route), j.getParams(), 2*time.Second)
	if err != nil {
		return fmt.Sprintf(`{"code": 10000, "msg": %s}`, err.Error())
	}

	return string(response)
}

// 获取初始化token
func (j *JstClientOptions) getInitToken() (string, error) {
	url := fmt.Sprintf("%s/%s", j.apiUrl, GET_INIT_TOKEN)
	param := map[string]string{
		"app_key":    j.appKey,
		"timestamp":  fmt.Sprintf("%d", time.Now().Unix()),
		"grant_type": "authorization_code",
		"charset":    "utf-8",
		"code":       util.GenerateRandomCode(6), // 随机码（随机创建六位字符串）自定义值
	}
	param["sign"] = util.GenerateSign(j.appSecret, param)
	response, err := util.HttpPostForm(url, param, 2*time.Second)
	if err != nil {
		return "", fmt.Errorf(`{"code": 10000, "msg": %s}`, err.Error())
	}

	return string(response), nil
}

// 刷新token
func (j *JstClientOptions) refreshToken() string {
	url := fmt.Sprintf("%s/%s", j.getApiUrl(), REFRESH_TOKEN)
	param := map[string]string{
		"app_key":    j.appKey,
		"timestamp":  fmt.Sprintf("%d", time.Now().Unix()),
		"grant_type": "refresh_token",
		"charset":    "utf-8",
		"scope":      "all",
	}
	param["sign"] = util.GenerateSign(j.appSecret, param)

	// code为0表示请求成功，其他情况都为请求失败
	response, err := util.HttpPostForm(url, param, 2*time.Second)
	if err != nil {
		return fmt.Sprintf(`{"code": 10000, "msg": %s}`, err.Error())
	}

	return string(response)
}

func (j *JstClientOptions) getPublicRequestParams() map[string]string {
	return map[string]string{
		"app_key":      j.appKey,
		"access_token": "b7e3b1e24e174593af8ca5c397e53dad",
		"timestamp":    fmt.Sprintf("%d", time.Now().Unix()),
		"charset":      "utf-8",
		"version":      "2",
	}
}

// 获取接口请求参数
func (j *JstClientOptions) getParams() map[string]string {
	param := j.getPublicRequestParams()
	bizParam := j.bizParam
	if len(j.bizParam) == 0 {
		bizParam = "{}"
	}
	param["biz"] = bizParam
	param["sign"] = util.GenerateSign(j.appSecret, param)

	return param
}

func (j *JstClientOptions) getApiUrl() string {
	apiUrl := "https://openapi.jushuitan.com"
	if j.sandbox {
		apiUrl = "https://dev-api.jushuitan.com"
	}

	return apiUrl
}

func (j *JstClientOptions) initStoreErpProvider() (intf.StoreErp, error) {
	erpPair, err := j.factory.NewStoreErpProvider(j.bizCode, intf.StoreConfig{
		AppKey:    j.appKey,
		AppSecret: j.appSecret,
		Sandbox:   j.sandbox,
	})
	return erpPair, err
}

// ShopsQuery 店铺查询
func (j *JstClientOptions) ShopsQuery() string {
	// 1、初始化支付实例
	tradePair, err := j.initStoreErpProvider()
	if err != nil {
		return fmt.Sprintf(`{"code": 10000, "msg": %s}`, err.Error())
	}

	return tradePair.ShopsQuery(j.bizParam)
}

func (j *JstClientOptions) WithAppKey(appKey string) Option {
	return func(option *JstClientOptions) {
		option.appKey = appKey
	}
}

func (j *JstClientOptions) WithAppSecret(appSecret string) Option {
	return func(option *JstClientOptions) {
		option.appSecret = appSecret
	}
}

func (j *JstClientOptions) WithSandbox(sandbox bool) Option {
	return func(option *JstClientOptions) {
		option.sandbox = sandbox
	}
}

func (j *JstClientOptions) WithBizParam(bizParam string) Option {
	return func(option *JstClientOptions) {
		option.bizParam = bizParam
	}
}

func (j *JstClientOptions) WithApiRoute(apiRoute string) Option {
	return func(option *JstClientOptions) {
		option.apiRoute = apiRoute
	}
}

func defaultRequestOptions() *JstClientOptions {
	return &JstClientOptions{ // 默认请求选项
		sandbox: false, // 默认生产环境
	}
}
