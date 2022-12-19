// 自动生成模板CompanyDiscount
package company

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// CompanyDiscount 结构体
type CompanyDiscount struct {
	global.GVA_MODEL
	CompanyId     *int     `json:"companyId" form:"companyId" gorm:"column:company_id;comment:;"`
	CompanyName   string   `json:"companyName" form:"companyName" gorm:"column:company_name;comment:;"`
	ProductId     uint     `json:"productId" form:"productId" gorm:"column:product_id;comment:;"`
	ProductCode   string   `json:"productCode" form:"productCode" gorm:"column:product_code;comment:;"`
	ProductNameCn string   `json:"productNameCn" form:"productNameCn" gorm:"column:product_name_cn;comment:;"`
	ProductNameEn string   `json:"productNameEn" form:"productNameEn" gorm:"column:product_name_en;comment:;"`
	Price         *float64 `json:"price" form:"price" gorm:"column:price;comment:;"`
	Discount      *float64 `json:"discount" form:"discount" gorm:"column:discount;comment:;"`
	CreatedBy     uint     `gorm:"column:created_by;comment:创建者"`
	UpdatedBy     uint     `gorm:"column:updated_by;comment:更新者"`
	DeletedBy     uint     `gorm:"column:deleted_by;comment:删除者"`
}

// TableName CompanyDiscount 表名
func (CompanyDiscount) TableName() string {
	return "company_discount"
}
