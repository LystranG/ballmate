import request from './request'

// 发起邀约
export function createInvitation(data) {
  return request.post('/invitations', data)
}

// 获取邀约详情
export function getInvitation(id) {
  return request.get(`/invitations/${id}`)
}

// 终止邀约
export function terminateInvitation(id) {
  return request.put(`/invitations/${id}/terminate`)
}

// 删除邀约
export function deleteInvitation(id) {
  return request.delete(`/invitations/${id}`)
}

// 获取邀约参与者列表
export function getParticipants(id) {
  return request.get(`/invitations/${id}/participants`)
}
