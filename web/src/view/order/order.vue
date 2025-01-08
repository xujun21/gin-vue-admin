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
        <el-form-item label="订单日期">

          <el-date-picker v-model="searchInfo.startOrder_date" format="DD-MM-YYYY" type="date" placeholder="搜索条件（起）" />
          —
          <el-date-picker v-model="searchInfo.endOrder_date" format="DD-MM-YYYY" type="date" placeholder="搜索条件（止）" />

        </el-form-item>
        <el-form-item label="订单号">
          <el-input v-model="searchInfo.ID" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="发票号">
          <el-input v-model="searchInfo.invoice_no" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="付款方式">
          <el-input v-model="searchInfo.payment_method" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="PO号码">
          <el-input v-model="searchInfo.po_number" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="发票日期">

          <el-date-picker v-model="searchInfo.startInvoice_date" format="DD-MM-YYYY" type="date" placeholder="搜索条件（起）" />
          —
          <el-date-picker v-model="searchInfo.endInvoice_date" format="DD-MM-YYYY" type="date" placeholder="搜索条件（止）" />

        </el-form-item>
        <el-form-item label="付款方">
          <el-input v-model="searchInfo.bill_to" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="收货方">
          <el-input v-model="searchInfo.ship_to" placeholder="搜索条件" />

        </el-form-item>
        <!--
        <el-form-item label="货物数量">

          <el-input v-model.number="searchInfo.startQuantity" placeholder="搜索条件（起）" />
          —
          <el-input v-model.number="searchInfo.endQuantity" placeholder="搜索条件（止）" />

        </el-form-item>
        <el-form-item label="货款小计">

          <el-input v-model.number="searchInfo.startSubtotal" placeholder="搜索条件（起）" />
          —
          <el-input v-model.number="searchInfo.endSubtotal" placeholder="搜索条件（止）" />

        </el-form-item>
        <el-form-item label="税金">

          <el-input v-model.number="searchInfo.startVat" placeholder="搜索条件（起）" />
          —
          <el-input v-model.number="searchInfo.endVat" placeholder="搜索条件（止）" />

        </el-form-item>
        <el-form-item label="自动折扣">

          <el-input v-model.number="searchInfo.startDiscount" placeholder="搜索条件（起）" />
          —
          <el-input v-model.number="searchInfo.endDiscount" placeholder="搜索条件（止）" />

        </el-form-item>
        <el-form-item label="订单总金额">

          <el-input v-model.number="searchInfo.startOrder_total" placeholder="搜索条件（起）" />
          —
          <el-input v-model.number="searchInfo.endOrder_total" placeholder="搜索条件（止）" />

        </el-form-item>
        <el-form-item label="付款人地址">
          <el-input v-model="searchInfo.bill_to_address" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="付款人城市">
          <el-input v-model="searchInfo.bill_to_city" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="付款人电话">
          <el-input v-model="searchInfo.bill_to_phone" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="付款人邮编">
          <el-input v-model="searchInfo.bill_to_postcode" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="收货人">
          <el-input v-model="searchInfo.ship_to_contact_name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="收货地址">
          <el-input v-model="searchInfo.ship_to_address" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="收货城市">
          <el-input v-model="searchInfo.ship_to_city" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="收货人电话">
          <el-input v-model="searchInfo.ship_to_phone" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="收货地邮编">
          <el-input v-model="searchInfo.ship_to_postcode" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="库存锁定" prop="is_locked">
          <el-select v-model="searchInfo.is_locked" clearable placeholder="请选择">
            <el-option
              key="true"
              label="是"
              value="true"
            />
            <el-option
              key="false"
              label="否"
              value="false"
            />
          </el-select>
        </el-form-item>
        <el-form-item label="手工折扣">

          <el-input v-model.number="searchInfo.startHand_discount" placeholder="搜索条件（起）" />
          —
          <el-input v-model.number="searchInfo.endHand_discount" placeholder="搜索条件（止）" />

        </el-form-item>
        <el-form-item label="客户ID">

          <el-input v-model.number="searchInfo.startCustomer_id" placeholder="搜索条件（起）" />
          —
          <el-input v-model.number="searchInfo.endCustomer_id" placeholder="搜索条件（止）" />

        </el-form-item>
        -->
        <el-form-item label="客户公司名">
          <el-input v-model="searchInfo.customer_company_name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="客户联系人">
          <el-input v-model="searchInfo.customer_contact_name" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button class="excel-btn" size="small" type="primary" icon="download" @click="handleExcelExport()">导出</el-button>
          <el-button size="small" icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
      <div class="gva-btn-list">
        <el-button size="small" type="primary" icon="plus" hidden @click="openDialog">新增</el-button>
        <el-button size="small" type="primary" icon="plus" @click="toDetail()">新增</el-button>
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
        <el-table-column align="left" label="付款人" prop="bill_to_contact_name" width="120" />
        <el-table-column align="left" label="付款人地址" prop="bill_to_address" width="120" />
        <el-table-column align="left" label="付款人城市" prop="bill_to_city" width="120" />
        <el-table-column align="left" label="付款人电话" prop="bill_to_phone" width="120" />
        <el-table-column align="left" label="付款人邮编" prop="bill_to_postcode" width="120" />
        <el-table-column align="left" label="收货人" prop="ship_to_contact_name" width="120" />
        <el-table-column align="left" label="收货地址" prop="ship_to_address" width="120" />
        <el-table-column align="left" label="收货城市" prop="ship_to_city" width="120" />
        <el-table-column align="left" label="收货人电话" prop="ship_to_phone" width="120" />
        <el-table-column align="left" label="收货地邮编" prop="ship_to_postcode" width="120" />
        <el-table-column align="left" label="库存锁定" prop="is_locked" width="120">
          <template #default="scope">{{ formatBoolean(scope.row.is_locked) }}</template>
        </el-table-column>
        <el-table-column align="left" label="手工折扣" prop="hand_discount" width="120" />
        <el-table-column align="left" label="客户ID" prop="customer_id" width="120" />
        <el-table-column align="left" label="客户联系人" prop="customer_contact_name" width="120" />
        -->
        <el-table-column align="center" label="操作" width="200">
          <template #default="scope">
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="toDetail(scope.row)">修改订单</el-button>
            <el-button type="primary" link icon="delete" size="small" @click="deleteRow(scope.row)">删除订单</el-button>
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="toItem(scope.row)">增减商品</el-button>
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="downloadConfirmation(scope.row)">订单确认</el-button>
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="exportDeliveryNote(scope.row)">运单确认</el-button>
            <el-button type="primary" link icon="edit" size="small" class="table-button" @click="downloadInvoice(scope.row)">发票打印</el-button>
          </template>
        </el-table-column>
        <el-table-column align="right" label="订单号" prop="ID" width="80" />
        <el-table-column sortable align="left" label="订单日期" width="110" prop="order_date">
          <template #default="scope">{{ formatDateOnly(scope.row.order_date) }}</template>
        </el-table-column>
        <el-table-column align="left" label="客户公司名" prop="customer_company_name" width="200" />
        <el-table-column align="right" label="商品数量" prop="quantity" width="80" />
        <el-table-column align="right" label="货款小计" prop="subtotal" width="100" :formatter="formatCurrency" />
        <el-table-column align="right" label="税金" prop="vat" width="100" :formatter="formatCurrency" />
        <el-table-column align="right" label="运费" prop="delivery_fee" width="80" :formatter="formatCurrency" />
        <el-table-column align="right" label="自动折扣" prop="discount" width="80" :formatter="formatCurrency" />
        <el-table-column align="right" label="手工折扣" prop="hand_discount" width="80" :formatter="formatCurrency" />
        <el-table-column align="right" label="订单总金额" prop="order_total" width="100" :formatter="formatCurrency" />
        <el-table-column align="left" label="发票日期" width="100" prop="invoice_date">
          <template #default="scope">{{ formatDateOnly(scope.row.invoice_date) }}</template>
        </el-table-column>
        <el-table-column align="left" label="发票号" prop="invoice_no" width="100" />
        <el-table-column align="left" label="付款方式" prop="payment_method" width="120" />
        <el-table-column align="left" label="PO号码" prop="po_number" width="100" />
        <el-table-column align="left" label="付款方" prop="bill_to" width="120" />
        <el-table-column align="left" label="收货方" prop="ship_to" width="120" />
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
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单日期:" prop="order_date">
          <el-date-picker v-model="formData.order_date" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
        </el-form-item>
        <el-form-item label="发票号:" prop="invoice_no">
          <el-input v-model="formData.invoice_no" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付款方式:" prop="payment_method">
          <el-input v-model="formData.payment_method" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="PO号码:" prop="po_number">
          <el-input v-model="formData.po_number" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="发票日期:" prop="invoice_date">
          <el-date-picker v-model="formData.invoice_date" type="date" style="width:100%" placeholder="选择日期" :clearable="true" />
        </el-form-item>
        <el-form-item label="付款方:" prop="bill_to">
          <el-input v-model="formData.bill_to" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收货方:" prop="ship_to">
          <el-input v-model="formData.ship_to" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="货物数量:" prop="quantity">
          <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="货款小计:" prop="subtotal">
          <el-input-number v-model="formData.subtotal" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="税金:" prop="vat">
          <el-input-number v-model="formData.vat" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="自动折扣:" prop="discount">
          <el-input-number v-model="formData.discount" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="订单总金额:" prop="order_total">
          <el-input-number v-model="formData.order_total" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="付款人:" prop="bill_to_contact_name">
          <el-input v-model="formData.bill_to_contact_name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付款人地址:" prop="bill_to_address">
          <el-input v-model="formData.bill_to_address" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付款人城市:" prop="bill_to_city">
          <el-input v-model="formData.bill_to_city" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付款人电话:" prop="bill_to_phone">
          <el-input v-model="formData.bill_to_phone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="付款人邮编:" prop="bill_to_postcode">
          <el-input v-model="formData.bill_to_postcode" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收货人:" prop="ship_to_contact_name">
          <el-input v-model="formData.ship_to_contact_name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收货地址:" prop="ship_to_address">
          <el-input v-model="formData.ship_to_address" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收货城市:" prop="ship_to_city">
          <el-input v-model="formData.ship_to_city" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收货人电话:" prop="ship_to_phone">
          <el-input v-model="formData.ship_to_phone" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收货地邮编:" prop="ship_to_postcode">
          <el-input v-model="formData.ship_to_postcode" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="库存锁定:" prop="is_locked">
          <el-switch v-model="formData.is_locked" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>
        <el-form-item label="手工折扣:" prop="hand_discount">
          <el-input-number v-model="formData.hand_discount" style="width:100%" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="客户ID:" prop="customer_id">
          <el-input v-model.number="formData.customer_id" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="客户公司名:" prop="customer_company_name">
          <el-input v-model="formData.customer_company_name" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="客户联系人:" prop="customer_contact_name">
          <el-input v-model="formData.customer_contact_name" :clearable="true" placeholder="请输入" />
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
  name: 'Order'
}
</script>

