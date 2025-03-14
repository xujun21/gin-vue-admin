package order

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/order"
	orderReq "github.com/flipped-aurora/gin-vue-admin/server/model/order/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

type OrderApi struct {
}

var ordService = service.ServiceGroupApp.OrderServiceGroup.OrderService

// CreateOrder 创建Order
// @Tags Order
// @Summary 创建Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body order.Order true "创建Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ord/createOrder [post]
func (ordApi *OrderApi) CreateOrder(c *gin.Context) {
	var ord order.Order
	err := c.ShouldBindJSON(&ord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ord.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Customer_company_name": {utils.NotEmpty()},
	}
	if err := utils.Verify(ord, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ordService.CreateOrder(ord); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteOrder 删除Order
// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body order.Order true "删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ord/deleteOrder [delete]
func (ordApi *OrderApi) DeleteOrder(c *gin.Context) {
	var ord order.Order
	err := c.ShouldBindJSON(&ord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ord.DeletedBy = utils.GetUserID(c)
	if err := ordService.DeleteOrder(ord); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteOrderByIds 批量删除Order
// @Tags Order
// @Summary 批量删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ord/deleteOrderByIds [delete]
func (ordApi *OrderApi) DeleteOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := ordService.DeleteOrderByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateOrder 更新Order
// @Tags Order
// @Summary 更新Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body order.Order true "更新Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ord/updateOrder [put]
func (ordApi *OrderApi) UpdateOrder(c *gin.Context) {
	var ord order.Order
	err := c.ShouldBindJSON(&ord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ord.UpdatedBy = utils.GetUserID(c)
	if err := ordService.UpdateOrder(ord); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindOrder 用id查询Order
// @Tags Order
// @Summary 用id查询Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query order.Order true "用id查询Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ord/findOrder [get]
func (ordApi *OrderApi) FindOrder(c *gin.Context) {
	var ord order.Order
	err := c.ShouldBindQuery(&ord)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reord, err := ordService.GetOrder(ord.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reord": reord}, c)
	}
}

// GetOrderList 分页获取Order列表
// @Tags Order
// @Summary 分页获取Order列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query orderReq.OrderSearch true "分页获取Order列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ord/getOrderList [get]
func (ordApi *OrderApi) GetOrderList(c *gin.Context) {
	var pageInfo orderReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ordService.GetOrderInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}

func (ordApi *OrderApi) ExportInvoiceExcel(c *gin.Context) {
	var excelInfo orderReq.OrderExcelInfo
	err := c.ShouldBindJSON(&excelInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	//if strings.Index(excelInfo.OrderId, "..") > -1 {
	//	response.FailWithMessage("包含非法字符", c)
	//	return
	//}
	filePath := global.GVA_CONFIG.Excel.Dir + "invoice" + strconv.Itoa(*excelInfo.OrderId) + ".xlsx"
	err = ordService.ExportInvoiceExcel(excelInfo.OrderId, filePath)
	if err != nil {
		global.GVA_LOG.Error("下载发票失败!", zap.Error(err))
		response.FailWithMessage("下载发票失败", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}

func (ordApi *OrderApi) ExportConfirmExcel(c *gin.Context) {
	var excelInfo orderReq.OrderExcelInfo
	err := c.ShouldBindJSON(&excelInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	filePath := global.GVA_CONFIG.Excel.Dir + "confirm" + strconv.Itoa(*excelInfo.OrderId) + ".xlsx"
	err = ordService.ExportConfirmExcel(excelInfo.OrderId, filePath)
	if err != nil {
		global.GVA_LOG.Error("下载确认书失败!", zap.Error(err))
		response.FailWithMessage("下载确认书失败", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}

func (ordApi *OrderApi) ExportDeliveryNoteExcel(c *gin.Context) {
	var excelInfo orderReq.OrderExcelInfo
	err := c.ShouldBindJSON(&excelInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	filePath := global.GVA_CONFIG.Excel.Dir + "delivery" + strconv.Itoa(*excelInfo.OrderId) + ".xlsx"
	err = ordService.ExportDeliveryNoteExcel(excelInfo.OrderId, filePath)
	if err != nil {
		global.GVA_LOG.Error("下载运单失败!", zap.Error(err))
		response.FailWithMessage("下载运单失败", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}

func (ordApi *OrderApi) ExportOrderExcel(c *gin.Context) {
	var pageInfo orderReq.OrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	filePath := global.GVA_CONFIG.Excel.Dir + "OrderList.xlsx"
	err = ordService.ExportOrderExcel(pageInfo, filePath)
	if err != nil {
		global.GVA_LOG.Error("下载订单列表失败!", zap.Error(err))
		response.FailWithMessage("下载订单列表失败", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}
