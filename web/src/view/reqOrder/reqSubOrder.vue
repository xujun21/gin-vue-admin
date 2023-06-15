<template>
  <div class="gva-table-box">
    <el-table
      ref="multipleTable"
      style="width: 100%"
      tooltip-effect="dark"
      :data="tableData"
      row-key="ID"
      @selection-change="handleSelectionChange"
    >
      <el-table-column type="selection" />
      <el-table-column align="left" label="操作" width="100">
        <template #default="scope">
          <el-button type="primary" link icon="edit" size="small" class="table-button" @click="modifyInQty(scope.row)">修改数量</el-button>
          <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
        </template>
      </el-table-column>
      <el-table-column sortable align="left" label="商品编号" prop="productCode" width="120" />
      <el-table-column align="left" label="商品名" width="400">
        <template #default="scope">{{ scope.row.productNameCn }}<br>{{ scope.row.productNameEn }}</template>
      </el-table-column>
      <el-table-column align="left" label="SELFLIFE" prop="selfLife" width="100" />
      <el-table-column align="left" label="采购数量" prop="inQty" width="100" />
      <el-table-column align="left" label="Barcode" prop="barcode" width="100" />
      <el-table-column align="left" label="Barcode(Case)" prop="barcodeCase" width="120" />
      <el-table-column align="left" label="Carton Size" prop="cartonSize" width="120" />
      <el-table-column align="right" label="CBM" prop="cbm" width="80" :formatter="formatCbm" />
      <el-table-column align="right" label="Weight" prop="weight" width="80" :formatter="formatNumber" />
      <el-table-column align="right" label="采购价格" prop="inPrice" width="80" :formatter="formatNumber" />
      <el-table-column align="left" label="创建日期" width="180">
        <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
      </el-table-column>
    </el-table>
    <div class="gva-pagination">
      <el-pagination
        layout="total, sizes, prev, pager, next, jumper"
        :current-page="page"
        :page-size="pageSize"
        :page-sizes="[10, 30, 50, 100]"
        :total="total"
        @current-change="handleCurrentChange"
        @size-change="handleSizeChange"
      />
    </div>
    <div class="gva-btn-list">
      <el-button size="small" type="primary" icon="plus" @click="openProductDialog">添加商品</el-button>
      <el-popover v-model:visible="deleteVisible" placement="top" width="160">
        <p>确定要删除吗？</p>
        <div style="text-align: right; margin-top: 8px;">
          <el-button size="small" type="primary" link @click="deleteVisible = false">取消</el-button>
          <el-button size="small" type="primary" @click="onDelete">确定</el-button>
        </div>
        <template #reference>
          <el-button icon="delete" size="small" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
        </template>
      </el-popover>
      <el-button size="small" type="primary" @click="back">返回</el-button>
    </div>

    <el-dialog v-model="productDialogVisible" :before-close="productBack" title="选择商品（显示所有商品）" style="width: 60%">
      <el-form :inline="true" :model="psearchInfo" class="demo-form-inline">
        <el-form-item>
          <el-input v-model="psearchInfo.code" placeholder="商品编号" style="width:100px;" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="psearchInfo.product_name_cn" placeholder="商品中文名" style="width:150px;" />
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="psearchInfo.product_name_en" placeholder="商品英文名" style="width:150px;" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="psearchInfo.startStore" placeholder="最小库存" style="width:80px;" />
        ～
          <el-input v-model="psearchInfo.endStore" placeholder="最大库存" style="width:80px;" />
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="ponSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="ponReset">重置</el-button>
        </el-form-item>
      </el-form>
      <el-table
        ref="pmultipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="ptableData"
        row-key="ID"
        @selection-change="phandleSelectionChange"
        @sort-change="psortChange"
      >
        <el-table-column type="selection" />
        <el-table-column sortable align="left" label="编号" prop="code" width="100" />
        <el-table-column align="left" label="商品名" width="400">
          <template #default="scope">{{ scope.row.product_name_cn }}<br>{{ scope.row.product_name_en }}</template>
        </el-table-column>
        <el-table-column align="right" label="库存量" prop="store" sortable width="100" />
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          layout="total, sizes, prev, pager, next, jumper"
          :current-page="ppage"
          :page-size="ppageSize"
          :page-sizes="[1, 30, 50, 100]"
          :total="ptotal"
          @current-change="phandleCurrentChange"
          @size-change="phandleSizeChange"
        />
      </div>
      <el-button size="small" type="primary" @click="chooseProducts">确定</el-button>
      <el-button size="small" type="primary" @click="productBack">取消</el-button>
    </el-dialog>

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="修改采购数量">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="100px">
        <el-form-item label="采购数量:" prop="inQty">
          <el-input v-model.number="formData.inQty" :clearable="true" placeholder="请输入采购数量" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

  </div>
