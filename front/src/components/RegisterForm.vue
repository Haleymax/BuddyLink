<template>
  <n-card class="auth-card" :bordered="false">
    <n-form 
      ref="formRef" 
      :model="formData" 
      :rules="rules"
      :show-label="false"
      size="large"
    >
      <n-form-item path="username">
        <n-input 
          v-model:value="formData.username" 
          placeholder="用户名"
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
                <path fill="currentColor" d="M20 4H4c-1.1 0-1.99.9-1.99 2L2 18c0 1.1.9 2 2 2h16c1.1 0 2-.9 2-2V6c0-1.1-.9-2-2-2zm0 4l-8 5-8-5V6l8 5 8-5v2z"/>
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
          @keydown.enter.prevent
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

      <n-form-item path="confirmPassword">
        <n-input 
          v-model:value="formData.confirmPassword" 
          type="password"
          placeholder="确认密码"
          size="large"
          show-password-on="click"
          @keydown.enter.prevent="handleSubmit"
        >
          <template #prefix>
            <n-icon color="rgba(255, 255, 255, 0.7)">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 1L3 5v6c0 5.55 3.84 10.74 9 12 5.16-1.26 9-6.45 9-12V5l-9-4zm0 10.99h7c-.53 4.12-3.28 7.79-7 8.94V12H5V6.3l7-3.11v8.8z"/>
              </svg>
            </n-icon>
          </template>
        </n-input>
      </n-form-item>

      <n-form-item path="verification_code">
        <n-input 
          v-model:value="formData.verification_code" 
          placeholder="请输入邮箱验证码"
          size="large"
          @keydown.enter.prevent="handleSubmit"
        >
          <template #prefix>
            <n-icon color="rgba(255, 255, 255, 0.7)">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z"/>
              </svg>
            </n-icon>
          </template>
          <template #suffix>
            <n-button 
              text 
              type="primary" 
              size="small"
              :loading="sendingCode"
              :disabled="!formData.email || countdown > 0"
              @click="handleSendCode"
            >
              {{ countdown > 0 ? `${countdown}s` : '获取验证码' }}
            </n-button>
          </template>
        </n-input>
      </n-form-item>

      <div class="form-options">
        <n-checkbox v-model:checked="agreeTerms" size="small">
          <span>我同意</span>
          <n-button text type="primary" size="small" style="padding: 0 4px;">
            服务条款
          </n-button>
          <span>和</span>
          <n-button text type="primary" size="small" style="padding: 0 4px;">
            隐私政策
          </n-button>
        </n-checkbox>
      </div>

      <n-button 
        type="primary" 
        size="large"
        block
        round
        :loading="loading"
        :disabled="!agreeTerms"
        @click="handleSubmit"
        class="submit-button"
      >
        <template #icon>
          <n-icon>
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z"/>
            </svg>
          </n-icon>
        </template>
        创建账户
      </n-button>
    </n-form>

    <!-- 分割线 -->
    <n-divider class="divider">
      或者
    </n-divider>

    <!-- 社交注册 -->
    <div class="social-login">
      <n-button quaternary circle size="large" class="social-btn">
        <template #icon>
          <n-icon size="20" color="#1890ff">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M22.56 12.25c0-.78-.07-1.53-.2-2.25H12v4.26h5.92c-.26 1.37-1.04 2.53-2.21 3.31v2.77h3.57c2.08-1.92 3.28-4.74 3.28-8.09z"/>
              <path fill="currentColor" d="M12 23c2.97 0 5.46-.98 7.28-2.66l-3.57-2.77c-.98.66-2.23 1.06-3.71 1.06-2.86 0-5.29-1.93-6.16-4.53H2.18v2.84C3.99 20.53 7.7 23 12 23z"/>
              <path fill="currentColor" d="M5.84 14.09c-.22-.66-.35-1.36-.35-2.09s.13-1.43.35-2.09V7.07H2.18C1.43 8.55 1 10.22 1 12s.43 3.45 1.18 4.93l2.85-2.22.81-.62z"/>
              <path fill="currentColor" d="M12 5.38c1.62 0 3.06.56 4.21 1.64l3.15-3.15C17.45 2.09 14.97 1 12 1 7.7 1 3.99 3.47 2.18 7.07l3.66 2.84c.87-2.6 3.3-4.53 6.16-4.53z"/>
            </svg>
          </n-icon>
        </template>
      </n-button>
      
      <n-button quaternary circle size="large" class="social-btn">
        <template #icon>
          <n-icon size="20" color="#07c160">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M8.5 12c0 1.5.5 3 1.5 4s2.5 1.5 4 1.5 3-.5 4-1.5 1.5-2.5 1.5-4-.5-3-1.5-4-2.5-1.5-4-1.5-3 .5-4 1.5-1.5 2.5-1.5 4zm6.5-8c-4.4 0-8 3.6-8 8s3.6 8 8 8 8-3.6 8-8-3.6-8-8-8z"/>
            </svg>
          </n-icon>
        </template>
      </n-button>
    </div>

    <!-- 切换到登录 -->
    <div class="switch-link">
      <span>已有账户？</span>
      <n-button text type="primary" @click="$emit('switch-form')">
        立即登录
      </n-button>
    </div>
  </n-card>
