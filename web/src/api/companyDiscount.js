import service from '@/utils/request'

// @Tags CompanyDiscount
// @Summary 创建CompanyDiscount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompanyDiscount true "创建CompanyDiscount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /compDiscount/createCompanyDiscount [post]
export const createCompanyDiscount = (data) => {
  return service({
    url: '/compDiscount/createCompanyDiscount',
    method: 'post',
    data
  })
}

// @Tags CompanyDiscount
// @Summary 删除CompanyDiscount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompanyDiscount true "删除CompanyDiscount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /compDiscount/deleteCompanyDiscount [delete]
export const deleteCompanyDiscount = (data) => {
  return service({
    url: '/compDiscount/deleteCompanyDiscount',
    method: 'delete',
    data
  })
}

// @Tags CompanyDiscount
// @Summary 删除CompanyDiscount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除CompanyDiscount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /compDiscount/deleteCompanyDiscount [delete]
export const deleteCompanyDiscountByIds = (data) => {
  return service({
    url: '/compDiscount/deleteCompanyDiscountByIds',
    method: 'delete',
    data
  })
}

// @Tags CompanyDiscount
// @Summary 更新CompanyDiscount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.CompanyDiscount true "更新CompanyDiscount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /compDiscount/updateCompanyDiscount [put]
export const updateCompanyDiscount = (data) => {
  return service({
    url: '/compDiscount/updateCompanyDiscount',
    method: 'put',
    data
  })
}

// @Tags CompanyDiscount
// @Summary 用id查询CompanyDiscount
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.CompanyDiscount true "用id查询CompanyDiscount"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /compDiscount/findCompanyDiscount [get]
export const findCompanyDiscount = (params) => {
  return service({
    url: '/compDiscount/findCompanyDiscount',
    method: 'get',
    params
  })
}

// @Tags CompanyDiscount
// @Summary 分页获取CompanyDiscount列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取CompanyDiscount列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /compDiscount/getCompanyDiscountList [get]
export const getCompanyDiscountList = (params) => {
  return service({
    url: '/compDiscount/getCompanyDiscountList',
    method: 'get',
    params
  })
}

export const addCompDiscountByProductIds = (data) => {
  return service({
    url: '/compDiscount/addCompDiscountByProductIds',
    method: 'put',
    data
  })
}
