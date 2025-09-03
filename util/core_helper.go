package util

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"sort"
	"time"
)

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

var (
	charset = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
	src     = rand.NewSource(time.Now().UnixNano())
)

// GenerateRandomCode 随机码（随机创建六位字符串）自定义值
func GenerateRandomCode(length int) string {
	b := make([]byte, length)
	charsetLen := len(charset)

	for i := range b {
		b[i] = charset[src.Int63()%int64(charsetLen)]
	}
	return string(b)
}