<script setup>
import {
  createOrder,
  deleteOrder,
  deleteOrderByIds,
  updateOrder,
  findOrder,
  getOrderList,
  exportInvoiceExcel,
  exportConfirmExcel,
  exportDeliveryNoteExcel,
  exportOrderExcel
} from '@/api/order'

// 全量引入格式化工具 请按需保留
import { formatDateOnly, formatCurrency } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()

const handleExcelExport = () => {
  exportOrderExcel({ ...searchInfo.value })
}

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
  order_date: new Date(),
  invoice_no: '',
  payment_method: '',
  po_number: '',
  invoice_date: new Date(),
  bill_to: '',
  ship_to: '',
  quantity: 0,
  subtotal: 0,
  vat: 0,
  discount: 0,
  order_total: 0,
  bill_to_contact_name: '',
  bill_to_address: '',
  bill_to_city: '',
  bill_to_phone: '',
  bill_to_postcode: '',
  ship_to_contact_name: '',
  ship_to_address: '',
  ship_to_city: '',
  ship_to_phone: '',
  ship_to_postcode: '',
  is_locked: false,
  hand_discount: 0,
  customer_id: 0,
  customer_company_name: '',
  customer_contact_name: '',
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
  if (searchInfo.value.is_locked === '') {
    searchInfo.value.is_locked = null
  }
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
  const table = await getOrderList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    deleteOrderFunc(row)
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
  const res = await deleteOrderByIds({ ids })
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
const updateOrderFunc = async(row) => {
  const res = await findOrder({ ID: row.ID })
  type.value = 'update'
  if (res.code === 0) {
    formData.value = res.data.reord
    dialogFormVisible.value = true
  }
}

