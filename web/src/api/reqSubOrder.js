import service from '@/utils/request'

// @Tags ReqSubOrder
// @Summary 创建ReqSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ReqSubOrder true "创建ReqSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /reqSubOrder/createReqSubOrder [post]
export const createReqSubOrder = (data) => {
  return service({
    url: '/reqSubOrder/createReqSubOrder',
    method: 'post',
    data
  })
}

// @Tags ReqSubOrder
// @Summary 删除ReqSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ReqSubOrder true "删除ReqSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /reqSubOrder/deleteReqSubOrder [delete]
export const deleteReqSubOrder = (data) => {
  return service({
    url: '/reqSubOrder/deleteReqSubOrder',
    method: 'delete',
    data
  })
}

// @Tags ReqSubOrder
// @Summary 删除ReqSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除ReqSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /reqSubOrder/deleteReqSubOrder [delete]
export const deleteReqSubOrderByIds = (data) => {
  return service({
    url: '/reqSubOrder/deleteReqSubOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags ReqSubOrder
// @Summary 更新ReqSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.ReqSubOrder true "更新ReqSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /reqSubOrder/updateReqSubOrder [put]
export const updateReqSubOrder = (data) => {
  return service({
    url: '/reqSubOrder/updateReqSubOrder',
    method: 'put',
    data
  })
}

// @Tags ReqSubOrder
// @Summary 用id查询ReqSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.ReqSubOrder true "用id查询ReqSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /reqSubOrder/findReqSubOrder [get]
export const findReqSubOrder = (params) => {
  return service({
    url: '/reqSubOrder/findReqSubOrder',
    method: 'get',
    params
  })
}

// @Tags ReqSubOrder
// @Summary 分页获取ReqSubOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取ReqSubOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /reqSubOrder/getReqSubOrderList [get]
export const getReqSubOrderList = (params) => {
  return service({
    url: '/reqSubOrder/getReqSubOrderList',
    method: 'get',
    params
  })
}

export const addReqSubOrderByProductIds = (data) => {
  return service({
    url: '/reqSubOrder/addReqSubOrderByProductIds',
    method: 'put',
    data
  })
}