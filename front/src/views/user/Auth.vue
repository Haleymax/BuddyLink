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

<style scoped>
.auth-container {
  min-height: 100vh;
  width: 100vw;
  position: fixed;
  top: 0;
  left: 0;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 50%, #18a058 100%);
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
}

/* 背景装饰 */
.background-decoration {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  pointer-events: none;
  overflow: hidden;
}

.background-decoration::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: radial-gradient(circle at 20% 80%, rgba(120, 119, 198, 0.3) 0%, transparent 50%),
              radial-gradient(circle at 80% 20%, rgba(255, 119, 198, 0.3) 0%, transparent 50%),
              radial-gradient(circle at 40% 40%, rgba(120, 219, 226, 0.3) 0%, transparent 50%);
  animation: backgroundShift 15s ease-in-out infinite;
}

@keyframes backgroundShift {
  0%, 100% {
    transform: scale(1) rotate(0deg);
    opacity: 0.6;
  }
  50% {
    transform: scale(1.1) rotate(180deg);
    opacity: 0.8;
  }
}

.circle {
  position: absolute;
  border-radius: 50%;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.15), rgba(255, 255, 255, 0.05));
  animation: float 8s ease-in-out infinite;
  backdrop-filter: blur(40px);
  border: 1px solid rgba(255, 255, 255, 0.1);
}

.circle-1 {
  width: 300px;
  height: 300px;
  top: -50px;
  left: -50px;
  animation-delay: 0s;
}

.circle-2 {
  width: 200px;
  height: 200px;
  top: 40%;
  right: -30px;
  animation-delay: 3s;
}

.circle-3 {
  width: 150px;
  height: 150px;
  bottom: -30px;
  left: 50%;
  transform: translateX(-50%);
  animation-delay: 6s;
}

/* 添加更多装饰圆圈 */
.background-decoration::after {
  content: '';
  position: absolute;
  width: 120px;
  height: 120px;
  border-radius: 50%;
  background: linear-gradient(45deg, rgba(255, 255, 255, 0.1), rgba(255, 255, 255, 0.05));
  top: 20%;
  right: 20%;
  animation: float 10s ease-in-out infinite reverse;
  backdrop-filter: blur(30px);
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px) scale(1);
    opacity: 0.6;
  }
  33% {
    transform: translateY(-30px) scale(1.05);
    opacity: 0.8;
  }
  66% {
    transform: translateY(-15px) scale(0.95);
    opacity: 0.7;
  }
}

/* 认证区域 */
.auth-wrapper {
  width: 100%;
  max-width: 450px;
  padding: 20px;
  z-index: 1;
  position: relative;
}

.auth-content {
  width: 100%;
}

/* 头部区域 */
.auth-header {
  text-align: center;
  margin-bottom: 40px;
  color: white;
}

.logo {
  margin-bottom: 16px;
  display: flex;
  justify-content: center;
}

.title {
  font-size: 32px;
  font-weight: 700;
  margin: 0 0 8px 0;
  background: linear-gradient(45deg, #fff, #f0f0f0);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.subtitle {
  font-size: 16px;
  margin: 0;
  opacity: 0.9;
  font-weight: 300;
  transition: all 0.3s ease;
}

/* 表单容器 */
.form-container {
  position: relative;
}

/* 表单卡片样式 */
.form-container :deep(.auth-card) {
  backdrop-filter: blur(25px);
  background: rgba(255, 255, 255, 0.15);
  border-radius: 24px;
  box-shadow: 0 25px 50px rgba(0, 0, 0, 0.15), 
              0 0 0 1px rgba(255, 255, 255, 0.2),
              inset 0 1px 0 rgba(255, 255, 255, 0.3);
  border: 1px solid rgba(255, 255, 255, 0.2);
  padding: 40px 32px;
  position: relative;
  overflow: hidden;
}

.form-container :deep(.auth-card::before) {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 1px;
  background: linear-gradient(90deg, transparent, rgba(255, 255, 255, 0.4), transparent);
}

/* 过渡动画 */
.fade-enter-active,
.fade-leave-active {
  transition: all 0.4s ease;
}

.fade-enter-from {
  opacity: 0;
  transform: translateY(20px) scale(0.95);
}

.fade-leave-to {
  opacity: 0;
  transform: translateY(-20px) scale(0.95);
}

/* 响应式设计 */
@media (max-width: 768px) {
  .auth-container {
    padding: 20px;
  }
  
  .auth-wrapper {
    max-width: 100%;
    padding: 0;
  }
  
  .form-container :deep(.auth-card) {
    padding: 32px 24px;
  }
  
  .title {
    font-size: 28px;
  }
  
  .subtitle {
    font-size: 14px;
  }
  
  .circle {
    display: none;
  }
  
  .background-decoration::after {
    display: none;
  }
}

@media (max-width: 480px) {
  .form-container :deep(.auth-card) {
    padding: 24px 20px;
    margin: 0 10px;
  }
  
  .title {
    font-size: 24px;
  }
}

/* 深色模式适配 */
@media (prefers-color-scheme: dark) {
  .auth-container {
    background: linear-gradient(135deg, #1a1a2e 0%, #16213e 50%, #0f3460 100%);
  }
  
  .form-container :deep(.auth-card) {
    background: rgba(255, 255, 255, 0.08);
    border: 1px solid rgba(255, 255, 255, 0.15);
  }
}
</style>
