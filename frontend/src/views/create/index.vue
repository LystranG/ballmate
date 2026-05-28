<template>
  <div class="create-page">
    <van-nav-bar title="发起约球" left-arrow @click-left="$router.back()" />

    <van-form @submit="onSubmit" class="form-content">
      <!-- 球类选择 -->
      <van-cell-group inset>
        <van-field name="sportType" label="球类" :rules="[{ validator: () => !!form.sportType, message: '请选择球类' }]">
          <template #input>
            <div class="tag-group">
              <van-tag
                v-for="sport in sportTypes"
                :key="sport"
                size="large"
                :type="form.sportType === sport ? 'primary' : 'default'"
                :plain="form.sportType !== sport"
                @click="form.sportType = sport"
              >
                {{ sport }}
              </van-tag>
            </div>
          </template>
        </van-field>
      </van-cell-group>

      <!-- 活动时间 -->
      <van-cell-group inset>
        <van-field
          v-model="displayTime"
          is-link
          readonly
          label="时间"
          placeholder="请选择活动时间"
          name="activityTime"
          :rules="[{ required: true, message: '请选择活动时间' }]"
          @click="showTimePicker = true"
        />
      </van-cell-group>

      <!-- 活动地点 -->
      <van-cell-group inset>
        <van-field
          v-model="form.address"
          is-link
          readonly
          label="地点"
          placeholder="请选择活动地点"
          name="location"
          :rules="[{ required: true, message: '请选择活动地点' }]"
          @click="showMap = true"
        />
      </van-cell-group>

      <!-- 人数选择 -->
      <van-cell-group inset>
        <van-field name="capacity" label="人数" :rules="[{ validator: () => !!form.capacity, message: '请选择人数' }]">
          <template #input>
            <div class="tag-group">
              <van-tag
                v-for="cap in capacityOptions"
                :key="cap"
                size="large"
                :type="form.capacity === cap ? 'primary' : 'default'"
                :plain="form.capacity !== cap"
                @click="form.capacity = cap"
              >
                {{ cap }}人
              </van-tag>
            </div>
          </template>
        </van-field>
      </van-cell-group>

      <!-- 提交按钮 -->
      <div class="submit-area">
        <van-button type="primary" block round native-type="submit" :loading="submitting">
          发布邀约
        </van-button>
        <van-button plain block round style="margin-top: 12px" @click="router.push('/')">
          返回主页
        </van-button>
      </div>
    </van-form>

    <!-- 时间选择弹出层 -->
    <van-popup v-model:show="showTimePicker" position="bottom" round>
      <van-picker-group
        title="选择活动时间"
        :tabs="['日期', '时间']"
        @confirm="onTimeConfirm"
        @cancel="showTimePicker = false"
      >
        <van-date-picker v-model="pickerDate" :min-date="minDate" />
        <van-time-picker v-model="pickerTime" />
      </van-picker-group>
    </van-popup>

    <!-- 地图选点弹出层 -->
    <van-popup v-model:show="showMap" position="bottom" round :style="{ height: '60%' }">
      <div class="map-popup-header">
        <van-button size="small" @click="showMap = false">取消</van-button>
        <span class="map-popup-title">选择地点</span>
        <van-button size="small" type="primary" @click="confirmLocation">确定</van-button>
      </div>
      <MapPicker @select="onMapSelect" />
    </van-popup>

    <TabBar />
  </div>
</template>

<script setup>
import { ref, reactive, computed } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import { createInvitation } from '@/api/invitation'
import MapPicker from '@/components/MapPicker.vue'
import TabBar from '@/components/TabBar.vue'

const router = useRouter()

// 球类选项
const sportTypes = ['篮球', '足球', '羽毛球', '乒乓球', '网球', '排球']
// 人数选项
const capacityOptions = [2, 4, 6, 8, 10]

const form = reactive({
  sportType: '',
  activityTime: '',
  longitude: null,
  latitude: null,
  address: '',
  capacity: null
})

const submitting = ref(false)
const showTimePicker = ref(false)
const showMap = ref(false)

// 时间选择器状态
const minDate = new Date()
const pickerDate = ref([
  String(minDate.getFullYear()),
  String(minDate.getMonth() + 1).padStart(2, '0'),
  String(minDate.getDate()).padStart(2, '0')
])
const pickerTime = ref([
  String(minDate.getHours()).padStart(2, '0'),
  String(minDate.getMinutes()).padStart(2, '0')
])

// 显示已选时间
const displayTime = computed(() => {
  if (!form.activityTime) return ''
  const d = new Date(form.activityTime)
  const month = d.getMonth() + 1
  const day = d.getDate()
  const hour = String(d.getHours()).padStart(2, '0')
  const min = String(d.getMinutes()).padStart(2, '0')
  return `${month}月${day}日 ${hour}:${min}`
})

// 临时地图选点数据
const tempLocation = ref(null)

// 时间确认
function onTimeConfirm() {
  const [year, month, day] = pickerDate.value
  const [hour, minute] = pickerTime.value
  form.activityTime = new Date(year, month - 1, day, hour, minute).toISOString()
  showTimePicker.value = false
}

// 地图选点回调
function onMapSelect(location) {
  tempLocation.value = location
}

// 确认地点选择
function confirmLocation() {
  if (!tempLocation.value) {
    showToast('请先在地图上点击选择地点')
    return
  }
  form.longitude = tempLocation.value.longitude
  form.latitude = tempLocation.value.latitude
  form.address = tempLocation.value.address || '已选择位置'
  showMap.value = false
}

// 提交表单
async function onSubmit() {
  submitting.value = true
  try {
    const payload = {
      sport_type: form.sportType,
      activity_time: form.activityTime,
      longitude: form.longitude,
      latitude: form.latitude,
      address: form.address,
      max_people: form.capacity
    }
    const res = await createInvitation(payload)
    showToast('发布成功')
    router.push(`/invitation/${res.data.id}`)
  } catch (err) {
    // 错误已由 request 拦截器处理
  } finally {
    submitting.value = false
  }
}
</script>

<style scoped>
.create-page {
  min-height: 100vh;
  background: #f7f8fa;
  padding-bottom: 60px;
}

.form-content {
  padding: 12px 0;
}

.tag-group {
  display: flex;
  flex-wrap: wrap;
  gap: 8px;
}

.submit-area {
  padding: 24px 16px;
}

.map-popup-header {
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px 16px;
  border-bottom: 1px solid #ebedf0;
}

.map-popup-title {
  font-size: 16px;
  font-weight: 500;
}
</style>
