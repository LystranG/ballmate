<template>
  <div class="map-picker">
    <div ref="mapContainer" class="map-container"></div>
    <div v-if="address" class="address-display">
      <van-icon name="location-o" />
      <span>{{ address }}</span>
    </div>
  </div>
</template>

<script setup>
import { ref, onMounted, onUnmounted } from 'vue'
import AMapLoader from '@amap/amap-jsapi-loader'
import { AMAP_KEY } from '@/config'
import { useGeolocation } from '@/composables/useGeolocation'

const emit = defineEmits(['select'])

const mapContainer = ref(null)
const address = ref('')
const { getPosition } = useGeolocation()
let map = null
let marker = null
let geocoder = null

// 初始化地图
onMounted(async () => {
  try {
    const AMap = await AMapLoader.load({
      key: AMAP_KEY,
      version: '2.0',
      plugins: ['AMap.Geocoder']
    })

    map = new AMap.Map(mapContainer.value, {
      zoom: 15,
      center: [104.06, 30.52] // 默认成都，定位成功后会移动
    })

    geocoder = new AMap.Geocoder()

    // 获取当前位置并移动地图中心
    getPosition().then(({ lat, lng }) => {
      map.setCenter([lng, lat])
    }).catch(() => {
      // 定位失败保持默认中心
    })

    // 点击地图选点
    map.on('click', (e) => {
      const { lng, lat } = e.lnglat
      placeMarker(AMap, lng, lat)
      // 先立即发送坐标，确保即使逆地理编码失败也能拿到位置
      emit('select', { longitude: lng, latitude: lat, address: '' })
      reverseGeocode(lng, lat)
    })
  } catch (err) {
    console.error('地图加载失败:', err)
  }
})

onUnmounted(() => {
  if (map) {
    map.destroy()
    map = null
  }
})

// 放置标记点
function placeMarker(AMap, lng, lat) {
  if (marker) {
    marker.setPosition([lng, lat])
  } else {
    marker = new AMap.Marker({ position: [lng, lat] })
    map.add(marker)
  }
  map.setCenter([lng, lat])
}

// 逆地理编码获取地址
function reverseGeocode(lng, lat) {
  if (!geocoder) return
  geocoder.getAddress([lng, lat], (status, result) => {
    if (status === 'complete' && result.regeocode) {
      address.value = result.regeocode.formattedAddress
      emit('select', {
        longitude: lng,
        latitude: lat,
        address: address.value
      })
    }
  })
}
</script>

<style scoped>
.map-picker {
  width: 100%;
}

.map-container {
  width: 100%;
  height: 300px;
  border-radius: 8px;
  overflow: hidden;
}

.address-display {
  display: flex;
  align-items: center;
  gap: 4px;
  margin-top: 8px;
  padding: 8px 12px;
  background: #f7f8fa;
  border-radius: 4px;
  font-size: 13px;
  color: #323233;
}
</style>
