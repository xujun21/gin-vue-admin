<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
          <el-input-number v-model="formData.price" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="折扣金额:" prop="discount">
          <el-input-number v-model="formData.discount" :precision="2" :clearable="true"></el-input-number>
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
  name: 'CompanyDiscount'
}
</script>

<script setup>
import {
  createCompanyDiscount,
  updateCompanyDiscount,
  findCompanyDiscount
} from '@/api/companyDiscount'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCompanyDiscount({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recompDiscount
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
