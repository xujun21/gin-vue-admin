<template>
  <div>
    <div class="gva-form-box">
      <el-form ref="elFormRef" :model="formData" label-position="right" :rules="rule" label-width="80px">
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
          <el-date-picker v-model="formData.exp_date" type="date" placeholder="选择日期" :clearable="true" />
        </el-form-item>
        <el-form-item label="商品价格:" prop="price">
          <el-input-number v-model="formData.price" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="商品税:" prop="vat">
          <el-input-number v-model="formData.vat" :precision="2" :clearable="true" />
        </el-form-item>
        <el-form-item label="商品库存量:" prop="store">
          <el-input v-model.number="formData.store" :clearable="true" placeholder="请输入" />
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
  name: 'Product'
}
</script>

<script setup>
import {
  createProduct,
  updateProduct,
  findProduct
} from '@/api/product'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from 'vue-router'
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
  code: '',
  product_name_cn: '',
  product_name_en: '',
  package: '',
  exp_date: new Date(),
  price: 0,
  vat: 0,
  store: 0,
})
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async() => {
  // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
  if (route.query.id) {
    const res = await findProduct({ ID: route.query.id })
    if (res.code === 0) {
      formData.value = res.data.reprod
      type.value = 'update'
    }
  } else {
    type.value = 'create'
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
