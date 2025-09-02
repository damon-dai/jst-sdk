/**
 * @Author: Damon
 * @Date: 2025/9/2 15:57
 * @Description: 客户端获取
 **/

package jst_sdk

import (
	"encoding/base64"
	"fmt"
	"github.com/damon-dai/jst-sdk/util"
	"time"
)

type JstClientOptions struct {
	appKey    string // 开发者应用Key
	appSecret string // 开发者应用密钥
	runmode   string // 环境：pro、线上，其他、沙箱环境
	apiUrl    string // 接口地址
}

type Option func(option *JstClientOptions)

func NewAgoraClient(opts ...Option) *JstClientOptions {
	options := defaultRequestOptions() // 默认的请求选项

	for _, optFunc := range opts {
		optFunc(options)
	}

	if len(options.appKey) == 0 || len(options.appSecret) == 0 {
		panic("初始化配置失败")
	}

	apiUrl := "https://openapi.jushuitan.com/open/shops/query"
	if options.runmode == "pro" {
		apiUrl = "https://openapi.jushuitan.com"
	}

	return &JstClientOptions{
		appKey:    options.appKey,
		appSecret: options.appSecret,
		runmode:   options.runmode,
		apiUrl:    apiUrl,
	}
}

// 获取初始化token
func (a *JstClientOptions) getInitToken(channelName string, uid uint32) (string, error) {
	util.HttpPostForm(fmt.Sprintf("%s/%s", a.apiUrl, GET_INIT_TOKEN), map[string]string{
		"app_key": a.appKey,
	}, 2*time.Second)

	return "", nil
}

// base64Encode 对字符串进行Base64编码
func base64Encode(s string) string {
	return base64.StdEncoding.EncodeToString([]byte(s))
}

func (a *JstClientOptions) WithAppKey(appKey string) Option {
	return func(option *JstClientOptions) {
		option.appKey = appKey
	}
}

func (a *JstClientOptions) WithAppSecret(appSecret string) Option {
	return func(option *JstClientOptions) {
		option.appSecret = appSecret
	}
}

func (a *JstClientOptions) WithRunmode(runmode string) Option {
	return func(option *JstClientOptions) {
		option.runmode = runmode
	}
}

func defaultRequestOptions() *JstClientOptions {
	return &JstClientOptions{ // 默认请求选项
		runmode: "pro", // 默认生产环境
	}
}
