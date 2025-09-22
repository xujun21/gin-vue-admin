package product

import (
	v1 "github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
}

// InitProductRouter 初始化 Product 路由信息
func (s *ProductRouter) InitProductRouter(Router *gin.RouterGroup) {
	prodRouter := Router.Group("prod").Use(middleware.OperationRecord())
	prodRouterWithoutRecord := Router.Group("prod")
	var prodApi = v1.ApiGroupApp.ProductApiGroup.ProductApi
	{
		prodRouter.POST("createProduct", prodApi.CreateProduct)             // 新建Product
		prodRouter.DELETE("deleteProduct", prodApi.DeleteProduct)           // 删除Product
		prodRouter.DELETE("deleteProductByIds", prodApi.DeleteProductByIds) // 批量删除Product
		prodRouter.PUT("updateProduct", prodApi.UpdateProduct)              // 更新Product
		prodRouter.POST("exportProductExcel", prodApi.ExportProductExcel)   // 导出产品
	}
	{
		prodRouterWithoutRecord.GET("findProduct", prodApi.FindProduct)                     // 根据ID获取Product
		prodRouterWithoutRecord.GET("getProductList", prodApi.GetProductList)               // 获取Product列表
		prodRouterWithoutRecord.GET("getDeletedProductList", prodApi.GetDeletedProductList) // 获取已删除的Product列表
	}
}
