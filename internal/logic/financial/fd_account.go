package financial

import (
	"context"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_dao"
	"github.com/SupenBysz/gf-admin-community/sys_model/sys_entity"
	"github.com/SupenBysz/gf-admin-community/sys_service"
	"github.com/SupenBysz/gf-admin-company-modules/co_consts"
	"github.com/SupenBysz/gf-admin-company-modules/co_interface"
	"github.com/SupenBysz/gf-admin-company-modules/co_model"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_do"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_entity"
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_enum"
	"github.com/kysion/base-library/base_hook"
	"github.com/kysion/base-library/base_model"
	"github.com/kysion/base-library/utility/format_utils"
	"github.com/kysion/base-library/utility/funs"
	"github.com/kysion/base-library/utility/kconv"

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
	data := kconv.Struct(info, &co_do.FdAccount{})

	data.Id = idgen.NextId()
	data.CreatedBy = sessionUser.Id
	data.CreatedAt = gtime.Now()
	data.UnionMainId = info.UnionMainId
	data.IsEnabled = 1
	data.LimitState = 0
	if info.CurrencyCode == "" {
		data.CurrencyCode = co_consts.Global.DefaultCurrency
	}

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

	return makeMore(ctx, s.dao.FdAccountDetail, *data), nil
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
]) UpdateAccountIsEnable(ctx context.Context, id int64, isEnabled int) (bool, error) {
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

// HasAccountByName 判断财务账号名是否存在
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
]) UpdateAccountLimitState(ctx context.Context, id int64, limitState int) (bool, error) {
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

	dataList := make([]TR, 0)

	for _, item := range data.Records {
		more := makeMore(ctx, s.dao.FdAccountDetail, item)
		dataList = append(dataList, more)
	}
	data.Records = dataList

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

	return makeMore(ctx, s.dao.FdAccountDetail, response), err
}

// GetAccountByUnionUserIdAndScene 根据union_user_id和业务类型找出财务账号，
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
]) GetAccountByUnionUserIdAndScene(ctx context.Context, unionUserId int64, accountType co_enum.AccountType, sceneType ...co_enum.SceneType) (response TR, err error) {
	if unionUserId == 0 {
		return response, sys_service.SysLogs().ErrorSimple(ctx, nil, s.modules.T(ctx, "error_Account_UnionUserId_NotNull"), s.dao.FdAccount.Table())
	}

	response = s.FactoryMakeResponseInstance()
	doWhere := s.dao.FdAccount.Ctx(ctx).Where(co_do.FdAccount{
		UnionUserId: unionUserId,
		AccountType: accountType.Code(),
	})

	if len(sceneType) > 0 {
		doWhere = doWhere.Where(co_do.FdAccount{
			SceneType: sceneType[0].Code(),
		})
	}
	err = doWhere.Scan(response.Data())

	if err != nil {
		err = sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_GetAccount_Failed"), s.dao.FdAccount.Table())
	}

	return makeMore(ctx, s.dao.FdAccountDetail, response), err
}

// ========================财务账号金额明细统计=================================

// GetAccountDetailById 根据财务账号id查询账单金额明细统计记录，如果主体id找不到财务账号的时候就创建财务账号
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
]) GetAccountDetailById(ctx context.Context, id int64) (res *co_model.FdAccountDetailRes, err error) {
	if id == 0 {
		return nil, gerror.New(s.modules.T(ctx, "error_AccountId_NonNull"))
	}
	data, err := daoctl.GetByIdWithError[co_model.FdAccountDetailRes](s.dao.FdAccountDetail.Ctx(ctx), id)

	if data == nil {
		account, err := s.GetAccountById(ctx, id)
		if err != nil {
			return nil, err
		}

		now := gtime.Now()
		return s.CreateAccountDetail(ctx, &co_model.FdAccountDetail{
			Id:                id,
			TodayAccountSum:   0,
			TodayUpdatedAt:    now,
			WeekAccountSum:    0,
			WeekUpdatedAt:     now,
			MonthAccountSum:   0,
			MonthUpdatedAt:    now,
			QuarterAccountSum: 0,
			QuarterUpdatedAt:  now,
			YearAccountSum:    0,
			YearUpdatedAt:     now,
			UnionMainId:       account.Data().UnionMainId,
			SysUserId:         account.Data().UnionUserId,
			Version:           0,
			SceneType:         account.Data().SceneType,
		})
	}

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_GetAccountDetailById_Failed"), s.dao.FdAccountDetail.Table())
	}

	return data, nil
}

