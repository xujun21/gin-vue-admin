import service from '@/utils/request'

// @Tags SubOrder
// @Summary 创建SubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SubOrder true "创建SubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /subOrd/createSubOrder [post]
export const createSubOrder = (data) => {
  return service({
    url: '/subOrd/createSubOrder',
    method: 'post',
    data
  })
}

// @Tags SubOrder
// @Summary 删除SubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SubOrder true "删除SubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /subOrd/deleteSubOrder [delete]
export const deleteSubOrder = (data) => {
  return service({
    url: '/subOrd/deleteSubOrder',
    method: 'delete',
    data
  })
}

// @Tags SubOrder
// @Summary 删除SubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除SubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /subOrd/deleteSubOrder [delete]
export const deleteSubOrderByIds = (data) => {
  return service({
    url: '/subOrd/deleteSubOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags SubOrder
// @Summary 更新SubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.SubOrder true "更新SubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /subOrd/updateSubOrder [put]
export const updateSubOrder = (data) => {
  return service({
    url: '/subOrd/updateSubOrder',
    method: 'put',
    data
  })
}

// @Tags SubOrder
// @Summary 用id查询SubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.SubOrder true "用id查询SubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /subOrd/findSubOrder [get]
export const findSubOrder = (params) => {
  return service({
    url: '/subOrd/findSubOrder',
    method: 'get',
    params
  })
}

// @Tags SubOrder
// @Summary 分页获取SubOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取SubOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /subOrd/getSubOrderList [get]
export const getSubOrderList = (params) => {
  return service({
    url: '/subOrd/getSubOrderList',
    method: 'get',
    params
  })
}

export const addSubOrderByProductIds = (data) => {
  return service({
    url: '/subOrd/addSubOrderByProductIds',
    method: 'put',
    data
  })
}
