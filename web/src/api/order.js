import service from '@/utils/request'
import { ElMessage } from 'element-plus'

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
  if (params['ID'] !== '0') {
    return service({
      url: '/ord/findOrder',
      method: 'get',
      params
    })
  }
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

export const exportInvoiceExcel = (orderId) => {
  service({
    url: '/ord/exportInvoiceExcel',
    method: 'post',
    data: {
      orderId: orderId
    },
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, 'invoice' + orderId + '.xlsx')
  })
}

export const exportConfirmExcel = (orderId) => {
  service({
    url: '/ord/exportConfirmExcel',
    method: 'post',
    data: {
      orderId: orderId
    },
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, 'confirm' + orderId + '.xlsx')
  })
}

export const exportDeliveryNoteExcel = (orderId) => {
  service({
    url: '/ord/exportDeliveryNoteExcel',
    method: 'post',
    data: {
      orderId: orderId
    },
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, 'delivery' + orderId + '.xlsx')
  })
}

export const exportOrderExcel = (params) => {
  service({
    url: '/ord/exportOrderExcel',
    method: 'post',
    params,
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, 'OrderList.xlsx')
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
