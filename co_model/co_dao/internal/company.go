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

// CompanyDao is the data access object for table co_company.
type CompanyDao struct {
	table   string         // table is the underlying table name of the DAO.
	group   string         // group is the database configuration group name of current DAO.
	columns CompanyColumns // columns contains all the column names of Table for convenient usage.
}

// CompanyColumns defines and stores column names for table co_company.
type CompanyColumns struct {
	Id            string // ID
	Name          string // 名称
	ContactName   string // 商务联系人
	ContactMobile string // 商务联系电话
	UserId        string // 管理员ID
	State         string // 状态：0未启用，1正常
	Remark        string // 备注
	CreatedBy     string // 创建者
	CreatedAt     string // 创建时间
	UpdatedBy     string // 更新者
	UpdatedAt     string // 更新时间
	DeletedBy     string // 删除者
	DeletedAt     string // 删除时间
	ParentId      string // 父级ID
}

// companyColumns holds the columns for table co_company.
var companyColumns = CompanyColumns{
	Id:            "id",
	Name:          "name",
	ContactName:   "contact_name",
	ContactMobile: "contact_mobile",
	UserId:        "user_id",
	State:         "state",
	Remark:        "remark",
	CreatedBy:     "created_by",
	CreatedAt:     "created_at",
	UpdatedBy:     "updated_by",
	UpdatedAt:     "updated_at",
	DeletedBy:     "deleted_by",
	DeletedAt:     "deleted_at",
	ParentId:      "parent_id",
}

// NewCompanyDao creates and returns a new DAO object for table data access.
func NewCompanyDao(proxy ...dao_interface.IDao) *CompanyDao {
	var dao *CompanyDao
	if len(proxy) > 0 {
		dao = &CompanyDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: companyColumns,
		}
		return dao
	}

	return &CompanyDao{
		group:   "default",
		table:   "co_company",
		columns: companyColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *CompanyDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *CompanyDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *CompanyDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *CompanyDao) Columns() CompanyColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *CompanyDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *CompanyDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) dao_interface.DaoConfig {
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
func (dao *CompanyDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
