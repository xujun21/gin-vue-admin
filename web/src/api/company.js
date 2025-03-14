import service from '@/utils/request'

// @Tags Company
// @Summary 创建Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Company true "创建Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comp/createCompany [post]
export const createCompany = (data) => {
  return service({
    url: '/comp/createCompany',
    method: 'post',
    data
  })
}

// @Tags Company
// @Summary 删除Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Company true "删除Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /comp/deleteCompany [delete]
export const deleteCompany = (data) => {
  return service({
    url: '/comp/deleteCompany',
    method: 'delete',
    data
  })
}

// @Tags Company
// @Summary 删除Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /comp/deleteCompany [delete]
export const deleteCompanyByIds = (data) => {
  return service({
    url: '/comp/deleteCompanyByIds',
    method: 'delete',
    data
  })
}

// @Tags Company
// @Summary 更新Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body model.Company true "更新Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /comp/updateCompany [put]
export const updateCompany = (data) => {
  return service({
    url: '/comp/updateCompany',
    method: 'put',
    data
  })
}

// @Tags Company
// @Summary 用id查询Company
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query model.Company true "用id查询Company"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /comp/findCompany [get]
export const findCompany = (params) => {
  return service({
    url: '/comp/findCompany',
    method: 'get',
    params
  })
}

// @Tags Company
// @Summary 分页获取Company列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query request.PageInfo true "分页获取Company列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /comp/getCompanyList [get]
export const getCompanyList = (params) => {
  return service({
    url: '/comp/getCompanyList',
    method: 'get',
    params
  })
}

export const exportCompanyExcel = (params) => {
  service({
    url: '/comp/exportCompanyExcel',
    method: 'post',
    params,
    responseType: 'blob'
  }).then((res) => {
    handleFileError(res, 'CompanyList.xlsx')
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
