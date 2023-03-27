// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package co_dao

import (
	"github.com/SupenBysz/gf-admin-company-modules/co_model/co_dao/internal"
	"github.com/kysion/base-library/utility/daoctl/dao_interface"
)

type FdInvoiceDao = dao_interface.TIDao[internal.FdInvoiceColumns]

func NewFdInvoice(dao ...dao_interface.IDao) FdInvoiceDao {
	return (FdInvoiceDao)(internal.NewFdInvoiceDao(dao...))
}

var (
	FdInvoice = NewFdInvoice()
)
