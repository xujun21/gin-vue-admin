package order

import (
	"fmt"
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	company2 "github.com/flipped-aurora/gin-vue-admin/server/model/company"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/shopspring/decimal"
	"github.com/xuri/excelize/v2"
	"gorm.io/gorm"
	"strconv"
	"time"
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
	// 查找suborder
	var subOrderListToDel []order.SubOrder
	if err = global.GVA_DB.Where("order_id = ?", ord.ID).Find(&subOrderListToDel).Error; err != nil {
		return err
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&order.Order{}).Where("id = ?", ord.ID).Update("deleted_by", ord.DeletedBy).Error; err != nil {
			return err
		}
		if err = tx.Delete(&ord).Error; err != nil {
			return err
		}

		for i, j := 0, len(subOrderListToDel); i < j; i++ {
			// 删除suborder
			if err = tx.Model(&order.SubOrder{}).Where("id = ?", subOrderListToDel[i].ID).Update("deleted_by", ord.DeletedBy).Error; err != nil {
				return err
			}
			if err = tx.Where("id = ?", subOrderListToDel[i].ID).Delete(&order.SubOrder{}).Error; err != nil {
				return err
			}
			// 返回库存
			if err = tx.Exec("update product set store = store + ? where id = ?", subOrderListToDel[i].Quantity, subOrderListToDel[i].Product_id).Error; err != nil {
				return err
			}
		}
		return nil
	})
	return nil
}

