// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package co_entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// License is the golang structure for table license.
type License struct {
	Id                        int64       `json:"id"                        description:""`
	IdcardFrontPath           string      `json:"idcardFrontPath"           description:"身份证头像面照片"`
	IdcardBackPath            string      `json:"idcardBackPath"            description:"身份证国徽面照片"`
	IdcardNo                  string      `json:"idcardNo"                  description:"身份证号"`
	IdcardExpiredDate         *gtime.Time `json:"idcardExpiredDate"         description:"身份证有效期"`
	IdcardAddress             string      `json:"idcardAddress"             description:"身份证户籍地址"`
	PersonContactName         string      `json:"personContactName"         description:"负责人，必须是自然人"`
	PersonContactMobile       string      `json:"personContactMobile"       description:"负责人，联系电话"`
	BusinessLicenseName       string      `json:"businessLicenseName"       description:"公司全称"`
	BusinessLicenseAddress    string      `json:"businessLicenseAddress"    description:"公司地址"`
	BusinessLicensePath       string      `json:"businessLicensePath"       description:"营业执照图片地址"`
	BusinessLicenseScope      string      `json:"businessLicenseScope"      description:"经营范围"`
	BusinessLicenseRegCapital string      `json:"businessLicenseRegCapital" description:"注册资本"`
	BusinessLicenseTermTime   string      `json:"businessLicenseTermTime"   description:"营业期限"`
	BusinessLicenseOrgCode    string      `json:"businessLicenseOrgCode"    description:"组织机构代码"`
	BusinessLicenseCreditCode string      `json:"businessLicenseCreditCode" description:"统一社会信用代码"`
	BusinessLicenseLegal      string      `json:"businessLicenseLegal"      description:"法人"`
	BusinessLicenseLegalPath  string      `json:"businessLicenseLegalPath"  description:"法人证照，如果法人不是自然人，则该项必填"`
	LatestAuditLogId          int64       `json:"latestAuditLogId"          description:"最新的审核记录ID"`
	State                     int         `json:"state"                     description:""`
	AuthType                  int         `json:"authType"                  description:""`
	Remark                    string      `json:"remark"                    description:""`
	UpdatedAt                 *gtime.Time `json:"updatedAt"                 description:""`
	CreatedAt                 *gtime.Time `json:"createdAt"                 description:""`
	DeletedAt                 *gtime.Time `json:"deletedAt"                 description:""`
	BrandName                 string      `json:"brandName"                 description:"品牌名称"`
}