</template>

<script>
export default {
  name: 'ReqSubOrder'
}
</script>

<script setup>
import {
  deleteReqSubOrder,
  deleteReqSubOrderByIds,
  updateReqSubOrderInQty,
  // findReqSubOrder,
  getReqSubOrderList,
  addReqSubOrderByProductIds
} from '@/api/reqSubOrder'

import {
  getProductList,
  findProduct
} from '@/api/product'

import { useRoute, useRouter } from 'vue-router'

// 全量引入格式化工具 请按需保留
import { formatDate, formatDateOnly, formatCurrency,formatCbm,formatNumber } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { fix } from 'mathjs'

import { useUserStore } from '@/pinia/modules/user'
// import {
//   getUploadSubOrderList,
//   doExecImport
// } from '@/api/uploadSubOrder'

const isDisabled = ref(true)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  reqOrderId: 0,
  product_id: 0,
  product_code: '',
  product_name_cn: '',
  product_name_en: '',
  inQty: 0,
  // exp_date: new Date(),
  // package: '',
  // price: 0,
  // quantity: 0,
  // sub_total: 0,
  // vat: 0,
  // sub_vat: 0,
  // discount: 0,
  // discount_total: 0,
  // leftStore: 0
})

var oldQuantity = 0

// =========== 表格控制部分 ===========
const ppage = ref(1)
const ptotal = ref(0)
const ppageSize = ref(10)
const ptableData = ref([])
const psearchInfo = ref({})

const route = useRoute()
const router = useRouter()
const reqOrderId = route.params.id

// 重置
const ponReset = () => {
  psearchInfo.value = {}
  getpTableData()
}

// 排序
const psortChange = ({ prop, reqOrder }) => {
  psearchInfo.value.sort = prop
  psearchInfo.value.reqOrder = reqOrder
  getpTableData()
}

// 搜索
const ponSubmit = () => {
  ppage.value = 1
  ppageSize.value = 10
  getpTableData()
}

// 分页
const phandleSizeChange = (val) => {
  ppageSize.value = val
  getpTableData()
}

// 修改页面容量
const phandleCurrentChange = (val) => {
  ppage.value = val
  getpTableData()
}

// 查询
const getpTableData = async() => {
  // psearchInfo.value['startStore'] = 1 // 库存至少为1

  const table = await getProductList({ page: ppage.value, pageSize: ppageSize.value, ...psearchInfo.value })
  if (table.code === 0) {
    ptableData.value = table.data.list
    ptotal.value = table.data.total
    ppage.value = table.data.page
    ppageSize.value = table.data.pageSize
  }
}

getpTableData()

// 选中商品
const chooseProducts = async() => {
  const ids = []
  if (pmultipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要添加入订单的商品'
    })
    return
  }
  pmultipleSelection.value &&
        pmultipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await addReqSubOrderByProductIds({ 'reqOrderId': reqOrderId, ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '添加成功，请继续选择其他商品或者回到订单修改数量'
    })
    // this.pmultipleTable.clearSelection()
    productDialogVisible.value = false
    getTableData()
    getpTableData()
  }
}
// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const psetOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
psetOptions()

// 多选数据
const pmultipleSelection = ref([])

// 多选
const phandleSelectionChange = (val) => {
  pmultipleSelection.value = val
}

// 验证规则
const rule = reactive({
  order_id: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  product_id: [{
    required: true,
    message: '',
    trigger: ['input', 'blur'],
  }],
  quantity: [{
    required: true,
    message: '至少为1',
    trigger: ['input', 'blur'],
  }],

})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
// const onReset = () => {
//   searchInfo.value = {}
//   getTableData()
// }

