// 自动生成模板SubOrder
package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// SubOrder 结构体
type SubOrder struct {
	global.GVA_MODEL
	Order_id        *int       `json:"order_id" form:"order_id" gorm:"column:order_id;comment:;"`
	Product_id      uint       `json:"product_id" form:"product_id" gorm:"column:product_id;comment:;"`
	Product_code    string     `json:"product_code" form:"product_code" gorm:"column:product_code;comment:;"`
	Product_name_cn string     `json:"product_name_cn" form:"product_name_cn" gorm:"column:product_name_cn;comment:;"`
	Product_name_en string     `json:"product_name_en" form:"product_name_en" gorm:"column:product_name_en;comment:;"`
	Exp_date        *time.Time `json:"exp_date" form:"exp_date" gorm:"column:exp_date;comment:;"`
	Package         string     `json:"package" form:"package" gorm:"column:package;comment:;"`
	Price           *float64   `json:"price" form:"price" gorm:"column:price;comment:;"`
	Quantity        uint       `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;"`
	Vat             *float64   `json:"vat" form:"vat" gorm:"column:vat;comment:;"`
	Sub_Vat         *float64   `json:"sub_vat" form:"sub_vat" gorm:"column:sub_vat;comment:;"`
	Sub_total       *float64   `json:"sub_total" form:"sub_total" gorm:"column:sub_total;comment:商品金额;"`
	Discount        *float64   `json:"discount" form:"discount" gorm:"column:discount;comment:;"`
	Discount_total  *float64   `json:"discount_total" form:"discount_total" gorm:"column:discount_total;comment:;"`
	CreatedBy       uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy       uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy       uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName SubOrder 表名
func (SubOrder) TableName() string {
	return "sub_order"
}