</template>

<script lang="ts" setup>
import { reactive, ref, onUnmounted } from 'vue';
import { useMessage, type FormInst, type FormRules } from 'naive-ui';
import type { RegisterFormData } from '../model/auth';
import { register, sendVerificationCode } from '../api/auth';

defineEmits<{
  'switch-form': []
}>();

const message = useMessage();
const formRef = ref<FormInst | null>(null);
const loading = ref(false);
const agreeTerms = ref(false);
const sendingCode = ref(false);
const countdown = ref(0);
let countdownTimer: number | null = null;

// 统一的表单数据结构
const formData = reactive<RegisterFormData & { confirmPassword: string }>({
  username: '',
  email: '',
  password: '',
  confirmPassword: '',
  verification_code: ''
});

// 优化后的验证规则
const rules: FormRules = {
  username: [
    {
      required: true,
      message: '请输入用户名',
      trigger: ['blur', 'input']
    },
    {
      min: 3,
      max: 20,
      message: '用户名长度应在 3-20 个字符之间',
      trigger: ['blur', 'input']
    },
    {
      pattern: /^[a-zA-Z0-9_\u4e00-\u9fa5]+$/,
      message: '用户名只能包含字母、数字、下划线和中文',
      trigger: ['blur', 'input']
    }
  ],
  email: [
    {
      required: true,
      message: '请输入邮箱地址',
      trigger: ['blur', 'input']
    },
    {
      type: 'email',
      message: '请输入有效的邮箱地址',
      trigger: ['blur', 'input']
    },
    {
      pattern: /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/,
      message: '邮箱格式不正确',
      trigger: ['blur', 'input']
    }
  ],
  password: [
    {
      required: true,
      message: '请输入密码',
      trigger: ['blur', 'input']
    },
    {
      min: 6,
      max: 20,
      message: '密码长度应在 6-20 个字符之间',
      trigger: ['blur', 'input']
    }
  ],
  verification_code: [
    {
      required: true,
      message: '请输入验证码',
      trigger: ['blur', 'input']
    },
    {
      min: 4,
      max: 8,
      message: '验证码长度应在 4-8 个字符之间',
      trigger: ['blur', 'input']
    }
  ],
  confirmPassword: [
    {
      required: true,
      message: '请确认密码',
      trigger: ['blur', 'input']
    },
    {
      validator: (_rule: any, value: string) => {
        if (value !== formData.password) {
          return new Error('两次输入的密码不一致');
        }
        return true;
      },
      trigger: ['blur', 'input']
    }
  ]
};

// 发送验证码
const handleSendCode = async () => {
  if (!formData.email) {
    message.error('请先输入邮箱地址');
    return;
  }

  // 验证邮箱格式
  const emailRegex = /^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$/;
  if (!emailRegex.test(formData.email)) {
    message.error('请输入有效的邮箱地址');
    return;
  }

  try {
    sendingCode.value = true;
    const res = await sendVerificationCode(formData.email);
    
    if (res.status) {
      message.success('验证码已发送到您的邮箱，请查收');
      
      // 开始倒计时（60秒）
      countdown.value = 60;
      countdownTimer = setInterval(() => {
        countdown.value--;
        if (countdown.value <= 0) {
          if (countdownTimer) {
            clearInterval(countdownTimer);
            countdownTimer = null;
          }
        }
      }, 1000);
    } else {
      message.error('验证码发送失败，请稍后重试');
    }
  } catch (error) {
    console.error('发送验证码错误:', error);
    message.error('验证码发送失败，请稍后重试');
  } finally {
    sendingCode.value = false;
  }
};

const handleSubmit = async () => {
  if (!agreeTerms.value) {
    message.warning('请同意服务条款和隐私政策');
    return;
  }
  
  if (!formRef.value) return;

  try {
    // 验证表单
    await formRef.value.validate();
    
    loading.value = true;

    // 准备注册数据（不包含 confirmPassword）
    const registerData: RegisterFormData = {
      username: formData.username,
      email: formData.email,
      password: formData.password,
      verification_code: formData.verification_code
    };

    const res = await register(registerData);
    if (res.status) {
      // 清空表单
      Object.assign(formData, {
        username: '',
        email: '',
        password: '',
        confirmPassword: '',
        verification_code: ''
      });
      agreeTerms.value = false;
      
      message.success('注册成功，请登录');
      // 在这里可以添加注册成功后的逻辑，比如跳转到登录页面
    } else {
      message.error('注册失败，请稍后再试');
    }
  } catch (error) {
    console.error('注册错误:', error);
    if (error instanceof Array) {
      // 表单验证错误
      message.error('请检查输入信息');
    } else {
      message.error('注册失败，请稍后再试');
    }
  } finally {
    loading.value = false;
  }
};

