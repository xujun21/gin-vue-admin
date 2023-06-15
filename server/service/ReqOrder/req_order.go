package ReqOrder

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	reqOrder "github.com/flipped-aurora/gin-vue-admin/server/model/ReqOrder"
	ReqOrderReq "github.com/flipped-aurora/gin-vue-admin/server/model/ReqOrder/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
)

type RequireOrderService struct {
}

// CreateRequireOrder 创建RequireOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (ReqOrdService *RequireOrderService) CreateRequireOrder(ReqOrd reqOrder.RequireOrder) (err error) {
	err = global.GVA_DB.Create(&ReqOrd).Error
	return err
}

// DeleteRequireOrder 删除RequireOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (ReqOrdService *RequireOrderService) DeleteRequireOrder(ReqOrd reqOrder.RequireOrder) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&reqOrder.RequireOrder{}).Where("id = ?", ReqOrd.ID).Update("deleted_by", ReqOrd.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&ReqOrd).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// DeleteRequireOrderByIds 批量删除RequireOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (ReqOrdService *RequireOrderService) DeleteRequireOrderByIds(ids request.IdsReq, deleted_by uint) (err error) {
	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&reqOrder.RequireOrder{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&reqOrder.RequireOrder{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

// UpdateRequireOrder 更新RequireOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (ReqOrdService *RequireOrderService) UpdateRequireOrder(ReqOrd reqOrder.RequireOrder) (err error) {
	err = global.GVA_DB.Save(&ReqOrd).Error
	return err
}

// GetRequireOrder 根据id获取RequireOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (ReqOrdService *RequireOrderService) GetRequireOrder(id uint) (ReqOrd reqOrder.RequireOrder, err error) {
	err = global.GVA_DB.Where("id = ?", id).First(&ReqOrd).Error
	return
}

// GetRequireOrderInfoList 分页获取RequireOrder记录
// Author [piexlmax](https://github.com/piexlmax)
func (ReqOrdService *RequireOrderService) GetRequireOrderInfoList(info ReqOrderReq.RequireOrderSearch) (list []reqOrder.RequireOrder, total int64, err error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GVA_DB.Model(&reqOrder.RequireOrder{})
	var ReqOrds []reqOrder.RequireOrder
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
		db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	}
	// if info.StartOrder_date != nil && info.EndOrder_date != nil {
	//     db = db.Where("order_date BETWEEN ? AND ? ",info.StartOrder_date,info.EndOrder_date)
	// }

	if info.PONumber != "" {
		db = db.Where("po_number like ?", "%"+info.PONumber+"%")
	}

	if info.StartQuantity != nil && info.EndQuantity != nil {
		db = db.Where("quantity BETWEEN ? AND ? ", info.StartQuantity, info.EndQuantity)
	}
	if info.StartTotalCbm != nil && info.EndTotalCbm != nil {
		db = db.Where("total_cbm BETWEEN ? AND ? ", info.StartTotalCbm, info.EndTotalCbm)
	}
	if info.StartTotalPrice != nil && info.EndTotalPrice != nil {
		db = db.Where("total_price BETWEEN ? AND ? ", info.StartTotalPrice, info.EndTotalPrice)
	}
	err = db.Count(&total).Error
	if err != nil {
		return
	}

	err = db.Limit(limit).Offset(offset).Find(&ReqOrds).Error
	return ReqOrds, total, err
}

func (ReqOrdService *RequireOrderService) ExportRequireExcel(reqOrderId *int, fileName string) (err error) {
	var ord reqOrder.RequireOrder

	if ord, err = ReqOrdService.GetRequireOrder(uint(*reqOrderId)); err != nil {
		return err
	}

	var subOrderList []reqOrder.ReqSubOrder
	if err = global.GVA_DB.Where("req_order_id = ?", reqOrderId).Order("id asc").Find(&subOrderList).Error; err != nil {
		return err
	}

	var xlsx *excelize.File
	if xlsx, err = excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "template_require.xlsx"); err != nil {
		return err
	}
	var sheet = xlsx.GetSheetName(0)

	xlsx.SetCellStr(sheet, "A2", ord.PONumber)

	i := 0
	j := 0

	var style [16]int
	var startingASCIINumber int = 65
	for k := 0; k < 16; k++ {
		style[k], _ = xlsx.GetCellStyle(sheet, string(rune(k+startingASCIINumber))+"4")
	}

	for i, j = 0, len(subOrderList); i < j; i++ {
		iStr := strconv.Itoa(i + 4)
		// if i%2 == 1 {
		// xlsx.SetCellStyle(sheet, "A"+iStr, "P"+iStr, strStyle)
		// xlsx.SetCellStyle(sheet, "S"+iStr, "U"+iStr, dateStyle)
		// xlsx.SetCellStyle(sheet, "V"+iStr, "X"+iStr, numStyle)
		// }
		for k := 0; k < 16; k++ {
			xlsx.SetCellStyle(sheet, string(rune(k+startingASCIINumber))+iStr, string(rune(k+startingASCIINumber))+iStr, style[k])
		}
		xlsx.SetCellStr(sheet, "A"+iStr, subOrderList[i].ProductCode)
		xlsx.SetCellStr(sheet, "B"+iStr, subOrderList[i].ProductNameCn+"\n"+subOrderList[i].ProductNameEn)
		xlsx.SetCellStr(sheet, "D"+iStr, subOrderList[i].SelfLife)
		xlsx.SetCellInt(sheet, "G"+iStr, *subOrderList[i].InQty)
		xlsx.SetCellStr(sheet, "H"+iStr, subOrderList[i].Barcode)
		xlsx.SetCellStr(sheet, "I"+iStr, subOrderList[i].BarcodeCase)
		xlsx.SetCellStr(sheet, "J"+iStr, subOrderList[i].CartonSize)
		if subOrderList[i].Cbm != nil {
			xlsx.SetCellFloat(sheet, "K"+iStr, *subOrderList[i].Cbm, 4, 64)
		}
		if subOrderList[i].Weight != nil {
			xlsx.SetCellFloat(sheet, "M"+iStr, *subOrderList[i].Weight, 2, 64)
		}
		if subOrderList[i].InPrice != nil {
			xlsx.SetCellFloat(sheet, "O"+iStr, *subOrderList[i].InPrice, 2, 64)
		}
	}
	iStr := strconv.Itoa(i + 4)
	xlsx.SetCellFormula(sheet, "G"+iStr, "sum(G4:G"+strconv.Itoa(i+3)+")")
	xlsx.SetCellFormula(sheet, "L"+iStr, "sum(L4:L"+strconv.Itoa(i+3)+")")
	xlsx.SetCellFormula(sheet, "N"+iStr, "sum(N4:N"+strconv.Itoa(i+3)+")")
	xlsx.SetCellFormula(sheet, "P"+iStr, "sum(P4:P"+strconv.Itoa(i+3)+")")

	styleL, _ := xlsx.GetCellStyle(sheet, "L2")
	xlsx.SetCellStyle(sheet, "L"+iStr, "L"+iStr, styleL)
	styleN, _ := xlsx.GetCellStyle(sheet, "N2")
	xlsx.SetCellStyle(sheet, "N"+iStr, "N"+iStr, styleN)
	styleP, _ := xlsx.GetCellStyle(sheet, "P2")
	xlsx.SetCellStyle(sheet, "P"+iStr, "P"+iStr, styleP)
	styleG, _ := xlsx.GetCellStyle(sheet, "G2")
	xlsx.SetCellStyle(sheet, "G"+iStr, "G"+iStr, styleG)

	err = xlsx.SaveAs(fileName)
	return err
}
