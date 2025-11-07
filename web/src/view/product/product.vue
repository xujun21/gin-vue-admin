<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
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
        <el-table-column align="left" label="商品图片" width="100">
          <template #default="scope">
            <CustomPic pic-type="file" :pic-src="scope.row.image?scope.row.image:'uploads/file/default.png'" />
          </template>
        </el-table-column>
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
        <!-- 优化后的图片显示区域，删除按钮为右上角叉号 -->
        <el-form-item label="商品图片:" prop="photo">
          <!-- 图片存在时：CustomPic控件 + 右上角删除叉号 -->
          <div v-if="formData.image" class="image-with-delete">
            <CustomPic pic-type="file" :pic-src="formData.image" />
            <el-button
              icon="close"
              size="mini"
              class="delete-icon"
              @click="formData.image = ''"
              title="删除图片"
            />
          </div>
          <!-- 图片不存在时：仅显示上传按钮 -->
          <el-upload
            v-else
            class="upload-demo"
            :action="`${path}/fileUploadAndDownload/upload`"
            :headers="{ 'x-token': userStore.token }"
            :on-success="(response, file) => { formData.image = response.data.file.url }"
            :limit="1"
            :file-list="[]"
          >
            <el-button size="small" type="primary">点击上传</el-button>
            <div slot="tip" class="el-upload__tip" style="margin-top: 5px;">支持JPG、PNG格式，建议尺寸比例1:1</div>
          </el-upload>
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

import { formatCurrency, formatDate, formatDateOnly } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

import { useUserStore } from '@/pinia/modules/user'
import CustomPic from '@/components/customPic/index.vue'

const userStore = useUserStore()
const path = ref(import.meta.env.VITE_BASE_API)

// 导出Excel
const handleExcelExport = (param) => {
  searchInfo.value.withPrice = param
  exportProductExcel({ ...searchInfo.value })
}

// 表单数据
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
  image: ''
})

// 验证规则
const rule = reactive({})

const elFormRef = ref()

// 表格控制
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 排序处理
const sortChange = ({ prop, order }) => {
  searchInfo.value.sort = prop
  searchInfo.value.order = order
  getTableData()
}

// 重置搜索
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索提交
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  getTableData()
}

// 分页大小变更
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 页码变更
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 获取表格数据
const getTableData = async() => {
  const table = await getProductList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

// 初始化加载表格数据
getTableData()

// 多选数据
const multipleSelection = ref([])
const handleSelectionChange = (val) => {
  multipleSelection.value = val
}

// 单行删除
const deleteRow = (row) => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(() => {
    deleteProductFunc(row)
  })
}

// 批量删除控制
const deleteVisible = ref(false)
const onDelete = async() => {
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({ type: 'warning', message: '请选择要删除的数据' })
    return
  }
  multipleSelection.value.forEach(item => ids.push(item.ID))
  const res = await deleteProductByIds({ ids })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '删除成功' })
    if (tableData.value.length === ids.length && page.value > 1) page.value--
    deleteVisible.value = false
    getTableData()
  }
}

// 弹窗操作类型（新增/编辑）
const type = ref('')

// 编辑商品
const updateProductFunc = async(row) => {
  const res = await findProduct({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reprod
    dialogFormVisible.value = true
  }
}

// 删除商品接口调用
const deleteProductFunc = async(row) => {
  const res = await deleteProduct({ ID: row.ID })
  if (res.code === 0) {
    ElMessage({ type: 'success', message: '删除成功' })
    if (tableData.value.length === 1 && page.value > 1) page.value--
    getTableData()
  }
}

// 弹窗显示控制
const dialogFormVisible = ref(false)

// 打开弹窗（新增）
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
    image: ''
  }
}

// 弹窗确认提交
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
      ElMessage({ type: 'success', message: '创建/更改成功' })
      closeDialog()
      getTableData()
    }
  })
}
</script>

<style scoped>
/* 图片与删除按钮容器样式 */
.image-with-delete {
  position: relative;
  display: inline-block; /* 确保容器大小与图片一致 */
  margin-bottom: 10px;
}

/* 右上角删除叉号样式 */
.delete-icon {
  position: absolute;
  top: -8px;
  right: -8px;
  width: 20px;
  height: 20px;
  padding: 0;
  background-color: rgba(255, 77, 79, 0.9);
  color: white;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  cursor: pointer;
  transition: background-color 0.2s;
}

.delete-icon:hover {
  background-color: #ff4d4f; /* hover时加深颜色 */
  color: white;
}

/* 上传提示文字样式 */
.el-upload__tip {
  font-size: 12px;
  color: #666;
}
</style>