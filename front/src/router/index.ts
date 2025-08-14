import { createRouter, createWebHistory, type RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '../stores/auth.store'

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
    meta: { requiresAuth: true }
  }
]

// 创建路由实例
const router = createRouter({
  history: createWebHistory('/'),
  routes
})

// 路由守卫
router.beforeEach((to, _from, next) => {
  const authStore = useAuthStore()
  
  console.log('Router guard - to:', to.path, 'token:', authStore.token)
  
  if (to.meta.requiresAuth && !authStore.token) {
    console.log('Route requires auth but no token found, redirecting to login')
    next('/login')
  } else if (to.path === '/login' && authStore.token) {
    console.log('Already authenticated, redirecting to home')
    next('/home')
  } else {
    next()
  }
})

export default router