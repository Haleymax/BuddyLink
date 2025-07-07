import { createRouter, createWebHistory, RouteRecordRaw } from 'vue-router'

// 定义路由
const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
]

// 创建路由实例
const router = createRouter({
  history: createWebHistory('/'),
  routes
})

export default router