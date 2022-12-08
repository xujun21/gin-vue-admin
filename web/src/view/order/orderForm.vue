<template>
  <div>
    <div class="gva-form-box1">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="120px">
        <el-row>
          <el-col :span="10">
            <el-form-item label="订单编号:" prop="ID">
              <el-input v-model="formData.ID" readonly placeholder="Auto Generated Order Number" />
            </el-form-item>
          </el-col>
          <el-col :span="14">
            <el-form-item label="订单日期:" prop="order_date">
              <el-date-picker v-model="formData.order_date" type="date" placeholder="选择日期" :clearable="true" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label="发票号:" prop="invoice_no">
              <el-input v-model="formData.invoice_no" :clearable="true" placeholder="Invoice No" />
            </el-form-item>
          </el-col>
          <el-col :span="14">
            <el-form-item label="发票日期:" prop="invoice_date">
              <el-date-picker v-model="formData.invoice_date" type="date" placeholder="选择日期" :clearable="true" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="14">
            <el-form-item label="客户公司:" prop="customer_company_name">
              <el-input v-model="formData.customer_company_name" readonly placeholder="Customer Company Name"><template #append>
                <el-button class="el-icon-search" @click="chooseCompany" />
              </template></el-input>
            </el-form-item>
          </el-col>
          <el-col :span="10" hidden>
            <el-form-item label="客户联系人:" prop="customer_contact_name">
              <el-input v-model="formData.customer_contact_name" :clearable="true" placeholder="Customer Contact Name" />
            </el-form-item>
            <el-form-item label="客户ID:" prop="customer_id">
              <el-input v-model="formData.customer_id" readonly="true" type="text" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label="付款方式:" prop="payment_method">
              <el-input v-model="formData.payment_method" :clearable="true" placeholder="Payment Method" />
            </el-form-item>
          </el-col>
          <el-col :span="14">
            <el-form-item label="PO号码:" prop="po_number">
              <el-input v-model="formData.po_number" :clearable="true" placeholder="PONumber" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label="付款方:" prop="bill_to">
              <el-input v-model="formData.bill_to" :clearable="true" placeholder="Bill To" />
            </el-form-item>
          </el-col>
          <el-col :span="14">
            <el-form-item label="付款人:" prop="bill_to_contact_name">
              <el-input v-model="formData.bill_to_contact_name" :clearable="true" placeholder="Bill To Contact Name" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col>
            <el-form-item label="付款人地址:" prop="bill_to_address">
              <el-input v-model="formData.bill_to_address" :clearable="true" placeholder="Bill To Address" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item label="付款人电话:" prop="bill_to_phone">
              <el-input v-model="formData.bill_to_phone" :clearable="true" placeholder="Bill To Phone" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="付款人城市:" prop="bill_to_city">
              <el-input v-model="formData.bill_to_city" :clearable="true" placeholder="Bill To City" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="付款人邮编:" prop="bill_to_postcode">
              <el-input v-model="formData.bill_to_postcode" :clearable="true" placeholder="Bill To Postcode" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="10">
            <el-form-item label="收货方:" prop="ship_to">
              <el-input v-model="formData.ship_to" :clearable="true" placeholder="Ship To" />
            </el-form-item>
          </el-col>
          <el-col :span="10">
            <el-form-item label="收货人:" prop="ship_to_contact_name">
              <el-input v-model="formData.ship_to_contact_name" :clearable="true" placeholder="Ship To Contact Name" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col>
            <el-form-item label="收货地址:" prop="ship_to_address">
              <el-input v-model="formData.ship_to_address" :clearable="true" placeholder="Ship To Address" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item label="收货人电话:" prop="ship_to_phone">
              <el-input v-model="formData.ship_to_phone" :clearable="true" placeholder="Ship To Phone" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="收货城市:" prop="ship_to_city">
              <el-input v-model="formData.ship_to_city" :clearable="true" placeholder="Ship To City" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="收货地邮编:" prop="ship_to_postcode">
              <el-input v-model="formData.ship_to_postcode" :clearable="true" placeholder="Ship To Postcode" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item label="货物数量:" prop="quantity">
              <el-input v-model.number="formData.quantity" :clearable="true" placeholder="自动计算" readonly />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="货款小计:" prop="subtotal">
              <el-input-number v-model="formData.subtotal" :precision="2" :clearable="true" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="税金:" prop="vat">
              <el-input-number v-model="formData.vat" :precision="2" :clearable="true" />
            </el-form-item>
          </el-col>
        </el-row>
        <el-row>
          <el-col :span="8">
            <el-form-item label="自动折扣:" prop="discount">
              <el-input-number v-model="formData.discount" :precision="2" :clearable="true" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="手工折扣:" prop="hand_discount">
              <el-input-number v-model="formData.hand_discount" :precision="2" :clearable="true" />
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="订单总金额:" prop="order_total">
              <el-input-number v-model="formData.order_total" :precision="2" :clearable="true" />
            </el-form-item>
          </el-col>
        </el-row>

        <el-form-item label="库存锁定:" prop="is_locked">
          <el-switch v-model="formData.is_locked" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable />
        </el-form-item>

        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>

  // eslint-disable-next-line vue/no-multiple-template-root
  <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="请选择下单客户">
    <div class="gva-table-box">
      <el-table ref="singleTable" highlight-current-row :data="tableData" style="width: 100%" tooltip-effect="dark" row-key="ID" @current-change="selectedChange">
        <el-table-column align="left" label="客户ID" prop="ID" width="120" />
        <el-table-column align="left" label="客户公司名" prop="company_name" />
      </el-table>
      <div class="gva-pagination">
        <el-pagination
          :current-page="page"
          :page-size="pageSize"
          :page-sizes="[10, 30, 50, 100]"
          :total="total"
          layout="total, sizes, prev, pager, next, jumper"
          @current-change="handleCurrentChange"
          @size-change="handleSizeChange"
        />
      </div>
    </div>

    <template #footer>
      <div class="dialog-footer">
        <el-button size="small" @click="closeDialog">取 消</el-button>
        <el-button size="small" type="primary" @click="enterDialog">确 定</el-button>
      </div>
    </template>
  </el-dialog>
