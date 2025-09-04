/**
 * @Author: Damon
 * @Date: 2025/9/4 14:19
 * @Description: 候鸟电商erp
 **/

package houniao

import "github.com/damon-dai/jst-sdk/intf"

// 实现接口
type hniaoErp struct {
}

func NewHniaoErp(conf intf.StoreConfig) *hniaoErp {
	hniaoClient := &hniaoErp{}

	return hniaoClient
}

// ShopsQuery 店铺查询
func (h *hniaoErp) ShopsQuery(bizParam string) string {
	return "{}"
}

// UploadItemSku 普通商品资料上传
func (h *hniaoErp) UploadItemSku(bizParam string) string {
	return "{}"
}
