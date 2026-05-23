<template>
  <div class="register-page">
    <van-nav-bar title="注册" left-arrow @click-left="$router.back()" />

    <van-form @submit="onSubmit">
      <van-cell-group inset>
        <van-field
          v-model="form.username"
          name="username"
          label="账号"
          placeholder="请输入账号"
          :rules="[{ required: true, message: '请输入账号' }]"
        />
        <van-field
          v-model="form.password"
          type="password"
          name="password"
          label="密码"
          placeholder="请输入密码"
          :rules="[{ required: true, message: '请输入密码' }]"
        />
        <van-field
          v-model="form.confirmPassword"
          type="password"
          name="confirmPassword"
          label="确认密码"
          placeholder="请确认密码"
          :rules="[
            { required: true, message: '请确认密码' },
            { validator: checkPassword, message: '两次密码输入不一致' }
          ]"
        />
        <van-field
          v-model="form.phone"
          type="tel"
          name="phone"
          label="手机号"
          placeholder="请输入手机号"
          :rules="[
            { required: true, message: '请输入手机号' },
            { pattern: /^\d{11}$/, message: '请输入正确的手机号' }
          ]"
        />
        <van-field
          v-model="form.location"
          is-link
          readonly
          name="location"
          label="常住地"
          placeholder="请选择常住地"
          @click="showArea = true"
        />
      </van-cell-group>

      <!-- 球类兴趣多选 -->
      <div class="sports-section">
        <div class="sports-label">球类兴趣</div>
        <van-checkbox-group v-model="selectedSports" direction="horizontal">
          <van-checkbox
            v-for="sport in sportsList"
            :key="sport"
            :name="sport"
            shape="round"
            class="sport-tag"
          >
            {{ sport }}
          </van-checkbox>
        </van-checkbox-group>
      </div>

      <div class="btn-wrapper">
        <van-button
          type="primary"
          block
          round
          native-type="submit"
          :loading="loading"
        >
          注册
        </van-button>
      </div>
    </van-form>

    <div class="link-wrapper">
      <span @click="$router.push('/login')">已有账号？去登录</span>
    </div>

    <!-- 省市区选择器 -->
    <van-popup v-model:show="showArea" position="bottom">
      <van-area
        :area-list="areaList"
        @confirm="onAreaConfirm"
        @cancel="showArea = false"
      />
    </van-popup>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import { areaList } from '@vant/area-data'
import { useUserStore } from '@/stores/user'
import { register } from '@/api/user'

const router = useRouter()
const userStore = useUserStore()
const loading = ref(false)
const showArea = ref(false)

const form = ref({
  username: '',
  password: '',
  confirmPassword: '',
  phone: '',
  location: ''
})

const selectedSports = ref([])

const sportsList = ['篮球', '足球', '羽毛球', '乒乓球', '网球', '排球', '棒球', '橄榄球']

// 校验两次密码一致
const checkPassword = (val) => val === form.value.password

// 省市区选择确认
const onAreaConfirm = ({ selectedOptions }) => {
  showArea.value = false
  form.value.location = selectedOptions.map(item => item.text).join('/')
}

const onSubmit = async () => {
  loading.value = true
  try {
    const data = {
      ...form.value,
      sports: selectedSports.value.join(',')
    }
    const res = await register(data)
    userStore.setToken(res.data.token)
    userStore.setUserInfo(res.data.user)
    showToast('注册成功')
    router.push('/')
  } catch (e) {
    // 错误由 Axios 拦截器统一处理
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.register-page {
  padding-bottom: 32px;
}

.sports-section {
  padding: 16px;
}

.sports-label {
  font-size: 14px;
  color: #646566;
  margin-bottom: 8px;
}

.sport-tag {
  margin: 4px;
}

.btn-wrapper {
  padding: 24px 16px;
}

.link-wrapper {
  text-align: center;
  color: #1989fa;
  font-size: 14px;
}

.link-wrapper span {
  cursor: pointer;
}
</style>
