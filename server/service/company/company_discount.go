package company

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/company"
	companyReq "github.com/flipped-aurora/gin-vue-admin/server/model/company/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/product"
	"gorm.io/gorm"
)

type CompanyDiscountService struct {
}

// CreateCompanyDiscount 创建CompanyDiscount记录
// Author [piexlmax](https://github.com/piexlmax)
func (compDiscountService *CompanyDiscountService) CreateCompanyDiscount(compDiscount company.CompanyDiscount) (err error) {
	err = global.GVA_DB.Create(&compDiscount).Error
	return err
}

// DeleteCompanyDiscount 删除CompanyDiscount记录
// Author [piexlmax](https://github.com/piexlmax)
func (compDiscountService *CompanyDiscountService) DeleteCompanyDiscount(compDiscount company.CompanyDiscount) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&company.CompanyDiscount{}).Where("id = ?", compDiscount.ID).Update("deleted_by", compDiscount.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&compDiscount).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteCompanyDiscountByIds 批量删除CompanyDiscount记录
// Author [piexlmax](https://github.com/piexlmax)
func (compDiscountService *CompanyDiscountService) DeleteCompanyDiscountByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&company.CompanyDiscount{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&company.CompanyDiscount{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateCompanyDiscount 更新CompanyDiscount记录
// Author [piexlmax](https://github.com/piexlmax)
func (compDiscountService *CompanyDiscountService) UpdateCompanyDiscount(compDiscount company.CompanyDiscount) (err error) {
	err = global.GVA_DB.Save(&compDiscount).Error
	return err
}

// GetCompanyDiscount 根据id获取CompanyDiscount记录
// Author [piexlmax](https://github.com/piexlmax)
func (compDiscountService *CompanyDiscountService) GetCompanyDiscount(id uint) (compDiscount company.CompanyDiscount, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&compDiscount).Error
	return
}

// GetCompanyDiscountInfoList 分页获取CompanyDiscount记录
// Author [piexlmax](https://github.com/piexlmax)
func (compDiscountService *CompanyDiscountService) GetCompanyDiscountInfoList(info companyReq.CompanyDiscountSearch) (list []company.CompanyDiscount, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&company.CompanyDiscount{})
	var compDiscounts []company.CompanyDiscount
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.CompanyId != nil {
		db = db.Where("company_id = ?", info.CompanyId)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	var OrderStr string
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

	err = db.Limit(limit).Offset(offset).Find(&compDiscounts).Error
	return compDiscounts, total, err
}

// AddCompDiscountByProductIds 根据商品IDS添加订单商品
func (compDiscountService *CompanyDiscountService) AddCompDiscountByProductIds(companyId *int, ids request.IdsReq, createdBy uint) (err error) {
	var prodList []product.Product
	err = global.GVA_DB.Where("id in ?", ids.Ids).Find(&prodList).Error
	if err != nil {
		return err
	} else {
		if len(prodList) > 0 {
			// 查询子订单表中，此订单下是否已经存在商品列表中的某些商品
			var repeatedCompDiscountList []company.CompanyDiscount
			err = global.GVA_DB.Where("company_id = ? and product_id in ?", companyId, ids.Ids).Find(&repeatedCompDiscountList).Error
			if err != nil {
				return err
			} else {
				// 查询公司表，获取公司名称
				var comp company.Company
				err = global.GVA_DB.Where("ID = ?", companyId).Find(&comp).Error
				if err != nil {
					return err
				} else {
					var compDiscounterList []company.CompanyDiscount
					for i, n := 0, len(prodList); i < n; i++ {
						var prod = prodList[i]
						// 判断此商品是否已经存在，存在则跳过，不存在才添加
						isExist := false
						for j, k := 0, len(repeatedCompDiscountList); j < k; j++ {
							if prod.ID == repeatedCompDiscountList[j].ProductId {
								isExist = true
								break
							}
						}
						if isExist {
							continue
						}
						var compDiscounter company.CompanyDiscount
						compDiscounter.CompanyId = companyId
						compDiscounter.ProductNameCn = prod.Product_name_cn
						compDiscounter.ProductId = prod.ID
						compDiscounter.ProductCode = prod.Code
						compDiscounter.Price = prod.Price
						compDiscounter.ProductNameEn = prod.Product_name_en
						compDiscounter.CompanyName = comp.Company_name
						compDiscounter.Price = prod.Price
						dis := float64(0)
						compDiscounter.Discount = &dis

						compDiscounterList = append(compDiscounterList, compDiscounter)
					}
					err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
						if err := global.GVA_DB.CreateInBatches(compDiscounterList, len(compDiscounterList)).Error; err != nil {
							return err
						}
						return nil
					})
				}
			}
		}
		return nil
	}
}
