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

type SubOrderApi struct {
}

var subOrdService = service.ServiceGroupApp.OrderServiceGroup.SubOrderService

// CreateSubOrder 创建SubOrder
// @Tags SubOrder
// @Summary 创建SubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body order.SubOrder true "创建SubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /subOrd/createSubOrder [post]
func (subOrdApi *SubOrderApi) CreateSubOrder(c *gin.Context) {
	var subOrd order.SubOrder
	err := c.ShouldBindJSON(&subOrd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	subOrd.CreatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Order_id":        {utils.NotEmpty()},
		"Product_id":      {utils.NotEmpty()},
		"Product_code":    {utils.NotEmpty()},
		"Product_name_cn": {utils.NotEmpty()},
	}
	if err := utils.Verify(subOrd, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := subOrdService.CreateSubOrder(subOrd); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
	// 更新订单
	if err := ordService.CalcOrder(subOrd.Order_id); err != nil {
		global.GVA_LOG.Error("更新订单失败!", zap.Error(err))
		response.FailWithMessage("更新订单失败", c)
	}
}

// DeleteSubOrder 删除SubOrder
// @Tags SubOrder
// @Summary 删除SubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body order.SubOrder true "删除SubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /subOrd/deleteSubOrder [delete]
func (subOrdApi *SubOrderApi) DeleteSubOrder(c *gin.Context) {
	var subOrd order.SubOrder
	err := c.ShouldBindJSON(&subOrd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取orderId
	subOrd, err = subOrdService.GetSubOrder(subOrd.ID)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	orderId := subOrd.Order_id

	subOrd.DeletedBy = utils.GetUserID(c)
	if err := subOrdService.DeleteSubOrder(subOrd); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}

	// 更新订单
	if err := ordService.CalcOrder(orderId); err != nil {
		global.GVA_LOG.Error("更新订单失败!", zap.Error(err))
		response.FailWithMessage("更新订单失败", c)
	}
}

// DeleteSubOrderByIds 批量删除SubOrder
// @Tags SubOrder
// @Summary 批量删除SubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /subOrd/deleteSubOrderByIds [delete]
func (subOrdApi *SubOrderApi) DeleteSubOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	// 获取orderId
	var orderId *int
	if len(IDS.Ids) > 0 {
		subOrder, _ := subOrdService.GetSubOrder(uint(IDS.Ids[0]))
		orderId = subOrder.Order_id
	} else {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
		return
	}

	deletedBy := utils.GetUserID(c)
	if err := subOrdService.DeleteSubOrderByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}

	// 更新订单
	if err := ordService.CalcOrder(orderId); err != nil {
		global.GVA_LOG.Error("更新订单失败!", zap.Error(err))
		response.FailWithMessage("更新订单失败", c)
	}
}

// UpdateSubOrder 更新SubOrder
// @Tags SubOrder
// @Summary 更新SubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body order.SubOrder true "更新SubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /subOrd/updateSubOrder [put]
func (subOrdApi *SubOrderApi) UpdateSubOrder(c *gin.Context) {
	var subOrd order.SubOrder
	err := c.ShouldBindJSON(&subOrd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	subOrd.UpdatedBy = utils.GetUserID(c)
	verify := utils.Rules{
		"Order_id":        {utils.NotEmpty()},
		"Product_id":      {utils.NotEmpty()},
		"Product_code":    {utils.NotEmpty()},
		"Product_name_cn": {utils.NotEmpty()},
	}
	if err := utils.Verify(subOrd, verify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := subOrdService.UpdateSubOrder(subOrd); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}

	// 更新订单
	if err := ordService.CalcOrder(subOrd.Order_id); err != nil {
		global.GVA_LOG.Error("更新订单失败!", zap.Error(err))
		response.FailWithMessage("更新订单失败", c)
	}
}

// FindSubOrder 用id查询SubOrder
// @Tags SubOrder
// @Summary 用id查询SubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query order.SubOrder true "用id查询SubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /subOrd/findSubOrder [get]
func (subOrdApi *SubOrderApi) FindSubOrder(c *gin.Context) {
	var subOrd order.SubOrder
	err := c.ShouldBindQuery(&subOrd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if resubOrd, err := subOrdService.GetSubOrder(subOrd.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"resubOrd": resubOrd}, c)
	}
}

// GetSubOrderList 分页获取SubOrder列表
// @Tags SubOrder
// @Summary 分页获取SubOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query orderReq.SubOrderSearch true "分页获取SubOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /subOrd/getSubOrderList [get]
func (subOrdApi *SubOrderApi) GetSubOrderList(c *gin.Context) {
	var pageInfo orderReq.SubOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := subOrdService.GetSubOrderInfoList(pageInfo); err != nil {
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

func (subOrdApi *SubOrderApi) AddSubOrderByProductIds(c *gin.Context) {
	var OrderIdAndIdsReq orderReq.OrderIdAndIdsReq

	err := c.ShouldBindJSON(&OrderIdAndIdsReq)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var IDS request.IdsReq
	IDS.Ids = OrderIdAndIdsReq.Ids

	orderStr := OrderIdAndIdsReq.OrderId
	orderId, _ := strconv.Atoi(orderStr)

	createdBy := utils.GetUserID(c)
	if err := subOrdService.AddSubOrderByProductIds(&orderId, IDS, createdBy); err != nil {
		global.GVA_LOG.Error("批量添加失败!", zap.Error(err))
		response.FailWithMessage("批量添加失败", c)
	} else {
		response.OkWithMessage("批量添加成功", c)
	}
	// 更新订单
	if err := ordService.CalcOrder(&orderId); err != nil {
		global.GVA_LOG.Error("更新订单失败!", zap.Error(err))
		response.FailWithMessage("更新订单失败", c)
	}
}
