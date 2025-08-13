<template>
  <div class="messages">
    <div class="messages-header">
      <h2>消息中心</h2>
      <div class="header-actions">
        <n-button type="primary" @click="markAllRead">
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
              </svg>
            </n-icon>
          </template>
          全部已读
        </n-button>
      </div>
    </div>

    <div class="messages-content">
      <n-tabs v-model:value="activeTab" type="line" size="large">
        <n-tab-pane name="all" tab="全部消息">
          <div class="message-list">
            <div 
              v-for="message in allMessages" 
              :key="message.id"
              class="message-item"
              :class="{ 'unread': !message.read }"
              @click="handleMessageClick(message)"
            >
              <div class="message-avatar">
                <n-avatar :src="message.avatar" :size="48">
                  {{ message.sender.charAt(0) }}
                </n-avatar>
                <div v-if="!message.read" class="unread-dot"></div>
              </div>
              <div class="message-content">
                <div class="message-header">
                  <div class="sender-name">{{ message.sender }}</div>
                  <div class="message-time">{{ formatTime(message.time) }}</div>
                </div>
                <div class="message-text">{{ message.content }}</div>
                <div class="message-type" :class="getTypeClass(message.type)">
                  {{ getTypeLabel(message.type) }}
                </div>
              </div>
              <div class="message-actions">
                <n-dropdown :options="messageMenuOptions" @select="(key: string) => handleMessageAction(key, message)">
                  <n-button quaternary circle size="small">
                    <template #icon>
                      <n-icon>
                        <svg viewBox="0 0 24 24">
                          <path fill="currentColor" d="M12 8c1.1 0 2-.9 2-2s-.9-2-2-2-2 .9-2 2 .9 2 2 2zm0 2c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2zm0 6c-1.1 0-2 .9-2 2s.9 2 2 2 2-.9 2-2-.9-2-2-2z"/>
                        </svg>
                      </n-icon>
                    </template>
                  </n-button>
                </n-dropdown>
              </div>
            </div>
          </div>
        </n-tab-pane>

        <n-tab-pane name="system" tab="系统通知">
          <div class="message-list">
            <div 
              v-for="message in systemMessages" 
              :key="message.id"
              class="message-item"
              :class="{ 'unread': !message.read }"
            >
              <div class="message-avatar">
                <n-icon size="48" color="#1890ff">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
                  </svg>
                </n-icon>
                <div v-if="!message.read" class="unread-dot"></div>
              </div>
              <div class="message-content">
                <div class="message-header">
                  <div class="sender-name">{{ message.sender }}</div>
                  <div class="message-time">{{ formatTime(message.time) }}</div>
                </div>
                <div class="message-text">{{ message.content }}</div>
              </div>
            </div>
          </div>
        </n-tab-pane>

        <n-tab-pane name="team" tab="团队消息">
          <div class="message-list">
            <div 
              v-for="message in teamMessages" 
              :key="message.id"
              class="message-item"
              :class="{ 'unread': !message.read }"
            >
              <div class="message-avatar">
                <n-avatar :src="message.avatar" :size="48">
                  {{ message.sender.charAt(0) }}
                </n-avatar>
                <div v-if="!message.read" class="unread-dot"></div>
              </div>
              <div class="message-content">
                <div class="message-header">
                  <div class="sender-name">{{ message.sender }}</div>
                  <div class="message-time">{{ formatTime(message.time) }}</div>
                </div>
                <div class="message-text">{{ message.content }}</div>
              </div>
            </div>
          </div>
        </n-tab-pane>
      </n-tabs>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import { useMessage } from 'naive-ui';
import '../../styles/Messages.css';

interface Message {
  id: string;
  sender: string;
  content: string;
  time: Date;
  read: boolean;
  type: 'system' | 'team' | 'notification';
  avatar?: string;
}

const message = useMessage();
const activeTab = ref('all');

// 消息数据
const messages = ref<Message[]>([
  {
    id: '1',
    sender: '系统通知',
    content: '您的账户安全设置已更新，请注意保护账户安全。',
    time: new Date(Date.now() - 1000 * 60 * 30), // 30分钟前
    read: false,
    type: 'system'
  },
  {
    id: '2',
    sender: '张小明',
    content: '项目进度更新：前端开发已完成 80%，预计本周内完成。',
    time: new Date(Date.now() - 1000 * 60 * 60 * 2), // 2小时前
    read: false,
    type: 'team',
    avatar: 'https://via.placeholder.com/48?text=张'
  },
  {
    id: '3',
    sender: '李小红',
    content: '设计稿已上传，请及时查看并提供反馈意见。',
    time: new Date(Date.now() - 1000 * 60 * 60 * 5), // 5小时前
    read: true,
    type: 'team',
    avatar: 'https://via.placeholder.com/48?text=李'
  },
  {
    id: '4',
    sender: '系统通知',
    content: '您有一个新的任务分配：优化用户界面交互体验。',
    time: new Date(Date.now() - 1000 * 60 * 60 * 24), // 1天前
    read: true,
    type: 'system'
  },
  {
    id: '5',
    sender: '王大力',
    content: '会议提醒：明天上午10点有项目评审会议，请准时参加。',
    time: new Date(Date.now() - 1000 * 60 * 60 * 24 * 2), // 2天前
    read: true,
    type: 'team',
    avatar: 'https://via.placeholder.com/48?text=王'
  }
]);

// 计算各类消息
const allMessages = computed(() => messages.value);
const systemMessages = computed(() => messages.value.filter(m => m.type === 'system'));
const teamMessages = computed(() => messages.value.filter(m => m.type === 'team'));

// 消息菜单选项
const messageMenuOptions = [
  {
    label: '标记为已读',
    key: 'mark-read'
  },
  {
    label: '删除消息',
    key: 'delete'
  }
];

// 格式化时间
const formatTime = (time: Date) => {
  const now = new Date();
  const diff = now.getTime() - time.getTime();
  const minutes = Math.floor(diff / (1000 * 60));
  const hours = Math.floor(diff / (1000 * 60 * 60));
  const days = Math.floor(diff / (1000 * 60 * 60 * 24));

  if (minutes < 60) {
    return `${minutes}分钟前`;
  } else if (hours < 24) {
    return `${hours}小时前`;
  } else {
    return `${days}天前`;
  }
};

// 获取消息类型样式
const getTypeClass = (type: string) => {
  switch (type) {
    case 'system':
      return 'type-system';
    case 'team':
      return 'type-team';
    case 'notification':
      return 'type-notification';
    default:
      return '';
  }
};

// 获取消息类型标签
const getTypeLabel = (type: string) => {
  switch (type) {
    case 'system':
      return '系统';
    case 'team':
      return '团队';
    case 'notification':
      return '通知';
    default:
      return '';
  }
};

// 处理消息点击
const handleMessageClick = (msg: Message) => {
  if (!msg.read) {
    msg.read = true;
    message.info(`已读消息：${msg.content.substring(0, 20)}...`);
  }
};

// 处理消息操作
const handleMessageAction = (action: string, msg: Message) => {
  switch (action) {
    case 'mark-read':
      msg.read = true;
      message.success('已标记为已读');
      break;
    case 'delete':
      const index = messages.value.findIndex(m => m.id === msg.id);
      if (index > -1) {
        messages.value.splice(index, 1);
        message.success('消息已删除');
      }
      break;
  }
};

// 标记全部已读
const markAllRead = () => {
  messages.value.forEach(msg => {
    msg.read = true;
  });
  message.success('所有消息已标记为已读');
};
</script>
