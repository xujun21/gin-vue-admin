package ReqOrder

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ReqOrder"
	ReqOrderReq "github.com/flipped-aurora/gin-vue-admin/server/model/ReqOrder/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/product"
	"gorm.io/gorm"
)

type ReqSubOrderService struct {
}

// CreateReqSubOrder 创建ReqSubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (reqSubOrderService *ReqSubOrderService) CreateReqSubOrder(reqSubOrder ReqOrder.ReqSubOrder) (err error) {
	err = global.GVA_DB.Create(&reqSubOrder).Error
	return err
}

// DeleteReqSubOrder 删除ReqSubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (reqSubOrderService *ReqSubOrderService) DeleteReqSubOrder(reqSubOrder ReqOrder.ReqSubOrder) (err error) {
	err = global.GVA_DB.Delete(&reqSubOrder).Error
	return err
}

// DeleteReqSubOrderByIds 批量删除ReqSubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (reqSubOrderService *ReqSubOrderService) DeleteReqSubOrderByIds(ids request.IdsReq) (err error) {
	err = global.GVA_DB.Delete(&[]ReqOrder.ReqSubOrder{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateReqSubOrder 更新ReqSubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (reqSubOrderService *ReqSubOrderService) UpdateReqSubOrder(reqSubOrder ReqOrder.ReqSubOrder) (err error) {
	err = global.GVA_DB.Save(&reqSubOrder).Error
	return err
}

func (reqSubOrderService *ReqSubOrderService) UpdateReqSubOrderInQty(reqOrderId uint, inQty *int) (err error) {
	var reqSubOrder ReqOrder.ReqSubOrder
	if err = global.GVA_DB.Where("id = ?", reqOrderId).First(&reqSubOrder).Error; err != nil {
		return err
	}

	reqSubOrder.InQty = inQty
	err = global.GVA_DB.Save(&reqSubOrder).Error
	return err
}

// GetReqSubOrder 根据id获取ReqSubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (reqSubOrderService *ReqSubOrderService) GetReqSubOrder(id uint) (reqSubOrder ReqOrder.ReqSubOrder, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&reqSubOrder).Error
	return
}

// GetReqSubOrderInfoList 分页获取ReqSubOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (reqSubOrderService *ReqSubOrderService) GetReqSubOrderInfoList(info ReqOrderReq.ReqSubOrderSearch) (list []ReqOrder.ReqSubOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)

	// 创建db
	db := global.GVA_DB.Model(&ReqOrder.ReqSubOrder{})
	var reqSubOrders []ReqOrder.ReqSubOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}

	if info.ReqOrderId != nil {
		db = db.Where("req_order_id = ?", info.ReqOrderId)
	}

	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&reqSubOrders).Error
	return reqSubOrders, total, err
}

// AddReqSubOrderByProductIds 根据商品IDS添加订单商品
func (reqSubOrdService *ReqSubOrderService) AddReqSubOrderByProductIds(reqOrderId *int, ids request.IdsReq, createdBy uint) (err error) {
	var reqOrd ReqOrder.RequireOrder
	if err = global.GVA_DB.Where("id = ?", reqOrderId).First(&reqOrd).Error; err != nil {
		return err
	}

	var prodList []product.Product
	err = global.GVA_DB.Where("id in ?", ids.Ids).Find(&prodList).Error
	if err != nil {
		return err
	} else {
		if len(prodList) > 0 {
			var reqSubOrderList []ReqOrder.ReqSubOrder
			// var prodListOfReduceStore []uint // 存放需要扣减库存的商品id
			for i, n := 0, len(prodList); i < n; i++ {
				var prod = prodList[i]

				var reqSubOrder ReqOrder.ReqSubOrder
				reqSubOrder.ReqOrderId = reqOrderId
				reqSubOrder.ProductNameCn = prod.Product_name_cn
				reqSubOrder.ProductNameEn = prod.Product_name_en
				reqSubOrder.ProductId = prod.ID
				reqSubOrder.ProductCode = prod.Code
				reqSubOrder.BarcodeCase = prod.BarcodeCase
				reqSubOrder.Barcode = prod.Barcode
				reqSubOrder.CartonSize = prod.CartonSize
				reqSubOrder.Cbm = prod.Cbm
				reqSubOrder.InPrice = prod.InPrice
				reqSubOrder.SelfLife = prod.SelfLife
				reqSubOrder.Weight = prod.Weight

				reqSubOrder.CreatedBy = createdBy
				reqSubOrderList = append(reqSubOrderList, reqSubOrder)
			}
			err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
				if err := global.GVA_DB.CreateInBatches(reqSubOrderList, len(reqSubOrderList)).Error; err != nil {
					return err
				}
				return nil
			})
		}
		return nil
	}
}
