// 自动生成模板ReqSubOrder
package ReqOrder

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// ReqSubOrder 结构体
type ReqSubOrder struct {
	global.GVA_MODEL
	BarcodeCase   string   `json:"barcodeCase" form:"barcodeCase" gorm:"column:barcode_case;comment:;size:191;"`
	Barcode       string   `json:"barcode" form:"barcode" gorm:"column:barcode;comment:;size:191;"`
	CartonSize    string   `json:"cartonSize" form:"cartonSize" gorm:"column:carton_size;comment:;size:191;"`
	Cbm           *float64 `json:"cbm" form:"cbm" gorm:"column:cbm;comment:;size:22;"`
	CreatedBy     uint     `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建者;size:20;"`
	DeletedBy     uint     `json:"deletedBy" form:"deletedBy" gorm:"column:deleted_by;comment:删除者;size:20;"`
	InPrice       *float64 `json:"inPrice" form:"inPrice" gorm:"column:in_price;comment:采购价;size:22;"`
	ProductCode   string   `json:"productCode" form:"productCode" gorm:"column:product_code;comment:;size:191;"`
	ProductId     uint     `json:"productId" form:"productId" gorm:"column:product_id;comment:;size:20;"`
	ProductNameCn string   `json:"productNameCn" form:"productNameCn" gorm:"column:product_name_cn;comment:;size:191;"`
	ProductNameEn string   `json:"productNameEn" form:"productNameEn" gorm:"column:product_name_en;comment:;size:191;"`
	InQty         *int     `json:"inQty" form:"inQty" gorm:"column:in_qty;default:0;comment:采购数量;size:20;"`
	ReqOrderId    *int     `json:"reqOrderId" form:"reqOrderId" gorm:"column:req_order_id;comment:;size:19;"`
	SelfLife      string   `json:"selfLife" form:"selfLife" gorm:"column:self_life;comment:;size:191;"`
	UpdatedBy     *int     `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新者;size:20;"`
	Weight        *float64 `json:"weight" form:"weight" gorm:"column:weight;comment:;size:22;"`
}

// TableName ReqSubOrder 表名
func (ReqSubOrder) TableName() string {
	return "req_sub_order"
}
