<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <!--
        <el-form-item label="创建时间">
          <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间" />
          —
          <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间" />
        </el-form-item>
        -->
        <el-form-item label="公司名称">
          <el-input v-model="searchInfo.customer_company_name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="联系人">
          <el-input v-model="searchInfo.customer_contact_name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="联系电话">
          <el-input v-model="searchInfo.phone" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="地址">
          <el-input v-model="searchInfo.address" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="邮编">
          <el-input v-model="searchInfo.postcode" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="备注">
          <el-input v-model="searchInfo.note" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="SAGE">
          <el-input v-model="searchInfo.sage" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button size="small" type="primary" icon="download" @click="onDownload">下载</el-button>
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
      >
        <el-table-column type="selection" width="55" />
        <!--
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="备注" prop="note" width="120" />
        -->
        <el-table-column align="left" label="公司名称 / 地址" width="400">
          <template #default="scope">{{ scope.row.company_name }} <br> {{ scope.row.address }}</template>
        </el-table-column>
        <el-table-column align="left" label="联系人 / 电话" width="140">
          <template #default="scope">{{ scope.row.contact_name }}  <br> {{ scope.row.phone }}</template>
        </el-table-column>
        <el-table-column align="left" label="城市" prop="city" width="120" />
        <el-table-column align="left" label="邮编" prop="postcode" width="120" />
        <el-table-column align="left" label="SAGE" prop="sage" width="120" />
        <el-table-column align="left" label="按钮组">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateCompanyFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除</el-button>
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="setDiscount(scope.row)">设置优惠</el-button>
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
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" :title="dialogTitle">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="公司名称:" prop="company_name">
          <el-input v-model="formData.company_name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系人:" prop="contact_name">
          <el-input v-model="formData.contact_name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="联系电话:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="城市:" prop="city">
          <el-input v-model="formData.city" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="邮编:" prop="postcode">
          <el-input v-model="formData.postcode" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:" prop="note">
          <el-input v-model="formData.note" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="SAGE:" prop="sage">
          <el-input v-model="formData.sage" :clearable="true" placeholder="请输入" />
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
  name: 'Company'
}
</script>

<script setup>
import {
  createCompany,
  deleteCompany,
  deleteCompanyByIds,
  updateCompany,
  findCompany,
  getCompanyList,
  exportCompanyExcel
} from '@/api/company'

// 全量引入格式化工具 请按需保留
// import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  company_name: '',
  contact_name: '',
  phone: '',
  address: '',
  city: '',
  postcode: '',
  note: '',
  sage: '',
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
const dialogTitle = ref('')

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
  const table = await getCompanyList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

const onDownload = () => {
  exportCompanyExcel({ ...searchInfo.value })
}

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
    deleteCompanyFunc(row)
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
  const res = await deleteCompanyByIds({ ids })
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
const updateCompanyFunc = async(row) => {
  const res = await findCompany({ ID: row.ID })
  type.value = 'update'
  dialogTitle.value = '修改客户信息：' + row.ID
  if (res.code === 0) {
    formData.value = res.data.recomp
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteCompanyFunc = async(row) => {
  const res = await deleteCompany({ ID: row.ID })
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
  dialogTitle.value = '新建客户'
  dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    company_name: '',
    contact_name: '',
    phone: '',
    address: '',
    city: '',
    postcode: '',
    note: '',
    sage: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createCompany(formData.value)
           break
         case 'update':
           res = await updateCompany(formData.value)
           break
         default:
           res = await createCompany(formData.value)
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

const router = useRouter()

const setDiscount = async(row) => {
  router.push({
    name: 'discountAdmin',
    params: {
      id: row.ID,
      // companyName: row.company_name
    },
  })
}

</script>

<style>
</style>
