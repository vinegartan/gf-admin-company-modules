package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type Team struct {
	Id                int64  `json:"id"                description:"ID"`
	Name              string `json:"name"              v:"required|max-length:128#名称不能为空|名称长度超128字符出限定范围" description:"团队名称，公司维度下唯一"`
	OwnerEmployeeId   int64  `json:"ownerEmployeeId"   description:"团队所有者/业务总监/业务经理/团队队长"`
	CaptainEmployeeId int64  `json:"captainEmployeeId" description:"团队队长编号/小组组长"`
	ParentId          int64  `json:"parentId" description:"团队或小组父级ID"`
	Remark            string `json:"remark"            description:"备注"`
}

type TeamRes struct {
	co_entity.CompanyTeam
	Owner   *EmployeeRes `json:"owner" dc:"团队所有者/业务总监/业务经理/团队队长"`
	Captain *EmployeeRes `json:"captain" dc:"团队队长编号/小组组长"`
	Company *CompanyRes  `json:"company" dc:"所属主体单位"`
	Parent  *TeamRes     `json:"parent" dc:"团队或小组父级ID"`
}

type TeamMemberRes struct {
	co_entity.CompanyTeamMember
	Employee   *EmployeeRes `json:"employee"   description:"成员"`
	InviteUser *EmployeeRes `json:"inviteUser" description:"邀约人"`
	UnionMain  *CompanyRes  `json:"unionMain"  description:"关联主体"`
}

type TeamListRes sys_model.CollectRes[*TeamRes]
type TeamMemberListRes sys_model.CollectRes[*TeamMemberRes]
