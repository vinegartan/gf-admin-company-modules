package audit

import (
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/kysion/base-library/utility/base_permission"
)

type PermissionTypeEnum = base_permission.IPermission

type permissionType struct {
	ViewDetail PermissionTypeEnum
	List       PermissionTypeEnum
	Update     PermissionTypeEnum
}

var (
	PermissionType = permissionType{
		ViewDetail: base_permission.New(5953151699124297, "ViewDetail", "查看资质审核信息", "查看某条资质审核信息"),
		List:       base_permission.New(5953151699124298, "List", "资质审核列表", "查看所有资质审核"),
		Update:     base_permission.New(5953151699124299, "Update", "更新资质审核信息", "更新某条资质审核信息"),
	}
	permissionTypeMap = gmap.NewStrAnyMapFrom(gconv.Map(PermissionType))
)
