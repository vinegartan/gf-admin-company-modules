// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// FdAccountDetail is the golang structure of table co_fd_account_detail for DAO operations like Where/Data.
type FdAccountDetail struct {
	g.Meta            `orm:"table:co_fd_account_detail, do:true"`
	Id                interface{} // 和财务账号 id保持一致
	TodayAccountSum   interface{} // 今日金额
	TodayUpdatedAt    *gtime.Time // 今日金额更新时间
	WeekAccountSum    interface{} // 本周金额
	WeekUpdatedAt     *gtime.Time // 本周金额更新时间
	MonthAccountSum   interface{} // 本月金额
	MonthUpdatedAt    *gtime.Time // 本月金额更新时间
	QuarterAccountSum interface{} // 本季度金额统计
	QuarterUpdatedAt  *gtime.Time // 本季度金额更新时间
	YearAccountSum    interface{} // 本年度金额统计
	YearUpdatedAt     *gtime.Time // 本年度金额更新时间
	UnionMainId       interface{} // 关联主体id
	SysUserId         interface{} // 关联用户id
	Version           interface{} // 乐观锁所需数据版本字段
}
