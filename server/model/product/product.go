// 自动生成模板Product
package product

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Product 结构体
type Product struct {
	global.GVA_MODEL
	Code            string     `json:"code" form:"code" gorm:"column:code;comment:;"`
	Product_name_cn string     `json:"product_name_cn" form:"product_name_cn" gorm:"column:product_name_cn;comment:;"`
	Product_name_en string     `json:"product_name_en" form:"product_name_en" gorm:"column:product_name_en;comment:;"`
	Package         string     `json:"package" form:"package" gorm:"column:package;comment:;"`
	Exp_date        *time.Time `json:"exp_date" form:"exp_date" gorm:"column:exp_date;comment:;"`
	Price           *float64   `json:"price" form:"price" gorm:"column:price;comment:;"`
	Vat             *float64   `json:"vat" form:"vat" gorm:"column:vat;comment:;"`
	Store           *int       `json:"store" form:"store" gorm:"column:store;comment:;"`
	Barcode         string     `json:"barcode" form:"barcode" gorm:"column:barcode;comment:;"`
	CreatedBy       uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy       uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy       uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Product 表名
func (Product) TableName() string {
	return "product"
}
