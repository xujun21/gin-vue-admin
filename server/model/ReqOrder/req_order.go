package ReqOrder

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// RequireOrder 结构体
type RequireOrder struct {
	global.GVA_MODEL
	// Order_date  *time.Time `json:"order_date" form:"order_date" gorm:"column:order_date;comment:;"`
	PONumber   string   `json:"poNumber" form:"poNumber" gorm:"column:po_number;comment:;"`
	Quantity   *int     `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;"`
	TotalCbm   *float64 `json:"totalCbm" form:"totalCbm" gorm:"column:total_cbm;comment:;"`
	TotalPrice *float64 `json:"totalPrice" form:"totalPrice" gorm:"column:total_price;comment:;"`
	CreatedBy  uint     `gorm:"column:created_by;comment:创建者"`
	UpdatedBy  uint     `gorm:"column:updated_by;comment:更新者"`
	DeletedBy  uint     `gorm:"column:deleted_by;comment:删除者"`
}

// TableName RequireOrder 表名
func (RequireOrder) TableName() string {
	return "require_order"
}
