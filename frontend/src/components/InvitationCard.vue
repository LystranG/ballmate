<template>
  <!-- 邀约卡片：mode=distance 显示距离，mode=status 显示状态标签（复用于我的邀约页） -->
  <div
    class="invitation-card"
    :class="{ 'is-inactive': isInactive }"
    @click="$emit('click', invitation)"
  >
    <div class="card-header">
      <span class="sport-type">{{ invitation.sport_type }}</span>
      <!-- mode=status 时显示状态标签，mode=distance 时显示距离 -->
      <van-tag
        v-if="mode === 'status'"
        :color="statusConfig[invitation.status]?.bg"
        :text-color="statusConfig[invitation.status]?.text"
      >
        {{ statusConfig[invitation.status]?.label }}
      </van-tag>
      <span v-else-if="invitation.distance != null" class="distance">
        {{ formatDistance(invitation.distance) }}
      </span>
    </div>

    <div class="card-body">
      <div class="info-row">
        <van-icon name="location-o" class="icon" />
        <span class="address">{{ truncateAddress(invitation.address) }}</span>
      </div>
      <div class="info-row">
        <van-icon name="clock-o" class="icon" />
        <span>{{ formatTime(invitation.activity_time) }}</span>
      </div>
      <div class="info-row">
        <van-icon name="friends-o" class="icon" />
        <span>{{ invitation.participant_count }}/{{ invitation.max_people }}人</span>
        <!-- 人数进度条 -->
        <van-progress
          class="progress"
          :percentage="Math.round((invitation.participant_count / invitation.max_people) * 100)"
          :show-pivot="false"
          stroke-width="4"
          :color="progressColor"
        />
      </div>
    </div>
  </div>
</template>

<script setup>
import { computed } from 'vue'
import { formatDistance, formatTime } from '@/utils/format'

const props = defineProps({
  invitation: {
    type: Object,
    required: true
  },
  /**
   * mode 控制卡片右上角显示内容：
   * - 'distance'（默认）：显示距离，用于附近邀约列表
   * - 'status'：显示状态标签，用于我的邀约列表
   */
  mode: {
    type: String,
    default: 'distance'
  }
})

defineEmits(['click'])

// 状态配置：颜色和标签（与详情页保持一致）
const statusConfig = {
  waiting: { bg: '#fffbe8', text: '#ed6a0c', label: '等待中' },
  gathered: { bg: '#e8f7e8', text: '#07c160', label: '已召集' },
  terminated: { bg: '#f5f5f5', text: '#969799', label: '已终止' },
  expired: { bg: '#f5f5f5', text: '#969799', label: '已过期' }
}

// 不可加入的邀约（已满/已终止/已过期）整体置灰
const isInactive = computed(() => {
  const s = props.invitation.status
  return s === 'gathered' || s === 'terminated' || s === 'expired'
})

// 人数进度条颜色：已满员显示绿色，否则显示主色
const progressColor = computed(() => {
  return props.invitation.participant_count >= props.invitation.max_people ? '#07c160' : '#1989fa'
})

// 地址截断：超过 16 字符显示省略号
function truncateAddress(address) {
  if (!address) return ''
  return address.length > 16 ? address.slice(0, 16) + '...' : address
}
</script>

<style scoped>
.invitation-card {
  margin: 8px 16px;
  padding: 12px 16px;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.08);
  cursor: pointer;
  transition: opacity 0.2s;
}

/* 不可加入的邀约整体置灰 */
.invitation-card.is-inactive {
  opacity: 0.55;
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.sport-type {
  font-size: 16px;
  font-weight: 600;
  color: #323233;
}

.distance {
  font-size: 13px;
  color: #1989fa;
  font-weight: 500;
}

.card-body {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.info-row {
  display: flex;
  align-items: center;
  gap: 6px;
  font-size: 13px;
  color: #646566;
}

.icon {
  font-size: 14px;
  color: #969799;
  flex-shrink: 0;
}

.address {
  flex: 1;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.progress {
  flex: 1;
  margin-left: 4px;
}
</style>
