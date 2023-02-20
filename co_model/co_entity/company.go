// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Company is the golang structure for table company.
type Company struct {
	Id            int64       `json:"id"            description:"ID"`
	Name          string      `json:"name"          description:"名称"`
	ContactName   string      `json:"contactName"   description:"商务联系人"`
	ContactMobile string      `json:"contactMobile" description:"商务联系电话"`
	UserId        int64       `json:"userId"        description:"管理员ID"`
	State         int         `json:"state"         description:"状态：0未启用，1正常"`
	Remark        string      `json:"remark"        description:"备注"`
	CreatedBy     int64       `json:"createdBy"     description:"创建者"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:"创建时间"`
	UpdatedBy     int64       `json:"updatedBy"     description:"更新者"`
	UpdatedAt     *gtime.Time `json:"updatedAt"     description:"更新时间"`
	DeletedBy     int64       `json:"deletedBy"     description:"删除者"`
	DeletedAt     *gtime.Time `json:"deletedAt"     description:"删除时间"`
	ParentId      int64       `json:"parentId"      description:"父级ID"`
}
