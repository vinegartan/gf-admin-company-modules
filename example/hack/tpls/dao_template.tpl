// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package co_dao

import (
	"github.com/SupenBysz/gf-admin-community/utility/daoctl/dao_interface"
	"{TplImportPrefix}/internal"
)

type {TplTableNameCamelCase} = dao_interface.TIDao[internal.{TplTableNameCamelCase}Columns]

func New{TplTableNameCamelCase}(dao ...dao_interface.IDao) {TplTableNameCamelCase} {
	return ({TplTableNameCamelCase})(internal.New{TplTableNameCamelCase}Dao(dao...))
}