// DeleteOrderByIds 批量删除Order记录
// Author [piexlmax](https://github.com/piexlmax)
func (ordService *OrderService) DeleteOrderByIds(ids request.IdsReq, deleted_by uint) (err error) {
	// 查找suborder
	var subOrderListToDel []order.SubOrder
	if err = global.GVA_DB.Where("order_id in ?", ids.Ids).Find(&subOrderListToDel).Error; err != nil {
		return err
	}

	err = global.GVA_DB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&order.Order{}).Where("id in ?", ids.Ids).Update("deleted_by", deleted_by).Error; err != nil {
			return err
		}
		if err := tx.Where("id in ?", ids.Ids).Delete(&order.Order{}).Error; err != nil {
			return err
		}

		for i, j := 0, len(subOrderListToDel); i < j; i++ {
			// 删除suborder
			if err = tx.Model(&order.SubOrder{}).Where("id = ?", subOrderListToDel[i].ID).Update("deleted_by", deleted_by).Error; err != nil {
				return err
			}
			if err = tx.Where("id = ?", subOrderListToDel[i].ID).Delete(&order.SubOrder{}).Error; err != nil {
				return err
			}
			// 返回库存
			if err = tx.Exec("update product set store = store + ? where id = ?", subOrderListToDel[i].Quantity, subOrderListToDel[i].Product_id).Error; err != nil {
				return err
			}
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
	//if info.StartCreatedAt != nil && info.EndCreatedAt != nil {
	//	db = db.Where("created_at BETWEEN ? AND ?", info.StartCreatedAt, info.EndCreatedAt)
	//}
	//if info.StartOrder_date != nil && info.EndOrder_date != nil {
	//	db = db.Where("order_date BETWEEN ? AND ? ", info.StartOrder_date, info.EndOrder_date)
	//}
	if info.StartCreatedAt != nil {
		db = db.Where("created_at >= ?", info.StartCreatedAt)
	}
	if info.EndCreatedAt != nil {
		db = db.Where("create_at <= ?", info.EndCreatedAt)
	}
	if info.StartOrder_date != nil {
		db = db.Where("order_date >= ? ", info.StartOrder_date)
	}
	if info.EndOrder_date != nil {
		db = db.Where("order_date <= ? ", info.EndOrder_date)
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

	var OrderStr string // 排序
	orderMap := make(map[string]bool)
	orderMap["order_date"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.OrderStr == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("id desc")
	}
	err = db.Limit(limit).Offset(offset).Find(&ords).Error
	return ords, total, err
}

// 计算订单各类金额
func (ordService *OrderService) CalcOrder(orderId *int) (err error) {
	var ord order.Order
	if err := global.GVA_DB.Where("id = ?", orderId).First(&ord).Error; err != nil {
		return err
	}

	var subOrderList []order.SubOrder
	if err = global.GVA_DB.Where("order_id = ?", orderId).Find(&subOrderList).Error; err != nil {
		return err
	}

	// 开始计算
	// order.sub_total = sum(sub_order.sub_total)
	var ordSubTotal, ordVatTotal, ordDiscountTotal float64
	ordSubTotal = 0
	ordVatTotal = 0
	ordDiscountTotal = 0

	var ordQtyTotal int
	ordQtyTotal = 0

	for i, j := 0, len(subOrderList); i < j; i++ {
		ordSubTotal += *subOrderList[i].Sub_total
		ordQtyTotal += *subOrderList[i].Quantity
		ordVatTotal += *subOrderList[i].Sub_Vat
		if subOrderList[i].Discount_total != nil {
			ordDiscountTotal += *subOrderList[i].Discount_total
		}
	}
	ord.Subtotal = &ordSubTotal
	ord.Quantity = &ordQtyTotal
	ord.Vat = &ordVatTotal
	ord.Discount = &ordDiscountTotal

	orderTotal, _ := decimal.NewFromFloat(*ord.Subtotal).
		Add(decimal.NewFromFloat(*ord.Vat)).
		Add(decimal.NewFromFloat(*ord.DeliveryFee)).
		Sub(decimal.NewFromFloat(*ord.Discount)).
		Sub(decimal.NewFromFloat(*ord.Hand_discount)).
		Sub(decimal.NewFromFloat(*ord.HandDiscountVat)).
		Float64()
	ord.Order_total = &orderTotal

	err = global.GVA_DB.Save(&ord).Error
	if err != nil {
		return err
	}
	return nil
}

func (ordService *OrderService) ExportInvoiceExcel(orderId *int, fileName string) (err error) {
	var ord order.Order
	if ord, err = ordService.GetOrder(uint(*orderId)); err != nil {
		return err
	}

	var subOrderList []order.SubOrder
	if err = global.GVA_DB.Where("order_id = ?", orderId).Order("product_code asc").Find(&subOrderList).Error; err != nil {
		return err
	}

	var xlsx *excelize.File
	if xlsx, err = excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "template.xlsx"); err != nil {
		return err
	}
	var sheet = xlsx.GetSheetName(0)

	xlsx.SetCellStr(sheet, "A6", ord.Bill_to)
	xlsx.SetCellStr(sheet, "A8", ord.Bill_to_address)
	xlsx.SetCellStr(sheet, "A10", ord.Bill_to_city)
	xlsx.SetCellStr(sheet, "A11", ord.Bill_to_postcode)
	xlsx.SetCellStr(sheet, "A12", ord.Bill_to_phone)
	xlsx.SetCellStr(sheet, "I6", ord.Ship_to)
	xlsx.SetCellStr(sheet, "I8", ord.Ship_to_address)
	xlsx.SetCellStr(sheet, "I10", ord.Ship_to_city)
	xlsx.SetCellStr(sheet, "I11", ord.Ship_to_postcode)
	xlsx.SetCellStr(sheet, "I12", ord.Ship_to_phone)
	datStyle, _ := xlsx.NewStyle(`{"custom_number_format":"dd/mm/yyyy;@"}`)
	xlsx.SetCellStyle(sheet, "U6", "U6", datStyle)
	xlsx.SetCellValue(sheet, "U6", returnDateStringOrNull(ord.Order_date))

	xlsx.SetCellStr(sheet, "U7", ord.Po_number)
	xlsx.SetCellStyle(sheet, "U8", "U8", datStyle)
	xlsx.SetCellValue(sheet, "U8", returnDateStringOrNull(ord.Invoice_date))
	xlsx.SetCellStr(sheet, "U9", ord.Invoice_no)
	xlsx.SetCellStr(sheet, "U10", ord.Payment_method)

	i := 0
	j := 0

	var strStyle, dateStyle, numStyle, moneyStyle int
	strStyle, _ = xlsx.GetCellStyle(sheet, "A15")
	dateStyle, _ = xlsx.GetCellStyle(sheet, "P15")
	numStyle, _ = xlsx.GetCellStyle(sheet, "R15")
	moneyStyle, _ = xlsx.GetCellStyle(sheet, "S15")

	for i, j = 0, len(subOrderList); i < j; i++ {
		iStr := strconv.Itoa(14 + i)

		if err != nil {
			fmt.Println(err)
		}
		if i%2 == 1 {
			xlsx.SetCellStyle(sheet, "A"+iStr, "N"+iStr, strStyle)
			xlsx.SetCellStyle(sheet, "O"+iStr, "Q"+iStr, dateStyle)
			xlsx.SetCellStyle(sheet, "R"+iStr, "R"+iStr, numStyle)
			xlsx.SetCellStyle(sheet, "S"+iStr, "X"+iStr, moneyStyle)
		}
		xlsx.MergeCell(sheet, "A"+iStr, "B"+iStr)
		xlsx.SetCellStr(sheet, "A"+iStr, subOrderList[i].Product_code)
		xlsx.MergeCell(sheet, "C"+iStr, "K"+iStr)
		xlsx.SetCellStr(sheet, "C"+iStr, subOrderList[i].Product_name_cn+" "+subOrderList[i].Product_name_en)
		xlsx.MergeCell(sheet, "L"+iStr, "N"+iStr)
		xlsx.SetCellStr(sheet, "L"+iStr, subOrderList[i].Package)
		xlsx.MergeCell(sheet, "O"+iStr, "Q"+iStr)
		xlsx.SetCellValue(sheet, "O"+iStr, subOrderList[i].Exp_date.Format("02/01/2006"))
		xlsx.SetCellInt(sheet, "R"+iStr, *subOrderList[i].Quantity)
		xlsx.MergeCell(sheet, "S"+iStr, "T"+iStr)
		xlsx.SetCellFloat(sheet, "S"+iStr, *subOrderList[i].Price, 2, 64)
		xlsx.MergeCell(sheet, "U"+iStr, "V"+iStr)
		//xlsx.SetCellFloat(sheet, "U"+iStr, *subOrderList[i].Sub_Vat, 2, 64)
		xlsx.SetCellFloat(sheet, "U"+iStr, *subOrderList[i].Vat, 2, 64) // 显示单件VAT

		xlsx.MergeCell(sheet, "W"+iStr, "X"+iStr)
		//xlsx.SetCellFloat(sheet, "W"+iStr, *subOrderList[i].Sub_total, 2, 64)
		tmpTotal, _ := decimal.NewFromFloat(*subOrderList[i].Sub_total).Add(decimal.NewFromFloat(*subOrderList[i].Sub_Vat)).Float64()
		xlsx.SetCellFloat(sheet, "W"+iStr, tmpTotal, 2, 64) // TOTAL=（单价+单件VAT）*数量
	}

	bigStrStyle, err := xlsx.NewStyle(`{"font":{"bold":false,"italic":false,"size":11}}`)
	bigNumStyle, err := xlsx.NewStyle(`{"font":{"bold":false,"italic":false,"size":11}}`)
	bigMoneyStyle, err := xlsx.NewStyle(`{"custom_number_format": "[$£-en-GB]#,##0.00"}`)
	orderTotalLStyle, err := xlsx.NewStyle(`{"font":{"bold":true,"italic":false,"size":12}}`)
	orderTotalVStyle, err := xlsx.NewStyle(`{"font":{"bold":true,"italic":false,"size":12},"custom_number_format": "[$£-en-GB]#,##0.00"}`)

	if err != nil {
		fmt.Println(err)
	}
	i++
	iStr := strconv.Itoa(14 + i)
	iStr3 := strconv.Itoa(14 + i + 3)
	xlsx.SetCellStyle(sheet, "R"+iStr, "R"+iStr, bigNumStyle)    // QTY
	xlsx.SetCellStyle(sheet, "S"+iStr, "S"+iStr3, bigStrStyle)   // SUB_XXX
	xlsx.SetCellStyle(sheet, "V"+iStr, "V"+iStr3, bigMoneyStyle) // SUB_XXX.Value

	xlsx.SetCellValue(sheet, "R"+iStr, *ord.Quantity)
	xlsx.MergeCell(sheet, "S"+iStr, "U"+iStr)
	xlsx.SetCellStr(sheet, "S"+iStr, "Subtotal:")
	xlsx.MergeCell(sheet, "V"+iStr, "X"+iStr)
	xlsx.SetCellValue(sheet, "V"+iStr, *ord.Subtotal)

	i++
	iStr = strconv.Itoa(14 + i)
	xlsx.MergeCell(sheet, "S"+iStr, "U"+iStr)
	xlsx.SetCellStr(sheet, "S"+iStr, "Vat:")
	xlsx.MergeCell(sheet, "V"+iStr, "X"+iStr)
	xlsx.SetCellValue(sheet, "V"+iStr, *ord.Vat)

	i++
	iStr = strconv.Itoa(14 + i)
	xlsx.MergeCell(sheet, "S"+iStr, "U"+iStr)
	xlsx.SetCellStr(sheet, "S"+iStr, "Discount:")
	xlsx.MergeCell(sheet, "V"+iStr, "X"+iStr)
	xlsx.SetCellValue(sheet, "V"+iStr, *ord.Discount+*ord.Hand_discount+*ord.HandDiscountVat)

	i++
	iStr = strconv.Itoa(14 + i)
	xlsx.MergeCell(sheet, "S"+iStr, "U"+iStr)
	xlsx.SetCellStr(sheet, "S"+iStr, "Delivery:")
	xlsx.MergeCell(sheet, "V"+iStr, "X"+iStr)
	xlsx.SetCellValue(sheet, "V"+iStr, *ord.DeliveryFee)

	i++
	iStr = strconv.Itoa(14 + i)
	xlsx.MergeCell(sheet, "R"+iStr, "U"+iStr)
	xlsx.SetCellStyle(sheet, "R"+iStr, "R"+iStr, orderTotalLStyle)
	xlsx.SetCellStr(sheet, "R"+iStr, "Order Total:")
	xlsx.MergeCell(sheet, "V"+iStr, "X"+iStr)
	xlsx.SetCellStyle(sheet, "V"+iStr, "V"+iStr, orderTotalVStyle)
	xlsx.SetCellValue(sheet, "V"+iStr, *ord.Order_total)

	err = xlsx.SaveAs(fileName)
	return err
}

