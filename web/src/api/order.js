import service from '@/utils/request'

// @Tags Order
// @Summary 创建Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "创建Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ord/createOrder [post]
export const createOrder = (data) => {
  return service({
    url: '/ord/createOrder',
    method: 'post',
    data
  })
}

// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ord/deleteOrder [delete]
export const deleteOrder = (data) => {
  return service({
    url: '/ord/deleteOrder',
    method: 'delete',
    data
  })
}

// @Tags Order
// @Summary 删除Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ord/deleteOrder [delete]
export const deleteOrderByIds = (data) => {
  return service({
    url: '/ord/deleteOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags Order
// @Summary 更新Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Order true "更新Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ord/updateOrder [put]
export const updateOrder = (data) => {
  return service({
    url: '/ord/updateOrder',
    method: 'put',
    data
  })
}

// @Tags Order
// @Summary 用id查询Order
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Order true "用id查询Order"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ord/findOrder [get]
export const findOrder = (params) => {
  return service({
    url: '/ord/findOrder',
    method: 'get',
    params
  })
}

// @Tags Order
// @Summary 分页获取Order列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Order列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ord/getOrderList [get]
export const getOrderList = (params) => {
  return service({
    url: '/ord/getOrderList',
    method: 'get',
    params
  })
}
