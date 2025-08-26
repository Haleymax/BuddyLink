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
        <n-button @click="refreshMessages">
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M17.65,6.35C16.2,4.9 14.21,4 12,4A8,8 0 0,0 4,12A8,8 0 0,0 12,20C15.73,20 18.84,17.45 19.73,14H17.65C16.83,16.33 14.61,18 12,18A6,6 0 0,1 6,12A6,6 0 0,1 12,6C13.66,6 15.14,6.69 16.22,7.78L13,11H20V4L17.65,6.35Z"/>
              </svg>
            </n-icon>
          </template>
          刷新
        </n-button>
      </div>
    </div>

    <div class="messages-content">
      <n-tabs v-model:value="activeTab" type="line" size="large">
        <n-tab-pane name="all" tab="全部消息">
          <div class="message-list">
            <MessageCard 
              v-for="message in allMessages" 
              :key="message.id"
              :message="message"
              @click="handleMessageClick"
              @action="handleMessageAction"
            />
          </div>
        </n-tab-pane>

        <n-tab-pane name="apply" tab="申请消息">
          <div class="message-list">
            <MessageCard 
              v-for="message in applyMessages" 
              :key="message.id"
              :message="message"
              @click="handleMessageClick"
              @action="handleMessageAction"
            />
          </div>
        </n-tab-pane>

        <n-tab-pane name="system" tab="系统消息">
          <div class="message-list">
            <MessageCard 
              v-for="message in systemMessages" 
              :key="message.id"
              :message="message"
              @click="handleMessageClick"
              @action="handleMessageAction"
            />
          </div>
        </n-tab-pane>
      </n-tabs>
    </div>

    <!-- 消息详情模态框 -->
    <n-modal 
      v-model:show="showDetailModal" 
      preset="card" 
      title="消息详情"
      style="width: 500px; max-width: 90vw;"
      :mask-closable="true"
    >
      <div v-if="selectedMessage" class="message-detail">
        <div class="detail-header">
          <div class="detail-sender">
            <n-avatar :size="64">
              <n-icon>
                <svg viewBox="0 0 24 24" v-if="selectedMessage.type === 'system'">
                  <path fill="currentColor" d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z"/>
                </svg>
                <svg viewBox="0 0 24 24" v-else>
                  <path fill="currentColor" d="M12,4A4,4 0 0,1 16,8A4,4 0 0,1 12,12A4,4 0 0,1 8,8A4,4 0 0,1 12,4M12,14C16.42,14 20,15.79 20,18V20H4V18C4,15.79 7.58,14 12,14Z"/>
                </svg>
              </n-icon>
            </n-avatar>
            <div class="sender-info">
              <div class="sender-name">{{ getSenderName(selectedMessage) }}</div>
              <div class="message-time">{{ formatTime(selectedMessage.created_at) }}</div>
              <n-tag :type="getTypeTagType(selectedMessage.type)" size="small">
                {{ getTypeLabel(selectedMessage.type) }}
              </n-tag>
            </div>
          </div>
        </div>
        
        <div class="detail-content">
          <p>{{ getMessageContent(selectedMessage) }}</p>
          <div v-if="selectedMessage.data && typeof selectedMessage.data === 'object'" class="message-data">
            <n-divider>详细信息</n-divider>
            <pre>{{ JSON.stringify(selectedMessage.data, null, 2) }}</pre>
          </div>
        </div>

        <div class="detail-actions">
          <n-space>
            <n-button 
              v-if="!selectedMessage.is_read" 
              type="primary" 
              @click="markAsRead(selectedMessage)"
            >
              标记已读
            </n-button>
            
            <n-button 
              v-if="canApprove(selectedMessage)" 
              type="success" 
              @click="approveMessage(selectedMessage)"
              :disabled="selectedMessage.status === 'approved'"
            >
              {{ selectedMessage.status === 'approved' ? '已同意' : '同意' }}
            </n-button>
            
            <n-button 
              v-if="canReject(selectedMessage)" 
              type="error" 
              @click="rejectMessage(selectedMessage)"
              :disabled="selectedMessage.status === 'rejected'"
            >
              {{ selectedMessage.status === 'rejected' ? '已拒绝' : '拒绝' }}
            </n-button>
            
            <n-button 
              type="error" 
              ghost 
              @click="deleteMessage(selectedMessage)"
            >
              删除
            </n-button>
          </n-space>
        </div>
      </div>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue';
import { useMessage } from 'naive-ui';
import { useAuthStore } from '../../stores/auth.store';
import type { BaseMessage } from '../../model/message';
import type { ApiResponse } from '../../api/apiResponse';
import { getMineMessage } from '../../api/message';
import MessageCard from '../../components/MessageCard.vue';
import '../../styles/Messages.css';

// 扩展 BaseMessage 接口以包含额外的显示字段
interface DisplayMessage extends BaseMessage {
  id: number;
  created_at: string;
  updated_at: string;
  is_read: boolean;
  status?: 'pending' | 'approved' | 'rejected';
}

const message = useMessage();
const authStore = useAuthStore();
const activeTab = ref('all');

// 响应式数据
const messages = ref<DisplayMessage[]>([]);
const loading = ref(false);
const showDetailModal = ref(false);
const selectedMessage = ref<DisplayMessage | null>(null);

