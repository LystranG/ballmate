import { defineStore } from 'pinia'
import { ref, computed } from 'vue'
import { getProfile } from '@/api/user'

export const useUserStore = defineStore('user', () => {
  const token = ref(localStorage.getItem('token') || '')
  const userInfo = ref(JSON.parse(localStorage.getItem('userInfo') || 'null'))

  const isLoggedIn = computed(() => !!token.value)

  // 设置 token 并同步到 localStorage
  function setToken(newToken) {
    token.value = newToken
    localStorage.setItem('token', newToken)
  }

  // 设置用户信息并同步到 localStorage
  function setUserInfo(info) {
    userInfo.value = info
    localStorage.setItem('userInfo', JSON.stringify(info))
  }

  // 应用初始化时，若有 token 但无 userInfo 则自动拉取
  async function fetchUserIfNeeded() {
    if (token.value && !userInfo.value) {
      try {
        const res = await getProfile()
        setUserInfo(res.data)
      } catch (e) {
        // token 失效时忽略，401 拦截器会处理
      }
    }
  }

  // 退出登录
  function logout() {
    token.value = ''
    userInfo.value = null
    localStorage.removeItem('token')
    localStorage.removeItem('userInfo')
  }

  return { token, userInfo, isLoggedIn, setToken, setUserInfo, fetchUserIfNeeded, logout }
})
