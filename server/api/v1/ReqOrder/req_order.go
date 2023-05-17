package ReqOrder

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/ReqOrder"
	ReqOrderReq "github.com/flipped-aurora/gin-vue-admin/server/model/ReqOrder/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type RequireOrderApi struct {
}

var ReqOrdService = service.ServiceGroupApp.ReqorderServiceGroup.RequireOrderService

// CreateRequireOrder 创建RequireOrder
// @Tags RequireOrder
// @Summary 创建RequireOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reqOrder.RequireOrder true "创建RequireOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ReqOrd/createRequireOrder [post]
func (ReqOrdApi *RequireOrderApi) CreateRequireOrder(c *gin.Context) {
	var ReqOrd ReqOrder.RequireOrder
	err := c.ShouldBindJSON(&ReqOrd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ReqOrd.CreatedBy = utils.GetUserID(c)
	if err := ReqOrdService.CreateRequireOrder(ReqOrd); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteRequireOrder 删除RequireOrder
// @Tags RequireOrder
// @Summary 删除RequireOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reqOrder.RequireOrder true "删除RequireOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ReqOrd/deleteRequireOrder [delete]
func (ReqOrdApi *RequireOrderApi) DeleteRequireOrder(c *gin.Context) {
	var ReqOrd ReqOrder.RequireOrder
	err := c.ShouldBindJSON(&ReqOrd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ReqOrd.DeletedBy = utils.GetUserID(c)
	if err := ReqOrdService.DeleteRequireOrder(ReqOrd); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteRequireOrderByIds 批量删除RequireOrder
// @Tags RequireOrder
// @Summary 批量删除RequireOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RequireOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /ReqOrd/deleteRequireOrderByIds [delete]
func (ReqOrdApi *RequireOrderApi) DeleteRequireOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := ReqOrdService.DeleteRequireOrderByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateRequireOrder 更新RequireOrder
// @Tags RequireOrder
// @Summary 更新RequireOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body reqOrder.RequireOrder true "更新RequireOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ReqOrd/updateRequireOrder [put]
func (ReqOrdApi *RequireOrderApi) UpdateRequireOrder(c *gin.Context) {
	var ReqOrd ReqOrder.RequireOrder
	err := c.ShouldBindJSON(&ReqOrd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	ReqOrd.UpdatedBy = utils.GetUserID(c)
	if err := ReqOrdService.UpdateRequireOrder(ReqOrd); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindRequireOrder 用id查询RequireOrder
// @Tags RequireOrder
// @Summary 用id查询RequireOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query reqOrder.RequireOrder true "用id查询RequireOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ReqOrd/findRequireOrder [get]
func (ReqOrdApi *RequireOrderApi) FindRequireOrder(c *gin.Context) {
	var ReqOrd ReqOrder.RequireOrder
	err := c.ShouldBindQuery(&ReqOrd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reReqOrd, err := ReqOrdService.GetRequireOrder(ReqOrd.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reReqOrd": reReqOrd}, c)
	}
}

// GetRequireOrderList 分页获取RequireOrder列表
// @Tags RequireOrder
// @Summary 分页获取RequireOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ReqOrderReq.RequireOrderSearch true "分页获取RequireOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ReqOrd/getRequireOrderList [get]
func (ReqOrdApi *RequireOrderApi) GetRequireOrderList(c *gin.Context) {
	var pageInfo ReqOrderReq.RequireOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := ReqOrdService.GetRequireOrderInfoList(pageInfo); err != nil {
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

func (ReqOrdApi *RequireOrderApi) ExportRequireExcel(c *gin.Context) {
	var excelInfo ReqOrderReq.ReqOrderExcelInfo
	err := c.ShouldBindJSON(&excelInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	filePath := global.GVA_CONFIG.Excel.Dir + "RequireOrder.xlsx"
	err = ReqOrdService.ExportRequireExcel(excelInfo.ReqOrderId, filePath)
	if err != nil {
		global.GVA_LOG.Error("下载采购单失败!", zap.Error(err))
		response.FailWithMessage("下载采购单失败", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)
}
