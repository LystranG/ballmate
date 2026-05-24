import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/login',
    component: () => import('@/views/login/index.vue'),
    meta: { guest: true }
  },
  {
    path: '/register',
    component: () => import('@/views/register/index.vue'),
    meta: { guest: true }
  },
  {
    path: '/',
    component: () => import('@/views/home/index.vue'),
    meta: { auth: true }
  },
  {
    path: '/profile',
    component: () => import('@/views/profile/index.vue'),
    meta: { auth: true }
  },
  {
    path: '/profile/edit',
    component: () => import('@/views/profile/edit.vue'),
    meta: { auth: true }
  },
  {
    path: '/profile/password',
    component: () => import('@/views/profile/password.vue'),
    meta: { auth: true }
  },
  {
    path: '/invitation/:id',
    component: () => import('@/views/invitation/detail.vue'),
    meta: { auth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 导航守卫
router.beforeEach((to, from, next) => {
  const token = localStorage.getItem('token')
  if (to.meta.auth && !token) {
    next('/login')
  } else if (to.meta.guest && token) {
    next('/')
  } else {
    next()
  }
})

export default router
