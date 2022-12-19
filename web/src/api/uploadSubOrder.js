import service from '@/utils/request'

// @Tags UploadSubOrder
// @Summary 创建UploadSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UploadSubOrder true "创建UploadSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /upSubOrd/createUploadSubOrder [post]
export const createUploadSubOrder = (data) => {
  return service({
    url: '/upSubOrd/createUploadSubOrder',
    method: 'post',
    data
  })
}

// @Tags UploadSubOrder
// @Summary 删除UploadSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UploadSubOrder true "删除UploadSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /upSubOrd/deleteUploadSubOrder [delete]
export const deleteUploadSubOrder = (data) => {
  return service({
    url: '/upSubOrd/deleteUploadSubOrder',
    method: 'delete',
    data
  })
}

// @Tags UploadSubOrder
// @Summary 删除UploadSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除UploadSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /upSubOrd/deleteUploadSubOrder [delete]
export const deleteUploadSubOrderByIds = (data) => {
  return service({
    url: '/upSubOrd/deleteUploadSubOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags UploadSubOrder
// @Summary 更新UploadSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.UploadSubOrder true "更新UploadSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /upSubOrd/updateUploadSubOrder [put]
export const updateUploadSubOrder = (data) => {
  return service({
    url: '/upSubOrd/updateUploadSubOrder',
    method: 'put',
    data
  })
}

// @Tags UploadSubOrder
// @Summary 用id查询UploadSubOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.UploadSubOrder true "用id查询UploadSubOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /upSubOrd/findUploadSubOrder [get]
export const findUploadSubOrder = (params) => {
  return service({
    url: '/upSubOrd/findUploadSubOrder',
    method: 'get',
    params
  })
}

// @Tags UploadSubOrder
// @Summary 分页获取UploadSubOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取UploadSubOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /upSubOrd/getUploadSubOrderList [get]
export const getUploadSubOrderList = (params) => {
  return service({
    url: '/upSubOrd/getUploadSubOrderList',
    method: 'get',
    params
  })
}

// 执行导入
export const doExecImport = (data) => {
  return service({
    url: '/upSubOrd/doExecImport',
    method: 'post',
    data
  })
}

