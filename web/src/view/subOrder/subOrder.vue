<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="商品编号">
         <el-input v-model="searchInfo.product_code" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="商品中文名">
         <el-input v-model="searchInfo.product_name_cn" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="商品英文名">
         <el-input v-model="searchInfo.product_name_en" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="过期时间">
            
            <el-date-picker v-model="searchInfo.startExp_date" type="datetime" placeholder="搜索条件（起）"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endExp_date" type="datetime" placeholder="搜索条件（止）"></el-date-picker>

        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
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
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column sortable align="left" label="订单号" prop="order_id" width="120" />
        <el-table-column sortable align="left" label="商品ID" prop="product_id" width="120" />
        <el-table-column sortable align="left" label="商品编号" prop="product_code" width="120" />
        <el-table-column align="left" label="商品中文名" prop="product_name_cn" width="120" />
        <el-table-column align="left" label="商品英文名" prop="product_name_en" width="120" />
         <el-table-column align="left" label="过期时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.exp_date) }}</template>
         </el-table-column>
        <el-table-column align="left" label="包装规格" prop="package" width="120" />
        <el-table-column align="left" label="价格" prop="price" width="120" />
        <el-table-column align="left" label="数量" prop="quantity" width="120" />
        <el-table-column align="left" label="税" prop="vat" width="120" />
        <el-table-column align="left" label="货款小计" prop="sub_total" width="120" />
        <el-table-column align="left" label="单件折扣" prop="discount" width="120" />
        <el-table-column align="left" label="折扣小计" prop="discount_total" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="updateSubOrderFunc(scope.row)">变更</el-button>
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
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="订单号:"  prop="order_id" >
          <el-input v-model.number="formData.order_id" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品ID:"  prop="product_id" >
          <el-input v-model.number="formData.product_id" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品编号:"  prop="product_code" >
          <el-input v-model="formData.product_code" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品中文名:"  prop="product_name_cn" >
          <el-input v-model="formData.product_name_cn" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品英文名:"  prop="product_name_en" >
          <el-input v-model="formData.product_name_en" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="过期时间:"  prop="exp_date" >
          <el-date-picker v-model="formData.exp_date" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="包装规格:"  prop="package" >
          <el-input v-model="formData.package" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="价格:"  prop="price" >
          <el-input-number v-model="formData.price"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="数量:"  prop="quantity" >
          <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="税:"  prop="vat" >
          <el-input-number v-model="formData.vat"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="货款小计:"  prop="sub_total" >
          <el-input-number v-model="formData.sub_total"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="单件折扣:"  prop="discount" >
          <el-input-number v-model="formData.discount"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="折扣小计:"  prop="discount_total" >
          <el-input-number v-model="formData.discount_total"  style="width:100%" :precision="2" :clearable="true"  />
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
  name: 'SubOrder'
}
</script>

<script setup>
import {
  createSubOrder,
  deleteSubOrder,
  deleteSubOrderByIds,
  updateSubOrder,
  findSubOrder,
  getSubOrderList
} from '@/api/subOrder'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
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
        })

// 验证规则
const rule = reactive({
               order_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               product_id : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               product_code : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               product_name_cn : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
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
  const table = await getSubOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
const setOptions = async () =>{
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
            deleteSubOrderFunc(row)
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
      const res = await deleteSubOrderByIds({ ids })
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
const updateSubOrderFunc = async(row) => {
    const res = await findSubOrder({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.resubOrd
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteSubOrderFunc = async (row) => {
    const res = await deleteSubOrder({ ID: row.ID })
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
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createSubOrder(formData.value)
                  break
                case 'update':
                  res = await updateSubOrder(formData.value)
                  break
                default:
                  res = await createSubOrder(formData.value)
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
