package company

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CompanyRouter struct {
}

// InitCompanyRouter 初始化 Company 路由信息
func (s *CompanyRouter) InitCompanyRouter(Router *gin.RouterGroup) {
	compRouter := Router.Group("comp").Use(middleware.OperationRecord())
	compRouterWithoutRecord := Router.Group("comp")
	var compApi = v1.ApiGroupApp.CompanyApiGroup.CompanyApi
	{
		compRouter.POST("createCompany", compApi.CreateCompany)             // 新建Company
		compRouter.DELETE("deleteCompany", compApi.DeleteCompany)           // 删除Company
		compRouter.DELETE("deleteCompanyByIds", compApi.DeleteCompanyByIds) // 批量删除Company
		compRouter.PUT("updateCompany", compApi.UpdateCompany)              // 更新Company
		compRouter.POST("exportCompanyExcel", compApi.ExportCompanyExcel)   // 下载客户信息
	}
	{
		compRouterWithoutRecord.GET("findCompany", compApi.FindCompany)       // 根据ID获取Company
		compRouterWithoutRecord.GET("getCompanyList", compApi.GetCompanyList) // 获取Company列表
	}
}
