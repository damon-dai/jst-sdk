package util

import (
	"fmt"
	"github.com/valyala/fasthttp"
	"log"
	"time"
)

func HttpPostForm(baseUrl string, formData map[string]string, timeout time.Duration) ([]byte, error) {
	// 创建请求对象
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)

	// 设置请求URL和方法
	req.SetRequestURI(baseUrl)
	req.Header.SetMethod(fasthttp.MethodPost)

	// 设置表单内容类型
	req.Header.SetContentType("application/x-www-form-urlencoded")

	// 构建表单数据
	args := fasthttp.AcquireArgs()
	defer fasthttp.ReleaseArgs(args)

	for key, value := range formData {
		args.Add(key, value)
	}

	// 将表单数据写入请求体
	args.WriteTo(req.BodyWriter())

	// 创建响应对象
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)

	if timeout <= 1*time.Second {
		timeout = 5 * time.Second // 默认超时时间 5 秒
	}
	// 创建客户端并设置超时
	client := &fasthttp.Client{
		ReadTimeout:  timeout,
		WriteTimeout: timeout,
	}

	// 发送请求
	err := client.Do(req, resp)
	if err != nil {
		return []byte{}, fmt.Errorf("request failed: %v", err)
	}

	// 获取响应
	log.Printf("http request post form url: %s, status_code: %d,  params: %+v; body: %s\n", baseUrl, resp.StatusCode(), formData, string(resp.Body()))
	return resp.Body(), nil
}
