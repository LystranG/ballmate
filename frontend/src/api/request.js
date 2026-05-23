import axios from 'axios'
import { showToast } from 'vant'
import router from '@/router'

const request = axios.create({
  baseURL: '/api/v1',
  timeout: 10000
})

// 请求拦截器：注入 token
request.interceptors.request.use(config => {
  const token = localStorage.getItem('token')
  if (token) {
    config.headers.Authorization = `Bearer ${token}`
  }
  return config
})

// 响应拦截器：统一错误处理
request.interceptors.response.use(
  response => response.data,
  error => {
    if (error.response?.status === 401) {
      localStorage.removeItem('token')
      router.push('/login')
      showToast('登录已过期，请重新登录')
    } else {
      showToast(error.response?.data?.message || '网络错误')
    }
    return Promise.reject(error)
  }
)

export default request
