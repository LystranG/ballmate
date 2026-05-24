<template>
  <div class="invitation-detail">
    <van-nav-bar title="邀约详情" left-arrow @click-left="$router.back()" />
    <div class="placeholder">
      <van-loading size="24px">加载中...</van-loading>
    </div>
  </div>
</template>

<script setup>
// 邀约详情页 — 后续 plan 实现完整功能
</script>

<style scoped>
.invitation-detail {
  min-height: 100vh;
  background: #f7f8fa;
}

.placeholder {
  display: flex;
  justify-content: center;
  padding-top: 40vh;
}
</style>
