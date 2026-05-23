<template>
  <div class="edit-page">
    <van-nav-bar title="编辑资料" left-arrow @click-left="router.back()" />

    <!-- 头像上传区域 -->
    <div class="avatar-section">
      <van-uploader :max-count="1" :after-read="onAfterRead" :before-read="beforeRead">
        <div class="avatar-trigger">
          <van-image
            v-if="userStore.userInfo?.avatar"
            round
            width="80px"
            height="80px"
            :src="userStore.userInfo.avatar"
            fit="cover"
          />
          <van-icon v-else name="user-circle-o" size="80" color="#c8c9cc" />
          <span class="avatar-text">更换头像</span>
        </div>
      </van-uploader>
    </div>

    <!-- 表单区域 -->
    <van-form @submit="onSubmit">
      <van-cell-group inset>
        <van-field v-model="form.nickname" label="昵称" placeholder="请输入昵称" />
        <van-field
          :model-value="userStore.userInfo?.phone"
          label="手机号"
          readonly
        />
        <van-field
          v-model="form.locationText"
          label="常住地"
          is-link
          readonly
          placeholder="请选择常住地"
          @click="showArea = true"
        />
      </van-cell-group>

      <!-- 球类兴趣 -->
      <div class="sports-section">
        <div class="sports-label">球类兴趣</div>
        <van-checkbox-group v-model="form.sports" direction="horizontal" class="sports-group">
          <van-checkbox
            v-for="sport in sportOptions"
            :key="sport"
            :name="sport"
            shape="round"
          >
            {{ sport }}
          </van-checkbox>
        </van-checkbox-group>
      </div>

      <div class="submit-btn">
        <van-button type="primary" block round :loading="loading" @click="onSubmit">
          保存
        </van-button>
      </div>
    </van-form>

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
import { ref, reactive, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import { areaList } from '@vant/area-data'
import { useUserStore } from '@/stores/user'
import { updateProfile, uploadAvatar } from '@/api/user'

const router = useRouter()
const userStore = useUserStore()
const loading = ref(false)
const showArea = ref(false)

const sportOptions = ['篮球', '足球', '羽毛球', '乒乓球', '网球', '排球', '棒球', '橄榄球']

const form = reactive({
  nickname: '',
  location: '',
  locationText: '',
  sports: []
})

// 初始化表单数据
onMounted(() => {
  const info = userStore.userInfo
  if (info) {
    form.nickname = info.nickname || ''
    form.location = info.location || ''
    form.locationText = info.location || ''
    form.sports = info.sports ? info.sports.split(',').filter(s => s) : []
  }
})

// 头像上传前校验
const beforeRead = (file) => {
  if (!/\.(jpg|jpeg|png|gif)$/i.test(file.name)) {
    showToast('请上传图片文件')
    return false
  }
  if (file.size > 5 * 1024 * 1024) {
    showToast('图片大小不能超过5MB')
    return false
  }
  return true
}

// 头像上传
const onAfterRead = async (file) => {
  const formData = new FormData()
  formData.append('avatar', file.file)
  try {
    const res = await uploadAvatar(formData)
    showToast('头像上传成功')
    userStore.setUserInfo({ ...userStore.userInfo, avatar: res.data.avatar })
  } catch (e) {
    showToast('上传失败，请重试')
  }
}

// 省市区选择确认
const onAreaConfirm = ({ selectedOptions }) => {
  showArea.value = false
  const text = selectedOptions.map(item => item.text).join('/')
  form.locationText = text
  form.location = text
}

// 保存资料
const onSubmit = async () => {
  loading.value = true
  try {
    const res = await updateProfile({
      nickname: form.nickname,
      location: form.location,
      sports: form.sports.join(',')
    })
    showToast('保存成功')
    userStore.setUserInfo(res.data)
    router.back()
  } catch (e) {
    // 错误由拦截器统一处理
  } finally {
    loading.value = false
  }
}
</script>

<style scoped>
.edit-page {
  background-color: #f7f8fa;
  min-height: 100vh;
}

.avatar-section {
  display: flex;
  justify-content: center;
  padding: 24px 0;
}

.avatar-trigger {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.avatar-text {
  margin-top: 8px;
  font-size: 14px;
  color: #1989fa;
}

.sports-section {
  padding: 16px;
  background-color: #fff;
  margin-top: 12px;
}

.sports-label {
  font-size: 14px;
  color: #646566;
  margin-bottom: 12px;
}

.sports-group {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.submit-btn {
  padding: 24px 16px;
}
</style>