// // 搜索
// const onSubmit = () => {
//   page.value = 1
//   pageSize.value = 10
//   getTableData()
// }

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  searchInfo.value['reqOrderId'] = reqOrderId
  // console.log(searchInfo)
  const table = await getReqSubOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async() => {
}

// 获取需要的字典 可能为空 按需保留
setOptions()

// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteReqSubOrderFunc(row)
  })
  getpTableData()
}

// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要删除的数据'
    })
    return
  }
  multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
  const res = await deleteReqSubOrderByIds({ ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === ids.length && page.value > 1) {
      page.value--
    }
    deleteVisible.value = false
    getTableData()
    getpTableData()
  }
}

// 行为控制标记（弹窗内部需要增还是改）
// const type = ref('')

// // 更新采购数量inQty
// const updateReqSubOrder = async(formData) => {
//   const res = await updateReqSubOrderInQty({ ID: formData.value.req_order_id, inQty: formData.value.inQty })
//   if (res.code === 0) {
//     ElMessage({
//       type: 'success',
//       message: '更新数量成功'
//     })
//     getTableData()
//     getpTableData()
//   }
// }

// 删除行
const deleteReqSubOrderFunc = async(row) => {
  const res = await deleteReqSubOrder({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
    getpTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

const productDialogVisible = ref(false)

// 打开修改采购数量弹窗
const modifyInQty = (row) => {
  // console.log(row)
  formData.value = {
    req_order_id: row.ID,
    inQty: row.inQty,
  }
  // console.log(formData)
  dialogFormVisible.value = true
}

// 打开商品选择的弹窗
const openProductDialog = () => {
  productDialogVisible.value = true
}
// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  isDisabled.value = true
  formData.value = {
    order_id: 0,
    product_id: 0,
    product_code: '',
    product_name_cn: '',
    product_name_en: '',
    exp_date: new Date(),
    package: '',
    price: 0,
    quantity: 0,
    vat: 0,
    sub_total: 0,
    discount: 0,
    discount_total: 0,
    inQty: 0,
  }
}
// 弹窗确定
const enterDialog = async() => {
  elFormRef.value?.validate(async(valid) => {
    if (!valid) return
    const res = await updateReqSubOrderInQty({ ID: formData.value.req_order_id, inQty: formData.value.inQty })
    // const res = await updateReqSubOrder(formData) // 更新采购数量
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '数量更改成功'
      })
      closeDialog()
      getTableData()
      getpTableData()
    }
  })
}

const productBack = () => {
  productDialogVisible.value = false
  isDisabled.value = true
}

const priceChanged = (price) => {
  if (price < 0) {
    ElMessage({
      type: 'error',
      message: '价格不能小于0'
    })
    isDisabled.value = true
    return
  } else {
    formData.value['sub_total'] = fix(price * formData.value['quantity'], 2)
    isDisabled.value = false
  }
}

const vatChanged = (vat) => {
  if (vat < 0) {
    ElMessage({
      type: 'error',
      message: '税不能小于0'
    })
    isDisabled.value = true
    return
  } else {
    formData.value['sub_vat'] = fix(vat * formData.value['quantity'], 2)
    isDisabled.value = false
  }
}

const discountChanged = (discount) => {
  if (discount < 0) {
    ElMessage({
      type: 'error',
      message: '折扣不能小于0'
    })
    isDisabled.value = true
    return
  } else {
    formData.value['discount_total'] = fix(discount * formData.value['quantity'], 2)
    isDisabled.value = false
  }
}

const qtyChanged = (qty) => {
  if (qty <= 0) {
    ElMessage({
      type: 'error',
      message: '数量至少为1'
    })
    isDisabled.value = true
    return
  } else {
    const totalQty = oldQuantity + formData.value['leftStore']
    if (qty > totalQty) {
      ElMessage({
        type: 'error',
        message: '数量不能大过剩余库存（' + totalQty + ')'
      })
      isDisabled.value = true
      return
    }
    formData.value.sub_total = fix(formData.value.price * qty, 2)
    formData.value.discount_total = fix(formData.value.discount * qty, 2)
    formData.value.sub_vat = fix(formData.value.vat * qty, 2)
    isDisabled.value = false
  }
}

// 返回按钮
const back = () => {
  router.go(-1)
}
</script>

<style>
</style>
