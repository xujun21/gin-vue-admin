<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
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
  name: 'Company'
}
</script>

<script setup>
import {
  createCompany,
  updateCompany,
  findCompany
} from '@/api/company'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findCompany({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.recomp
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
