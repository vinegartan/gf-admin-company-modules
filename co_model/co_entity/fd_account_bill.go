// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// FdAccountBill is the golang structure for table fd_account_bill.
type FdAccountBill struct {
	Id            int64       `json:"id"            description:"ID"`
	FromUserId    int64       `json:"fromUserId"    description:"交易发起方UserID，如果是系统则固定为-1"`
	ToUserId      int64       `json:"toUserId"      description:"交易对象UserID"`
	FdAccountId   int64       `json:"fdAccountId"   description:"财务账户ID"`
	BeforeBalance int64       `json:"beforeBalance" description:"交易前账户余额"`
	Amount        int64       `json:"amount"        description:"交易金额"`
	AfterBalance  int64       `json:"afterBalance"  description:"交易后账户余额"`
	UnionOrderId  int64       `json:"unionOrderId"  description:"关联业务订单ID"`
	InOutType     int         `json:"inOutType"     description:"收支类型：1收入，2支出"`
	TradeType     int         `json:"tradeType"     description:"交易类型，1转账、2消费、4退款、8佣金、16保证金、32诚意金、64手续费/服务费、128提现、256充值、512营收，8192其它"`
	TradeAt       *gtime.Time `json:"tradeAt"       description:"交易时间"`
	Remark        string      `json:"remark"        description:"备注信息"`
	TradeState    int         `json:"tradeState"    description:"交易状态：1待支付、2支付中、4已支付、8支付失败、16交易完成、"`
	CreatedAt     *gtime.Time `json:"createdAt"     description:""`
	DeletedAt     *gtime.Time `json:"deletedAt"     description:""`
}
