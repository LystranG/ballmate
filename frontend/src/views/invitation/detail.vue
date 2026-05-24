<template>
  <div class="detail-page">
    <van-nav-bar
      title="邀约详情"
      left-arrow
      @click-left="router.back()"
    />

    <div class="detail-content" v-if="invitation" :style="!isCreator ? 'padding-bottom: 80px' : ''">
      <!-- 状态徽章 -->
      <div class="status-section">
        <van-tag
          size="large"
          :color="statusConfig[invitation.status]?.bg"
          :text-color="statusConfig[invitation.status]?.text"
        >
          {{ statusConfig[invitation.status]?.label }}
        </van-tag>
      </div>

      <!-- 邀约信息 -->
      <van-cell-group inset class="info-group">
        <van-cell title="球类" :value="invitation.sport_type" />
        <van-cell title="活动时间" :value="formatTime(invitation.activity_time)" />
        <van-cell title="活动地点" :value="invitation.address" />
        <van-cell
          title="参与人数"
          :value="`${invitation.participant_count}/${invitation.max_people}人`"
        />
      </van-cell-group>

      <!-- 参与人列表（仅发起人可见） -->
      <van-cell-group
        v-if="isCreator && participants.length > 0"
        inset
        class="info-group"
        title="参与人列表"
      >
        <van-cell
          v-for="item in participants"
          :key="item.user_id"
          :title="item.nickname"
          :value="item.phone"
        />
      </van-cell-group>

      <!-- 操作按钮（仅发起人可见） -->
      <div class="action-section" v-if="isCreator">
        <!-- 等待中状态：终止邀约按钮 -->
        <van-button
          v-if="invitation.status === 'waiting'"
          type="danger"
          block
          @click="handleTerminate"
        >
          终止邀约
        </van-button>

        <!-- 已终止状态：删除邀约按钮 -->
        <van-button
          v-if="invitation.status === 'terminated'"
          type="danger"
          block
          @click="handleDelete"
        >
          删除邀约
        </van-button>
      </div>
    </div>

    <!-- 加载状态 -->
    <div class="loading-wrapper" v-else>
      <van-loading size="24px" vertical>加载中...</van-loading>
    </div>

    <!-- 底部固定操作栏（非创建者可见） -->
    <div class="bottom-action" v-if="!isCreator && invitation">
      <!-- 未加入：显示加入按钮 -->
      <van-button
        v-if="!invitation.has_joined"
        type="primary"
        block
        round
        :disabled="!canJoin"
        @click="handleJoin"
      >
        {{ canJoin ? '加入邀约' : (invitation.status === 'waiting' ? '已满员' : '已' + (statusConfig[invitation.status]?.label || '结束')) }}
      </van-button>

      <!-- 已加入：显示退出按钮 -->
      <van-button
        v-else
        type="warning"
        block
        round
        :disabled="!canLeave"
        @click="handleLeave"
      >
        退出邀约
      </van-button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue'
import { useRoute, useRouter } from 'vue-router'
import { showDialog, showToast } from 'vant'
import { useUserStore } from '@/stores/user'
import {
  getInvitation,
  terminateInvitation,
  deleteInvitation,
  getParticipants,
  joinInvitation,
  leaveInvitation
} from '@/api/invitation'

const route = useRoute()
const router = useRouter()
const userStore = useUserStore()

const invitation = ref(null)
const participants = ref([])

// 状态配置：颜色和标签
const statusConfig = {
  waiting: { bg: '#fffbe8', text: '#ed6a0c', label: '等待中' },
  gathered: { bg: '#e8f7e8', text: '#07c160', label: '已召集' },
  terminated: { bg: '#f5f5f5', text: '#969799', label: '已终止' },
  expired: { bg: '#f5f5f5', text: '#969799', label: '已过期' }
}

// 判断当前用户是否为发起人
const isCreator = computed(() => {
  return userStore.userInfo?.id === invitation.value?.creator_id
})

// 是否可以加入：waiting 状态且未满员
const canJoin = computed(() => {
  if (!invitation.value) return false
  return invitation.value.status === 'waiting' &&
    invitation.value.participant_count < invitation.value.max_people
})

// 是否可以退出：waiting 或 gathered 状态均可退出
const canLeave = computed(() => {
  if (!invitation.value) return false
  const s = invitation.value.status
  return s === 'waiting' || s === 'gathered'
})

// 格式化时间为 MM月DD日 HH:mm
function formatTime(timeStr) {
  if (!timeStr) return ''
  const date = new Date(timeStr)
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  return `${month}月${day}日 ${hours}:${minutes}`
}

// 加载邀约详情
async function loadDetail() {
  try {
    const res = await getInvitation(route.params.id)
    invitation.value = res.data
    // 如果是发起人，加载参与人列表
    if (userStore.userInfo?.id === res.data.creator_id) {
      loadParticipants()
    }
  } catch (error) {
    showToast('加载失败')
  }
}

// 加载参与人列表
async function loadParticipants() {
  try {
    const res = await getParticipants(route.params.id)
    participants.value = res.data || []
  } catch (error) {
    // 参与人加载失败不阻塞页面
  }
}

// 终止邀约
function handleTerminate() {
  showDialog({
    title: '终止邀约',
    message: '确定要终止这个邀约吗？终止后参与人将无法继续加入。',
    showCancelButton: true
  }).then(async () => {
    try {
      await terminateInvitation(route.params.id)
      showToast('已终止')
      invitation.value.status = 'terminated'
    } catch (error) {
      showToast('操作失败')
    }
  }).catch(() => {})
}

// 删除邀约
function handleDelete() {
  showDialog({
    title: '删除邀约',
    message: '确定要删除这个邀约吗？删除后将不再显示。',
    showCancelButton: true
  }).then(async () => {
    try {
      await deleteInvitation(route.params.id)
      showToast('已删除')
      router.back()
    } catch (error) {
      showToast('操作失败')
    }
  }).catch(() => {})
}

// 加入邀约
function handleJoin() {
  showDialog({
    title: '加入邀约',
    message: '确定要加入这个邀约吗？',
    showCancelButton: true
  }).then(async () => {
    try {
      await joinInvitation(route.params.id)
      showToast('已加入')
      loadDetail()
    } catch (error) {
      const msg = error?.response?.data?.message || '操作失败'
      showToast(msg)
    }
  }).catch(() => {})
}

// 退出邀约
function handleLeave() {
  showDialog({
    title: '退出邀约',
    message: '确定要退出这个邀约吗？',
    showCancelButton: true
  }).then(async () => {
    try {
      await leaveInvitation(route.params.id)
      showToast('已退出')
      loadDetail()
    } catch (error) {
      const msg = error?.response?.data?.message || '操作失败'
      showToast(msg)
    }
  }).catch(() => {})
}

onMounted(() => {
  loadDetail()
})
</script>

<style scoped>
.detail-page {
  min-height: 100vh;
  background: #f7f8fa;
}

.detail-content {
  padding: 16px 0;
}

.status-section {
  padding: 0 16px 12px;
}

.info-group {
  margin-bottom: 12px;
}

.action-section {
  padding: 16px;
}

.loading-wrapper {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60vh;
}

/* 底部固定操作栏 */
.bottom-action {
  position: fixed;
  bottom: 0;
  left: 0;
  right: 0;
  padding: 12px 16px;
  background: #fff;
  box-shadow: 0 -2px 8px rgba(0, 0, 0, 0.06);
  z-index: 100;
}
</style>
