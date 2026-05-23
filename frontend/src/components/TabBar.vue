<template>
  <van-tabbar v-model="active" active-color="#1989fa" inactive-color="#646566">
    <van-tabbar-item icon="home-o" to="/">首页</van-tabbar-item>
    <van-tabbar-item icon="add-o" to="/create">发起约球</van-tabbar-item>
    <van-tabbar-item icon="user-o" to="/profile">我的</van-tabbar-item>
  </van-tabbar>
</template>

<script setup>
import { ref, watch } from 'vue'
import { useRoute } from 'vue-router'

const route = useRoute()
const active = ref(0)

// 根据当前路由设置 active tab
const routeMap = { '/': 0, '/create': 1, '/profile': 2 }

watch(
  () => route.path,
  (path) => {
    if (path in routeMap) {
      active.value = routeMap[path]
    }
  },
  { immediate: true }
)
</script>
