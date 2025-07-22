<template>
  <div class="settings">
    <div class="settings-header">
      <h2>系统设置</h2>
    </div>

    <div class="settings-content">
      <n-grid :cols="1" :y-gap="24">
        <!-- 外观设置 -->
        <n-gi>
          <n-card title="外观设置" :bordered="false">
            <n-form label-placement="left" label-width="120px" size="large">
              <n-form-item label="主题模式">
                <n-radio-group v-model:value="themeMode" name="theme">
                  <n-radio-button value="light">浅色模式</n-radio-button>
                  <n-radio-button value="dark">深色模式</n-radio-button>
                  <n-radio-button value="auto">跟随系统</n-radio-button>
                </n-radio-group>
              </n-form-item>
              <n-form-item label="主题色">
                <div class="color-picker">
                  <div 
                    v-for="color in themeColors" 
                    :key="color.value"
                    class="color-item"
                    :class="{ active: primaryColor === color.value }"
                    :style="{ backgroundColor: color.value }"
                    @click="primaryColor = color.value"
                  >
                    <n-icon v-if="primaryColor === color.value" color="white">
                      <svg viewBox="0 0 24 24">
                        <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
                      </svg>
                    </n-icon>
                  </div>
                </div>
              </n-form-item>
              <n-form-item label="语言设置">
                <n-select
                  v-model:value="language"
                  :options="languageOptions"
                  style="width: 200px;"
                />
              </n-form-item>
              <n-form-item label="紧凑模式">
                <n-switch v-model:value="compactMode" />
                <span class="setting-desc">启用后界面元素间距更小</span>
              </n-form-item>
            </n-form>
          </n-card>
        </n-gi>

        <!-- 通知设置 -->
        <n-gi>
          <n-card title="通知设置" :bordered="false">
            <n-form label-placement="left" label-width="120px" size="large">
              <n-form-item label="桌面通知">
                <n-switch v-model:value="desktopNotifications" />
                <span class="setting-desc">允许显示桌面通知</span>
              </n-form-item>
              <n-form-item label="声音提醒">
                <n-switch v-model:value="soundNotifications" />
                <span class="setting-desc">新消息时播放提示音</span>
              </n-form-item>
              <n-form-item label="邮件通知">
                <n-checkbox-group v-model:value="emailNotificationTypes">
                  <n-space vertical>
                    <n-checkbox value="system" label="系统通知" />
                    <n-checkbox value="task" label="任务提醒" />
                    <n-checkbox value="message" label="消息推送" />
                    <n-checkbox value="weekly" label="周报总结" />
                  </n-space>
                </n-checkbox-group>
              </n-form-item>
              <n-form-item label="推送频率">
                <n-radio-group v-model:value="notificationFrequency">
                  <n-space vertical>
                    <n-radio value="realtime">实时推送</n-radio>
                    <n-radio value="hourly">每小时汇总</n-radio>
                    <n-radio value="daily">每日汇总</n-radio>
                  </n-space>
                </n-radio-group>
              </n-form-item>
            </n-form>
          </n-card>
        </n-gi>

        <!-- 安全设置 -->
        <n-gi>
          <n-card title="安全设置" :bordered="false">
            <n-form label-placement="left" label-width="120px" size="large">
              <n-form-item label="登录验证">
                <n-switch v-model:value="twoFactorAuth" />
                <span class="setting-desc">启用两步验证增强安全性</span>
              </n-form-item>
              <n-form-item label="会话超时">
                <n-select
                  v-model:value="sessionTimeout"
                  :options="timeoutOptions"
                  style="width: 200px;"
                />
              </n-form-item>
              <n-form-item label="登录设备管理">
                <n-button type="primary" ghost>
                  查看活跃设备
                </n-button>
                <span class="setting-desc">管理您的登录设备</span>
              </n-form-item>
              <n-form-item label="密码策略">
                <n-checkbox-group v-model:value="passwordRequirements">
                  <n-space vertical>
                    <n-checkbox value="length" label="至少8位字符" disabled checked />
                    <n-checkbox value="uppercase" label="包含大写字母" />
                    <n-checkbox value="numbers" label="包含数字" />
                    <n-checkbox value="symbols" label="包含特殊字符" />
                  </n-space>
                </n-checkbox-group>
              </n-form-item>
            </n-form>
          </n-card>
        </n-gi>

        <!-- 数据设置 -->
        <n-gi>
          <n-card title="数据与隐私" :bordered="false">
            <n-form label-placement="left" label-width="120px" size="large">
              <n-form-item label="数据导出">
                <n-button type="primary" ghost @click="handleDataExport">
                  导出我的数据
                </n-button>
                <span class="setting-desc">下载您的个人数据副本</span>
              </n-form-item>
              <n-form-item label="使用统计">
                <n-switch v-model:value="allowAnalytics" />
                <span class="setting-desc">帮助我们改善产品体验</span>
              </n-form-item>
              <n-form-item label="个性化推荐">
                <n-switch v-model:value="personalizedContent" />
                <span class="setting-desc">基于使用习惯提供个性化内容</span>
              </n-form-item>
              <n-form-item label="数据清理">
                <n-button type="warning" ghost @click="handleDataCleanup">
                  清理缓存数据
                </n-button>
                <span class="setting-desc">清理本地缓存和临时文件</span>
              </n-form-item>
            </n-form>
          </n-card>
        </n-gi>

        <!-- 危险操作 -->
        <n-gi>
          <n-card title="账户管理" :bordered="false">
            <n-alert type="warning" style="margin-bottom: 24px;">
              以下操作具有风险性，请谨慎操作
            </n-alert>
            <n-form label-placement="left" label-width="120px" size="large">
              <n-form-item label="注销账户">
                <n-button type="error" ghost @click="showDeleteConfirm = true">
                  删除账户
                </n-button>
                <span class="setting-desc">永久删除您的账户和所有数据</span>
              </n-form-item>
            </n-form>
          </n-card>
        </n-gi>
      </n-grid>
    </div>

    <!-- 删除确认对话框 -->
    <n-modal v-model:show="showDeleteConfirm" preset="dialog" title="确认删除账户">
      <template #header>
        <n-icon color="#f5222d" size="24">
          <svg viewBox="0 0 24 24">
            <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
          </svg>
        </n-icon>
        确认删除账户
      </template>
      <p>此操作将永久删除您的账户和所有相关数据，且无法恢复。</p>
      <p>请输入 <strong>DELETE</strong> 来确认删除：</p>
      <n-input v-model:value="deleteConfirmText" placeholder="请输入 DELETE" />
      <template #action>
        <n-button @click="showDeleteConfirm = false">取消</n-button>
        <n-button 
          type="error" 
          :disabled="deleteConfirmText !== 'DELETE'"
          @click="handleDeleteAccount"
        >
          确认删除
        </n-button>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref } from 'vue';