func (ordService *OrderService) ExportConfirmExcel(orderId *int, fileName string) (err error) {
	var ord order.Order
	if ord, err = ordService.GetOrder(uint(*orderId)); err != nil {
		return err
	}

	var subOrderList []order.SubOrder
	if err = global.GVA_DB.Where("order_id = ?", orderId).Order("product_code asc").Find(&subOrderList).Error; err != nil {
		return err
	}

	var xlsx *excelize.File
	if xlsx, err = excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "template.xlsx"); err != nil {
		return err
	}
	var sheet = xlsx.GetSheetName(0)

	xlsx.SetCellStr(sheet, "A6", ord.Bill_to)
	xlsx.SetCellStr(sheet, "A8", ord.Bill_to_address)
	xlsx.SetCellStr(sheet, "A10", ord.Bill_to_city)
	xlsx.SetCellStr(sheet, "A11", ord.Bill_to_postcode)
	xlsx.SetCellStr(sheet, "A12", ord.Bill_to_phone)
	xlsx.SetCellStr(sheet, "I6", ord.Ship_to)
	xlsx.SetCellStr(sheet, "I8", ord.Ship_to_address)
	xlsx.SetCellStr(sheet, "I10", ord.Ship_to_city)
	xlsx.SetCellStr(sheet, "I11", ord.Ship_to_postcode)
	xlsx.SetCellStr(sheet, "I12", ord.Ship_to_phone)
	datStyle, _ := xlsx.NewStyle(`{"custom_number_format":"dd/mm/yyyy;@"}`)
	xlsx.SetCellStyle(sheet, "U6", "U6", datStyle)
	xlsx.SetCellValue(sheet, "U6", returnDateStringOrNull(ord.Order_date))

	xlsx.SetCellStr(sheet, "U7", ord.Po_number)
	xlsx.SetCellStyle(sheet, "U8", "U8", datStyle)
	xlsx.SetCellValue(sheet, "U8", returnDateStringOrNull(ord.Invoice_date))
	xlsx.SetCellStr(sheet, "U9", ord.Invoice_no)
	xlsx.SetCellStr(sheet, "U10", ord.Payment_method)

	i := 0
	j := 0

	var strStyle, dateStyle, numStyle, moneyStyle int
	strStyle, _ = xlsx.GetCellStyle(sheet, "A15")
	dateStyle, _ = xlsx.GetCellStyle(sheet, "P15")
	numStyle, _ = xlsx.GetCellStyle(sheet, "R15")
	moneyStyle, _ = xlsx.GetCellStyle(sheet, "S15")

	for i, j = 0, len(subOrderList); i < j; i++ {
		iStr := strconv.Itoa(14 + i)

		if err != nil {
			fmt.Println(err)
		}
		if i%2 == 1 {
			xlsx.SetCellStyle(sheet, "A"+iStr, "N"+iStr, strStyle)
			xlsx.SetCellStyle(sheet, "O"+iStr, "Q"+iStr, dateStyle)
			xlsx.SetCellStyle(sheet, "R"+iStr, "R"+iStr, numStyle)
			xlsx.SetCellStyle(sheet, "S"+iStr, "X"+iStr, moneyStyle)
		}
		xlsx.MergeCell(sheet, "A"+iStr, "B"+iStr)
		xlsx.SetCellStr(sheet, "A"+iStr, subOrderList[i].Product_code)
		xlsx.MergeCell(sheet, "C"+iStr, "K"+iStr)
		xlsx.SetCellStr(sheet, "C"+iStr, subOrderList[i].Product_name_cn+" "+subOrderList[i].Product_name_en)
		xlsx.MergeCell(sheet, "L"+iStr, "N"+iStr)
		xlsx.SetCellStr(sheet, "L"+iStr, subOrderList[i].Package)
		xlsx.MergeCell(sheet, "O"+iStr, "Q"+iStr)
		xlsx.SetCellValue(sheet, "O"+iStr, subOrderList[i].Exp_date.Format("02/01/2006"))
		xlsx.SetCellInt(sheet, "R"+iStr, *subOrderList[i].Quantity)
		xlsx.MergeCell(sheet, "S"+iStr, "T"+iStr)
		xlsx.SetCellFloat(sheet, "S"+iStr, *subOrderList[i].Price, 2, 64)
		xlsx.MergeCell(sheet, "U"+iStr, "V"+iStr)
		//xlsx.SetCellFloat(sheet, "U"+iStr, *subOrderList[i].Sub_Vat, 2, 64)
		xlsx.SetCellFloat(sheet, "U"+iStr, *subOrderList[i].Vat, 2, 64) // 显示单件VAT

		xlsx.MergeCell(sheet, "W"+iStr, "X"+iStr)
		//xlsx.SetCellFloat(sheet, "W"+iStr, *subOrderList[i].Sub_total, 2, 64)
		tmpTotal, _ := decimal.NewFromFloat(*subOrderList[i].Sub_total).Add(decimal.NewFromFloat(*subOrderList[i].Sub_Vat)).Float64()
		xlsx.SetCellFloat(sheet, "W"+iStr, tmpTotal, 2, 64) // TOTAL=（单价+单件VAT）*数量
	}

	bigStrStyle, err := xlsx.NewStyle(`{"font":{"bold":false,"italic":false,"size":11}}`)
	bigNumStyle, err := xlsx.NewStyle(`{"font":{"bold":false,"italic":false,"size":11}}`)
	bigMoneyStyle, err := xlsx.NewStyle(`{"custom_number_format": "[$£-en-GB]#,##0.00"}`)
	orderTotalLStyle, err := xlsx.NewStyle(`{"font":{"bold":true,"italic":false,"size":12}}`)
	orderTotalVStyle, err := xlsx.NewStyle(`{"font":{"bold":true,"italic":false,"size":12},"custom_number_format": "[$£-en-GB]#,##0.00"}`)

	if err != nil {
		fmt.Println(err)
	}
	i++
	iStr := strconv.Itoa(14 + i)
	iStr3 := strconv.Itoa(14 + i + 3)
	xlsx.SetCellStyle(sheet, "R"+iStr, "R"+iStr, bigNumStyle)    // QTY
	xlsx.SetCellStyle(sheet, "S"+iStr, "S"+iStr3, bigStrStyle)   // SUB_XXX
	xlsx.SetCellStyle(sheet, "V"+iStr, "V"+iStr3, bigMoneyStyle) // SUB_XXX.Value

	xlsx.SetCellValue(sheet, "R"+iStr, *ord.Quantity)
	xlsx.MergeCell(sheet, "S"+iStr, "U"+iStr)
	xlsx.SetCellStr(sheet, "S"+iStr, "Subtotal:")
	xlsx.MergeCell(sheet, "V"+iStr, "X"+iStr)
	xlsx.SetCellValue(sheet, "V"+iStr, *ord.Subtotal)

	i++
	iStr = strconv.Itoa(14 + i)
	xlsx.MergeCell(sheet, "S"+iStr, "U"+iStr)
	xlsx.SetCellStr(sheet, "S"+iStr, "Vat:")
	xlsx.MergeCell(sheet, "V"+iStr, "X"+iStr)
	xlsx.SetCellValue(sheet, "V"+iStr, *ord.Vat)

	i++
	iStr = strconv.Itoa(14 + i)
	xlsx.MergeCell(sheet, "S"+iStr, "U"+iStr)
	xlsx.SetCellStr(sheet, "S"+iStr, "Discount:")
	xlsx.MergeCell(sheet, "V"+iStr, "X"+iStr)
	xlsx.SetCellValue(sheet, "V"+iStr, *ord.Discount+*ord.Hand_discount+*ord.HandDiscountVat)

	i++
	iStr = strconv.Itoa(14 + i)
	xlsx.MergeCell(sheet, "S"+iStr, "U"+iStr)
	xlsx.SetCellStr(sheet, "S"+iStr, "Delivery:")
	xlsx.MergeCell(sheet, "V"+iStr, "X"+iStr)
	xlsx.SetCellValue(sheet, "V"+iStr, *ord.DeliveryFee)

	i++
	iStr = strconv.Itoa(14 + i)
	xlsx.MergeCell(sheet, "R"+iStr, "U"+iStr)
	xlsx.SetCellStyle(sheet, "R"+iStr, "R"+iStr, orderTotalLStyle)
	xlsx.SetCellStr(sheet, "R"+iStr, "Order Total:")
	xlsx.MergeCell(sheet, "V"+iStr, "X"+iStr)
	xlsx.SetCellStyle(sheet, "V"+iStr, "V"+iStr, orderTotalVStyle)
	xlsx.SetCellValue(sheet, "V"+iStr, *ord.Order_total)

	// 表头要变一下
	xlsx.SetCellStr(sheet, "A1", "ORDER CONFIRMATION")

	err = xlsx.SaveAs(fileName)
	return err
}

