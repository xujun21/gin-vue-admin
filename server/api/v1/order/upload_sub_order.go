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
	"github.com/gin-gonic/gin/binding"
	"go.uber.org/zap"
	"strconv"
)

type UploadSubOrderApi struct {
}

var upSubOrdService = service.ServiceGroupApp.OrderServiceGroup.UploadSubOrderService

// CreateUploadSubOrder 创建UploadSubOrder
// @Tags UploadSubOrder
// @Summary 创建UploadSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body order.UploadSubOrder true "创建UploadSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /upSubOrd/createUploadSubOrder [post]
func (upSubOrdApi *UploadSubOrderApi) CreateUploadSubOrder(c *gin.Context) {
	var upSubOrd order.UploadSubOrder
	err := c.ShouldBindJSON(&upSubOrd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	upSubOrd.CreatedBy = utils.GetUserID(c)
	if err := upSubOrdService.CreateUploadSubOrder(upSubOrd); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

func (upSubOrdApi *UploadSubOrderApi) UploadSubOrderExcel(c *gin.Context) {
	_, header, err := c.Request.FormFile("file")
	ordId := c.Request.Form.Get("orderId")
	orderId, _ := strconv.Atoi(ordId)

	if err != nil {
		global.GVA_LOG.Error("接收文件失败!", zap.Error(err))
		response.FailWithMessage("接收文件失败", c)
		return
	}
	_ = c.SaveUploadedFile(header, global.GVA_CONFIG.Excel.Dir+ordId+".xlsx")
	response.OkWithMessage("上传成功", c)

	// 清除此订单下的残留导入数据
	err = upSubOrdService.DeleteUploadSubOrderByOrderId(&orderId, utils.GetUserID(c))
	if err != nil {
		global.GVA_LOG.Error("清除上次残留导入数据失败!", zap.Error(err))
		response.FailWithMessage("清除上次残留导入数据失败", c)
		return
	}
	// 导入数据库
	err = upSubOrdService.ParseExcel2InfoList(&orderId)
	if err != nil {
		global.GVA_LOG.Error("保存数据失败!", zap.Error(err))
		response.FailWithMessage("保存数据失败", c)
		return
	}
}

// DeleteUploadSubOrder 删除UploadSubOrder
// @Tags UploadSubOrder
// @Summary 删除UploadSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body order.UploadSubOrder true "删除UploadSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /upSubOrd/deleteUploadSubOrder [delete]
func (upSubOrdApi *UploadSubOrderApi) DeleteUploadSubOrder(c *gin.Context) {
	var upSubOrd order.UploadSubOrder
	err := c.ShouldBindJSON(&upSubOrd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	upSubOrd.DeletedBy = utils.GetUserID(c)
	if err := upSubOrdService.DeleteUploadSubOrder(upSubOrd); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteUploadSubOrderByIds 批量删除UploadSubOrder
// @Tags UploadSubOrder
// @Summary 批量删除UploadSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UploadSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /upSubOrd/deleteUploadSubOrderByIds [delete]
func (upSubOrdApi *UploadSubOrderApi) DeleteUploadSubOrderByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := upSubOrdService.DeleteUploadSubOrderByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateUploadSubOrder 更新UploadSubOrder
// @Tags UploadSubOrder
// @Summary 更新UploadSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body order.UploadSubOrder true "更新UploadSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /upSubOrd/updateUploadSubOrder [put]
func (upSubOrdApi *UploadSubOrderApi) UpdateUploadSubOrder(c *gin.Context) {
	var upSubOrd order.UploadSubOrder
	err := c.ShouldBindJSON(&upSubOrd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	upSubOrd.UpdatedBy = utils.GetUserID(c)
	if err := upSubOrdService.UpdateUploadSubOrder(upSubOrd); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindUploadSubOrder 用id查询UploadSubOrder
// @Tags UploadSubOrder
// @Summary 用id查询UploadSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query order.UploadSubOrder true "用id查询UploadSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /upSubOrd/findUploadSubOrder [get]
func (upSubOrdApi *UploadSubOrderApi) FindUploadSubOrder(c *gin.Context) {
	var upSubOrd order.UploadSubOrder
	err := c.ShouldBindQuery(&upSubOrd)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reupSubOrd, err := upSubOrdService.GetUploadSubOrder(upSubOrd.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reupSubOrd": reupSubOrd}, c)
	}
}

// GetUploadSubOrderList 分页获取UploadSubOrder列表
// @Tags UploadSubOrder
// @Summary 分页获取UploadSubOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query orderReq.UploadSubOrderSearch true "分页获取UploadSubOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /upSubOrd/getUploadSubOrderList [get]
func (upSubOrdApi *UploadSubOrderApi) GetUploadSubOrderList(c *gin.Context) {
	var pageInfo orderReq.UploadSubOrderSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := upSubOrdService.GetUploadSubOrderInfoList(pageInfo); err != nil {
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

func (upSubOrdApi *UploadSubOrderApi) DoExecImport(c *gin.Context) {
	var orderId orderReq.OrderIdReq
	err := c.ShouldBindWith(&orderId, binding.JSON)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	if err := upSubOrdService.DoExecImport(orderId); err != nil {
		global.GVA_LOG.Error("导入失败!", zap.Error(err))
		response.FailWithMessage("导入失败", c)
	} else {
		response.OkWithMessage("导入成功", c)
	}

	// 更新订单
	ordId, _ := strconv.Atoi(orderId.OrderId)
	if err = ordService.CalcOrder(&ordId); err != nil {
		global.GVA_LOG.Error("更新订单失败!", zap.Error(err))
		response.FailWithMessage("更新订单失败", c)
	}

}
