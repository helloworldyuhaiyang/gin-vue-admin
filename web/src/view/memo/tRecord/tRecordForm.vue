<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="记录名字:" prop="recordName">
          <el-input v-model="formData.recordName" :clearable="true"  placeholder="请输入记录名字" />
       </el-form-item>
        <el-form-item label="第三方存储的地址:" prop="resourceUrl">
          <el-input v-model="formData.resourceUrl" :clearable="true"  placeholder="请输入第三方存储的地址" />
       </el-form-item>
        <el-form-item label="记录位置:" prop="location">
          <el-input v-model="formData.location" :clearable="true"  placeholder="请输入记录位置" />
       </el-form-item>
        <el-form-item label="记录文件大小:" prop="rSize">
          <el-input v-model.number="formData.rSize" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="创建者:" prop="createdBy">
          <el-input v-model.number="formData.createdBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="更新者:" prop="updatedBy">
          <el-input v-model.number="formData.updatedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="删除者:" prop="deletedBy">
          <el-input v-model.number="formData.deletedBy" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createTRecord,
  updateTRecord,
  findTRecord
} from '@/api/memo/tRecord'

defineOptions({
    name: 'TRecordForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            recordName: '',
            resourceUrl: '',
            location: '',
            rSize: undefined,
            createdBy: undefined,
            updatedBy: undefined,
            deletedBy: undefined,
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findTRecord({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
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
               res = await createTRecord(formData.value)
               break
             case 'update':
               res = await updateTRecord(formData.value)
               break
             default:
               res = await createTRecord(formData.value)
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
