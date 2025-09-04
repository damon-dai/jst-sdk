package jst_sdk

import (
	"fmt"
	"github.com/damon-dai/jst-sdk/intf"
	"github.com/damon-dai/jst-sdk/strategy/houniao"
	"github.com/damon-dai/jst-sdk/strategy/jushuitan"
)

// Factory 选择具体的电商erp实现
type factory struct{}

const (
	JST_ERP   = "B101" // 聚水潭
	HNIAO_ERP = "B102" // 候鸟
)

// NewStoreErpProvider 创建电商erp实例
func (f *factory) NewStoreErpProvider(bizCode string, conf intf.StoreConfig) (intf.StoreErp, error) {
	switch bizCode {
	case JST_ERP:
		// 聚水潭
		return jushuitan.NewJstErp(conf), nil
	case HNIAO_ERP:
		// 候鸟
		return houniao.NewHniaoErp(conf), nil
	default:
		return nil, fmt.Errorf("不支持的业务编码: %s", bizCode)
	}
}
