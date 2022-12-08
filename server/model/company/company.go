// 自动生成模板Company
package company

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Company 结构体
type Company struct {
	global.GVA_MODEL
	Company_name string `json:"company_name" form:"company_name" gorm:"column:company_name;comment:;"`
	Contact_name string `json:"contact_name" form:"contact_name" gorm:"column:contact_name;comment:;"`
	Phone        string `json:"phone" form:"phone" gorm:"column:phone;comment:;"`
	Address      string `json:"address" form:"address" gorm:"column:address;comment:;"`
	City         string `json:"city" form:"city" gorm:"column:city;comment:;"`
	Postcode     string `json:"postcode" form:"postcode" gorm:"column:postcode;comment:;"`
	Note         string `json:"note" form:"note" gorm:"column:note;comment:;"`
	Sage         string `json:"sage" form:"sage" gorm:"column:sage;comment:;"`
	CreatedBy    uint   `gorm:"column:created_by;comment:创建者"`
	UpdatedBy    uint   `gorm:"column:updated_by;comment:更新者"`
	DeletedBy    uint   `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Company 表名
func (Company) TableName() string {
	return "company"
}