// 删除行
const deleteOrderFunc = async(row) => {
  const res = await deleteOrder({ ID: row.ID })
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

const toDetail = async(row) => {
  if (row) {
    router.push({
      name: 'orderDetail',
      params: {
        id: row.ID,
      },
    })
  } else {
    router.push({
      name: 'orderDetail',
      params: {
        id: 0,
      }
    })
  }
}

// 修改订单商品
const toItem = async(row) => {
  router.push({
    name: 'subOrder',
    params: {
      id: row.ID,
    },
  })
}

// 关闭弹窗
const closeDialog = () => {
  dialogFormVisible.value = false
  formData.value = {
    order_date: new Date(),
    invoice_no: '',
    payment_method: '',
    po_number: '',
    invoice_date: new Date(),
    bill_to: '',
    ship_to: '',
    quantity: 0,
    subtotal: 0,
    vat: 0,
    discount: 0,
    order_total: 0,
    bill_to_contact_name: '',
    bill_to_address: '',
    bill_to_city: '',
    bill_to_phone: '',
    bill_to_postcode: '',
    ship_to_contact_name: '',
    ship_to_address: '',
    ship_to_city: '',
    ship_to_phone: '',
    ship_to_postcode: '',
    is_locked: false,
    hand_discount: 0,
    customer_id: 0,
    customer_company_name: '',
    customer_contact_name: '',
  }
}
// 弹窗确定
const enterDialog = async() => {
     elFormRef.value?.validate(async(valid) => {
       if (!valid) return
       let res
       switch (type.value) {
         case 'create':
           res = await createOrder(formData.value)
           break
         case 'update':
           res = await updateOrder(formData.value)
           break
         default:
           res = await createOrder(formData.value)
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

// 下载
const downloadInvoice = (row) => {
  exportInvoiceExcel(row.ID)
}

const downloadConfirmation = (row) => {
  exportConfirmExcel(row.ID)
}

const exportDeliveryNote = (row) => {
  exportDeliveryNoteExcel(row.ID)
}
</script>

<style>
</style>
