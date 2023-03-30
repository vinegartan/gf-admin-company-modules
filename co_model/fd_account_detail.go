package co_model

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/base_model"
)

// FdAccountDetail 财务账号相关金额统计
type FdAccountDetail struct {
	Id                int64       `json:"id"                description:"和财务账号 id保持一致"`
	TodayAccountSum   int         `json:"todayAccountSum"   description:"今日金额"`
	TodayUpdatedAt    *gtime.Time `json:"todayUpdatedAt"    description:"今日金额更新时间"`
	WeekAccountSum    int         `json:"weekAccountSum"    description:"本周金额"`
	WeekUpdatedAt     *gtime.Time `json:"weekUpdatedAt"     description:"本周金额更新时间"`
	MonthAccountSum   int         `json:"monthAccountSum"   description:"本月金额"`
	MonthUpdatedAt    *gtime.Time `json:"monthUpdatedAt"    description:"本月金额更新时间"`
	QuarterAccountSum int         `json:"quarterAccountSum" description:"本季度金额统计"`
	QuarterUpdatedAt  *gtime.Time `json:"quarterUpdatedAt"  description:"本季度金额更新时间"`
	YearAccountSum    int64       `json:"yearAccountSum"    description:"本年度金额统计"`
	YearUpdatedAt     *gtime.Time `json:"yearUpdatedAt"     description:"本年度金额更新时间"`
	UnionMainId       int64       `json:"unionMainId"       description:"关联主体id"`
	SysUserId         int64       `json:"sysUserId"         description:"关联用户id"`
	Version           int         `json:"version"           description:"乐观锁所需数据版本字段"`
	SceneType         int         `json:"sceneType"         description:"场景类型：0不限、1充电佣金收入"`
}

type FdAccountDetailRes struct {
	co_entity.FdAccountDetail
}

type FdAccountDetailListRes base_model.CollectRes[FdAccountDetailRes]

func (m *FdAccountDetailRes) Data() *FdAccountDetailRes {
	return m
}

type IFdAccountDetailRes interface {
	Data() *FdAccountDetailRes
}