func (ordService *OrderService) ExportDeliveryNoteExcel(orderId *int, fileName string) (err error) {
	var ord order.Order
	if ord, err = ordService.GetOrder(uint(*orderId)); err != nil {
		return err
	}

	var subOrderList []order.SubOrder
	if err = global.GVA_DB.Where("order_id = ?", orderId).Order("product_code asc").Find(&subOrderList).Error; err != nil {
		return err
	}

	var xlsx *excelize.File
	if xlsx, err = excelize.OpenFile(global.GVA_CONFIG.Excel.Dir + "template_delivery.xlsx"); err != nil {
		return err
	}
	var sheet = xlsx.GetSheetName(0)

	xlsx.SetCellStr(sheet, "A6", ord.Bill_to)
	xlsx.SetCellStr(sheet, "A8", ord.Bill_to_address)
	xlsx.SetCellStr(sheet, "A10", ord.Bill_to_city)
	xlsx.SetCellStr(sheet, "A11", ord.Bill_to_postcode)
	xlsx.SetCellStr(sheet, "A12", ord.Bill_to_phone)
	xlsx.SetCellStr(sheet, "I6", ord.Ship_to)
	xlsx.SetCellStr(sheet, "I8", ord.Ship_to_address)
	xlsx.SetCellStr(sheet, "I10", ord.Ship_to_city)
	xlsx.SetCellStr(sheet, "I11", ord.Ship_to_postcode)
	xlsx.SetCellStr(sheet, "I12", ord.Ship_to_phone)
	datStyle, _ := xlsx.NewStyle(`{"custom_number_format":"dd/mm/yyyy;@"}`)
	xlsx.SetCellStyle(sheet, "U6", "U6", datStyle)
	xlsx.SetCellValue(sheet, "U6", returnDateStringOrNull(ord.Order_date))

	xlsx.SetCellStr(sheet, "U7", ord.Po_number)
	xlsx.SetCellStyle(sheet, "U8", "U8", datStyle)

	xlsx.SetCellValue(sheet, "U8", returnDateStringOrNull(ord.Invoice_date))

	xlsx.SetCellStr(sheet, "U9", ord.Invoice_no)
	xlsx.SetCellStr(sheet, "U10", ord.Payment_method)

	i := 0
	j := 0

	var strStyle, dateStyle, numStyle int
	strStyle, _ = xlsx.GetCellStyle(sheet, "A15")
	dateStyle, _ = xlsx.GetCellStyle(sheet, "S15")
	numStyle, _ = xlsx.GetCellStyle(sheet, "V15")

	for i, j = 0, len(subOrderList); i < j; i++ {
		iStr := strconv.Itoa(14 + i)

		if err != nil {
			fmt.Println(err)
		}
		if i%2 == 1 {
			xlsx.SetCellStyle(sheet, "A"+iStr, "R"+iStr, strStyle)
			xlsx.SetCellStyle(sheet, "S"+iStr, "U"+iStr, dateStyle)
			xlsx.SetCellStyle(sheet, "V"+iStr, "X"+iStr, numStyle)
		}
		xlsx.MergeCell(sheet, "A"+iStr, "C"+iStr)
		xlsx.SetCellStr(sheet, "A"+iStr, subOrderList[i].Product_code)
		xlsx.MergeCell(sheet, "D"+iStr, "O"+iStr)
		xlsx.SetCellStr(sheet, "D"+iStr, subOrderList[i].Product_name_cn+" "+subOrderList[i].Product_name_en)
		xlsx.MergeCell(sheet, "P"+iStr, "R"+iStr)
		xlsx.SetCellStr(sheet, "P"+iStr, subOrderList[i].Package)
		xlsx.MergeCell(sheet, "S"+iStr, "U"+iStr)
		xlsx.SetCellValue(sheet, "S"+iStr, subOrderList[i].Exp_date.Format("02/01/2006"))
		xlsx.MergeCell(sheet, "V"+iStr, "X"+iStr)
		xlsx.SetCellInt(sheet, "V"+iStr, *subOrderList[i].Quantity)
	}

	bigNumStyle, err := xlsx.NewStyle(`{"font":{"bold":false,"italic":false,"size":11}}`)

	if err != nil {
		fmt.Println(err)
	}
	i++
	iStr := strconv.Itoa(14 + i)

	xlsx.MergeCell(sheet, "V"+iStr, "X"+iStr)
	xlsx.SetCellStyle(sheet, "V"+iStr, "X"+iStr, bigNumStyle) // QTY
	xlsx.SetCellValue(sheet, "V"+iStr, *ord.Quantity)

	err = xlsx.SaveAs(fileName)
	return err
}

