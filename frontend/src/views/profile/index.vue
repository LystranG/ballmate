<template>
  <div class="profile-page">
    <div class="content">
      <p>个人中心（开发中）</p>
    </div>
    <TabBar />
  </div>
</template>

<script setup>
import TabBar from '@/components/TabBar.vue'
</script>

<style scoped>
.profile-page {
  padding-bottom: 50px;
}

.content {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 60vh;
  color: #969799;
  font-size: 16px;
}
</style>
