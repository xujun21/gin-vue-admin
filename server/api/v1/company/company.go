package company

import (
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

type CompanyApi struct {
}

var compService = service.ServiceGroupApp.CompanyServiceGroup.CompanyService

// CreateCompany 创建Company
// @Tags Company
// @Summary 创建Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body company.Company true "创建Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comp/createCompany [post]
func (compApi *CompanyApi) CreateCompany(c *gin.Context) {
	var comp company.Company
	err := c.ShouldBindJSON(&comp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	comp.CreatedBy = utils.GetUserID(c)
	if err := compService.CreateCompany(comp); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteCompany 删除Company
// @Tags Company
// @Summary 删除Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body company.Company true "删除Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /comp/deleteCompany [delete]
func (compApi *CompanyApi) DeleteCompany(c *gin.Context) {
	var comp company.Company
	err := c.ShouldBindJSON(&comp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	comp.DeletedBy = utils.GetUserID(c)
	if err := compService.DeleteCompany(comp); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteCompanyByIds 批量删除Company
// @Tags Company
// @Summary 批量删除Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /comp/deleteCompanyByIds [delete]
func (compApi *CompanyApi) DeleteCompanyByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := compService.DeleteCompanyByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateCompany 更新Company
// @Tags Company
// @Summary 更新Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body company.Company true "更新Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /comp/updateCompany [put]
func (compApi *CompanyApi) UpdateCompany(c *gin.Context) {
	var comp company.Company
	err := c.ShouldBindJSON(&comp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	comp.UpdatedBy = utils.GetUserID(c)
	if err := compService.UpdateCompany(comp); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindCompany 用id查询Company
// @Tags Company
// @Summary 用id查询Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query company.Company true "用id查询Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /comp/findCompany [get]
func (compApi *CompanyApi) FindCompany(c *gin.Context) {
	var comp company.Company
	err := c.ShouldBindQuery(&comp)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if recomp, err := compService.GetCompany(comp.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"recomp": recomp}, c)
	}
}

// GetCompanyList 分页获取Company列表
// @Tags Company
// @Summary 分页获取Company列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query companyReq.CompanySearch true "分页获取Company列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comp/getCompanyList [get]
func (compApi *CompanyApi) GetCompanyList(c *gin.Context) {
	var pageInfo companyReq.CompanySearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := compService.GetCompanyInfoList(pageInfo); err != nil {
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
