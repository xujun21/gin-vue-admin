<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <!--
        <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间" />
        </el-form-item>-->
        <el-form-item>
          <el-input v-model="searchInfo.code" placeholder="商品编号" />
        </el-form-item>
        <el-form-item>
          <el-input v-model="searchInfo.product_name_cn" placeholder="商品中文名" style="width:300px;" />
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="searchInfo.product_name_en" placeholder="商品英文名" style="width:300px;" />
        </el-form-item>
        <el-form-item label="">
          <el-input v-model="searchInfo.package" placeholder="商品包装规格" />
        </el-form-item>
        <el-form-item label="商品保质期">

          <el-date-picker v-model="searchInfo.startExp_date" format="DD-MM-YYYY" type="date" placeholder="搜索条件（起）" style="width:150px;" />
          —
          <el-date-picker v-model="searchInfo.endExp_date" format="DD-MM-YYYY" type="date" placeholder="搜索条件（止）" style="width:150px;" />

        </el-form-item>
        <el-form-item label="商品价格">

          <el-input v-model.number="searchInfo.startPrice" placeholder="起" style="width:100px;" />
          —
          <el-input v-model.number="searchInfo.endPrice" placeholder="止" style="width:100px;" />

        </el-form-item>
        <el-form-item label="商品税">

          <el-input v-model.number="searchInfo.startVat" placeholder="起" style="width:100px;" />
          —
          <el-input v-model.number="searchInfo.endVat" placeholder="止" style="width:100px;" />

        </el-form-item>
        <el-form-item label="商品库存量">

          <el-input v-model.number="searchInfo.startStore" placeholder="起" style="width:100px;" />
          —
          <el-input v-model.number="searchInfo.endStore" placeholder="止" style="width:100px;" />

        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button class="excel-btn" size="small" type="primary" icon="download" @click="handleExcelExport(1)">导出(含价格)</el-button>
          <el-button class="excel-btn" size="small" type="primary" icon="download" @click="handleExcelExport(0)">导出(无价格)</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" @click="openDialog">新增</el-button>
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
      </div>
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
      >
        <el-table-column type="selection" width="40" />
        <el-table-column sortable align="left" label="编号" prop="code" width="100" />
        <el-table-column align="left" label="商品名" prop="product_name_cn" width="400">
          <template #default="scope">{{ scope.row.product_name_cn }} {{ scope.row.product_name_en }}</template>
        </el-table-column>
        <el-table-column align="left" label="包装规格" prop="package" width="100" />
        <el-table-column align="left" label="保质期" width="100">
          <template #default="scope">{{ formatDateOnly(scope.row.exp_date) }}</template>
        </el-table-column>
        <el-table-column align="right" label="商品价格" prop="price" width="80" :formatter="formatCurrency" />
        <el-table-column align="right" label="商品税" prop="vat" width="80" :formatter="formatCurrency" />
        <el-table-column align="right" label="库存量" prop="store" width="80" />
        <el-table-column align="left" label="日期" width="180">
          <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateProductFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
          </template>
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
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="120px">
        <el-form-item label="商品编号:" prop="code">
          <el-input v-model="formData.code" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品中文名:" prop="product_name_cn">
          <el-input v-model="formData.product_name_cn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品英文名:" prop="product_name_en">
          <el-input v-model="formData.product_name_en" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品包装规格:" prop="package">
          <el-input v-model="formData.package" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品保质期:" prop="exp_date">
          <el-date-picker v-model="formData.exp_date" format="DD-MM-YYYY" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
        </el-form-item>
        <el-form-item label="商品价格:" prop="price">
          <el-input-number v-model="formData.price" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="商品税:" prop="vat">
          <el-input-number v-model="formData.vat" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="商品库存量:" prop="store">
          <el-input v-model.number="formData.store" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品条码:" prop="barcode">
          <el-input v-model.number="formData.barcode" :clearable="true" placeholder="请输入" />
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
  name: 'Product'
}
</script>

<script setup>
import {
  createProduct,
  deleteProduct,
  deleteProductByIds,
  updateProduct,
  findProduct,
  getProductList,
  exportProductExcel
} from '@/api/product'

// 全量引入格式化工具 请按需保留
import { formatCurrency, formatDate, formatDateOnly, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

const handleExcelExport = (param) => {
  searchInfo.value.withPrice = param
  exportProductExcel({ ...searchInfo.value })
}

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  code: '',
  product_name_cn: '',
  product_name_en: '',
  package: '',
  exp_date: new Date(),
  price: 0,
  vat: 0,
  store: 0,
  barcode: '',
})

// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

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
  const table = await getProductList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteProductFunc(row)
  })
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
  const res = await deleteProductByIds({ ids })
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
  }
}

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateProductFunc = async(row) => {
  const res = await findProduct({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reprod
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteProductFunc = async(row) => {
  const res = await deleteProduct({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '删除成功'
    })
    if (tableData.value.length === 1 && page.value > 1) {
      page.value--
    }
    getTableData()
  }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
  type.value = 'create'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    code: '',
    product_name_cn: '',
    product_name_en: '',
    package: '',
    exp_date: new Date(),
    price: 0,
    vat: 0,
    store: 0,
    barcode: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createProduct(formData.value)
           break
         case 'update':
           res = await updateProduct(formData.value)
           break
         default:
           res = await createProduct(formData.value)
           break
       }
       if (res.code === 0) {
         ElMessage({
           type: 'success',
           message: '创建/更改成功'
         })
         closeDialog()
         getTableData()
       }
     })
}
</script>

<style>
</style>
