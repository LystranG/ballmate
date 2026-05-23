<template>
  <div class="login-page">
    <h1 class="app-title">掌上约球</h1>

    <van-form @submit="onSubmit">
      <van-cell-group inset>
        <van-field
          v-model="form.username"
          name="username"
          placeholder="请输入账号"
          :rules="[{ required: true, message: '请输入账号' }]"
        />
        <van-field
          v-model="form.password"
          type="password"
          name="password"
          placeholder="请输入密码"
          :rules="[{ required: true, message: '请输入密码' }]"
        />
      </van-cell-group>

      <div class="btn-wrapper">
        <van-button
          type="primary"
          block
          round
          native-type="submit"
          :loading="loading"
        >
          登录
        </van-button>
      </div>
    </van-form>

    <div class="link-wrapper">
      <span @click="$router.push('/register')">还没有账号？去注册</span>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { useUserStore } from '@/stores/user'
import { login } from '@/api/user'

const router = useRouter()
const userStore = useUserStore()
const loading = ref(false)

const form = ref({
  username: '',
  password: ''
})

const onSubmit = async () => {
  loading.value = true
  try {
    const res = await login(form.value)
    userStore.setToken(res.data.token)
    userStore.setUserInfo(res.data.user)
    router.push('/')
  } catch (e) {
    // 错误由 Axios 拦截器统一处理
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.login-page {
  padding: 32px 16px;
}

.app-title {
  text-align: center;
  font-size: 24px;
  font-weight: 600;
  margin-top: 32px;
  margin-bottom: 32px;
}

.btn-wrapper {
  padding: 24px 16px;
}

.link-wrapper {
  text-align: center;
  margin-top: 16px;
  color: #1989fa;
  font-size: 14px;
}

.link-wrapper span {
  cursor: pointer;
}
</style>