func (ordService *OrderService) ExportOrderExcel(info orderReq.OrderSearch, fileName string) (err error) {
	// 创建db
	db := global.GVA_DB.Model(&order.Order{})
	var ords []order.Order
	// 如果有条件搜索 下方会自动创建搜索语句
	if info.StartCreatedAt != nil {
		db = db.Where("created_at >= ?", info.StartCreatedAt)
	}
	if info.EndCreatedAt != nil {
		db = db.Where("create_at <= ?", info.EndCreatedAt)
	}
	if info.StartOrder_date != nil {
		db = db.Where("order_date >= ? ", info.StartOrder_date)
	}
	if info.EndOrder_date != nil {
		db = db.Where("order_date <= ? ", info.EndOrder_date)
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

	var OrderStr string // 排序
	orderMap := make(map[string]bool)
	orderMap["order_date"] = true
	if orderMap[info.Sort] {
		OrderStr = info.Sort
		if info.OrderStr == "descending" {
			OrderStr = OrderStr + " desc"
		}
		db = db.Order(OrderStr)
	} else {
		db = db.Order("id desc")
	}
	err = db.Find(&ords).Error

	// excel
	var xlsx *excelize.File
	xlsx = excelize.NewFile()
	var sheet = xlsx.GetSheetName(0)

	err = xlsx.SetSheetRow(sheet, "A1", &[]string{
		"Text10",
		"CustomerName",
		"Text11",
		"Date",
		"Text12",
		"Text13",
		"SumOfExtPrice2",
		"Text14",
		"SumOfVat"})

	for i, j := 0, len(ords); i < j; i++ {
		var company company2.Company
		ord := ords[i]
		err = global.GVA_DB.Where("id = ?", ord.Customer_id).First(&company).Error
		if err != nil {
			return err
		}
		axis := fmt.Sprintf("A%d", i+2)

		price2, _ := decimal.NewFromFloat(*ord.Subtotal).Sub(decimal.NewFromFloat(*ord.Discount)).Sub(decimal.NewFromFloat(*ord.Hand_discount)).Float64()
		var t12 string
		if decimal.NewFromFloat(*ord.Vat).Cmp(decimal.NewFromInt32(0)) == 1 {
			t12 = "T1"
		} else {
			t12 = "T2"
		}
		err = xlsx.SetSheetRow(sheet, axis, &[]interface{}{
			"SI",
			company.Sage,
			4000,
			returnDateStringOrNull(ord.Invoice_date),
			ord.Invoice_no,
			"GOODS",
			price2,
			t12,
			*ord.Vat - *ord.HandDiscountVat,
		})
	}

	err = xlsx.SaveAs(fileName)
	return err
}

func returnDateStringOrNull(val *time.Time) string {
	if val == nil {
		return ""
	} else {
		return val.Format("02/01/2006")
	}
}