// 计算各类消息
const allMessages = computed(() => messages.value);
const applyMessages = computed(() => messages.value.filter(m => m.type === 'apply'));
const systemMessages = computed(() => messages.value.filter(m => m.type === 'system'));

// 获取发送者名称
const getSenderName = (msg: DisplayMessage) => {
  switch (msg.type) {
    case 'system':
      return '系统通知';
    case 'apply':
      return `用户 ${msg.sender_id}`;
    default:
      return `用户 ${msg.sender_id}`;
  }
};

// 获取消息内容
const getMessageContent = (msg: DisplayMessage) => {
  if (typeof msg.data === 'string') {
    return msg.data;
  } else if (typeof msg.data === 'object' && msg.data !== null) {
    if (msg.data.data) {
      return msg.data.data;
    }
    return JSON.stringify(msg.data);
  }
  return '暂无内容';
};

// 获取消息类型标签
const getTypeLabel = (type: string) => {
  switch (type) {
    case 'system':
      return '系统';
    case 'apply':
      return '申请';
    case 'comment':
      return '评论';
    case 'like':
      return '点赞';
    case 'follow':
      return '关注';
    default:
      return '';
  }
};

// 格式化时间
const formatTime = (timeStr: string) => {
  const time = new Date(timeStr);
  const now = new Date();
  const diff = now.getTime() - time.getTime();
  const minutes = Math.floor(diff / (1000 * 60));
  const hours = Math.floor(diff / (1000 * 60 * 60));
  const days = Math.floor(diff / (1000 * 60 * 60 * 24));

  if (minutes < 60) {
    return `${minutes}分钟前`;
  } else if (hours < 24) {
    return `${hours}小时前`;
  } else if (days < 30) {
    return `${days}天前`;
  } else {
    return time.toLocaleDateString('zh-CN');
  }
};

// 获取标签类型
const getTypeTagType = (type: string) => {
  switch (type) {
    case 'system':
      return 'info';
    case 'apply':
      return 'warning';
    case 'comment':
      return 'success';
    case 'like':
      return 'error';
    case 'follow':
      return 'default';
    default:
      return 'default';
  }
};

// 判断是否可以同意
const canApprove = (msg: DisplayMessage) => {
  return msg.type === 'apply' && msg.receiver_id === authStore.user?.id;
};

// 判断是否可以拒绝
const canReject = (msg: DisplayMessage) => {
  return msg.type === 'apply' && msg.receiver_id === authStore.user?.id;
};

// 处理消息点击 - 现在从 MessageCard 组件接收事件
const handleMessageClick = (msg: DisplayMessage) => {
  selectedMessage.value = msg;
  showDetailModal.value = true;
  
  if (!msg.is_read) {
    markAsRead(msg);
  }
};

// 处理消息操作 - 现在从 MessageCard 组件接收事件
const handleMessageAction = (action: string, msg: DisplayMessage) => {
  switch (action) {
    case 'mark-read':
      markAsRead(msg);
      break;
    case 'view-detail':
      handleMessageClick(msg);
      break;
    case 'delete':
      deleteMessage(msg);
      break;
  }
};

// 标记为已读
const markAsRead = (msg: DisplayMessage) => {
  msg.is_read = true;
  message.success('已标记为已读');
  // TODO: 调用API标记为已读
};

// 同意申请
const approveMessage = (msg: DisplayMessage) => {
  msg.status = 'approved';
  message.success('已同意申请');
  // TODO: 调用API处理申请
};

// 拒绝申请
const rejectMessage = (msg: DisplayMessage) => {
  msg.status = 'rejected';
  message.warning('已拒绝申请');
  // TODO: 调用API处理申请
};

// 删除消息
const deleteMessage = (msg: DisplayMessage) => {
  const index = messages.value.findIndex(m => m.id === msg.id);
  if (index > -1) {
    messages.value.splice(index, 1);
    message.success('消息已删除');
    if (selectedMessage.value?.id === msg.id) {
      showDetailModal.value = false;
    }
  }
  // TODO: 调用API删除消息
};

// 标记全部已读
const markAllRead = () => {
  messages.value.forEach(msg => {
    msg.is_read = true;
  });
  message.success('所有消息已标记为已读');
  // TODO: 调用API标记全部已读
};

// 刷新消息
const refreshMessages = async () => {
  await loadMessages();
  message.success('消息列表已刷新');
};

// 加载消息数据
const loadMessages = async () => {
  try {
    loading.value = true;
    
    // 使用 getMineMessage API 获取消息
    const response = await getMineMessage(
      authStore.token ?? '',
      authStore.user?.id ?? 0,
      '',  // 所有类型
      ''   // 所有状态
    ) as unknown as ApiResponse;

    if (response.code === 200) {
      // 处理API返回的数据
      if (Array.isArray(response.data)) {
        messages.value = response.data.map((msg: any) => ({
          ...msg,
          is_read: msg.is_read || false,
          status: msg.status || 'pending'
        }));
      } else if (response.data && typeof response.data === 'object') {
        messages.value = [{
          ...response.data,
          is_read: response.data.is_read || false,
          status: response.data.status || 'pending'
        }];
      } else {
        messages.value = [];
      }
    } else {
      message.error(response.message || '获取消息失败');
    }
  } catch (error) {
    console.error('加载消息失败:', error);
    message.error('加载消息失败，使用模拟数据');
  } finally {
    loading.value = false;
  }
};


onMounted(() => {
  loadMessages();
});
</script>
