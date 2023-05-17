import service from '@/utils/request'
import { ElMessage } from 'element-plus'

// @Tags RequireOrder
// @Summary 创建RequireOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RequireOrder true "创建RequireOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ReqOrd/createRequireOrder [post]
export const createRequireOrder = (data) => {
  return service({
    url: '/ReqOrd/createRequireOrder',
    method: 'post',
    data
  })
}

// @Tags RequireOrder
// @Summary 删除RequireOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RequireOrder true "删除RequireOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ReqOrd/deleteRequireOrder [delete]
export const deleteRequireOrder = (data) => {
  return service({
    url: '/ReqOrd/deleteRequireOrder',
    method: 'delete',
    data
  })
}

// @Tags RequireOrder
// @Summary 删除RequireOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除RequireOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /ReqOrd/deleteRequireOrder [delete]
export const deleteRequireOrderByIds = (data) => {
  return service({
    url: '/ReqOrd/deleteRequireOrderByIds',
    method: 'delete',
    data
  })
}

// @Tags RequireOrder
// @Summary 更新RequireOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.RequireOrder true "更新RequireOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /ReqOrd/updateRequireOrder [put]
export const updateRequireOrder = (data) => {
  return service({
    url: '/ReqOrd/updateRequireOrder',
    method: 'put',
    data
  })
}

// @Tags RequireOrder
// @Summary 用id查询RequireOrder
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.RequireOrder true "用id查询RequireOrder"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /ReqOrd/findRequireOrder [get]
export const findRequireOrder = (params) => {
  return service({
    url: '/ReqOrd/findRequireOrder',
    method: 'get',
    params
  })
}

// @Tags RequireOrder
// @Summary 分页获取RequireOrder列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取RequireOrder列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /ReqOrd/getRequireOrderList [get]
export const getRequireOrderList = (params) => {
  return service({
    url: '/ReqOrd/getRequireOrderList',
    method: 'get',
    params
  })
}

export const exportRequireExcel = (reqOrderId) => {
  // console.log(reqOrderId)
  service({
    url: '/ReqOrd/exportRequireExcel',
    method: 'post',
    data: {
      reqOrderId: reqOrderId
    },
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, 'RequireOrder' + reqOrderId + '.xlsx')
  })
}

const handleFileError = (res, fileName) => {
  // console.log(typeof (res.data))
  if (typeof (res.data) !== 'undefined') {
    if (res.data.type === 'application/json') {
      const reader = new FileReader()
      reader.onload = function() {
        const message = JSON.parse(reader.result).msg
        ElMessage({
          showClose: true,
          message: message,
          type: 'error'
        })
      }
      reader.readAsText(new Blob([res.data]))
    }
  } else {
    var downloadUrl = window.URL.createObjectURL(new Blob([res]))
    var a = document.createElement('a')
    a.style.display = 'none'
    a.href = downloadUrl
    // console.log(downloadUrl)
    a.download = fileName
    var event = new MouseEvent('click')
    a.dispatchEvent(event)
  }
}
