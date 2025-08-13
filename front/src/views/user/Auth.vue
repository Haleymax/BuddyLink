<template>
  <div class="auth-container">
    <div class="background-decoration">
      <div class="circle circle-1"></div>
      <div class="circle circle-2"></div>
      <div class="circle circle-3"></div>
    </div>
    <div class="auth-wrapper">
      <div class="auth-content">
        <!-- Logo和标题 -->
        <div class="auth-header">
          <div class="logo">
            <n-icon size="48" color="rgba(255, 255, 255, 0.9)">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
              </svg>
            </n-icon>
          </div>
          <h1 class="title">BuddyLink</h1>
          <p class="subtitle">{{ isLogin ? '欢迎回来，请登录您的账户' : '创建您的新账户，开启精彩旅程' }}</p>
        </div>

        <!-- 表单切换 -->
        <div class="form-container">
          <transition name="fade" mode="out-in">
            <LoginForm v-if="isLogin" @switch-form="switchToRegister" />
            <RegisterForm v-else @switch-form="switchToLogin" />
          </transition>
        </div>
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup name="Auth">
import { ref, onMounted } from 'vue';
import { useRoute, useRouter } from 'vue-router';
import LoginForm from '../../components/LoginForm.vue';
import RegisterForm from '../../components/RegisterForm.vue';
import '@/styles/Auth.css';

const route = useRoute();
const router = useRouter();
const isLogin = ref(true);

onMounted(() => {
  // 根据路由决定显示登录还是注册表单
  isLogin.value = route.path === '/login';
});

const switchToLogin = () => {
  isLogin.value = true;
  router.push('/login');
};

const switchToRegister = () => {
  isLogin.value = false;
  router.push('/register');
};
</script>


