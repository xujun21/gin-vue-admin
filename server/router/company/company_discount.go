package company

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type CompanyDiscountRouter struct {
}

// InitCompanyDiscountRouter 初始化 CompanyDiscount 路由信息
func (s *CompanyDiscountRouter) InitCompanyDiscountRouter(Router *gin.RouterGroup) {
	compDiscountRouter := Router.Group("compDiscount").Use(middleware.OperationRecord())
	compDiscountRouterWithoutRecord := Router.Group("compDiscount")
	var compDiscountApi = v1.ApiGroupApp.CompanyApiGroup.CompanyDiscountApi
	{
		compDiscountRouter.POST("createCompanyDiscount", compDiscountApi.CreateCompanyDiscount)             // 新建CompanyDiscount
		compDiscountRouter.DELETE("deleteCompanyDiscount", compDiscountApi.DeleteCompanyDiscount)           // 删除CompanyDiscount
		compDiscountRouter.DELETE("deleteCompanyDiscountByIds", compDiscountApi.DeleteCompanyDiscountByIds) // 批量删除CompanyDiscount
		compDiscountRouter.PUT("updateCompanyDiscount", compDiscountApi.UpdateCompanyDiscount)              // 更新CompanyDiscount
		compDiscountRouter.PUT("addCompDiscountByProductIds", compDiscountApi.AddCompDiscountByProductIds)
	}
	{
		compDiscountRouterWithoutRecord.GET("findCompanyDiscount", compDiscountApi.FindCompanyDiscount)       // 根据ID获取CompanyDiscount
		compDiscountRouterWithoutRecord.GET("getCompanyDiscountList", compDiscountApi.GetCompanyDiscountList) // 获取CompanyDiscount列表
	}
}