// 组件销毁时清理定时器
onUnmounted(() => {
  if (countdownTimer) {
    clearInterval(countdownTimer);
    countdownTimer = null;
  }
});
</script>

<style scoped>
/* 表单样式 */
.auth-card :deep(.n-form-item) {
  margin-bottom: 24px;
}

.auth-card :deep(.n-input) {
  border-radius: 16px;
  border: 1px solid rgba(255, 255, 255, 0.2);
  background: rgba(255, 255, 255, 0.1);
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
}

.auth-card :deep(.n-input .n-input__input-el) {
  background: transparent;
  color: white;
  padding-left: 50px;
  padding-right: 16px;
  text-indent: 0;
  line-height: 1.5;
  font-size: 16px;
}

.auth-card :deep(.n-input .n-input__input-el::placeholder) {
  color: rgba(255, 255, 255, 0.6);
  transition: all 0.3s ease;
  transform: translateY(0);
}

.auth-card :deep(.n-input .n-input__input-el:focus::placeholder) {
  opacity: 0.4;
  transform: translateY(-2px);
}

.auth-card :deep(.n-input:hover) {
  border-color: rgba(255, 255, 255, 0.4);
  background: rgba(255, 255, 255, 0.15);
  box-shadow: 0 0 0 2px rgba(255, 255, 255, 0.1);
}

.auth-card :deep(.n-input.n-input--focus) {
  border-color: rgba(255, 255, 255, 0.6);
  background: rgba(255, 255, 255, 0.2);
  box-shadow: 0 0 0 2px rgba(255, 255, 255, 0.2);
}

.auth-card :deep(.n-input .n-input__prefix) {
  padding-left: 16px;
  padding-right: 10px;
  width: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
}

.auth-card :deep(.n-input .n-input__prefix .n-icon) {
  color: rgba(255, 255, 255, 0.7) !important;
}

/* 表单选项 */
.form-options {
  margin-bottom: 32px;
}

.form-options :deep(.n-checkbox .n-checkbox__label) {
  color: rgba(255, 255, 255, 0.9);
}

.form-options :deep(.n-checkbox .n-checkbox__label .n-button) {
  color: rgba(255, 255, 255, 0.8);
}

.form-options :deep(.n-checkbox .n-checkbox__label .n-button:hover) {
  color: white;
}

/* 提交按钮 */
.submit-button {
  height: 48px;
  font-size: 16px;
  font-weight: 600;
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.2), rgba(255, 255, 255, 0.1));
  border: 1px solid rgba(255, 255, 255, 0.3);
  color: white;
  backdrop-filter: blur(10px);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.15);
  transition: all 0.3s ease;
  margin-bottom: 24px;
}

.submit-button:hover {
  transform: translateY(-2px);
  background: linear-gradient(135deg, rgba(255, 255, 255, 0.3), rgba(255, 255, 255, 0.2));
  border-color: rgba(255, 255, 255, 0.4);
  box-shadow: 0 12px 30px rgba(0, 0, 0, 0.2);
}

.submit-button:active {
  transform: translateY(0px);
}

.submit-button:disabled {
  opacity: 0.5;
  cursor: not-allowed;
  transform: none;
}

/* 分割线 */
.divider {
  margin: 24px 0;
}

.divider :deep(.n-divider__title) {
  color: rgba(255, 255, 255, 0.8);
  font-size: 14px;
}

.divider :deep(.n-divider__line) {
  background-color: rgba(255, 255, 255, 0.2);
}

/* 社交登录 */
.social-login {
  display: flex;
  justify-content: center;
  gap: 16px;
  margin-bottom: 24px;
}

.social-btn {
  width: 48px;
  height: 48px;
  background: rgba(255, 255, 255, 0.1);
  border: 1px solid rgba(255, 255, 255, 0.2);
  backdrop-filter: blur(10px);
  transition: all 0.3s ease;
}

.social-btn:hover {
  transform: translateY(-2px);
  background: rgba(255, 255, 255, 0.15);
  border-color: rgba(255, 255, 255, 0.3);
  box-shadow: 0 8px 20px rgba(0, 0, 0, 0.15);
}

/* 切换链接 */
.switch-link {
  text-align: center;
  color: rgba(255, 255, 255, 0.8);
  font-size: 14px;
}

.switch-link span {
  margin-right: 8px;
}

.switch-link :deep(.n-button) {
  color: white;
  font-weight: 500;
}
</style>
