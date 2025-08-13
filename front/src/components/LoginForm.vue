<template>
  <n-card class="auth-card" :bordered="false">
    <n-form 
      ref="formRef" 
      :model="formData" 
      :rules="rules"
      :show-label="false"
      size="large"
    >
      <n-form-item path="email">
        <n-input 
          v-model:value="formData.email" 
          placeholder="邮箱地址"
          size="large"
          @keydown.enter.prevent
        >
          <template #prefix>
            <n-icon color="rgba(255, 255, 255, 0.7)">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
              </svg>
            </n-icon>
          </template>
        </n-input>
      </n-form-item>

      <n-form-item path="password">
        <n-input 
          v-model:value="formData.password" 
          type="password"
          placeholder="密码"
          size="large"
          show-password-on="click"
          @keydown.enter.prevent="handleSubmit"
        >
          <template #prefix>
            <n-icon color="rgba(255, 255, 255, 0.7)">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M18 8h-1V6c0-2.76-2.24-5-5-5S7 3.24 7 6v2H6c-1.1 0-2 .9-2 2v10c0 1.1.9 2 2 2h12c1.1 0 2-.9 2-2V10c0-1.1-.9-2-2-2zm-6 9c-1.1 0-2-.9-2-2s.9-2 2-2 2 .9 2 2-.9 2-2 2zm3.1-9H8.9V6c0-1.71 1.39-3.1 3.1-3.1 1.71 0 3.1 1.39 3.1 3.1v2z"/>
              </svg>
            </n-icon>
          </template>
        </n-input>
      </n-form-item>

      <div class="form-options">
        <n-checkbox v-model:checked="rememberMe" size="small">
          记住我
        </n-checkbox>
        <n-button text type="primary" size="small">
          忘记密码？
        </n-button>
      </div>

      <n-button 
        type="primary" 
        size="large"
        block
        round
        :loading="loading"
        @click="handleSubmit"
        class="submit-button"
      >
        <template #icon>
          <n-icon>
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M10.09 15.59L11.5 17l5-5-5-5-1.41 1.41L12.67 11H3v2h9.67l-2.58 2.59zM19 3H5c-1.11 0-2 .9-2 2v4h2V5h14v14H5v-4H3v4c0 1.1.89 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2z"/>
            </svg>
          </n-icon>
        </template>
        登录
      </n-button>
    </n-form>

    <!-- 切换到注册 -->
    <div class="switch-link">
      <span>还没有账户？</span>
      <n-button text type="primary" @click="$emit('switch-form')">
        立即注册
      </n-button>
    </div>
  </n-card>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue';
import { c, useMessage, type FormInst, type FormRules } from 'naive-ui';
import type { LoginFormData } from '../model/auth';
import { login } from '../api/auth';
import '../styles/LoginForm.css';
import { useAuthStore } from '../stores/auth.store';
import router from '../router';

defineEmits<{
  'switch-form': []
}>();

const message = useMessage();
const formRef = ref<FormInst | null>(null);
const loading = ref(false);
const rememberMe = ref(false);
const auth = useAuthStore();

const formData = reactive<LoginFormData>({
  email: '',
  password: ''
});

const rules: FormRules = {
  email: [
    { required: true, message: '请输入邮箱地址', trigger: ['blur', 'input'] },
    { type: 'email', message: '请输入有效的邮箱地址', trigger: ['blur', 'input'] }
  ],
  password: [
    { required: true, message: '请输入密码', trigger: ['blur', 'input'] },
    { min: 6, max: 20, message: '密码长度应在 6-20 个字符之间', trigger: ['blur', 'input'] }
  ]
};

const handleSubmit = async () => {
  if (!formRef.value) return;
  
  try {
    await formRef.value.validate();
    
    loading.value = true;
    console.log('登录数据:', formData);
    
    const res = await login(formData);

    console.log('登录响应:', res);
    

    if (res.status) {
      message.success('登录成功');
      const token = res.token;
      console.log('Token:', token);
      auth.SetToken(token);
      if (rememberMe.value) {
        localStorage.setItem('rememberMe', 'true');
      }
      // 跳转到 /home （路由里 name 定义为 'Home'）
      await router.push({ name: 'Home' });
      return;
    } else {
      message.error('登录失败，请重试');
    }
  } catch (error) {
    console.error('登录错误:', error);
    if (error instanceof Array) {
      // 表单验证错误
      message.error('请检查输入信息');
    } else {
      message.error('登录失败，请重试');
    }
  } finally {
    loading.value = false;
  }
};
</script>
