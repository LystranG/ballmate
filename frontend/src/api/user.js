import request from './request'

// 用户登录
export function login(data) {
  return request.post('/auth/login', data)
}

// 用户注册
export function register(data) {
  return request.post('/auth/register', data)
}

// 获取个人信息
export function getProfile() {
  return request.get('/user/profile')
}

// 更新个人资料
export function updateProfile(data) {
  return request.put('/user/profile', data)
}

// 修改密码
export function updatePassword(data) {
  return request.put('/user/password', data)
}

// 上传头像
export function uploadAvatar(formData) {
  return request.post('/user/avatar', formData)
}
