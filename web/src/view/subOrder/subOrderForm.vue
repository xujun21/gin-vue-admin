<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单号:" prop="order_id">
          <el-input v-model.number="formData.order_id" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品ID:" prop="product_id">
          <el-input v-model.number="formData.product_id" :clearable="false" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品编号:" prop="product_code">
          <el-input v-model="formData.product_code" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品中文名:" prop="product_name_cn">
          <el-input v-model="formData.product_name_cn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品英文名:" prop="product_name_en">
          <el-input v-model="formData.product_name_en" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="过期时间:" prop="exp_date">
          <el-date-picker v-model="formData.exp_date" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="包装规格:" prop="package">
          <el-input v-model="formData.package" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="价格:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="数量:" prop="quantity">
          <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="税:" prop="vat">
          <el-input-number v-model="formData.vat" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="货款小计:" prop="sub_total">
          <el-input-number v-model="formData.sub_total" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="单件折扣:" prop="discount">
          <el-input-number v-model="formData.discount" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="折扣小计:" prop="discount_total">
          <el-input-number v-model="formData.discount_total" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item>
          <el-button size="small" type="primary" @click="save">保存</el-button>
          <el-button size="small" type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
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
  updateSubOrder,
  findSubOrder
} from '@/api/subOrder'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findSubOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.resubOrd
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
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
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
