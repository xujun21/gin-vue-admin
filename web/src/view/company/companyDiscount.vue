<template>
  <div>
    <!--
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    -->
    <div class="gva-table-box">
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
      >
        <el-table-column type="selection" width="55" />
        <!--
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="客户ID" prop="companyId" width="120" />
        <el-table-column align="left" label="客户公司名" prop="companyName" width="120" />
        <el-table-column align="left" label="商品ID" prop="productId" width="120" />
        -->
        <el-table-column align="left" label="商品编码" prop="productCode" width="120" />
        <el-table-column align="left" label="商品中文名" prop="productNameCn" width="200" />
        <el-table-column align="left" label="商品英文名" prop="productNameEn" width="200" />
        <el-table-column align="left" label="商品价格" prop="price" width="120" :formatter="formatCurrency" />
        <el-table-column align="left" label="折扣金额" prop="discount" width="120" :formatter="formatCurrency" />
        <el-table-column align="left" label="操作">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateCompanyDiscountFunc(scope.row)">变更</el-button>
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
      </div>
    </div>

    <el-dialog v-model="productDialogVisible" :before-close="productBack" title="选择商品">
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
        <el-table-column type="selection" width="40" />
        <el-table-column sortable align="left" label="编号" prop="code" width="100" />
        <el-table-column align="left" label="商品中文名" prop="product_name_cn" width="200" />
        <el-table-column align="left" label="商品英文名" prop="product_name_en" width="200" />
        <el-table-column align="left" label="包装规格" prop="package" width="100" />
        <el-table-column align="left" label="保质期" width="100">
          <template #default="scope">{{ formatDateOnly(scope.row.exp_date) }}</template>
        </el-table-column>
        <el-table-column align="right" label="商品价格" prop="price" width="80" :formatter="formatCurrency" />
        <el-table-column align="right" label="商品税" prop="vat" width="80" :formatter="formatCurrency" />
        <el-table-column align="right" label="库存量" prop="store" width="80" />
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

    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="修改商品折扣">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="100px">
        <el-form-item label="商品编号:" prop="product_code">
          <el-input v-model="formData.productCode" readonly />
        </el-form-item>
        <el-form-item label="商品中文名:" prop="product_name_cn">
          <el-input v-model="formData.productNameCn" readonly />
        </el-form-item>
        <el-form-item label="商品英文名:" prop="product_name_en">
          <el-input v-model="formData.productNameEn" readonly />
        </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input v-model="formData.price" style="width:50%" readonly>
            <template #prepend>£</template>
          </el-input>
        </el-form-item>
        <el-form-item label="单件折扣:" prop="discount">
          <el-input-number v-model="formData.discount" style="width:100%" :precision="2" :step="0.1" @change="discountChanged"><template #prepend>£</template></el-input-number>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" :disabled="isDisabled" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <!-- <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="客户ID:" prop="companyId">
          <el-input v-model.number="formData.companyId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="客户公司名:" prop="companyName">
          <el-input v-model="formData.companyName" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品ID:" prop="productId">
          <el-input v-model.number="formData.productId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品编码:" prop="productCode">
          <el-input v-model="formData.productCode" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品中文名:" prop="productNameCn">
          <el-input v-model="formData.productNameCn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品英文名:" prop="productNameEn">
          <el-input v-model="formData.productNameEn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品价格:" prop="price">
          <el-input-number v-model="formData.price" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="折扣金额:" prop="discount">
          <el-input-number v-model="formData.discount" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button size="small" @click="closeDialog">取 消</el-button>
          <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog> -->
  </div>
</template>

<script>
export default {
  name: 'CompanyDiscount'
}
</script>

<script setup>
import {
  createCompanyDiscount,
  deleteCompanyDiscount,
  deleteCompanyDiscountByIds,
  updateCompanyDiscount,
  findCompanyDiscount,
  getCompanyDiscountList,
  addCompDiscountByProductIds
} from '@/api/companyDiscount'

// import {
//   findCompany
// } from '@/api/company'

import {
  getProductList
} from '@/api/product'

// 全量引入格式化工具 请按需保留
import { formatCurrency, formatDateOnly, formatCurrencyInput } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRoute, useRouter } from 'vue-router'

const route = useRoute()
const router = useRouter()
const companyId = route.params.id
// const company_name = route.params.companyName
const isDisabled = ref(true)

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  companyId: 0,
  companyName: '',
  productId: 0,
  productCode: '',
  productNameCn: '',
  productNameEn: '',
  price: 0,
  discount: 0,
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

// const company = ref({})

// 查询
const getTableData = async() => {
  searchInfo.value['companyId'] = companyId
  // const res = await findCompany({ ID: companyId })
  // if (res.code === 0) {
  //   company.value = res.data.recomp
  //   // for (const key in company.value) {
  //   //   console.log(key)
  //   //   if (key === 'company_name') {
  //   //     company_name = company.value[key]
  //   //   }
  //   // }
  //   // console.log(company_name)
  // }

  const table = await getCompanyDiscountList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteCompanyDiscountFunc(row)
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
  const res = await deleteCompanyDiscountByIds({ ids })
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
const updateCompanyDiscountFunc = async(row) => {
  const res = await findCompanyDiscount({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.recompDiscount
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteCompanyDiscountFunc = async(row) => {
  const res = await deleteCompanyDiscount({ ID: row.ID })
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
// const openDialog = () => {
//   type.value = 'create'
//   dialogFormVisible.value = true
// }

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    companyId: 0,
    companyName: '',
    productId: 0,
    productCode: '',
    productNameCn: '',
    productNameEn: '',
    price: 0,
    discount: 0,
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createCompanyDiscount(formData.value)
           break
         case 'update':
           res = await updateCompanyDiscount(formData.value)
           break
         default:
           res = await createCompanyDiscount(formData.value)
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

// =========== 表格控制部分 ===========
const ppage = ref(1)
const ptotal = ref(0)
const ppageSize = ref(10)
const ptableData = ref([])
const psearchInfo = ref({})

// 重置
const ponReset = () => {
  psearchInfo.value = {}
  getpTableData()
}

// 排序
const psortChange = ({ prop, order }) => {
  psearchInfo.value.sort = prop
  psearchInfo.value.order = order
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
  const res = await addCompDiscountByProductIds({ 'companyId': companyId, ids })
  if (res.code === 0) {
    ElMessage({
      type: 'success',
      message: '添加成功，请继续选择其他商品或者修改折扣金额'
    })
    // this.pmultipleTable.clearSelection()
    productDialogVisible.value = false
    getTableData()
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

const productDialogVisible = ref(false)

// 打开商品选择的弹窗
const openProductDialog = () => {
  productDialogVisible.value = true
}

const productBack = () => {
  productDialogVisible.value = false
  isDisabled.value = true
}

const discountChanged = (val) => {
  if (val < 0) {
    ElMessage({
      type: 'error',
      message: '折扣最少为0'
    })
    isDisabled.value = true
    return
  } else if (val > formData.value['price']) {
    ElMessage({
      type: 'error',
      message: '折扣不能大于价格'
    })
    isDisabled.value = true
    return
  }
  isDisabled.value = false
}

</script>

<style>
</style>
