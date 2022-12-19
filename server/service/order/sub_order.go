package order

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/company"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/product"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"strconv"
)

type SubOrderService struct {
}

// CreateSubOrder 创建SubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (subOrdService *SubOrderService) CreateSubOrder(subOrd order.SubOrder) (err error) {
	err = global.GVA_DB.Create(&subOrd).Error
	return err
}

// DeleteSubOrder 删除SubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (subOrdService *SubOrderService) DeleteSubOrder(subOrd order.SubOrder) (err error) {
	// 查询subOrder，找到即将删除的数据，记录各sub_order.quantity
	var subOrderOfDel order.SubOrder
	if err := global.GVA_DB.Where("id = ?", subOrd.ID).First(&subOrderOfDel).Error; err != nil {
		return err
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err = tx.Model(&order.SubOrder{}).Where("id = ?", subOrd.ID).Update("deleted_by", subOrd.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&subOrd).Error; err != nil {
			return err
		}
		// 返回库存
		if err = tx.Exec("update product set store = store + ? where id = ?", subOrderOfDel.Quantity, subOrderOfDel.Product_id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteSubOrderByIds 批量删除SubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (subOrdService *SubOrderService) DeleteSubOrderByIds(ids request.IdsReq, deleted_by uint) (err error) {
	// 查询subOrder，找到即将删除的数据，记录各sub_order.quantity
	var subOrderListOfDel []order.SubOrder
	if err := global.GVA_DB.Where("id in ?", ids.Ids).Find(&subOrderListOfDel).Error; err != nil {
		return err
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 删除记录
		if err = tx.Model(&order.SubOrder{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&order.SubOrder{}).Error; err != nil {
			return err
		}
		// 返回库存
		for i, j := 0, len(subOrderListOfDel); i < j; i++ {
			if err = tx.Exec("update product set store = store + ? where id = ?", subOrderListOfDel[i].Quantity, subOrderListOfDel[i].Product_id).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

// UpdateSubOrder 更新SubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (subOrdService *SubOrderService) UpdateSubOrder(subOrd order.SubOrder) (err error) {
	// 获取更新前subOrder中的productQuantity
	var oldSubOrder order.SubOrder
	if err := global.GVA_DB.Where("id = ?", subOrd.ID).First(&oldSubOrder).Error; err != nil {
		return err
	}
	oldQuantity := oldSubOrder.Quantity

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		// 更新subOrder
		if err = global.GVA_DB.Save(&subOrd).Error; err != nil {
			return err
		}
		// 更新商品库存
		if err = global.GVA_DB.Exec("update product set store = store + ? - ? where id = ?", oldQuantity, subOrd.Quantity, subOrd.Product_id).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// GetSubOrder 根据id获取SubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (subOrdService *SubOrderService) GetSubOrder(id uint) (subOrd order.SubOrder, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&subOrd).Error
	return
}

// GetSubOrderInfoList 分页获取SubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (subOrdService *SubOrderService) GetSubOrderInfoList(info orderReq.SubOrderSearch) (list []order.SubOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&order.SubOrder{})
	var subOrds []order.SubOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Product_code != "" {
		db = db.Where("product_code LIKE ?", "%"+info.Product_code+"%")
	}
	if info.Product_name_cn != "" {
		db = db.Where("product_name_cn LIKE ?", "%"+info.Product_name_cn+"%")
	}
	if info.Product_name_en != "" {
		db = db.Where("product_name_en LIKE ?", "%"+info.Product_name_en+"%")
	}
	if info.StartExp_date != nil && info.EndExp_date != nil {
		db = db.Where("exp_date BETWEEN ? AND ? ", info.StartExp_date, info.EndExp_date)
	}
	if info.Order_id != nil {
		db = db.Where("order_id = ?", info.Order_id)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string // 排序
	orderMap := make(map[string]bool)
	orderMap["product_code"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("id desc")
	}
	err = db.Limit(limit).Offset(offset).Find(&subOrds).Error
	return subOrds, total, err
}

// AddSubOrderByProductIds 根据商品IDS添加订单商品
func (subOrdService *SubOrderService) AddSubOrderByProductIds(orderId *int, ids request.IdsReq, createdBy uint) (err error) {
	var ord order.Order
	if err = global.GVA_DB.Where("id = ?", orderId).First(&ord).Error; err != nil {
		return err
	}

	var prodList []product.Product
	err = global.GVA_DB.Where("id in ?", ids.Ids).Find(&prodList).Error
	if err != nil {
		return err
	} else {
		if len(prodList) > 0 {
			// 查询子订单表中，此订单下是否已经存在商品列表中的某些商品
			var repeatedSubOrderList []order.SubOrder
			err = global.GVA_DB.Where("order_id = ? and product_id in ?", orderId, ids.Ids).Find(&repeatedSubOrderList).Error
			if err != nil {
				return err
			} else {
				var subOrderList []order.SubOrder
				var prodListOfReduceStore []uint // 存放需要扣减库存的商品id
				for i, n := 0, len(prodList); i < n; i++ {
					var prod = prodList[i]
					// 判断此商品是否已经存在，存在则跳过，不存在才添加
					isExist := false
					for j, k := 0, len(repeatedSubOrderList); j < k; j++ {
						if prod.ID == repeatedSubOrderList[j].Product_id {
							isExist = true
							break
						}
					}
					if isExist {
						continue
					}
					var subOrder order.SubOrder
					subOrder.Order_id = orderId
					subOrder.Product_name_cn = prod.Product_name_cn
					subOrder.Product_id = prod.ID
					subOrder.Product_code = prod.Code
					subOrder.Price = prod.Price
					subOrder.Product_name_en = prod.Product_name_en
					subOrder.Sub_total = prod.Price // 默认为1个
					subOrder.Exp_date = prod.Exp_date
					subOrder.Package = prod.Package
					subOrder.Quantity = 1
					subOrder.Vat = prod.Vat
					subOrder.Sub_Vat = prod.Vat // 1个

					subOrder.CreatedBy = createdBy

					var compDiscountList []company.CompanyDiscount
					if err = global.GVA_DB.Where("company_id = ? and product_id = ?",
						ord.Customer_id, prod.ID).Find(&compDiscountList).Error; err != nil {
						return err
					}

					if len(compDiscountList) > 0 {
						if len(compDiscountList) > 1 {
							return errors.New("数据错误！单一客户的某一商品折扣记录超过1" + strconv.Itoa(len(compDiscountList)))
						}
						if len(compDiscountList) == 1 {
							subOrder.Discount = compDiscountList[0].Discount // quantity = 1
							subOrder.Discount_total = compDiscountList[0].Discount
						}
					} else {
						zero, _ := decimal.Zero.Float64()
						subOrder.Discount = &zero
						subOrder.Discount_total = &zero
					}

					subOrderList = append(subOrderList, subOrder)
					// 扣减库存
					prodListOfReduceStore = append(prodListOfReduceStore, prod.ID)
				}
				err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
					if err := global.GVA_DB.CreateInBatches(subOrderList, len(subOrderList)).Error; err != nil {
						return err
					}

					if err = global.GVA_DB.Exec("update product set store = store - 1 where ID in ?", prodListOfReduceStore).Error; err != nil {
						return err
					}
					return nil
				})
			}
		}
		return nil
	}
}
