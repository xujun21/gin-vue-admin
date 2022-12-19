// 自动生成模板UploadSubOrder
package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// UploadSubOrder 结构体
type UploadSubOrder struct {
	global.GVA_MODEL
	OrderId       *int       `json:"orderId" form:"orderId" gorm:"column:order_id;comment:;"`
	ProductId     uint       `json:"productId" form:"productId" gorm:"column:product_id;comment:;"`
	ProductCode   string     `json:"productCode" form:"productCode" gorm:"column:product_code;comment:;"`
	ProductNameCn string     `json:"productNameCn" form:"productNameCn" gorm:"column:product_name_cn;comment:;"`
	ProductNameEn string     `json:"productNameEn" form:"productNameEn" gorm:"column:product_name_en;comment:;"`
	Package       string     `json:"package" form:"package" gorm:"column:package;comment:;"`
	ExpDate       *time.Time `json:"expDate" form:"expDate" gorm:"column:exp_date;comment:;"`
	Quantity      uint       `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;"`
	Store         *int       `json:"store" form:"store" gorm:"column:store;comment:;"`
	CreatedBy     uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy     uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy     uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName UploadSubOrder 表名
func (UploadSubOrder) TableName() string {
	return "upload_sub_order"
}
