<template>
  <div class="profile-page">
    <!-- 顶部 header 区域 -->
    <div class="profile-header">
      <van-image
        v-if="userStore.userInfo?.avatar"
        round
        width="64px"
        height="64px"
        :src="userStore.userInfo.avatar"
        fit="cover"
      />
      <van-icon v-else name="user-circle-o" size="64" color="#c8c9cc" />
      <div class="profile-info">
        <span class="nickname">{{ userStore.userInfo?.nickname || userStore.userInfo?.username }}</span>
        <span class="username">账号：{{ userStore.userInfo?.username }}</span>
      </div>
    </div>

    <!-- 菜单区域 -->
    <van-cell-group class="menu-group">
      <van-cell title="编辑资料" is-link @click="router.push('/profile/edit')" />
      <van-cell title="修改密码" is-link @click="router.push('/profile/password')" />
      <van-cell
        title="退出登录"
        :title-style="{ color: '#ee0a24' }"
        @click="handleLogout"
      />
    </van-cell-group>

    <TabBar />
  </div>
</template>

<script setup>
import { onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { showConfirmDialog } from 'vant'
import TabBar from '@/components/TabBar.vue'
import { useUserStore } from '@/stores/user'
import { getProfile } from '@/api/user'

const router = useRouter()
const userStore = useUserStore()

// 页面加载时获取用户信息
onMounted(async () => {
  try {
    const res = await getProfile()
    userStore.setUserInfo(res.data)
  } catch (e) {
    // 错误由拦截器统一处理
  }
})

// 退出登录
const handleLogout = () => {
  showConfirmDialog({
    title: '退出登录',
    message: '确定要退出登录吗？'
  }).then(() => {
    userStore.logout()
    router.push('/login')
  }).catch(() => {
    // 取消操作
  })
}
</script>

<style scoped>
.profile-page {
  padding-bottom: 50px;
}

.profile-header {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 24px 16px;
  background-color: #f7f8fa;
}

.profile-info {
  display: flex;
  flex-direction: column;
  align-items: center;
  margin-top: 12px;
}

.nickname {
  font-size: 18px;
  font-weight: 600;
  color: #323233;
}

.username {
  font-size: 14px;
  color: #969799;
  margin-top: 4px;
}

.menu-group {
  margin-top: 12px;
}
</style>
