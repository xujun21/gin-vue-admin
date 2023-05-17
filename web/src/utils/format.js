import { formatTimeToStr } from '@/utils/date'
import { getDict } from '@/utils/dictionary'

export const formatBoolean = (bool) => {
  if (bool !== null) {
    return bool ? '是' : '否'
  } else {
    return ''
  }
}
export const formatDate = (time) => {
  if (time !== null && time !== '') {
    var date = new Date(time)
    return formatTimeToStr(date, 'dd-MM-yyyy hh:mm:ss')
  } else {
    return ''
  }
}

export const filterDict = (value, options) => {
  const rowLabel = options && options.filter(item => item.value === value)
  return rowLabel && rowLabel[0] && rowLabel[0].label
}

export const getDictFunc = async(type) => {
  const dicts = await getDict(type)
  return dicts
}

export const formatDateOnly = (time) => {
  if (time !== null && time !== '') {
    var date = new Date(time)
    return formatTimeToStr(date, 'dd-MM-yyyy')
  } else {
    return ''
  }
}

// el-column展示时数据格式化
export const formatCurrency = (row, column) => {
  var data = row[column.property] ? row[column.property] : 0
  return '£ ' + parseFloat(data).toFixed(2)
}

// el-input展示时数据格式化
export const formatCurrencyInput = (value) => {
  console.log(parseFloat(val).toFixed(2))
  return parseFloat(val).toFixed(2)
}

export const formatCbm = (row, column)=> {
  var data = row[column.property] ? row[column.property] : 0
  return parseFloat(data).toFixed(4)
}

export const formatNumber = (row, column)=> {
  var data = row[column.property] ? row[column.property] : 0
  return parseFloat(data).toFixed(2)
}

