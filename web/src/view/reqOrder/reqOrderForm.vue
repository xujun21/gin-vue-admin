<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="订单日期:" prop="order_date">
          <el-date-picker v-model="formData.order_date" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="采购数量:" prop="quantity">
          <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="T.CBM:" prop="totalCbm">
          <el-input-number v-model="formData.totalCbm" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="总价:" prop="totalPrice">
          <el-input-number v-model="formData.totalPrice" :precision="2" :clearable="true"></el-input-number>
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
  name: 'RequireOrder'
}
</script>

<script setup>
import {
  createRequireOrder,
  updateRequireOrder,
  findRequireOrder
} from '@/api/reqOrder'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            order_date: new Date(),
            quantity: 0,
            totalCbm: 0,
            totalPrice: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findRequireOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reReqOrd
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
               res = await createRequireOrder(formData.value)
               break
             case 'update':
               res = await updateRequireOrder(formData.value)
               break
             default:
               res = await createRequireOrder(formData.value)
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
