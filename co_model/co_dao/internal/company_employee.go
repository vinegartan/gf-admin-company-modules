// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/kysion/base-library/utility/daoctl"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// CompanyEmployeeDao is the data access object for table co_company_employee.
type CompanyEmployeeDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns CompanyEmployeeColumns // columns contains all the column names of Table for convenient usage.
}

// CompanyEmployeeColumns defines and stores column names for table co_company_employee.
type CompanyEmployeeColumns struct {
	Id           string // ID，保持与USERID一致
	No           string // 工号
	Avatar       string // 头像
	Name         string // 姓名
	Mobile       string // 手机号
	UnionMainId  string // 所属主体
	State        string // 状态： -1已离职，0待确认，1已入职
	LastActiveIp string // 最后活跃IP
	HiredAt      string // 入职时间
	CreatedBy    string //
	CreatedAt    string //
	UpdatedBy    string //
	UpdatedAt    string //
	DeletedBy    string //
	DeletedAt    string //
}

// companyEmployeeColumns holds the columns for table co_company_employee.
var companyEmployeeColumns = CompanyEmployeeColumns{
	Id:           "id",
	No:           "no",
	Avatar:       "avatar",
	Name:         "name",
	Mobile:       "mobile",
	UnionMainId:  "union_main_id",
	State:        "state",
	LastActiveIp: "last_active_ip",
	HiredAt:      "hired_at",
	CreatedBy:    "created_by",
	CreatedAt:    "created_at",
	UpdatedBy:    "updated_by",
	UpdatedAt:    "updated_at",
	DeletedBy:    "deleted_by",
	DeletedAt:    "deleted_at",
}

// NewCompanyEmployeeDao creates and returns a new DAO object for table data access.
func NewCompanyEmployeeDao(proxy ...dao_interface.IDao) *CompanyEmployeeDao {
	var dao *CompanyEmployeeDao
	if len(proxy) > 0 {
		dao = &CompanyEmployeeDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: companyEmployeeColumns,
		}
		return dao
	}

	return &CompanyEmployeeDao{
		group:   "default",
		table:   "co_company_employee",
		columns: companyEmployeeColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CompanyEmployeeDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CompanyEmployeeDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *CompanyEmployeeDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *CompanyEmployeeDao) Columns() CompanyEmployeeColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CompanyEmployeeDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *CompanyEmployeeDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) dao_interface.DaoConfig {
	daoConfig := dao_interface.DaoConfig{
		Dao:   dao,
		DB:    dao.DB(),
		Table: dao.table,
		Group: dao.group,
		Model: dao.DB().Model(dao.Table()).Safe().Ctx(ctx),
	}

	if len(cacheOption) == 0 {
		daoConfig.CacheOption = daoctl.MakeDaoCache(dao.Table())
		daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
	} else {
		if cacheOption[0] != nil {
			daoConfig.CacheOption = cacheOption[0]
			daoConfig.Model = daoConfig.Model.Cache(*daoConfig.CacheOption)
		}
	}

	daoConfig.Model = daoctl.RegisterDaoHook(daoConfig.Model)

	return daoConfig
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *CompanyEmployeeDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
