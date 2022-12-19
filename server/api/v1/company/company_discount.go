package company

import (
	"strconv"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/company"
	companyReq "github.com/flipped-aurora/gin-vue-admin/server/model/company/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type CompanyDiscountApi struct {
}

var compDiscountService = service.ServiceGroupApp.CompanyServiceGroup.CompanyDiscountService

// CreateCompanyDiscount 创建CompanyDiscount
// @Tags CompanyDiscount
// @Summary 创建CompanyDiscount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body company.CompanyDiscount true "创建CompanyDiscount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /compDiscount/createCompanyDiscount [post]
func (compDiscountApi *CompanyDiscountApi) CreateCompanyDiscount(c *gin.Context) {
	var compDiscount company.CompanyDiscount
	err := c.ShouldBindJSON(&compDiscount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	compDiscount.CreatedBy = utils.GetUserID(c)
	if err := compDiscountService.CreateCompanyDiscount(compDiscount); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCompanyDiscount 删除CompanyDiscount
// @Tags CompanyDiscount
// @Summary 删除CompanyDiscount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body company.CompanyDiscount true "删除CompanyDiscount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /compDiscount/deleteCompanyDiscount [delete]
func (compDiscountApi *CompanyDiscountApi) DeleteCompanyDiscount(c *gin.Context) {
	var compDiscount company.CompanyDiscount
	err := c.ShouldBindJSON(&compDiscount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	compDiscount.DeletedBy = utils.GetUserID(c)
	if err := compDiscountService.DeleteCompanyDiscount(compDiscount); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCompanyDiscountByIds 批量删除CompanyDiscount
// @Tags CompanyDiscount
// @Summary 批量删除CompanyDiscount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CompanyDiscount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /compDiscount/deleteCompanyDiscountByIds [delete]
func (compDiscountApi *CompanyDiscountApi) DeleteCompanyDiscountByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := compDiscountService.DeleteCompanyDiscountByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCompanyDiscount 更新CompanyDiscount
// @Tags CompanyDiscount
// @Summary 更新CompanyDiscount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body company.CompanyDiscount true "更新CompanyDiscount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /compDiscount/updateCompanyDiscount [put]
func (compDiscountApi *CompanyDiscountApi) UpdateCompanyDiscount(c *gin.Context) {
	var compDiscount company.CompanyDiscount
	err := c.ShouldBindJSON(&compDiscount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	compDiscount.UpdatedBy = utils.GetUserID(c)
	if err := compDiscountService.UpdateCompanyDiscount(compDiscount); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCompanyDiscount 用id查询CompanyDiscount
// @Tags CompanyDiscount
// @Summary 用id查询CompanyDiscount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query company.CompanyDiscount true "用id查询CompanyDiscount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /compDiscount/findCompanyDiscount [get]
func (compDiscountApi *CompanyDiscountApi) FindCompanyDiscount(c *gin.Context) {
	var compDiscount company.CompanyDiscount
	err := c.ShouldBindQuery(&compDiscount)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recompDiscount, err := compDiscountService.GetCompanyDiscount(compDiscount.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recompDiscount": recompDiscount}, c)
	}
}

// GetCompanyDiscountList 分页获取CompanyDiscount列表
// @Tags CompanyDiscount
// @Summary 分页获取CompanyDiscount列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query companyReq.CompanyDiscountSearch true "分页获取CompanyDiscount列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /compDiscount/getCompanyDiscountList [get]
func (compDiscountApi *CompanyDiscountApi) GetCompanyDiscountList(c *gin.Context) {
	var pageInfo companyReq.CompanyDiscountSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := compDiscountService.GetCompanyDiscountInfoList(pageInfo); err != nil {
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

func (compDiscountApi *CompanyDiscountApi) AddCompDiscountByProductIds(c *gin.Context) {
	var CompanyIdAndIdsReq companyReq.CompanyIdAndIdsReq

	err := c.ShouldBindJSON(&CompanyIdAndIdsReq)

	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	var IDS request.IdsReq
	IDS.Ids = CompanyIdAndIdsReq.Ids

	companyIdStr := CompanyIdAndIdsReq.CompanyId
	companyId, _ := strconv.Atoi(companyIdStr)

	createdBy := utils.GetUserID(c)
	if err := compDiscountService.AddCompDiscountByProductIds(&companyId, IDS, createdBy); err != nil {
		global.GVA_LOG.Error("批量添加失败!", zap.Error(err))
		response.FailWithMessage("批量添加失败", c)
	} else {
		response.OkWithMessage("批量添加成功", c)
	}
}