// CreateAccountDetail 创建财务账单金额明细统计记录
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
]) CreateAccountDetail(ctx context.Context, info *co_model.FdAccountDetail) (res *co_model.FdAccountDetailRes, err error) {
	// 关联用户id是否正确
	user, err := daoctl.GetByIdWithError[sys_entity.SysUser](sys_dao.SysUser.Ctx(ctx), info.SysUserId)
	if user == nil || err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_Financial_UnionUserId_Failed"), sys_dao.SysUser.Table())
	}

	// 生产随机id
	data := co_do.FdAccountDetail{}
	gconv.Struct(info, &data)

	// 插入财务账号金额明细
	_, err = s.dao.FdAccountDetail.Ctx(ctx).Insert(data)
	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "error_AccountDetail_Save_Failed"), s.dao.FdAccountDetail.Table())
	}

	return s.GetAccountDetailById(ctx, gconv.Int64(data.Id))
}

// Increment 收入
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
]) Increment(ctx context.Context, id int64, amount int) (bool, error) {
	ret, err := s.updateAccountDetailAmount(ctx, id, amount, co_enum.Financial.InOutType.In)

	return ret > 0, err
}

// Decrement 支出
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
]) Decrement(ctx context.Context, id int64, amount int) (bool, error) {
	ret, err := s.updateAccountDetailAmount(ctx, id, amount, co_enum.Financial.InOutType.Out)

	return ret > 0, err
}

// UpdateAccountDetailAmount 修改财务账户余额(上下文, id, 需要修改的钱数目, 收支类型)
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
]) updateAccountDetailAmount(ctx context.Context, id int64, amount int, inOutType co_enum.FinancialInOutType) (int64, error) {
	// 先通过财务账号id查询账号出来，然后查询出来的当前财务账号版本为修改条件
	account, err := s.GetAccountDetailById(ctx, id) // 如果不存在，会创建
	if err != nil {
		return 0, err
	}

	// 版本
	version := account.Data().Version

	db := s.dao.FdAccountDetail.Ctx(ctx)

	now := gtime.Now()

	data := co_do.FdAccountDetail{
		// gdb.Raw是字符串类型，该类型的参数将会直接作为SQL片段嵌入到提交到底层的SQL语句中，不会被自动转换为字符串参数类型、也不会被当做预处理参数
		// Increment  自增
		// Decrement  自减
		Version:          gdb.Raw(s.dao.FdAccountDetail.Columns().Version + "+1"),
		TodayUpdatedAt:   now,
		WeekUpdatedAt:    now,
		MonthUpdatedAt:   now,
		QuarterUpdatedAt: now,
		YearUpdatedAt:    now,
	}
	operator := " + "
	if (inOutType.Code() & co_enum.Financial.InOutType.Out.Code()) == co_enum.Financial.InOutType.Out.Code() { // 支出
		operator = " - "
	}

	// 判断是否是今日统计
	if account.FdAccountDetail.TodayUpdatedAt.Format("Y-m-d") != now.Format("Y-m-d") {
		data.TodayAccountSum = amount
	} else {
		data.TodayAccountSum = gdb.Raw(s.dao.FdAccountDetail.Columns().TodayAccountSum + operator + gconv.String(amount))
	}

	if account.WeekUpdatedAt.Format("Y-W") != now.Format("Y-W") {
		data.WeekAccountSum = amount
	} else {
		data.WeekAccountSum = gdb.Raw(s.dao.FdAccountDetail.Columns().WeekAccountSum + operator + gconv.String(amount))
	}

	if account.MonthUpdatedAt.Format("Y-m") != now.Format("Y-m") {
		data.MonthAccountSum = amount
	} else {
		data.MonthAccountSum = gdb.Raw(s.dao.FdAccountDetail.Columns().MonthAccountSum + operator + gconv.String(amount))
	}

	// 季度
	quarter := format_utils.GetQuarter(account.QuarterUpdatedAt)
	quarter2 := format_utils.GetQuarter(now)
	if account.QuarterUpdatedAt.Year() == now.Year() && quarter != quarter2 {
		data.QuarterAccountSum = amount
	} else {
		data.QuarterAccountSum = gdb.Raw(s.dao.FdAccountDetail.Columns().QuarterAccountSum + operator + gconv.String(amount))
	}

	if account.YearUpdatedAt.Year() != now.Year() {
		data.YearAccountSum = amount
	} else {
		data.YearAccountSum = gdb.Raw(s.dao.FdAccountDetail.Columns().YearAccountSum + operator + gconv.String(amount))
	}

	result, err := db.Data(data).Where(co_do.FdAccountDetail{
		Id:      id,
		Version: version,
	}).Update()

	if err != nil {
		return 0, err
	}

	affected, err := result.RowsAffected()

	return affected, err
}

