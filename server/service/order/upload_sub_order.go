package order

import (
	"errors"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/company"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	product2 "github.com/flipped-aurora/gin-vue-admin/server/model/product"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"strconv"
)

type UploadSubOrderService struct {
}

// CreateUploadSubOrder 创建UploadSubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (upSubOrdService *UploadSubOrderService) CreateUploadSubOrder(upSubOrd order.UploadSubOrder) (err error) {
	err = global.GVA_DB.Create(&upSubOrd).Error
	return err
}

// DeleteUploadSubOrder 删除UploadSubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (upSubOrdService *UploadSubOrderService) DeleteUploadSubOrder(upSubOrd order.UploadSubOrder) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&order.UploadSubOrder{}).Where("id = ?", upSubOrd.ID).Update("deleted_by", upSubOrd.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&upSubOrd).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteUploadSubOrderByIds 批量删除UploadSubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (upSubOrdService *UploadSubOrderService) DeleteUploadSubOrderByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&order.UploadSubOrder{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&order.UploadSubOrder{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (upSubOrdService *UploadSubOrderService) DeleteUploadSubOrderByOrderId(orderId *int, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&order.UploadSubOrder{}).Where("order_id = ?", orderId).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("order_id = ?", orderId).Delete(&order.UploadSubOrder{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateUploadSubOrder 更新UploadSubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (upSubOrdService *UploadSubOrderService) UpdateUploadSubOrder(upSubOrd order.UploadSubOrder) (err error) {
	err = global.GVA_DB.Save(&upSubOrd).Error
	return err
}

func (upSubOrdService *UploadSubOrderService) ParseExcel2InfoList(orderId *int) (err error) {
	fixedHeader := []string{"CODE", "PRODUCT NAME", "PACKING", "EXP", "QTY"}
	file, err := excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + strconv.Itoa(*orderId) + ".xlsx")
	if err != nil {
		return err
	}

	rows, err := file.Rows(file.GetSheetName(0))
	if err != nil {
		return err
	}

	bFoundData := false
	for rows.Next() {
		row, err := rows.Columns()
		if err != nil {
			return err
		}

		// 判断数据有效性
		if !bFoundData {
			if row[0] == fixedHeader[0] {
				bFoundData = true
			}
			continue
		}

		if len(row) > 4 {
			var quantity int
			quantity, _ = strconv.Atoi(row[4])
			if quantity > 0 {
				var uploadSubOrder order.UploadSubOrder
				uploadSubOrder.ProductCode = row[0]
				var product product2.Product
				if err := global.GVA_DB.Where("code = ?", row[0]).Find(&product).Error; err != nil {
					return err
				}
				uploadSubOrder.ProductNameCn = product.Product_name_cn
				uploadSubOrder.ProductId = product.ID
				uploadSubOrder.Package = product.Package
				uploadSubOrder.ExpDate = product.Exp_date
				uploadSubOrder.ProductNameEn = product.Product_name_en
				uploadSubOrder.Store = product.Store
				uploadSubOrder.Quantity = &quantity
				uploadSubOrder.OrderId = orderId

				if err := global.GVA_DB.Create(&uploadSubOrder).Error; err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// GetUploadSubOrder 根据id获取UploadSubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (upSubOrdService *UploadSubOrderService) GetUploadSubOrder(id uint) (upSubOrd order.UploadSubOrder, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&upSubOrd).Error
	return
}

// GetUploadSubOrderInfoList 分页获取UploadSubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (upSubOrdService *UploadSubOrderService) GetUploadSubOrderInfoList(info orderReq.UploadSubOrderSearch) (list []order.UploadSubOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&order.UploadSubOrder{})
	var upSubOrds []order.UploadSubOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.ProductCode != "" {
		db = db.Where("product_code LIKE ?", "%"+info.ProductCode+"%")
	}
	if info.OrderId != nil {
		db = db.Where("order_id = ?", info.OrderId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["productCode"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&upSubOrds).Error
	return upSubOrds, total, err
}

func (upSubOrdService *UploadSubOrderService) DoExecImport(orderId orderReq.OrderIdReq) (err error) {
	var uploadSubOrd []order.UploadSubOrder
	if err := global.GVA_DB.Where("order_id = ? and deleted_by = 0", orderId.OrderId).Find(&uploadSubOrd).Error; err != nil {
		return err
	}

	// 导入
	var ord order.Order
	if err = global.GVA_DB.Where("id = ?", orderId.OrderId).First(&ord).Error; err != nil {
		return err
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		for i, j := 0, len(uploadSubOrd); i < j; i++ {
			var productId = uploadSubOrd[i].ProductId
			// 查询商品是否存在
			var prod product2.Product
			if err = global.GVA_DB.Where("id = ?", productId).Find(&prod).Error; err != nil {
				return err
			}
			if prod.ID == 0 || prod.Code == "" {
				// 商品不存在，跳转下一个商品
				continue
			}

			// 查询子订单表中，此订单下是否已经存在此商品
			//var repeatedSubOrderList []order.SubOrder
			//err = global.GVA_DB.Where("order_id = ? and product_id = ?", orderId.OrderId, productId).Find(&repeatedSubOrderList).Error
			//if len(repeatedSubOrderList) > 0 {
			//	continue // 已经存在，跳过此商品，不导入
			//}

			var subOrder order.SubOrder // 需要创建的subOrder
			ordId, _ := strconv.Atoi(orderId.OrderId)
			subOrder.Order_id = &ordId
			subOrder.Product_name_cn = prod.Product_name_cn
			subOrder.Product_id = prod.ID
			subOrder.Product_code = prod.Code
			subOrder.Price = prod.Price
			subOrder.Product_name_en = prod.Product_name_en
			quantity := uploadSubOrd[i].Quantity
			subTotal, _ := decimal.NewFromFloat(*prod.Price).Mul(decimal.NewFromInt(int64(*quantity))).Float64()
			subOrder.Sub_total = &subTotal
			subOrder.Exp_date = prod.Exp_date
			subOrder.Package = prod.Package
			subOrder.Quantity = uploadSubOrd[i].Quantity
			subOrder.Vat = prod.Vat
			subVat, _ := decimal.NewFromFloat(*prod.Vat).Mul(decimal.NewFromInt(int64(*quantity))).Float64()
			subOrder.Sub_Vat = &subVat

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
					subOrder.Discount = compDiscountList[0].Discount
					discount, _ := decimal.NewFromFloat(*compDiscountList[0].Discount).Mul(decimal.NewFromInt(int64(*quantity))).Float64()
					subOrder.Discount_total = &discount
				}
			}
			if err := global.GVA_DB.Create(&subOrder).Error; err != nil {
				return err
			}

			if err = global.GVA_DB.Exec("update product set store = store - ? where ID = ?", uploadSubOrd[i].Quantity, prod.ID).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}
