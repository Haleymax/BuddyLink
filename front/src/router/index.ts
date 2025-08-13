import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'

// 定义路由
const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    redirect: '/login'   // 默认重定向到登陆页
  },
  {
    path: '/login',
    name: 'Login',
    component: ()=> import('@/views/user/Auth.vue')
  },
  {
    path: '/register',
    name: 'Register',
    component: ()=> import('@/views/user/Auth.vue')
  },
  {
    path: '/home',
    name: 'Home',
    component: () => import('@/views/Home.vue'),
  }
]

// 创建路由实例
const router = createRouter({
  history: createWebHistory('/'),
  routes
})

export default router