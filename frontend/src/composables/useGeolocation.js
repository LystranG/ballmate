import { ref } from 'vue'

/**
 * useGeolocation composable
 * 封装浏览器 Geolocation API，提供响应式位置数据和错误状态
 */
export function useGeolocation() {
  const latitude = ref(null)
  const longitude = ref(null)
  const error = ref(null)
  const loading = ref(false)

  /**
   * 获取当前位置，返回 Promise
   * 成功 resolve { lat, lng }，失败 reject 错误码字符串
   */
  function getPosition() {
    return new Promise((resolve, reject) => {
      // 检查浏览器是否支持定位
      if (!navigator.geolocation) {
        error.value = 'NOT_SUPPORTED'
        reject(error.value)
        return
      }

      loading.value = true
      navigator.geolocation.getCurrentPosition(
        (position) => {
          latitude.value = position.coords.latitude
          longitude.value = position.coords.longitude
          loading.value = false
          resolve({ lat: latitude.value, lng: longitude.value })
        },
        (err) => {
          loading.value = false
          // err.code: 1=PERMISSION_DENIED, 2=POSITION_UNAVAILABLE, 3=TIMEOUT
          switch (err.code) {
            case 1:
              error.value = 'PERMISSION_DENIED'
              break
            case 2:
              error.value = 'POSITION_UNAVAILABLE'
              break
            case 3:
              error.value = 'TIMEOUT'
              break
            default:
              error.value = 'UNKNOWN'
          }
          reject(error.value)
        },
        {
          enableHighAccuracy: true,
          timeout: 10000,
          maximumAge: 300000 // 5分钟缓存，避免频繁请求
        }
      )
    })
  }

  return { latitude, longitude, error, loading, getPosition }
}
