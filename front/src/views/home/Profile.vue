<template>
  <div class="profile">
    <div class="profile-header">
      <h2>个人资料</h2>
    </div>

    <div class="profile-content">
      <n-grid :cols="3" :x-gap="24">
        <!-- 左侧个人信息卡片 -->
        <n-gi :span="1">
          <n-card class="profile-card" :bordered="false">
            <div class="profile-avatar">
              <div class="avatar-container" @click="handleAvatarUpload">
                <n-avatar :size="120" src="https://via.placeholder.com/120?text=User">
                  用户
                </n-avatar>
                <div class="avatar-overlay">
                  <n-icon size="28" color="white">
                    <svg viewBox="0 0 24 24">
                      <path fill="currentColor" d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z"/>
                    </svg>
                  </n-icon>
                  <span class="upload-text">点击更换头像</span>
                </div>
              </div>
              <input
                ref="fileInput"
                type="file"
                accept="image/*"
                style="display: none;"
                @change="handleFileChange"
              />
            </div>
            <div class="profile-info">
              <h3>{{ userInfo.name }}</h3>
              <p class="profile-title">{{ userInfo.title }}</p>
              <p class="profile-email">{{ userInfo.email }}</p>
              <div class="profile-stats">
                <div class="stat-item">
                  <div class="stat-number">{{ userInfo.projects }}</div>
                  <div class="stat-label">项目数</div>
                </div>
                <div class="stat-item">
                  <div class="stat-number">{{ userInfo.tasks }}</div>
                  <div class="stat-label">任务数</div>
                </div>
                <div class="stat-item">
                  <div class="stat-number">{{ userInfo.score }}</div>
                  <div class="stat-label">积分</div>
                </div>
              </div>
            </div>
          </n-card>
        </n-gi>

        <!-- 右侧表单 -->
        <n-gi :span="2">
          <n-card title="基本信息" :bordered="false">
            <n-form
              ref="formRef"
              :model="formData"
              :rules="rules"
              label-placement="left"
              label-width="100px"
              size="large"
            >
              <n-grid :cols="2" :x-gap="24">
                <n-gi>
                  <n-form-item label="姓名" path="name">
                    <n-input v-model:value="formData.name" placeholder="请输入姓名" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="职位" path="title">
                    <n-input v-model:value="formData.title" placeholder="请输入职位" />
                  </n-form-item>
                </n-gi>
                <n-gi :span="2">
                  <n-form-item label="邮箱" path="email">
                    <n-input v-model:value="formData.email" placeholder="请输入邮箱" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="电话" path="phone">
                    <n-input v-model:value="formData.phone" placeholder="请输入电话号码" />
                  </n-form-item>
                </n-gi>
                <n-gi>
                  <n-form-item label="部门" path="department">
                    <n-select
                      v-model:value="formData.department"
                      :options="departmentOptions"
                      placeholder="请选择部门"
                    />
                  </n-form-item>
                </n-gi>
                <n-gi :span="2">
                  <n-form-item label="地址" path="address">
                    <n-input v-model:value="formData.address" placeholder="请输入地址" />
                  </n-form-item>
                </n-gi>
                <n-gi :span="2">
                  <n-form-item label="个人简介" path="bio">
                    <n-input
                      v-model:value="formData.bio"
                      type="textarea"
                      placeholder="请输入个人简介"
                      :rows="4"
                    />
                  </n-form-item>
                </n-gi>
              </n-grid>

              <div class="form-actions">
                <n-button type="primary" size="large" @click="handleSave" :loading="loading">
                  保存更改
                </n-button>
                <n-button size="large" @click="handleReset">
                  重置
                </n-button>
              </div>
            </n-form>
          </n-card>
        </n-gi>
      </n-grid>

      <!-- 其他设置 -->
      <n-grid :cols="2" :x-gap="24" style="margin-top: 24px;">
        <n-gi>
          <n-card title="账户设置" :bordered="false">
            <n-list>
              <n-list-item>
                <div class="setting-item">
                  <div class="setting-info">
                    <div class="setting-title">修改密码</div>
                    <div class="setting-desc">更新您的登录密码</div>
                  </div>
                  <n-button type="primary" ghost size="small">
                    修改
                  </n-button>
                </div>
              </n-list-item>
              <n-list-item>
                <div class="setting-item">
                  <div class="setting-info">
                    <div class="setting-title">两步验证</div>
                    <div class="setting-desc">增强账户安全性</div>
                  </div>
                  <n-switch v-model:value="twoFactorEnabled" />
                </div>
              </n-list-item>
              <n-list-item>
                <div class="setting-item">
                  <div class="setting-info">
                    <div class="setting-title">邮箱通知</div>
                    <div class="setting-desc">接收重要邮件通知</div>
                  </div>
                  <n-switch v-model:value="emailNotifications" />
                </div>
              </n-list-item>
            </n-list>
          </n-card>
        </n-gi>
        <n-gi>
          <n-card title="隐私设置" :bordered="false">
            <n-list>
              <n-list-item>
                <div class="setting-item">
                  <div class="setting-info">
                    <div class="setting-title">个人资料可见性</div>
                    <div class="setting-desc">控制谁能查看您的资料</div>
                  </div>
                  <n-select
                    v-model:value="profileVisibility"
                    :options="visibilityOptions"
                    style="width: 120px;"
                  />
                </div>
              </n-list-item>
              <n-list-item>
                <div class="setting-item">
                  <div class="setting-info">
                    <div class="setting-title">活动状态</div>
                    <div class="setting-desc">显示您的在线状态</div>
                  </div>
                  <n-switch v-model:value="showActivity" />
                </div>
              </n-list-item>
              <n-list-item>
                <div class="setting-item">
                  <div class="setting-info">
                    <div class="setting-title">数据分析</div>
                    <div class="setting-desc">允许收集使用数据以改善服务</div>
                  </div>
                  <n-switch v-model:value="allowAnalytics" />
                </div>
              </n-list-item>
            </n-list>
          </n-card>
        </n-gi>
      </n-grid>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, reactive } from 'vue';
