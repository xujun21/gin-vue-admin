package product

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/model/product"
	productReq "github.com/flipped-aurora/gin-vue-admin/server/model/product/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ProductApi struct {
}

var prodService = service.ServiceGroupApp.ProductServiceGroup.ProductService

// CreateProduct 创建Product
// @Tags Product
// @Summary 创建Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body product.Product true "创建Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /prod/createProduct [post]
func (prodApi *ProductApi) CreateProduct(c *gin.Context) {
	var prod product.Product
	err := c.ShouldBindJSON(&prod)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	prod.CreatedBy = utils.GetUserID(c)
	if err := prodService.CreateProduct(prod); err != nil {
		global.GVA_LOG.Error("创建失败", zap.Error(err))
		response.FailWithMessage("创建失败: "+err.Error(), c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteProduct 删除Product
// @Tags Product
// @Summary 删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body product.Product true "删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /prod/deleteProduct [delete]
func (prodApi *ProductApi) DeleteProduct(c *gin.Context) {
	var prod product.Product
	err := c.ShouldBindJSON(&prod)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	prod.DeletedBy = utils.GetUserID(c)
	if err := prodService.DeleteProduct(prod); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// RestoreProduct 恢复Product
// @Tags Product
// @Summary 恢复Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body product.Product true "恢复Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"恢复成功"}"
// @Router /prod/restoreProduct [put]
func (prodApi *ProductApi) RestoreProduct(c *gin.Context) {
	var prod product.Product
	err := c.ShouldBindJSON(&prod)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	prod.UpdatedBy = utils.GetUserID(c)
	if err := prodService.RestoreProduct(prod); err != nil {
		global.GVA_LOG.Error("恢复失败!", zap.Error(err))
		response.FailWithMessage("恢复失败", c)
	} else {
		response.OkWithMessage("恢复成功", c)
	}
}

// DeleteProductByIds 批量删除Product
// @Tags Product
// @Summary 批量删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /prod/deleteProductByIds [delete]
func (prodApi *ProductApi) DeleteProductByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	deletedBy := utils.GetUserID(c)
	if err := prodService.DeleteProductByIds(IDS, deletedBy); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateProduct 更新Product
// @Tags Product
// @Summary 更新Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body product.Product true "更新Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /prod/updateProduct [put]
func (prodApi *ProductApi) UpdateProduct(c *gin.Context) {
	var prod product.Product
	err := c.ShouldBindJSON(&prod)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	prod.UpdatedBy = utils.GetUserID(c)
	if err := prodService.UpdateProduct(prod); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindProduct 用id查询Product
// @Tags Product
// @Summary 用id查询Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query product.Product true "用id查询Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /prod/findProduct [get]
func (prodApi *ProductApi) FindProduct(c *gin.Context) {
	var prod product.Product
	err := c.ShouldBindQuery(&prod)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reprod, err := prodService.GetProduct(prod.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reprod": reprod}, c)
	}
}

// GetProductList 分页获取Product列表
// @Tags Product
// @Summary 分页获取Product列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query productReq.ProductSearch true "分页获取Product列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /prod/getProductList [get]
func (prodApi *ProductApi) GetProductList(c *gin.Context) {
	var pageInfo productReq.ProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := prodService.GetProductInfoList(pageInfo); err != nil {
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

// GetDeletedProductList 分页获取已经删除的Product列表
// @Tags Product
// @Summary 分页获取Product列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query productReq.ProductSearch true "分页获取Product列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /prod/getDeletedProductList [get]
func (prodApi *ProductApi) GetDeletedProductList(c *gin.Context) {
	var pageInfo productReq.ProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := prodService.GetDeletedProductInfoList(pageInfo); err != nil {
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

func (prodApi *ProductApi) ExportProductExcel(c *gin.Context) {
	var pageInfo productReq.ProductSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}

	filePath := global.GVA_CONFIG.Excel.Dir + "product" + ".xlsx"
	err = prodService.ExportProductExcel(pageInfo, filePath)
	if err != nil {
		global.GVA_LOG.Error("下载商品表失败!", zap.Error(err))
		response.FailWithMessage("下载商品表失败", c)
		return
	}
	c.Writer.Header().Add("success", "true")
	c.File(filePath)

}
