import service from '@/utils/request'

// @Tags Product
// @Summary 创建Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "创建Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /prod/createProduct [post]
export const createProduct = (data) => {
  return service({
    url: '/prod/createProduct',
    method: 'post',
    data
  })
}

// @Tags Product
// @Summary 删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /prod/deleteProduct [delete]
export const deleteProduct = (data) => {
  return service({
    url: '/prod/deleteProduct',
    method: 'delete',
    data
  })
}

// @Tags Product
// @Summary 恢复Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "恢复Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"恢复成功"}"
// @Router /prod/restoreProduct [post]
export const restoreProduct = (data) => {
  return service({
    url: '/prod/restoreProduct',
    method: 'put',
    data
  })
}

// @Tags Product
// @Summary 删除Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /prod/deleteProduct [delete]
export const deleteProductByIds = (data) => {
  return service({
    url: '/prod/deleteProductByIds',
    method: 'delete',
    data
  })
}

// @Tags Product
// @Summary 更新Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Product true "更新Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /prod/updateProduct [put]
export const updateProduct = (data) => {
  return service({
    url: '/prod/updateProduct',
    method: 'put',
    data
  })
}

// @Tags Product
// @Summary 用id查询Product
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Product true "用id查询Product"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /prod/findProduct [get]
export const findProduct = (params) => {
  return service({
    url: '/prod/findProduct',
    method: 'get',
    params
  })
}

// @Tags Product
// @Summary 分页获取Product列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Product列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /prod/getProductList [get]
export const getProductList = (params) => {
  // console.log(params)
  return service({
    url: '/prod/getProductList',
    method: 'get',
    params
  })
}

// @Tags Product
// @Summary 分页获取DeletedProduct列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Product列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /prod/getDeletedProductList [get]
export const getDeletedProductList = (params) => {
  // console.log(params)
  return service({
    url: '/prod/getDeletedProductList',
    method: 'get',
    params
  })
}

export const exportProductExcel = (params) => {
  // console.log(params)
  service({
    url: '/prod/exportProductExcel',
    method: 'post',
    params,
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, 'product' + '.xlsx')
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
