package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"gorm.io/gorm"
)

type OrderService struct {
}

// CreateOrder 创建Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordService *OrderService) CreateOrder(ord order.Order) (err error) {
	err = global.GVA_DB.Create(&ord).Error
	return err
}

// DeleteOrder 删除Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordService *OrderService) DeleteOrder(ord order.Order) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&order.Order{}).Where("id = ?", ord.ID).Update("deleted_by", ord.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&ord).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteOrderByIds 批量删除Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordService *OrderService) DeleteOrderByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&order.Order{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&order.Order{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateOrder 更新Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordService *OrderService) UpdateOrder(ord order.Order) (err error) {
	err = global.GVA_DB.Save(&ord).Error
	return err
}

// GetOrder 根据id获取Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordService *OrderService) GetOrder(id uint) (ord order.Order, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ord).Error
	return
}

// GetOrderInfoList 分页获取Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordService *OrderService) GetOrderInfoList(info orderReq.OrderSearch) (list []order.Order, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&order.Order{})
	var ords []order.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.StartOrder_date != nil && info.EndOrder_date != nil {
		db = db.Where("order_date BETWEEN ? AND ? ", info.StartOrder_date, info.EndOrder_date)
	}
	if info.Invoice_no != "" {
		db = db.Where("invoice_no LIKE ?", "%"+info.Invoice_no+"%")
	}
	if info.Payment_method != "" {
		db = db.Where("payment_method LIKE ?", "%"+info.Payment_method+"%")
	}
	if info.Po_number != "" {
		db = db.Where("po_number LIKE ?", "%"+info.Po_number+"%")
	}
	if info.StartInvoice_date != nil && info.EndInvoice_date != nil {
		db = db.Where("invoice_date BETWEEN ? AND ? ", info.StartInvoice_date, info.EndInvoice_date)
	}
	if info.Bill_to != "" {
		db = db.Where("bill_to LIKE ?", "%"+info.Bill_to+"%")
	}
	if info.Ship_to != "" {
		db = db.Where("ship_to LIKE ?", "%"+info.Ship_to+"%")
	}
	if info.StartQuantity != nil && info.EndQuantity != nil {
		db = db.Where("quantity BETWEEN ? AND ? ", info.StartQuantity, info.EndQuantity)
	}
	if info.StartSubtotal != nil && info.EndSubtotal != nil {
		db = db.Where("subtotal BETWEEN ? AND ? ", info.StartSubtotal, info.EndSubtotal)
	}
	if info.StartVat != nil && info.EndVat != nil {
		db = db.Where("vat BETWEEN ? AND ? ", info.StartVat, info.EndVat)
	}
	if info.StartDiscount != nil && info.EndDiscount != nil {
		db = db.Where("discount BETWEEN ? AND ? ", info.StartDiscount, info.EndDiscount)
	}
	if info.StartOrder_total != nil && info.EndOrder_total != nil {
		db = db.Where("order_total BETWEEN ? AND ? ", info.StartOrder_total, info.EndOrder_total)
	}
	if info.Bill_to_address != "" {
		db = db.Where("bill_to_address LIKE ?", "%"+info.Bill_to_address+"%")
	}
	if info.Bill_to_city != "" {
		db = db.Where("bill_to_city LIKE ?", "%"+info.Bill_to_city+"%")
	}
	if info.Bill_to_phone != "" {
		db = db.Where("bill_to_phone LIKE ?", "%"+info.Bill_to_phone+"%")
	}
	if info.Bill_to_postcode != "" {
		db = db.Where("bill_to_postcode LIKE ?", "%"+info.Bill_to_postcode+"%")
	}
	if info.Ship_to_contact_name != "" {
		db = db.Where("ship_to_contact_name LIKE ?", "%"+info.Ship_to_contact_name+"%")
	}
	if info.Ship_to_address != "" {
		db = db.Where("ship_to_address LIKE ?", "%"+info.Ship_to_address+"%")
	}
	if info.Ship_to_city != "" {
		db = db.Where("ship_to_city LIKE ?", "%"+info.Ship_to_city+"%")
	}
	if info.Ship_to_phone != "" {
		db = db.Where("ship_to_phone LIKE ?", "%"+info.Ship_to_phone+"%")
	}
	if info.Ship_to_postcode != "" {
		db = db.Where("ship_to_postcode LIKE ?", "%"+info.Ship_to_postcode+"%")
	}
	if info.Is_locked != nil {
		db = db.Where("is_locked = ?", info.Is_locked)
	}
	if info.StartHand_discount != nil && info.EndHand_discount != nil {
		db = db.Where("hand_discount BETWEEN ? AND ? ", info.StartHand_discount, info.EndHand_discount)
	}
	if info.StartCustomer_id != nil && info.EndCustomer_id != nil {
		db = db.Where("customer_id BETWEEN ? AND ? ", info.StartCustomer_id, info.EndCustomer_id)
	}
	if info.Customer_company_name != "" {
		db = db.Where("customer_company_name LIKE ?", "%"+info.Customer_company_name+"%")
	}
	if info.Customer_contact_name != "" {
		db = db.Where("customer_contact_name LIKE ?", "%"+info.Customer_contact_name+"%")
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&ords).Error
	return ords, total, err
}
