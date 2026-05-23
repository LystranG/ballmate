<template>
  <div class="home-page">
    <div class="content">
      <p>首页（开发中）</p>
    </div>
    <TabBar />
  </div>
</template>

<script setup>
import TabBar from '@/components/TabBar.vue'
</script>

<style scoped>
.home-page {
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
