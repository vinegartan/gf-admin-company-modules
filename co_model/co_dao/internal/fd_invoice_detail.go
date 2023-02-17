// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"context"

	"github.com/SupenBysz/gf-admin-community/utility/daoctl"
	"github.com/SupenBysz/gf-admin-community/utility/daoctl/dao_interface"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

// FdInvoiceDetailDao is the data access object for table co_fd_invoice_detail.
type FdInvoiceDetailDao struct {
	table   string                 // table is the underlying table name of the DAO.
	group   string                 // group is the database configuration group name of current DAO.
	columns FdInvoiceDetailColumns // columns contains all the column names of Table for convenient usage.
}

// FdInvoiceDetailColumns defines and stores column names for table co_fd_invoice_detail.
type FdInvoiceDetailColumns struct {
	Id            string // ID
	TaxNumber     string // 纳税识别号
	TaxName       string // 纳税人名称
	BillIds       string // 账单ID组
	Amount        string // 开票金额，单位精度：分
	Rate          string // 税率，如3% 则填入3
	RateMount     string // 税额，单位精度：分
	Remark        string // 发布内容描述
	Type          string // 发票类型：1电子发票，2纸质发票
	State         string // 状态：1待审核、2待开票、4开票失败、8已开票、16已撤销
	AuditUserIds  string // 审核者UserID，多个用逗号隔开
	MakeType      string // 出票类型：1普通发票、2增值税专用发票、3专业发票
	MakeUserId    string // 出票人UserID，如果是系统出票则默认-1
	MakeAt        string // 出票时间
	CourierName   string // 快递名称，限纸质发票
	CourierNumber string // 快递编号，限纸质发票
	FdInvoiceId   string // 发票抬头ID
	AuditUserId   string // 审核者UserID
	AuditReplyMsg string // 审核回复，仅审核不通过时才有值
	AuditAt       string // 审核时间
	CreatedAt     string //
	UpdatedAt     string //
	DeletedAt     string //
	UserId        string // 申请者用户ID
	UnionMainId   string // 主体ID：运营商ID、服务商ID、商户ID、消费者ID
	Email         string // 发票收件邮箱，限电子发票
}

// fdInvoiceDetailColumns holds the columns for table co_fd_invoice_detail.
var fdInvoiceDetailColumns = FdInvoiceDetailColumns{
	Id:            "id",
	TaxNumber:     "tax_number",
	TaxName:       "tax_name",
	BillIds:       "bill_ids",
	Amount:        "amount",
	Rate:          "rate",
	RateMount:     "rate_mount",
	Remark:        "remark",
	Type:          "type",
	State:         "state",
	AuditUserIds:  "audit_user_ids",
	MakeType:      "make_type",
	MakeUserId:    "make_user_id",
	MakeAt:        "make_at",
	CourierName:   "courier_name",
	CourierNumber: "courier_number",
	FdInvoiceId:   "fd_invoice_id",
	AuditUserId:   "audit_user_id",
	AuditReplyMsg: "audit_reply_msg",
	AuditAt:       "audit_at",
	CreatedAt:     "created_at",
	UpdatedAt:     "updated_at",
	DeletedAt:     "deleted_at",
	UserId:        "user_id",
	UnionMainId:   "union_main_id",
	Email:         "email",
}

// NewFdInvoiceDetailDao creates and returns a new DAO object for table data access.
func NewFdInvoiceDetailDao(proxy ...dao_interface.IDao) *FdInvoiceDetailDao {
	var dao *FdInvoiceDetailDao
	if len(proxy) > 0 {
		dao = &FdInvoiceDetailDao{
			group:   proxy[0].Group(),
			table:   proxy[0].Table(),
			columns: fdInvoiceDetailColumns,
		}
		return dao
	}

	return &FdInvoiceDetailDao{
		group:   "default",
		table:   "co_fd_invoice_detail",
		columns: fdInvoiceDetailColumns,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *FdInvoiceDetailDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *FdInvoiceDetailDao) Table() string {
	return dao.table
}

// Group returns the configuration group name of database of current dao.
func (dao *FdInvoiceDetailDao) Group() string {
	return dao.group
}

// Columns returns all column names of current dao.
func (dao *FdInvoiceDetailDao) Columns() FdInvoiceDetailColumns {
	return dao.columns
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *FdInvoiceDetailDao) Ctx(ctx context.Context, cacheOption ...*gdb.CacheOption) *gdb.Model {
	return dao.DaoConfig(ctx, cacheOption...).Model
}

func (dao *FdInvoiceDetailDao) DaoConfig(ctx context.Context, cacheOption ...*gdb.CacheOption) dao_interface.DaoConfig {
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
func (dao *FdInvoiceDetailDao) Transaction(ctx context.Context, f func(ctx context.Context, tx gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
