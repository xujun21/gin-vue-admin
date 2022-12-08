package product

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/product"
	productReq "github.com/flipped-aurora/gin-vue-admin/server/model/product/request"
	"gorm.io/gorm"
)

type ProductService struct {
}

// CreateProduct 创建Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (prodService *ProductService) CreateProduct(prod product.Product) (err error) {
	err = global.GVA_DB.Create(&prod).Error
	return err
}

// DeleteProduct 删除Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (prodService *ProductService) DeleteProduct(prod product.Product) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&product.Product{}).Where("id = ?", prod.ID).Update("deleted_by", prod.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&prod).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteProductByIds 批量删除Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (prodService *ProductService) DeleteProductByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&product.Product{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&product.Product{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateProduct 更新Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (prodService *ProductService) UpdateProduct(prod product.Product) (err error) {
	err = global.GVA_DB.Save(&prod).Error
	return err
}

// GetProduct 根据id获取Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (prodService *ProductService) GetProduct(id uint) (prod product.Product, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&prod).Error
	return
}

// GetProductInfoList 分页获取Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (prodService *ProductService) GetProductInfoList(info productReq.ProductSearch) (list []product.Product, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&product.Product{})
	var prods []product.Product
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	if info.Code != "" {
		db = db.Where("code LIKE ?", "%"+info.Code+"%")
	}
	if info.Product_name_cn != "" {
		db = db.Where("product_name_cn LIKE ?", "%"+info.Product_name_cn+"%")
	}
	if info.Product_name_en != "" {
		db = db.Where("product_name_en LIKE ?", "%"+info.Product_name_en+"%")
	}
	if info.Package != "" {
		db = db.Where("package LIKE ?", "%"+info.Package+"%")
	}
	if info.StartExp_date != nil && info.EndExp_date != nil {
		db = db.Where("exp_date BETWEEN ? AND ? ", info.StartExp_date, info.EndExp_date)
	}
	if info.StartPrice != nil && info.EndPrice != nil {
		db = db.Where("price BETWEEN ? AND ? ", info.StartPrice, info.EndPrice)
	}
	if info.StartVat != nil && info.EndVat != nil {
		db = db.Where("vat BETWEEN ? AND ? ", info.StartVat, info.EndVat)
	}
	if info.StartStore != nil && info.EndStore != nil {
		db = db.Where("store BETWEEN ? AND ? ", info.StartStore, info.EndStore)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}
	var OrderStr string
	orderMap := make(map[string]bool)
	orderMap["code"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	}

	err = db.Limit(limit).Offset(offset).Find(&prods).Error
	return prods, total, err
}
