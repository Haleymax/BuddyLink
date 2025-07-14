<template>
  <n-layout class="login-container">
    <n-layout-content>
      <n-space justify="center" align="center" class="">
        <n-card title="用户登录" style="width: 400px;">
          <n-form ref="formRef" :model="formDate" :rules="rules">
            <n-form-item label="用户名" path="username">
              <n-input 
                v-model:value="formDate.username" 
                placeholder="请输入用户名" 
                @keydown.enter.prevent
              />
            </n-form-item>

            <n-form-item label="密码" path="password">
              <n-input 
                v-model:value="formDate.password" 
                type="password" 
                placeholder="请输入密码" 
                @keydown.enter.prevent
              />
            </n-form-item>

            <n-form-item>
              <n-checkbox v-model:checked="rememberMe">记住我</n-checkbox>
            </n-form-item>

            <n-button 
              type="primary" 
              block
              :loading="loading"
              @click="handleLogin"
            >
              登录
            </n-button>
          </n-form>
        </n-card>
      </n-space>
    </n-layout-content>
  </n-layout>
</template>

<script lang="ts" name="Login">
import { defineComponent, ref } from 'vue';
import { useMessage, type FormInst } from 'naive-ui';

export default defineComponent({
  setup() {
    const message = useMessage();
    const formRef = ref<FormInst | null>(null);
    const loading = ref(false);
    const rememberMe = ref(false);

    const formDate = ref({
      username: '',
      password: ''
    });

    const rules = {
      username: [
        { required: true, message: '请输入用户名', trigger: 'blur' },
        { min: 3, max: 20, message: '用户名长度在3到20个字符', trigger: 'blur' }
      ],
      password: [
        { required: true, message: '请输入密码', trigger: 'blur' },
        { min: 6, max: 20, message: '密码长度在6到20个字符', trigger: 'blur' }
      ]
    };

    const handleLogin = async () => {
      try {
        await formRef.value?.validate();
        loading.value = true;
        // 模拟登录请求
        await new Promise((resolve) => setTimeout(resolve, 1000));
        message.success('登录成功');
        // 在这里处理登录成功后的逻辑，比如跳转到首页
      } catch (error) {
        message.error('登录失败，请重试');
      } finally {
        loading.value = false;
      }
    };

    return {
      formRef,
      loading,
      rememberMe,
      formDate,
      rules,
      handleLogin
    };
  }
});
</script>

<style scoped>
.login-container {
  height: 100vh;
  background: #f5f7fa;
}

.login-box {
  height: 100%;
}
</style>