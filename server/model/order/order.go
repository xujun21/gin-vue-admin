// 自动生成模板Order
package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"time"
)

// Order 结构体
type Order struct {
	global.GVA_MODEL
	Order_date            *time.Time `json:"order_date" form:"order_date" gorm:"column:order_date;comment:;"`
	Invoice_no            string     `json:"invoice_no" form:"invoice_no" gorm:"column:invoice_no;comment:;"`
	Payment_method        string     `json:"payment_method" form:"payment_method" gorm:"column:payment_method;comment:;"`
	Po_number             string     `json:"po_number" form:"po_number" gorm:"column:po_number;comment:;"`
	Invoice_date          *time.Time `json:"invoice_date" form:"invoice_date" gorm:"column:invoice_date;comment:;"`
	Bill_to               string     `json:"bill_to" form:"bill_to" gorm:"column:bill_to;comment:;"`
	Ship_to               string     `json:"ship_to" form:"ship_to" gorm:"column:ship_to;comment:;"`
	Quantity              uint       `json:"quantity" form:"quantity" gorm:"column:quantity;comment:;"`
	Subtotal              *float64   `json:"subtotal" form:"subtotal" gorm:"column:subtotal;comment:单价*箱数，不含税;"`
	Vat                   *float64   `json:"vat" form:"vat" gorm:"column:vat;comment:;"`
	DeliveryFee           *float64   `json:"delivery_fee" form:"delivery_fee" gorm:"column:delivery_fee;comment:运费;"`
	Discount              *float64   `json:"discount" form:"discount" gorm:"column:discount;comment:;"`
	Order_total           *float64   `json:"order_total" form:"order_total" gorm:"column:order_total;comment:;"`
	Bill_to_contact_name  string     `json:"bill_to_contact_name" form:"bill_to_contact_name" gorm:"column:bill_to_contact_name;comment:;"`
	Bill_to_address       string     `json:"bill_to_address" form:"bill_to_address" gorm:"column:bill_to_address;comment:;"`
	Bill_to_city          string     `json:"bill_to_city" form:"bill_to_city" gorm:"column:bill_to_city;comment:;"`
	Bill_to_phone         string     `json:"bill_to_phone" form:"bill_to_phone" gorm:"column:bill_to_phone;comment:;"`
	Bill_to_postcode      string     `json:"bill_to_postcode" form:"bill_to_postcode" gorm:"column:bill_to_postcode;comment:;"`
	Ship_to_contact_name  string     `json:"ship_to_contact_name" form:"ship_to_contact_name" gorm:"column:ship_to_contact_name;comment:;"`
	Ship_to_address       string     `json:"ship_to_address" form:"ship_to_address" gorm:"column:ship_to_address;comment:;"`
	Ship_to_city          string     `json:"ship_to_city" form:"ship_to_city" gorm:"column:ship_to_city;comment:;"`
	Ship_to_phone         string     `json:"ship_to_phone" form:"ship_to_phone" gorm:"column:ship_to_phone;comment:;"`
	Ship_to_postcode      string     `json:"ship_to_postcode" form:"ship_to_postcode" gorm:"column:ship_to_postcode;comment:;"`
	Is_locked             *bool      `json:"is_locked" form:"is_locked" gorm:"column:is_locked;comment:;"`
	Hand_discount         *float64   `json:"hand_discount" form:"hand_discount" gorm:"column:hand_discount;comment:;"`
	Customer_id           *int       `json:"customer_id" form:"customer_id" gorm:"column:customer_id;comment:;"`
	Customer_company_name string     `json:"customer_company_name" form:"customer_company_name" gorm:"column:customer_company_name;comment:;"`
	Customer_contact_name string     `json:"customer_contact_name" form:"customer_contact_name" gorm:"column:customer_contact_name;comment:;"`
	CreatedBy             uint       `gorm:"column:created_by;comment:创建者"`
	UpdatedBy             uint       `gorm:"column:updated_by;comment:更新者"`
	DeletedBy             uint       `gorm:"column:deleted_by;comment:删除者"`
}

// TableName Order 表名
func (Order) TableName() string {
	return "order"
}
