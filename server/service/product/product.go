package product

import (
	"errors"
	"fmt"
	"image"
	"os"
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
	// 先判断是否存在相同编码的商品(包括被软删除的商品)
	// 如果存在则报错，商品编码不能重复
	// 如果不存在则创建新商品
	var count int64
	err = global.GVA_DB.Unscoped().Model(&product.Product{}).Where("code = ?", prod.Code).Count(&count).Error
	if err != nil {
		return err
	}
	if count > 0 {
		return errors.New("商品编码 " + prod.Code + " 已存在，不能重复添加")
	}

	if prod.CreatedAt.Before(time.Date(1900, 1, 1, 0, 0, 0, 0, time.UTC)) {
		prod.CreatedAt = time.Now()
	}
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

func (prodService *ProductService) RestoreProduct(prod product.Product) (err error) {
	err = global.GVA_DB.Unscoped().Model(&product.Product{}).Where("id = ?", prod.ID).First(&prod).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Model(&product.Product{}).Where("id = ?", prod.ID).Update("deleted_by", 0).Error
	if err != nil {
		return err
	}
	err = global.GVA_DB.Unscoped().Model(&product.Product{}).Where("id = ?", prod.ID).Update("deleted_at", nil).Error
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

// GetDeletedProductInfoList 分页获取Product记录
// Author [piexlmax](https://github.com/piexlmax)
func (prodService *ProductService) GetDeletedProductInfoList(info productReq.ProductSearch) (list []product.Product, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Unscoped().Model(&product.Product{})
	// db := global.GVA_DB.Model(&product.Product{})
	var prods []product.Product

	db = db.Where("deleted_at IS NOT NULL")

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
	// 查询数据（逻辑不变）
	db := global.GVA_DB.Model(&product.Product{})
	var prods []product.Product
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
	orderMap := map[string]bool{"code": true}
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.Order == "descending" {
			OrderStr += " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("id desc")
	}

	if err = db.Find(&prods).Error; err != nil {
		return err
	}

	// 创建Excel（核心调整部分）
	var xlsx *excelize.File
	i, j := 0, 0

	if info.RequireOrd != nil && *info.RequireOrd == 1 {
		// 分支1：使用采购商品模板
		if xlsx, err = excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "template_product_require.xlsx"); err != nil {
			return err
		}
		sheet := xlsx.GetSheetName(0)
		const ColSize int = 11              // 原有列数，右移后变为12列（+1列图片）
		var style1, style2 [ColSize + 1]int // 样式数组+1（适配新增的第一列）
		startingASCIINumber := 65           // 'A'的ASCII码

		// // 读取样式（原有列样式右移）
		// for k := 0; k < ColSize; k++ {
		// 	// 原有第k列（A开始）样式移到第k+1列（B开始）
		// 	style1[k+1], _ = xlsx.GetCellStyle(sheet, string(rune(k+startingASCIINumber))+"4")
		// 	style2[k+1], _ = xlsx.GetCellStyle(sheet, string(rune(k+startingASCIINumber))+"5")
		// }
		// 第一列（图片列）使用A4/A5的样式
		// style1[0], _ = xlsx.GetCellStyle(sheet, "A4")
		// style2[0], _ = xlsx.GetCellStyle(sheet, "A5")

		// 调整第一列宽度以容纳82px图片（1px≈0.14列宽单位）
		xlsx.SetColWidth(sheet, "A", "A", 82*0.14)

		for i, j = 0, len(prods); i < j; i++ {
			iStr := strconv.Itoa(4 + i)
			index := i % 2

			// 设置行样式（包含新增的第一列）
			for k := 0; k <= ColSize; k++ {
				col := string(rune(k + startingASCIINumber))
				if index == 0 {
					xlsx.SetCellStyle(sheet, col+iStr, col+iStr, style1[k])
				} else {
					xlsx.SetCellStyle(sheet, col+iStr, col+iStr, style2[k])
				}
			}

			// 1. 第一列（A列）添加82x82图片
			if prods[i].Image != "" {
				imgPath := prods[i].Image
				config, err := getImageScaleConfig(imgPath, 82, 82)
				if err != nil {
					fmt.Printf("图片处理警告: %v (路径: %s)\n", err, imgPath)
				} else {
					if err := xlsx.AddPicture(sheet, "A"+iStr, imgPath, config); err != nil {
						fmt.Printf("添加图片失败: %v (路径: %s)\n", err, imgPath)
					}
				}
			}

			// 2. 原有列右移一列（A→B, B→C, ..., K→L）
			xlsx.SetCellStr(sheet, "B"+iStr, prods[i].Code)
			xlsx.SetCellStr(sheet, "C"+iStr, prods[i].Product_name_cn+"\n"+prods[i].Product_name_en)
			xlsx.SetCellStr(sheet, "D"+iStr, prods[i].Package)
			xlsx.SetCellInt(sheet, "E"+iStr, *prods[i].Store)
			xlsx.SetCellStr(sheet, "F"+iStr, prods[i].SelfLife)
			xlsx.SetCellStr(sheet, "G"+iStr, prods[i].Barcode)
			xlsx.SetCellStr(sheet, "H"+iStr, prods[i].BarcodeCase)
			xlsx.SetCellStr(sheet, "I"+iStr, prods[i].CartonSize)
			if prods[i].Cbm != nil {
				xlsx.SetCellFloat(sheet, "J"+iStr, *prods[i].Cbm, 4, 64)
			}
			if prods[i].Weight != nil {
				xlsx.SetCellFloat(sheet, "K"+iStr, *prods[i].Weight, 2, 64)
			}
			if prods[i].InPrice != nil {
				xlsx.SetCellFloat(sheet, "L"+iStr, *prods[i].InPrice, 2, 64)
			}
		}
	} else {
		// 分支2：使用普通商品模板
		if xlsx, err = excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "template_product.xlsx"); err != nil {
			return err
		}
		sheet := xlsx.GetSheetName(0)

		// 样式数组调整（原有列右移）
		var strStyle, dateStyle, numStyle, moneyStyle [2]int
		strStyle[0], _ = xlsx.GetCellStyle(sheet, "C4")   // 原有A列样式→新B列
		dateStyle[0], _ = xlsx.GetCellStyle(sheet, "E4")  // 原有D列→新E列
		numStyle[0], _ = xlsx.GetCellStyle(sheet, "F4")   // 原有E列→新F列
		moneyStyle[0], _ = xlsx.GetCellStyle(sheet, "G4") // 原有F列→新G列
		strStyle[1], _ = xlsx.GetCellStyle(sheet, "C5")
		dateStyle[1], _ = xlsx.GetCellStyle(sheet, "E5")
		numStyle[1], _ = xlsx.GetCellStyle(sheet, "F5")
		moneyStyle[1], _ = xlsx.GetCellStyle(sheet, "G5")

		// 调整第一列宽度以容纳82px图片
		xlsx.SetColWidth(sheet, "A", "A", 82*0.14)

		for i, j = 0, len(prods); i < j; i++ {
			iStr := strconv.Itoa(4 + i)
			index := i % 2

			// 设置行样式（原有列右移）
			xlsx.SetCellStyle(sheet, "B"+iStr, "D"+iStr, strStyle[index])   // 原A-C→新B-D
			xlsx.SetCellStyle(sheet, "E"+iStr, "E"+iStr, dateStyle[index])  // 原D→新E
			xlsx.SetCellStyle(sheet, "F"+iStr, "F"+iStr, numStyle[index])   // 原E→新F
			xlsx.SetCellStyle(sheet, "G"+iStr, "G"+iStr, moneyStyle[index]) // 原F→新G
			xlsx.SetCellStyle(sheet, "H"+iStr, "H"+iStr, strStyle[index])   // 原G→新H
			xlsx.SetCellStyle(sheet, "I"+iStr, "I"+iStr, moneyStyle[index]) // 原H→新I
			// 第一列（图片列）样式
			xlsx.SetCellStyle(sheet, "A"+iStr, "A"+iStr, strStyle[index])

			// 1. 第一列（A列）添加82x82图片
			if prods[i].Image != "" {
				imgPath := prods[i].Image
				config, err := getImageScaleConfig(imgPath, 82, 82)
				if err != nil {
					fmt.Printf("图片处理警告: %v (路径: %s)\n", err, imgPath)
				} else {
					if err := xlsx.AddPicture(sheet, "A"+iStr, imgPath, config); err != nil {
						fmt.Printf("添加图片失败: %v (路径: %s)\n", err, imgPath)
					}
				}
			}

			// 2. 原有列右移一列（A→B, B→C, ..., H→I）
			xlsx.SetCellStr(sheet, "B"+iStr, prods[i].Code)
			xlsx.SetCellStr(sheet, "C"+iStr, prods[i].Product_name_cn+" "+prods[i].Product_name_en)
			xlsx.SetCellStr(sheet, "D"+iStr, prods[i].Package)
			if prods[i].Exp_date != nil {
				xlsx.SetCellValue(sheet, "E"+iStr, prods[i].Exp_date.Format("02/01/2006"))
			}
			xlsx.SetCellInt(sheet, "F"+iStr, *prods[i].Store)
			xlsx.SetCellValue(sheet, "G"+iStr, *prods[i].Price)
			xlsx.SetCellValue(sheet, "H"+iStr, prods[i].Barcode)
			xlsx.SetCellValue(sheet, "I"+iStr, *prods[i].Vat)
		}

		// 处理无价格情况（调整后列索引）
		if *info.WithPrice == 0 {
			xlsx.RemoveCol(sheet, "H") // 原G列→新H列
			xlsx.RemoveCol(sheet, "G") // 原F列→新G列
		}
	}

	return xlsx.SaveAs(fileName)
}

// 辅助函数：计算图片缩放比例，生成82x82像素的配置
func getImageScaleConfig(imgPath string, targetWidth, targetHeight int) (string, error) {
	// 打开图片文件
	file, err := os.Open(imgPath)
	if err != nil {
		return "", fmt.Errorf("打开图片失败: %w", err)
	}
	defer file.Close()

	// 获取原图尺寸
	imgConfig, _, err := image.DecodeConfig(file)
	if err != nil {
		return "", fmt.Errorf("解析图片尺寸失败: %w", err)
	}

	// 计算缩放比例（目标尺寸/原图尺寸）
	xScale := float64(targetWidth) / float64(imgConfig.Width)
	yScale := float64(targetHeight) / float64(imgConfig.Height)

	// 生成配置字符串（关闭宽高比锁定，强制缩放）
	return fmt.Sprintf(`{
		"x_offset": 3,          
		"y_offset": 8,          
		"x_scale": %f,          
		"y_scale": %f,          
		"print_obj": true,      
		"lock_aspect_ratio": false  
	}`, xScale, yScale), nil
}
