package ReqOrder

import (
	"strconv"

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

type ReqSubOrderApi struct {
}

var reqSubOrderService = service.ServiceGroupApp.ReqorderServiceGroup.ReqSubOrderService

// CreateReqSubOrder 创建ReqSubOrder
// @Tags ReqSubOrder
// @Summary 创建ReqSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ReqOrder.ReqSubOrder true "创建ReqSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /reqSubOrder/createReqSubOrder [post]
func (reqSubOrderApi *ReqSubOrderApi) CreateReqSubOrder(c *gin.Context) {
	var reqSubOrder ReqOrder.ReqSubOrder
	err := c.ShouldBindJSON(&reqSubOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := reqSubOrderService.CreateReqSubOrder(reqSubOrder); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteReqSubOrder 删除ReqSubOrder
// @Tags ReqSubOrder
// @Summary 删除ReqSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ReqOrder.ReqSubOrder true "删除ReqSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /reqSubOrder/deleteReqSubOrder [delete]
func (reqSubOrderApi *ReqSubOrderApi) DeleteReqSubOrder(c *gin.Context) {
	var reqSubOrder ReqOrder.ReqSubOrder
	err := c.ShouldBindJSON(&reqSubOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := reqSubOrderService.DeleteReqSubOrder(reqSubOrder); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteReqSubOrderByIds 批量删除ReqSubOrder
// @Tags ReqSubOrder
// @Summary 批量删除ReqSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ReqSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /reqSubOrder/deleteReqSubOrderByIds [delete]
func (reqSubOrderApi *ReqSubOrderApi) DeleteReqSubOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := reqSubOrderService.DeleteReqSubOrderByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateReqSubOrder 更新ReqSubOrder
// @Tags ReqSubOrder
// @Summary 更新ReqSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body ReqOrder.ReqSubOrder true "更新ReqSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /reqSubOrder/updateReqSubOrder [put]
func (reqSubOrderApi *ReqSubOrderApi) UpdateReqSubOrder(c *gin.Context) {
	var reqSubOrder ReqOrder.ReqSubOrder
	err := c.ShouldBindJSON(&reqSubOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := reqSubOrderService.UpdateReqSubOrder(reqSubOrder); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindReqSubOrder 用id查询ReqSubOrder
// @Tags ReqSubOrder
// @Summary 用id查询ReqSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ReqOrder.ReqSubOrder true "用id查询ReqSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /reqSubOrder/findReqSubOrder [get]
func (reqSubOrderApi *ReqSubOrderApi) FindReqSubOrder(c *gin.Context) {
	var reqSubOrder ReqOrder.ReqSubOrder
	err := c.ShouldBindQuery(&reqSubOrder)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if rereqSubOrder, err := reqSubOrderService.GetReqSubOrder(reqSubOrder.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"rereqSubOrder": rereqSubOrder}, c)
	}
}

// GetReqSubOrderList 分页获取ReqSubOrder列表
// @Tags ReqSubOrder
// @Summary 分页获取ReqSubOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query ReqOrderReq.ReqSubOrderSearch true "分页获取ReqSubOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /reqSubOrder/getReqSubOrderList [get]
func (reqSubOrderApi *ReqSubOrderApi) GetReqSubOrderList(c *gin.Context) {
	var pageInfo ReqOrderReq.ReqSubOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := reqSubOrderService.GetReqSubOrderInfoList(pageInfo); err != nil {
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

func (reqSubOrdApi *ReqSubOrderApi) AddReqSubOrderByProductIds(c *gin.Context) {
	var ReqOrderIdAndIdsReq ReqOrderReq.ReqOrderIdAndIdsReq

	err := c.ShouldBindJSON(&ReqOrderIdAndIdsReq)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var IDS request.IdsReq
	IDS.Ids = ReqOrderIdAndIdsReq.Ids

	reqOrderStr := ReqOrderIdAndIdsReq.ReqOrderId
	reqOrderId, _ := strconv.Atoi(reqOrderStr)

	createdBy := utils.GetUserID(c)
	if err := reqSubOrderService.AddReqSubOrderByProductIds(&reqOrderId, IDS, createdBy); err != nil {
		global.GVA_LOG.Error("批量添加失败!", zap.Error(err))
		response.FailWithMessage("批量添加失败", c)
	} else {
		response.OkWithMessage("批量添加成功", c)
	}
	// 更新订单
	// if err := reqSubOrderService.CalcOrder(&reqOrderId); err != nil {
	// 	global.GVA_LOG.Error("更新订单失败!", zap.Error(err))
	// 	response.FailWithMessage("更新订单失败", c)
	// } else {
	// 	println("更新订单成功")
	// }
}
