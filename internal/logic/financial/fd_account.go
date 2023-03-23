package financial

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
	"reflect"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/kysion/base-library/utility/daoctl"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/yitter/idgenerator-go/idgen"
)

type sFdAccount[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	TR co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
] struct {
	base_hook.ResponseFactoryHook[TR]
	modules co_interface.IModules[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		TR,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]
	dao *co_dao.XDao
}

func NewFdAccount[
	ITCompanyRes co_model.ICompanyRes,
	ITEmployeeRes co_model.IEmployeeRes,
	ITTeamRes co_model.ITeamRes,
	TR co_model.IFdAccountRes,
	ITFdAccountBillRes co_model.IFdAccountBillRes,
	ITFdBankCardRes co_model.IFdBankCardRes,
	ITFdCurrencyRes co_model.IFdCurrencyRes,
	ITFdInvoiceRes co_model.IFdInvoiceRes,
	ITFdInvoiceDetailRes co_model.IFdInvoiceDetailRes,
](modules co_interface.IModules[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	TR,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) co_interface.IFdAccount[TR] {
	result := &sFdAccount[
		ITCompanyRes,
		ITEmployeeRes,
		ITTeamRes,
		TR,
		ITFdAccountBillRes,
		ITFdBankCardRes,
		ITFdCurrencyRes,
		ITFdInvoiceRes,
		ITFdInvoiceDetailRes,
	]{
		modules: modules,
		dao:     modules.Dao(),
	}

	result.ResponseFactoryHook.RegisterResponseFactory(result.FactoryMakeResponseInstance)

	return result
}

// FactoryMakeResponseInstance 响应实例工厂方法
func (s *sFdAccount[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	TR,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) FactoryMakeResponseInstance() TR {
	var ret co_model.IFdAccountRes
	ret = &co_model.FdAccountRes{}
	return ret.(TR)
}

// CreateAccount 创建财务账号
func (s *sFdAccount[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	TR,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) CreateAccount(ctx context.Context, info co_model.FdAccountRegister) (response TR, err error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser
	// 检查指定参数是否为空
	if err := g.Validator().Data(info).Run(ctx); err != nil {
		return response, err
	}

	// 关联用户id是否正确
	user, err := daoctl.GetByIdWithError[sys_entity.SysUser](sys_dao.SysUser.Ctx(ctx), info.UnionUserId)
	if user == nil || err != nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Financial_UnionUserId_Failed"), sys_dao.SysUser.Table())
	}

	// 判断货币代码是否符合标准
	currency, err := s.modules.Currency().GetCurrencyByCurrencyCode(ctx, info.CurrencyCode)
	if err != nil || reflect.ValueOf(currency).IsNil() {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Financial_CurrencyCode_Failed"), s.dao.FdCurrency.Table())
	}
	if currency.Data().IsLegalTender != 1 {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_PleaseUse_Legal_Currency"), s.dao.FdCurrency.Table())
	}
	// 生产随机id
	data := co_do.FdAccount{}
	gconv.Struct(info, &data)
	data.Id = idgen.NextId()
	data.CreatedBy = sessionUser.Id
	data.CreatedAt = gtime.Now()

	// 插入财务账号
	_, err = s.dao.FdAccount.Ctx(ctx).Insert(data)
	if err != nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Account_Save_Failed"), s.dao.FdAccount.Table())
	}

	return s.GetAccountById(ctx, gconv.Int64(data.Id))
}

// GetAccountById 根据ID获取财务账号
func (s *sFdAccount[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	TR,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetAccountById(ctx context.Context, id int64) (response TR, err error) {
	if id == 0 {
		return response, gerror.New(s.modules.T(ctx, "error_AccountId_NonNull"))
	}
	data, err := daoctl.GetByIdWithError[TR](s.dao.FdAccount.Ctx(ctx), id)

	if err != nil {
		return response, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_GetAccountById_Failed"), s.dao.FdAccount.Table())
	}

	return *data, nil
}

// UpdateAccountIsEnable 修改财务账号状态（是否启用：0禁用 1启用）
func (s *sFdAccount[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	TR,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) UpdateAccountIsEnable(ctx context.Context, id int64, isEnabled int64) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	account, err := daoctl.GetByIdWithError[co_entity.FdAccount](s.dao.FdAccount.Ctx(ctx), id)
	if account == nil || err != nil {
		return false, gerror.New(s.modules.T(ctx, "{#Account}{#error_Data_NotFound}"))
	}

	_, err = s.dao.FdAccount.Ctx(ctx).Where(co_do.FdAccount{Id: id}).Update(co_do.FdAccount{
		IsEnabled: isEnabled,
		UpdatedBy: sessionUser.Id,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

// HasAccountByName 根据账户名查询财务账户
func (s *sFdAccount[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	TR,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) HasAccountByName(ctx context.Context, name string) (response TR, err error) {
	response = s.FactoryMakeResponseInstance()

	err = s.dao.FdAccount.Ctx(ctx).Where(co_do.FdAccount{Name: name}).Scan(response.Data())

	if err != nil {
		var ret TR
		return ret, err
	}

	return response, nil
}

// UpdateAccountLimitState 修改财务账号的限制状态 （0不限制，1限制支出、2限制收入）
func (s *sFdAccount[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	TR,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) UpdateAccountLimitState(ctx context.Context, id int64, limitState int64) (bool, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	_, err := s.dao.FdAccount.Ctx(ctx).Where(co_do.FdAccount{Id: id}).Update(co_do.FdAccount{
		LimitState: limitState,
		UpdatedBy:  sessionUser.Id,
	})
	if err != nil {
		return false, err
	}

	return true, nil
}

// QueryAccountListByUserId 获取指定用户的所有财务账号
func (s *sFdAccount[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	TR,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) QueryAccountListByUserId(ctx context.Context, userId int64) (*base_model.CollectRes[TR], error) {
	if userId == 0 {
		return nil, gerror.New("用户id不能为空")
	}

	data, err := daoctl.Query[TR](s.dao.FdAccount.Ctx(ctx).Where(co_do.FdAccount{UnionUserId: userId}), nil, false)

	if err != nil || len(data.Records) <= 0 {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_ThisUser_NotHas_Account"), s.dao.FdAccount.Table())
	}

	return data, nil
}

// UpdateAccountBalance 修改财务账户余额(上下文, 财务账号id, 需要修改的钱数目, 版本, 收支类型)
func (s *sFdAccount[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	TR,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) UpdateAccountBalance(ctx context.Context, accountId int64, amount int64, version int, inOutType int) (int64, error) {
	sessionUser := sys_service.SysSession().Get(ctx).JwtClaimsUser

	db := s.dao.FdAccount.Ctx(ctx)

	data := co_do.FdAccount{
		Version: gdb.Raw(s.dao.FdAccount.Columns().Version + "+1"),
	}

	if inOutType == 1 { // 收入
		// 余额 = 之前的余额 + 本次交易的余额
		data.Balance = gdb.Raw(s.dao.FdAccount.Columns().Balance + "+" + gconv.String(amount))
		data.UpdatedBy = sessionUser.Id

	} else if inOutType == 2 { // 支出
		// 余额 = 之前的余额 - 本次交易的余额
		data.Balance = gdb.Raw(s.dao.FdAccount.Columns().Balance + "-" + gconv.String(amount))
		data.UpdatedBy = sessionUser.Id
	}

	result, err := db.Data(data).Where(co_do.FdAccount{
		Id:      accountId,
		Version: version,
	}).Update()

	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()

	return affected, err
}

// GetAccountByUnionUserIdAndCurrencyCode 根据用户union_user_id和货币代码currency_code获取财务账号
func (s *sFdAccount[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	TR,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetAccountByUnionUserIdAndCurrencyCode(ctx context.Context, unionUserId int64, currencyCode string) (response TR, err error) {
	if unionUserId == 0 {
		return response, gerror.New(s.modules.T(ctx, "error_Account_UnionUserId_NotNull"))
	}

	response = s.FactoryMakeResponseInstance()

	// 查找指定用户名下指定货币类型的财务账号
	err = s.dao.FdAccount.Ctx(ctx).Where(co_do.FdAccount{
		UnionUserId:  unionUserId,
		CurrencyCode: currencyCode,
	}).Scan(response.Data())

	return response, err
}

// GetAccountByUnionUserIdAndScene 根据union_user_id和业务类型找出财务账号，如果主体id找不到财务账号的时候就创建财务账号
func (s *sFdAccount[
	ITCompanyRes,
	ITEmployeeRes,
	ITTeamRes,
	TR,
	ITFdAccountBillRes,
	ITFdBankCardRes,
	ITFdCurrencyRes,
	ITFdInvoiceRes,
	ITFdInvoiceDetailRes,
]) GetAccountByUnionUserIdAndScene(ctx context.Context, unionUserId int64, sceneType ...int) (response TR, err error) {
	if unionUserId == 0 {
		return response, gerror.New(s.modules.T(ctx, "error_Account_UnionUserId_NotNull"))
	}

	response = s.FactoryMakeResponseInstance()
	doWhere := s.dao.FdAccount.Ctx(ctx).Where(co_do.FdAccount{
		UnionUserId: unionUserId,
	})

	if len(sceneType) > 0 {
		doWhere = doWhere.Where(co_do.FdAccount{
			SceneType: sceneType[0],
		})
	}
	err = doWhere.Scan(response.Data())

	// var res co_model.FdAccountRes

	// return s.GetAccountById(ctx, res.Id)

	return response, err

	//if data == nil { //如果主体id找不到财务账号的时候就创建财务账号  （不应该在这里）
	//	s.CreateAccount(ctx,co_model.FdAccountRegister{
	//		UnionLicenseId:     0,
	//		UnionUserId:        0,
	//		Name:               "",
	//		CurrencyCode:       "",
	//		IsEnabled:          0,
	//		LimitState:         0,
	//		PrecisionOfBalance: 0,
	//		Version:            0,
	//		SceneType:          0,
	//		AccountType:        0,
	//		AccountNumber:      "",
	//	})
	//}

}
