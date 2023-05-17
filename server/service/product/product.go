package product

import (
	"fmt"
	"strconv"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/product"
	productReq "github.com/flipped-aurora/gin-vue-admin/server/model/product/request"
	"github.com/xuri/excelize/v2"
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
	if prod.CreatedAt.Before(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)) {
		prod.CreatedAt = time.Now()
	}
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
	if info.StartStore != nil {
		db = db.Where("store >= ?", info.StartStore)
	}
	if info.EndStore != nil {
		db = db.Where("store <= ?", info.EndStore)
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
	} else {
		db = db.Order("id desc")
	}

	err = db.Limit(limit).Offset(offset).Find(&prods).Error
	return prods, total, err
}

func (prodService *ProductService) ExportProductExcel(info productReq.ProductSearch, fileName string) (err error) {
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
	if info.StartStore != nil {
		db = db.Where("store >= ?", info.StartStore)
	}
	if info.EndStore != nil {
		db = db.Where("store <= ?", info.EndStore)
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
	} else {
		db = db.Order("id desc")
	}

	//err = db.Limit(limit).Offset(offset).Find(&prods).Error
	err = db.Find(&prods).Error
	if err != nil {
		return err
	}

	// 创建EXCEL
	var xlsx *excelize.File
	if xlsx, err = excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "template_product.xlsx"); err != nil {
		return err
	}
	var sheet = xlsx.GetSheetName(0)

	i := 0
	j := 0

	var strStyle, dateStyle, numStyle, moneyStyle [2]int
	strStyle[0], _ = xlsx.GetCellStyle(sheet, "A4")
	dateStyle[0], _ = xlsx.GetCellStyle(sheet, "D4")
	numStyle[0], _ = xlsx.GetCellStyle(sheet, "E4")
	moneyStyle[0], _ = xlsx.GetCellStyle(sheet, "F4")
	strStyle[1], _ = xlsx.GetCellStyle(sheet, "A5")
	dateStyle[1], _ = xlsx.GetCellStyle(sheet, "D5")
	numStyle[1], _ = xlsx.GetCellStyle(sheet, "E5")
	moneyStyle[1], _ = xlsx.GetCellStyle(sheet, "F5")

	for i, j = 0, len(prods); i < j; i++ {
		iStr := strconv.Itoa(4 + i)

		if err != nil {
			fmt.Println(err)
		}
		index := i % 2
		xlsx.SetCellStyle(sheet, "A"+iStr, "C"+iStr, strStyle[index])
		xlsx.SetCellStyle(sheet, "D"+iStr, "D"+iStr, dateStyle[index])
		xlsx.SetCellStyle(sheet, "E"+iStr, "E"+iStr, numStyle[index])
		xlsx.SetCellStyle(sheet, "F"+iStr, "F"+iStr, moneyStyle[index])
		xlsx.SetCellStyle(sheet, "G"+iStr, "G"+iStr, strStyle[index])

		xlsx.SetCellStr(sheet, "A"+iStr, prods[i].Code)
		xlsx.SetCellStr(sheet, "B"+iStr, prods[i].Product_name_cn+" "+prods[i].Product_name_en)
		xlsx.SetCellStr(sheet, "C"+iStr, prods[i].Package)
		if prods[i].Exp_date != nil {
			xlsx.SetCellValue(sheet, "D"+iStr, prods[i].Exp_date.Format("02/01/2006"))
		}
		xlsx.SetCellInt(sheet, "E"+iStr, *prods[i].Store)
		xlsx.SetCellValue(sheet, "F"+iStr, *prods[i].Price)
		xlsx.SetCellValue(sheet, "G"+iStr, prods[i].Barcode)
	}

	if *info.WithPrice == 0 {
		xlsx.RemoveCol(sheet, "G")
		xlsx.RemoveCol(sheet, "F")
	}

	err = xlsx.SaveAs(fileName)
	return err
}
