package intf

// StoreErp 电商erp接口
type StoreErp interface {
	ShopsQuery(bizParam string) string    // 店铺查询
	UploadItemSku(bizParam string) string // 普通商品资料上传
	// ...
}

type StoreConfig struct {
	AppKey    string
	AppSecret string
	Sandbox   bool
}
