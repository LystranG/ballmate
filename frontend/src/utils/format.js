/**
 * formatDistance 格式化距离显示
 * < 1km 显示米数（四舍五入），>= 1km 显示公里数（保留1位小数）
 * @param {number} distanceKm - 距离（公里）
 * @returns {string} 格式化后的距离字符串，如 "500m" 或 "1.2km"
 */
export function formatDistance(distanceKm) {
  if (distanceKm < 1) {
    return `${Math.round(distanceKm * 1000)}m`
  }
  return `${distanceKm.toFixed(1)}km`
}

/**
 * formatTime 格式化时间为 MM月DD日 HH:mm
 * @param {string} timeStr - ISO 时间字符串
 * @returns {string} 格式化后的时间字符串
 */
export function formatTime(timeStr) {
  if (!timeStr) return ''
  const date = new Date(timeStr)
  const month = String(date.getMonth() + 1).padStart(2, '0')
  const day = String(date.getDate()).padStart(2, '0')
  const hours = String(date.getHours()).padStart(2, '0')
  const minutes = String(date.getMinutes()).padStart(2, '0')
  return `${month}月${day}日 ${hours}:${minutes}`
}