import { useMessage } from 'naive-ui';

const message = useMessage();

// 外观设置
const themeMode = ref('light');
const primaryColor = ref('#18a058');
const language = ref('zh-CN');
const compactMode = ref(false);

// 通知设置
const desktopNotifications = ref(true);
const soundNotifications = ref(false);
const emailNotificationTypes = ref(['system', 'task']);
const notificationFrequency = ref('realtime');

// 安全设置
const twoFactorAuth = ref(false);
const sessionTimeout = ref('30');
const passwordRequirements = ref(['length']);

// 数据设置
const allowAnalytics = ref(false);
const personalizedContent = ref(true);

// 删除账户
const showDeleteConfirm = ref(false);
const deleteConfirmText = ref('');

// 主题色选项
const themeColors = [
  { label: '默认绿', value: '#18a058' },
  { label: '蓝色', value: '#1890ff' },
  { label: '紫色', value: '#722ed1' },
  { label: '橙色', value: '#fa8c16' },
  { label: '红色', value: '#f5222d' },
  { label: '青色', value: '#13c2c2' }
];

// 语言选项
const languageOptions = [
  { label: '简体中文', value: 'zh-CN' },
  { label: 'English', value: 'en-US' },
  { label: '日本語', value: 'ja-JP' }
];

// 会话超时选项
const timeoutOptions = [
  { label: '15分钟', value: '15' },
  { label: '30分钟', value: '30' },
  { label: '1小时', value: '60' },
  { label: '4小时', value: '240' },
  { label: '1天', value: '1440' }
];

// 数据导出
const handleDataExport = () => {
  message.loading('正在准备导出数据...');
  setTimeout(() => {
    message.success('数据导出完成，请查看下载文件');
  }, 2000);
};

// 数据清理
const handleDataCleanup = () => {
  message.loading('正在清理缓存数据...');
  setTimeout(() => {
    message.success('缓存数据清理完成');
  }, 1500);
};

// 删除账户
const handleDeleteAccount = () => {
  message.loading('正在删除账户...');
  setTimeout(() => {
    message.error('账户删除失败：演示环境不允许删除');
    showDeleteConfirm.value = false;
    deleteConfirmText.value = '';
  }, 1500);
};
</script>

<style scoped>
.settings {
  max-width: 1000px;
  margin: 0 auto;
}

.settings-header {
  margin-bottom: 24px;
}

.settings-header h2 {
  margin: 0;
  font-size: 24px;
  font-weight: 600;
  color: #262626;
}

.settings-content :deep(.n-card) {
  border-radius: 12px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.setting-desc {
  margin-left: 12px;
  font-size: 14px;
  color: #8c8c8c;
}

.color-picker {
  display: flex;
  gap: 8px;
  flex-wrap: wrap;
}

.color-item {
  width: 32px;
  height: 32px;
  border-radius: 6px;
  cursor: pointer;
  display: flex;
  align-items: center;
  justify-content: center;
  transition: all 0.3s ease;
  border: 2px solid transparent;
}

.color-item:hover {
  transform: scale(1.1);
}

.color-item.active {
  border-color: #d9d9d9;
  transform: scale(1.1);
}

.settings-content :deep(.n-form-item) {
  margin-bottom: 24px;
}

.settings-content :deep(.n-form-item:last-child) {
  margin-bottom: 0;
}

.settings-content :deep(.n-checkbox) {
  margin-bottom: 8px;
}

.settings-content :deep(.n-radio) {
  margin-bottom: 8px;
}

@media (max-width: 768px) {
  .settings-content :deep(.n-form) {
    label-width: 80px;
  }
  
  .color-picker {
    justify-content: flex-start;
  }
  
  .color-item {
    width: 28px;
    height: 28px;
  }
}
</style>
