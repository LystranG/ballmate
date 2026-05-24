<template>
  <div class="my-invitations-page">
    <van-nav-bar title="我的邀约" left-arrow @click-left="router.back()" />

    <!-- Tab 切换：我发起的 / 我参与的 -->
    <van-tabs v-model:active="activeTab" @change="onTabChange">
      <!-- Tab 0：我发起的 -->
      <van-tab title="我发起的">
        <van-list
          v-model:loading="createdState.loading"
          :finished="createdState.finished"
          finished-text="没有更多了"
          @load="loadCreated"
        >
          <InvitationCard
            v-for="item in createdState.list"
            :key="item.id"
            :invitation="item"
            mode="status"
            @click="goDetail(item.id)"
          />
          <!-- 空状态 -->
          <van-empty
            v-if="!createdState.loading && createdState.list.length === 0 && createdState.finished"
            description="暂无发起的邀约"
          />
        </van-list>
      </van-tab>

      <!-- Tab 1：我参与的 -->
      <van-tab title="我参与的">
        <van-list
          v-model:loading="joinedState.loading"
          :finished="joinedState.finished"
          finished-text="没有更多了"
          @load="loadJoined"
        >
          <InvitationCard
            v-for="item in joinedState.list"
            :key="item.id"
            :invitation="item"
            mode="status"
            @click="goDetail(item.id)"
          />
          <!-- 空状态 -->
          <van-empty
            v-if="!joinedState.loading && joinedState.list.length === 0 && joinedState.finished"
            description="暂无参与的邀约"
          />
        </van-list>
      </van-tab>
    </van-tabs>
  </div>
</template>

<script setup>
import { reactive, ref } from 'vue'
import { useRouter } from 'vue-router'
import { showToast } from 'vant'
import InvitationCard from '@/components/InvitationCard.vue'
import { getMyInvitations } from '@/api/invitation'

const router = useRouter()
const activeTab = ref(0)

// 每个 Tab 独立维护自己的列表状态，切换时不重置对方数据
const createdState = reactive({
  list: [],
  loading: false,
  finished: false,
  page: 1
})

const joinedState = reactive({
  list: [],
  loading: false,
  finished: false,
  page: 1
})

// 加载"我发起的"邀约
async function loadCreated() {
  try {
    const res = await getMyInvitations({ type: 'created', page: createdState.page, page_size: 10 })
    const { list, total } = res.data
    createdState.list.push(...list)
    createdState.page++
    // 已加载全部数据则标记 finished
    if (createdState.list.length >= total) {
      createdState.finished = true
    }
  } catch (e) {
    showToast('加载失败，请重试')
    createdState.finished = true
  } finally {
    createdState.loading = false
  }
}

// 加载"我参与的"邀约
async function loadJoined() {
  try {
    const res = await getMyInvitations({ type: 'joined', page: joinedState.page, page_size: 10 })
    const { list, total } = res.data
    joinedState.list.push(...list)
    joinedState.page++
    if (joinedState.list.length >= total) {
      joinedState.finished = true
    }
  } catch (e) {
    showToast('加载失败，请重试')
    joinedState.finished = true
  } finally {
    joinedState.loading = false
  }
}

// Tab 切换时不重置另一个 Tab 的数据（Vant Tabs lazy-render 默认行为）
function onTabChange() {
  // 无需额外处理，各 Tab 的 van-list 独立管理自身加载状态
}

// 跳转邀约详情
function goDetail(id) {
  router.push(`/invitation/${id}`)
}
</script>

<style scoped>
.my-invitations-page {
  min-height: 100vh;
  background-color: #f7f8fa;
}
</style>