import { useMessage, type FormInst, type FormRules } from 'naive-ui';

const message = useMessage();
const formRef = ref<FormInst | null>(null);
const fileInput = ref<HTMLInputElement | null>(null);
const loading = ref(false);

// 用户信息
const userInfo = reactive({
  name: 'John Doe',
  title: '前端开发工程师',
  email: 'john.doe@example.com',
  projects: 12,
  tasks: 48,
  score: 2580
});

// 表单数据
const formData = reactive({
  name: 'John Doe',
  title: '前端开发工程师',
  email: 'john.doe@example.com',
  phone: '+86 138 0013 8000',
  department: 'tech',
  address: '北京市朝阳区某某街道123号',
  bio: '热爱技术，专注于前端开发和用户体验设计。拥有5年Web开发经验，熟悉Vue、React等主流框架。'
});

// 设置状态
const twoFactorEnabled = ref(false);
const emailNotifications = ref(true);
const profileVisibility = ref('public');
const showActivity = ref(true);
const allowAnalytics = ref(false);

// 部门选项
const departmentOptions = [
  { label: '技术部', value: 'tech' },
  { label: '产品部', value: 'product' },
  { label: '设计部', value: 'design' },
  { label: '运营部', value: 'operation' },
  { label: '人事部', value: 'hr' }
];

// 可见性选项
const visibilityOptions = [
  { label: '公开', value: 'public' },
  { label: '团队', value: 'team' },
  { label: '私人', value: 'private' }
];

// 表单验证规则
const rules: FormRules = {
  name: [
    { required: true, message: '请输入姓名', trigger: 'blur' }
  ],
  email: [
    { required: true, message: '请输入邮箱', trigger: 'blur' },
    { type: 'email', message: '请输入有效的邮箱地址', trigger: 'blur' }
  ],
  phone: [
    { pattern: /^1[3-9]\d{9}$/, message: '请输入有效的手机号码', trigger: 'blur' }
  ]
};

// 保存更改
const handleSave = async () => {
  if (!formRef.value) return;

  try {
    await formRef.value.validate();
    loading.value = true;
    
    // 模拟API调用
    await new Promise(resolve => setTimeout(resolve, 1000));
    
    // 更新用户信息
    Object.assign(userInfo, {
      name: formData.name,
      title: formData.title,
      email: formData.email
    });
    
    message.success('个人资料更新成功');
  } catch (error) {
    message.error('请检查输入信息');
  } finally {
    loading.value = false;
  }
};