// 添加财务账号附加数据 - 明细信息

// makeMore 按需加载附加数据
func makeMore[TR co_model.IFdAccountRes](ctx context.Context, dao co_dao.FdAccountDetailDao, info TR) TR {
	if reflect.ValueOf(info).IsNil() {
		return info
	}

	funs.AttrMake[TR](ctx,
		"id",
		func() TR {
			g.Try(ctx, func(ctx context.Context) {
				accountDetail, err := daoctl.GetByIdWithError[co_entity.FdAccountDetail](dao.Ctx(ctx), info.Data().FdAccount.Id)
				if err != nil {
					return
				}

				info.Data().Detail = accountDetail
			})
			return info
		},
	)
	return info
}

// QueryDetailByUnionUserIdAndSceneType  获取用户指定业务场景的财务账号金额明细统计记录|列表
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
]) QueryDetailByUnionUserIdAndSceneType(ctx context.Context, unionUserId int64, sceneType co_enum.SceneType) (*base_model.CollectRes[co_model.FdAccountDetailRes], error) {
	if unionUserId == 0 {
		return nil, gerror.New(s.modules.T(ctx, "error_Financial_UnionUserId_Failed"))
	}

	// 这是有缓存的情况，但是实际不能缓存
	// result, err := daoctl.Query[co_model.FdAccountDetailRes](s.dao.FdAccountDetail.Ctx(ctx).Where(co_do.FdAccountDetail{
	//    SysUserId: unionUserId,
	//    SceneType: sceneType,
	// }), nil, false)

	result, err := daoctl.Query[co_model.FdAccountDetailRes](s.dao.FdAccountDetail.Ctx(ctx), &base_model.SearchParams{
		Filter: append(make([]base_model.FilterInfo, 0),
			base_model.FilterInfo{
				Field: s.dao.FdAccountDetail.Columns().SysUserId,
				Where: "=",
				Value: unionUserId,
			},
			base_model.FilterInfo{
				Field: s.dao.FdAccountDetail.Columns().SceneType,
				Where: "=",
				Value: sceneType.Code(),
			},
		),
	}, false)

	if err != nil {
		return nil, sys_service.SysLogs().ErrorSimple(ctx, err, s.modules.T(ctx, "{#AccountDetail}{#error_Data_Get_Failed}"), s.dao.FdAccountDetail.Table())
	}

	return result, nil
}
