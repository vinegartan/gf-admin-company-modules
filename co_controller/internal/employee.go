package internal

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/api_v1"
	"github.com/SupenBysz/gf-admin-community/sys_model"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-community/utility/funs"
	"github.com/SupenBysz/gf-admin-company-modules/api/co_company_api"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface/i_controller"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_permission"
	"github.com/kysion/base-library/base_model"
)

type EmployeeController[
	ITCompanyRes co_model.ICompanyRes,
	TIRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] struct {
	modules co_interface.IModules[
		ITCompanyRes,
		TIRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]
	dao *co_dao.XDao
}

func Employee[
	ITCompanyRes co_model.ICompanyRes,
	TIRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	ITFdAccountRes co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) i_controller.IEmployee[TIRes] {
	return &EmployeeController[
		ITCompanyRes,
		TIRes,
		ITTeamRes,
		ITFdAccountRes,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]{
		modules: modules,
		dao:     modules.Dao(),
	}
}

func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetEmployeeById(ctx context.Context, req *co_company_api.GetEmployeeByIdReq) (TIRes, error) {
	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			return c.modules.Employee().GetEmployeeById(c.makeMore(ctx), req.Id)
		},
		co_permission.Employee.PermissionType(c.modules).ViewDetail,
	)
}

// GetEmployeeDetailById 获取员工详情信息
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetEmployeeDetailById(ctx context.Context, req *co_company_api.GetEmployeeDetailByIdReq) (res TIRes, err error) {
	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			return c.modules.Employee().GetEmployeeDetailById(c.makeMore(ctx), req.Id)
		},
		co_permission.Employee.PermissionType(c.modules).MoreDetail,
	)
}

// HasEmployeeByName 员工名称是否存在
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) HasEmployeeByName(ctx context.Context, req *co_company_api.HasEmployeeByNameReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Employee().HasEmployeeByName(ctx, req.Name, req.UnionNameId, req.ExcludeId) == true, nil
		},
	)
}

// HasEmployeeByNo 员工工号是否存在
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) HasEmployeeByNo(ctx context.Context, req *co_company_api.HasEmployeeByNoReq) (api_v1.BoolRes, error) {
	unionMainId := sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			return c.modules.Employee().HasEmployeeByNo(ctx, req.No, unionMainId, req.ExcludeId) == true, nil
		},
	)
}

// QueryEmployeeList 查询员工列表
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) QueryEmployeeList(ctx context.Context, req *co_company_api.QueryEmployeeListReq) (*base_model.CollectRes[TIRes], error) {
	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[TIRes], error) {
			return c.modules.Employee().QueryEmployeeList(c.makeMore(ctx), &req.SearchParams)
		},
		co_permission.Employee.PermissionType(c.modules).List,
	)
}

// CreateEmployee 创建员工信息
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) CreateEmployee(ctx context.Context, req *co_company_api.CreateEmployeeReq) (TIRes, error) {
	req.UnionMainId = sys_service.SysSession().Get(ctx).JwtClaimsUser.UnionMainId

	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			ret, err := c.modules.Employee().CreateEmployee(c.makeMore(ctx), &req.Employee)
			return ret, err
		},
		co_permission.Employee.PermissionType(c.modules).Create,
	)
}

// UpdateEmployee 更新员工信息
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) UpdateEmployee(ctx context.Context, req *co_company_api.UpdateEmployeeReq) (TIRes, error) {
	return funs.CheckPermission(ctx,
		func() (TIRes, error) {
			ret, err := c.modules.Employee().UpdateEmployee(c.makeMore(ctx), &req.Employee)
			return ret, err
		},
		co_permission.Employee.PermissionType(c.modules).Update,
	)
}

// DeleteEmployee 删除员工信息
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) DeleteEmployee(ctx context.Context, req *co_company_api.DeleteEmployeeReq) (api_v1.BoolRes, error) {
	return funs.CheckPermission(ctx,
		func() (api_v1.BoolRes, error) {
			ret, err := c.modules.Employee().DeleteEmployee(ctx, req.Id)
			return ret == true, err
		},
		co_permission.Employee.PermissionType(c.modules).Delete,
	)
}

// GetEmployeeListByRoleId 根据角色ID获取所有所属员工
func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetEmployeeListByRoleId(ctx context.Context, req *co_company_api.GetEmployeeListByRoleIdReq) (*base_model.CollectRes[TIRes], error) {
	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[TIRes], error) {
			return c.modules.Employee().GetEmployeeListByRoleId(c.makeMore(ctx), req.Id)
		},
		co_permission.Employee.PermissionType(c.modules).ViewDetail,
	)
}

func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetEmployeeListByTeamId(ctx context.Context, req *co_company_api.GetEmployeeListByTeamId) (*base_model.CollectRes[TIRes], error) {
	return funs.CheckPermission(ctx,
		func() (*base_model.CollectRes[TIRes], error) {
			return c.modules.Employee().GetEmployeeListByTeamId(c.makeMore(ctx), req.TeamId)
		},
		co_permission.Team.PermissionType(c.modules).MemberDetail,
	)
}

func (c *EmployeeController[
	ITCompanyRes,
	TIRes,
	ITTeamRes,
	ITFdAccountRes,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) makeMore(ctx context.Context) context.Context {
	ctx = funs.AttrBuilder[co_model.EmployeeRes, []co_model.Team](ctx, c.dao.Employee.Columns().UnionMainId)
	ctx = funs.AttrBuilder[co_model.EmployeeRes, *co_model.EmployeeRes](ctx, c.dao.Employee.Columns().Id)

	// 因为需要附加公共模块user的数据，所以也要添加有关sys_user的附加数据订阅
	ctx = funs.AttrBuilder[sys_model.SysUser, *sys_entity.SysUserDetail](ctx, sys_dao.SysUser.Columns().Id)
	return ctx
}