</template>

<script>
export default {
  name: 'Order'
}
</script>

<script setup>
import {
  createOrder,
  updateOrder,
  findOrder
} from '@/api/order'

import {
  getCompanyList
} from '@/api/company'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref()

const companySelected = ref({ id: 0, company_name: '', contact_name: '' })
// const company = ref({CreatedAt: '', CreatedBy: 0, DeletedBy: 0, ID: 0, UpdatedAt: '', UpdatedBy: '',
//                 address: '', city: '', company_name: '', contact_name: '', note: '', phone: '', postcode: '', sage: '' })
// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getCompanyData()
}

const handleCurrentChange = (val) => {
  page.value = val
  getCompanyData()
}

const selectedChange = (row) => {
  for (const key in row) {
    if (key === 'ID') {
      companySelected.value['id'] = row[key]
    } else if (key === 'company_name') {
      companySelected.value['company_name'] = row[key]
    } else if (key === 'contact_name') {
      companySelected.value['contact_name'] = row[key]
    }
  }
}

const dialogFormVisible = ref(false)
const type = ref('')
const formData = ref({
  ID: 0,
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

const getCompanyData = async() => {
  const table = await getCompanyList({
    page: page.value,
    pageSize: pageSize.value,
    ...searchInfo.value,
  })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getCompanyData()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.params.id) {
    const res = await findOrder({ ID: route.params.id })
    if (res.code === 0) {
      formData.value = res.data.reord
      console.log(res.data.reord)
      type.value = 'update'
    }
  } else {
    type.value = 'create'
    console.log(type.value)
  }
}

init()
// 保存按钮
const save = async() => {
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
        }
      })
}

// 返回按钮
const back = () => {
  router.go(-1)
}

const closeDialog = () => {
  dialogFormVisible.value = false
}

const chooseCompany = () => {
  dialogFormVisible.value = true
}

const enterDialog = () => {
  if (companySelected.value['id'] === 0) {
    // 提示需要选择一行
  } else {
    console.log(companySelected)
    formData.value['customer_company_name'] = companySelected.value['company_name']
    formData.value['customer_contact_name'] = companySelected.value['contact_name']
    formData.value['customer_id'] = companySelected.value['id']
    closeDialog()
  }
}
</script>

<style>

</style>
