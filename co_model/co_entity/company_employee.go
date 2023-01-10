// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// CompanyEmployee is the golang structure for table company_employee.
type CompanyEmployee struct {
	Id           int64       `json:"id"           description:"ID，保持与USERID一致"`
	No           string      `json:"no"           description:"工号"`
	Avatar       string      `json:"avatar"       description:"头像"`
	Name         string      `json:"name"         description:"姓名"`
	Mobile       string      `json:"mobile"       description:"手机号"`
	UnionMainId  int64       `json:"unionMainId"  description:"所属主体"`
	State        int         `json:"state"        description:"状态： -1已离职，0待确认，1已入职"`
	LastActiveIp string      `json:"lastActiveIp" description:"最后活跃IP"`
	HiredAt      *gtime.Time `json:"hiredAt"      description:"入职时间"`
	CreatedBy    int64       `json:"createdBy"    description:""`
	CreatedAt    *gtime.Time `json:"createdAt"    description:""`
	UpdatedBy    int64       `json:"updatedBy"    description:""`
	UpdatedAt    *gtime.Time `json:"updatedAt"    description:""`
	DeletedBy    int64       `json:"deletedBy"    description:""`
	DeletedAt    *gtime.Time `json:"deletedAt"    description:""`
}
