package co_consts

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/kysion/base-library/utility/base_permission"
)

type global struct {
	// 默认货币类型
	DefaultCurrency string
}

var (
	Global = global{}

	PermissionTree []base_permission.IPermission

	FinancialPermissionTree []base_permission.IPermission
)

func init() {
	defaultCurrency, _ := g.Cfg().Get(context.Background(), "service.defaultCurrency")
	Global.DefaultCurrency = defaultCurrency.String()
}
