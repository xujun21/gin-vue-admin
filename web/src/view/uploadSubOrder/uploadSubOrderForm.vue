<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="商品ID:" prop="productId">
          <el-input v-model.number="formData.productId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品代码:" prop="productCode">
          <el-input v-model="formData.productCode" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品中文名:" prop="productNameCn">
          <el-input v-model="formData.productNameCn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="商品英文名:" prop="productNameEn">
          <el-input v-model="formData.productNameEn" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="包装规格:" prop="package">
          <el-input v-model="formData.package" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="保质期:" prop="expDate">
          <el-date-picker v-model="formData.expDate" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="数量:" prop="quantity">
          <el-input v-model.number="formData.quantity" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="剩余库存:" prop="store">
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
  name: 'UploadSubOrder'
}
</script>

<script setup>
import {
  createUploadSubOrder,
  updateUploadSubOrder,
  findUploadSubOrder
} from '@/api/uploadSubOrder'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            productId: 0,
            productCode: '',
            productNameCn: '',
            productNameEn: '',
            package: '',
            expDate: new Date(),
            quantity: 0,
            store: 0,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findUploadSubOrder({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reupSubOrd
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
               res = await createUploadSubOrder(formData.value)
               break
             case 'update':
               res = await updateUploadSubOrder(formData.value)
               break
             default:
               res = await createUploadSubOrder(formData.value)
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
