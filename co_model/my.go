package co_model

import (
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
)

type MyCompanyRes co_entity.Company

type MyTeamRes struct {
	TeamRes
	MemberItems []*EmployeeRes `json:"memberItems" dc:"团队或小组成员"`
}

type MyTeamListRes []MyTeamRes

type MyProfileRes struct {
	User     *sys_model.SysUser `json:"user" dc:"员工信息"`
	Employee *EmployeeRes       `json:"employee" dc:"员工信息"`
}

type AccountBillRes struct {
	Account co_entity.FdAccount
	Bill    *FdAccountBillListRes
}

type MyAccountBillRes []AccountBillRes
