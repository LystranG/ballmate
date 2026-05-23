<template>
  <div class="password-page">
    <van-nav-bar title="修改密码" left-arrow @click-left="router.back()" />

    <van-form @submit="onSubmit">
      <van-cell-group inset>
        <van-field
          v-model="form.oldPassword"
          type="password"
          label="原密码"
          placeholder="请输入原密码"
          :rules="[{ required: true, message: '请输入原密码' }]"
        />
        <van-field
          v-model="form.newPassword"
          type="password"
          label="新密码"
          placeholder="请输入新密码"
          :rules="[{ required: true, message: '请输入新密码' }]"
        />
        <van-field
          v-model="form.confirmPassword"
          type="password"
          label="确认新密码"
          placeholder="请再次输入新密码"
          :rules="[{ required: true, message: '请确认新密码' }]"
        />
      </van-cell-group>

      <div class="submit-btn">
        <van-button type="primary" block round :loading="loading" @click="onSubmit">
          确认修改
        </van-button>
      </div>
    </van-form>
  </div>
</template>

<script setup>
import { ref, reactive } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import { updatePassword } from '@/api/user'

const router = useRouter()
const loading = ref(false)

const form = reactive({
  oldPassword: '',
  newPassword: '',
  confirmPassword: ''
})

const onSubmit = async () => {
  // 前端校验两次密码是否一致
  if (form.newPassword !== form.confirmPassword) {
    showToast('两次密码输入不一致')
    return
  }

  loading.value = true
  try {
    await updatePassword({
      oldPassword: form.oldPassword,
      newPassword: form.newPassword,
      confirmPassword: form.confirmPassword
    })
    showToast('修改成功')
    router.back()
  } catch (e) {
    // 错误由拦截器统一处理
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.password-page {
  background-color: #f7f8fa;
  min-height: 100vh;
}

.submit-btn {
  padding: 24px 16px;
}
</style>
