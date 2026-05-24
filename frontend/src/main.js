import { createApp } from 'vue'
import { createPinia } from 'pinia'
import router from './router'
import App from './App.vue'
import { useUserStore } from './stores/user'

// 全局引入 Vant 样式
import 'vant/lib/index.css'

const app = createApp(App)
app.use(createPinia())
app.use(router)

// 应用启动时自动拉取用户信息（若 token 存在但 userInfo 缺失）
const userStore = useUserStore()
userStore.fetchUserIfNeeded()

app.mount('#app')