// 重置表单
const handleReset = () => {
  Object.assign(formData, {
    name: userInfo.name,
    title: userInfo.title,
    email: userInfo.email,
    phone: '+86 138 0013 8000',
    department: 'tech',
    address: '北京市朝阳区某某街道123号',
    bio: '热爱技术，专注于前端开发和用户体验设计。拥有5年Web开发经验，熟悉Vue、React等主流框架。'
  });
  message.info('表单已重置');
};

// 头像上传处理
const handleAvatarUpload = () => {
  fileInput.value?.click();
};

// 文件选择处理
const handleFileChange = (e: Event) => {
  const target = e.target as HTMLInputElement;
  const file = target.files?.[0];
  
  if (file) {
    // 检查文件类型
    if (!file.type.startsWith('image/')) {
      message.error('请选择图片文件');
      return;
    }
    
    // 检查文件大小 (5MB限制)
    if (file.size > 5 * 1024 * 1024) {
      message.error('图片大小不能超过5MB');
      return;
    }
    
    // 创建FileReader预览图片
    const reader = new FileReader();
    reader.onload = () => {
      // 这里应该上传到服务器，目前只是更新预览
      message.success('头像上传成功');
      // 实际项目中应该调用API上传文件
      // uploadAvatar(file);
    };
    reader.readAsDataURL(file);
    
    // 清空input value，允许重复选择同一文件
    target.value = '';
  }
};
</script>

<style scoped>
.profile {
  max-width: 1200px;
  margin: 0 auto;
}

.profile-header {
  margin-bottom: 24px;
}

.profile-header h2 {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
  color: #262626;
}

.profile-content :deep(.n-card) {
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.profile-card {
  text-align: center;
}

.profile-avatar {
  position: relative;
  display: inline-block;
  margin-bottom: 24px;
}

.avatar-container {
  position: relative;
  display: inline-block;
  cursor: pointer;
  border-radius: 50%;
  overflow: hidden;
  transition: all 0.3s ease;
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
}

.avatar-container:hover {
  transform: scale(1.08);
  box-shadow: 0 8px 25px rgba(0, 0, 0, 0.25);
}

.avatar-container:hover .avatar-overlay {
  opacity: 1;
}

.avatar-overlay {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background: rgba(0, 0, 0, 0.6);
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  opacity: 0;
  transition: opacity 0.3s ease;
  border-radius: 50%;
}

.upload-text {
  color: white;
  font-size: 13px;
  font-weight: 500;
  margin-top: 6px;
  text-align: center;
}

.profile-info h3 {
  font-size: 20px;
  font-weight: 600;
  color: #262626;
  margin: 0 0 8px 0;
}

.profile-title {
  color: #595959;
  font-size: 16px;
  margin: 0 0 4px 0;
}

.profile-email {
  color: #8c8c8c;
  font-size: 14px;
  margin: 0 0 20px 0;
}

.profile-stats {
  display: flex;
  justify-content: space-around;
  padding: 20px 0;
  border-top: 1px solid #f0f0f0;
}

.stat-item {
  text-align: center;
}

.stat-number {
  font-size: 24px;
  font-weight: 600;
  color: #262626;
  margin-bottom: 4px;
}

.stat-label {
  font-size: 12px;
  color: #8c8c8c;
}

.form-actions {
  display: flex;
  gap: 16px;
  justify-content: flex-end;
  margin-top: 24px;
  padding-top: 24px;
  border-top: 1px solid #f0f0f0;
}

.setting-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
}

.setting-info {
  flex: 1;
}

.setting-title {
  font-size: 16px;
  font-weight: 500;
  color: #262626;
  margin-bottom: 4px;
}

.setting-desc {
  font-size: 14px;
  color: #8c8c8c;
}

.profile-content :deep(.n-list-item) {
  padding: 16px 0;
}

.profile-content :deep(.n-list-item:not(:last-child)) {
  border-bottom: 1px solid #f0f0f0;
}

@media (max-width: 768px) {
  .profile-content :deep(.n-gi[span="1"]) {
    grid-column: span 3;
  }
  
  .profile-content :deep(.n-gi[span="2"]) {
    grid-column: span 3;
  }
  
  .form-actions {
    flex-direction: column;
  }
  
  .setting-item {
    flex-direction: column;
    align-items: flex-start;
    gap: 12px;
  }
}
</style>
