<template>
  <div class="home-page">
    <van-nav-bar title="附近约球" />

    <!-- 球类筛选标签栏（横向滚动） -->
    <div class="sport-tabs">
      <div
        v-for="sport in sportOptions"
        :key="sport.value"
        class="sport-tab"
        :class="{ active: selectedSport === sport.value }"
        @click="onSportChange(sport.value)"
      >
        {{ sport.label }}
      </div>
    </div>

    <!-- 排序切换 -->
    <div class="sort-bar">
      <span class="sort-label">排序：</span>
      <button
        class="sort-btn"
        :class="{ active: sortBy === 'distance' }"
        @click="onSortChange('distance')"
      >
        距离优先
      </button>
      <button
        class="sort-btn"
        :class="{ active: sortBy === 'time' }"
        @click="onSortChange('time')"
      >
        时间优先
      </button>
    </div>

    <!-- 定位中提示 -->
    <div v-if="locating" class="locating-tip">
      <van-loading size="20px" />
      <span>正在获取位置...</span>
    </div>

    <!-- 定位失败提示 + 手动选点 -->
    <div v-else-if="locationError && !hasLocation" class="location-error">
      <van-empty image="location" description="无法获取位置" />
      <van-button type="primary" size="small" @click="showMapPicker = true">
        手动选择位置
      </van-button>
    </div>

    <!-- 邀约列表（定位成功后渲染，per Pitfall 3） -->
    <van-list
      v-if="hasLocation"
      v-model:loading="loading"
      :finished="finished"
      finished-text="没有更多了"
      @load="onLoad"
    >
      <InvitationCard
        v-for="item in list"
        :key="item.id"
        :invitation="item"
        mode="distance"
        @click="goDetail(item.id)"
      />
      <!-- 空状态 -->
      <van-empty
        v-if="finished && list.length === 0"
        description="附近暂无邀约"
      />
    </van-list>

    <!-- 地图选点弹窗（定位失败时使用） -->
    <van-popup
      v-model:show="showMapPicker"
      position="bottom"
      :style="{ height: '60%' }"
    >
      <div class="map-picker-header">
        <span>选择位置</span>
        <van-icon name="cross" @click="showMapPicker = false" />
      </div>
      <MapPicker @select="onMapSelect" />
    </van-popup>

    <TabBar />
  </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import TabBar from '@/components/TabBar.vue'
import InvitationCard from '@/components/InvitationCard.vue'
import MapPicker from '@/components/MapPicker.vue'
import { useGeolocation } from '@/composables/useGeolocation'
import { getNearbyInvitations } from '@/api/invitation'

const router = useRouter()
const { getPosition } = useGeolocation()

// 位置状态
const hasLocation = ref(false)
const locating = ref(false)
const locationError = ref(false)
const currentLat = ref(null)
const currentLng = ref(null)
const showMapPicker = ref(false)

// 筛选和排序
const selectedSport = ref('')
const sortBy = ref('distance')

// 球类选项
const sportOptions = [
  { label: '全部', value: '' },
  { label: '篮球', value: '篮球' },
  { label: '足球', value: '足球' },
  { label: '羽毛球', value: '羽毛球' },
  { label: '乒乓球', value: '乒乓球' },
  { label: '网球', value: '网球' },
  { label: '排球', value: '排球' }
]

// 列表状态
const list = ref([])
const loading = ref(false)
const finished = ref(false)
const page = ref(1)
const total = ref(0)
const PAGE_SIZE = 10

// 上拉加载更多
async function onLoad() {
  try {
    const res = await getNearbyInvitations({
      lat: currentLat.value,
      lng: currentLng.value,
      page: page.value,
      page_size: PAGE_SIZE,
      sort_by: sortBy.value,
      sport_type: selectedSport.value
    })
    const data = res.data
    list.value.push(...(data.list || []))
    total.value = data.total || 0
    page.value++
    // 判断是否已加载全部
    if (list.value.length >= total.value) {
      finished.value = true
    }
  } catch (err) {
    showToast('加载失败，请重试')
    finished.value = true
  } finally {
    loading.value = false
  }
}

// 重置列表并重新加载（切换筛选/排序时调用）
function resetList() {
  list.value = []
  page.value = 1
  total.value = 0
  finished.value = false
  loading.value = true
  onLoad()
}

// 切换球类筛选
function onSportChange(sport) {
  if (selectedSport.value === sport) return
  selectedSport.value = sport
  if (hasLocation.value) resetList()
}

// 切换排序方式
function onSortChange(sort) {
  if (sortBy.value === sort) return
  sortBy.value = sort
  if (hasLocation.value) resetList()
}

// 跳转邀约详情
function goDetail(id) {
  router.push(`/invitation/${id}`)
}

// 地图选点确认（定位失败时的回退）
function onMapSelect({ latitude, longitude }) {
  currentLat.value = latitude
  currentLng.value = longitude
  hasLocation.value = true
  locationError.value = false
  showMapPicker.value = false
  resetList()
}

// 页面挂载时自动定位
onMounted(async () => {
  locating.value = true
  try {
    const { lat, lng } = await getPosition()
    currentLat.value = lat
    currentLng.value = lng
    hasLocation.value = true
  } catch (err) {
    locationError.value = true
    showToast('定位失败，请手动选择位置')
  } finally {
    locating.value = false
  }
})
</script>

<style scoped>
.home-page {
  min-height: 100vh;
  background: #f7f8fa;
  padding-bottom: 60px;
}

/* 球类标签横向滚动栏 */
.sport-tabs {
  display: flex;
  overflow-x: auto;
  padding: 10px 12px;
  gap: 8px;
  background: #fff;
  border-bottom: 1px solid #ebedf0;
  scrollbar-width: none;
}

.sport-tabs::-webkit-scrollbar {
  display: none;
}

.sport-tab {
  flex-shrink: 0;
  padding: 4px 14px;
  border-radius: 16px;
  font-size: 13px;
  color: #646566;
  background: #f7f8fa;
  cursor: pointer;
  transition: all 0.2s;
  white-space: nowrap;
}

.sport-tab.active {
  background: #1989fa;
  color: #fff;
}

/* 排序切换栏 */
.sort-bar {
  display: flex;
  align-items: center;
  padding: 8px 16px;
  background: #fff;
  border-bottom: 1px solid #ebedf0;
  gap: 8px;
}

.sort-label {
  font-size: 13px;
  color: #969799;
}

.sort-btn {
  padding: 3px 12px;
  border-radius: 12px;
  font-size: 13px;
  border: 1px solid #ebedf0;
  background: #f7f8fa;
  color: #646566;
  cursor: pointer;
  transition: all 0.2s;
}

.sort-btn.active {
  border-color: #1989fa;
  color: #1989fa;
  background: #e8f3ff;
}

/* 定位中提示 */
.locating-tip {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  padding: 40px 16px;
  color: #969799;
  font-size: 14px;
}

/* 定位失败提示 */
.location-error {
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 20px 16px;
  gap: 12px;
}

/* 地图选点弹窗头部 */
.map-picker-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 12px 16px;
  font-size: 15px;
  font-weight: 500;
  border-bottom: 1px solid #ebedf0;
}
</style>
