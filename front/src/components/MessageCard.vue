<template>
  <div 
    class="message-card"
    :class="{ 'unread': !message.isRead }"
    @click="handleCardClick"
  >
    <div class="message-avatar">
      <n-avatar :size="40">
        <n-icon>
          <svg viewBox="0 0 24 24" v-if="message.type === 'system'">
            <path fill="currentColor" d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z"/>
          </svg>
          <svg viewBox="0 0 24 24" v-else-if="message.type === 'apply'">
            <path fill="currentColor" d="M16,4C18.21,4 20,5.79 20,8C20,10.21 18.21,12 16,12C13.79,12 12,10.21 12,8C12,5.79 13.79,4 16,4M16,14C18.67,14 24,15.34 24,18V20H8V18C8,15.34 13.33,14 16,14Z"/>
          </svg>
          <svg viewBox="0 0 24 24" v-else>
            <path fill="currentColor" d="M12,4A4,4 0 0,1 16,8A4,4 0 0,1 12,12A4,4 0 0,1 8,8A4,4 0 0,1 12,4M12,14C16.42,14 20,15.79 20,18V20H4V18C4,15.79 7.58,14 12,14Z"/>
          </svg>
        </n-icon>
      </n-avatar>
      <div v-if="!message.isRead" class="unread-dot"></div>
    </div>
    
    <div class="message-content">
      <div class="message-header">
        <div class="sender-name">{{ getSenderName() }}</div>
        <div class="message-type" :class="getTypeClass()">
          {{ getTypeLabel() }}
        </div>
      </div>
      <div class="message-text">{{ getMessagePreview() }}</div>
      <div class="message-time">{{ formatTime(message.created_at) }}</div>
    </div>
    
    <div class="message-indicator">
      <n-icon v-if="!message.isRead" size="8" color="#ff4d4f">
        <svg viewBox="0 0 24 24">
          <path fill="currentColor" d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z"/>
        </svg>
      </n-icon>
      <n-icon size="16" color="#8c8c8c">
        <svg viewBox="0 0 24 24">
          <path fill="currentColor" d="M8.59,16.58L13.17,12L8.59,7.41L10,6L16,12L10,18L8.59,16.58Z"/>
        </svg>
      </n-icon>
    </div>

    <!-- 消息详情模态框 -->
    <n-modal 
      v-model:show="showDetailModal" 
      preset="card" 
      title="消息详情"
      style="width: 600px; max-width: 90vw;"
      :mask-closable="true"
    >
      <div class="message-detail">
        <div class="detail-header">
          <div class="detail-sender">
            <n-avatar :size="64">
              <n-icon>
                <svg viewBox="0 0 24 24" v-if="message.type === 'system'">
                  <path fill="currentColor" d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2Z"/>
                </svg>
                <svg viewBox="0 0 24 24" v-else-if="message.type === 'apply'">
                  <path fill="currentColor" d="M16,4C18.21,4 20,5.79 20,8C20,10.21 18.21,12 16,12C13.79,12 12,10.21 12,8C12,5.79 13.79,4 16,4M16,14C18.67,14 24,15.34 24,18V20H8V18C8,15.34 13.33,14 16,14Z"/>
                </svg>
                <svg viewBox="0 0 24 24" v-else>
                  <path fill="currentColor" d="M12,4A4,4 0 0,1 16,8A4,4 0 0,1 12,12A4,4 0 0,1 8,8A4,4 0 0,1 12,4M12,14C16.42,14 20,15.79 20,18V20H4V18C4,15.79 7.58,14 12,14Z"/>
                </svg>
              </n-icon>
            </n-avatar>
            <div class="sender-info">
              <div class="sender-name-detail">{{ getSenderName() }}</div>
              <div class="message-time-detail">{{ formatTime(message.created_at) }}</div>
              <n-tag :type="getTypeTagType()" size="small">
                {{ getTypeLabel() }}
              </n-tag>
              <n-tag v-if="message.action" :type="getActionTagType()" size="small">
                {{ getActionLabel() }}
              </n-tag>
            </div>
          </div>
        </div>
        
        <div class="detail-content">
          <p class="message-full-content">{{ getMessageContent() }}</p>
          <div v-if="message.data && typeof message.data === 'object'" class="message-data">
            <n-divider>详细信息</n-divider>
            <div class="data-content">
              <pre>{{ JSON.stringify(message.data, null, 2) }}</pre>
            </div>
          </div>
        </div>

        <div class="detail-actions">
          <n-space>
            <n-button 
              v-if="!message.isRead" 
              type="primary" 
              @click="handleMarkRead"
            >
              <template #icon>
                <n-icon>
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12,2A7,7 0 0,1 19,9C19,11.38 17.81,13.47 16,14.74V17A1,1 0 0,1 15,18H9A1,1 0 0,1 8,17V14.74C6.19,13.47 5,11.38 5,9A7,7 0 0,1 12,2M9,21H15A1,1 0 0,1 15,22H9A1,1 0 0,1 9,21Z"/>
                  </svg>
                </n-icon>
              </template>
              标记已读
            </n-button>
            
            <!-- 基于 action 的功能按钮 -->
            <n-button v-if="shouldShowCreateButton" type="info" @click="handleCreate">
              <template #icon>
                <n-icon>
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M19,13H13V19H11V13H5V11H11V5H13V11H19V13Z"/>
                  </svg>
                </n-icon>
              </template>
              创建操作
            </n-button>
            
            <n-button v-if="shouldShowUpdateButton" type="warning" @click="handleUpdate">
              <template #icon>
                <n-icon>
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z"/>
                  </svg>
                </n-icon>
              </template>
              更新操作
            </n-button>
            
            <n-button v-if="shouldShowDoneButton" type="success" @click="handleDone">
              <template #icon>
                <n-icon>
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M9,20.42L2.79,14.21L5.62,11.38L9,14.77L18.88,4.88L21.71,7.71L9,20.42Z"/>
                  </svg>
                </n-icon>
              </template>
              完成操作
            </n-button>
            
            <!-- 基于消息类型的按钮 -->
            <n-button v-if="canApprove()" type="success" @click="handleApprove">
              <template #icon>
                <n-icon>
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M9,20.42L2.79,14.21L5.62,11.38L9,14.77L18.88,4.88L21.71,7.71L9,20.42Z"/>
                  </svg>
                </n-icon>
              </template>
              同意申请
            </n-button>
            
            <n-button v-if="canReject()" type="error" @click="handleReject">
              <template #icon>
                <n-icon>
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M19,6.41L17.59,5L12,10.59L6.41,5L5,6.41L10.59,12L5,17.59L6.41,19L12,13.41L17.59,19L19,17.59L13.41,12L19,6.41Z"/>
                  </svg>
                </n-icon>
              </template>
              拒绝申请
            </n-button>
            
            <n-button @click="handleReply">
              <template #icon>
                <n-icon>
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M10,9V5L3,12L10,19V14.9C15,14.9 18.5,16.5 21,20C20,15 17,10 10,9Z"/>
                  </svg>
                </n-icon>
              </template>
              回复
            </n-button>
            
            <n-button quaternary type="error" @click="handleDelete">
              <template #icon>
                <n-icon>
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M9,3V4H4V6H5V19A2,2 0 0,0 7,21H17A2,2 0 0,0 19,19V6H20V4H15V3H9M7,6H17V19H7V6M9,8V17H11V8H9M13,8V17H15V8H13Z"/>
                  </svg>
                </n-icon>
              </template>
              删除消息
            </n-button>
          </n-space>
        </div>
      </div>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed } from 'vue';
import type { BaseMessage } from '../model/message';
import '../styles/MessageCard.css';

interface Props {
  message: BaseMessage;
}

interface Emits {
  (e: 'mark-read', message: BaseMessage): void;
  (e: 'approve', message: BaseMessage): void;
  (e: 'reject', message: BaseMessage): void;
  (e: 'reply', message: BaseMessage): void;
  (e: 'delete', message: BaseMessage): void;
  (e: 'create', message: BaseMessage): void;
  (e: 'update', message: BaseMessage): void;
  (e: 'done', message: BaseMessage): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

const showDetailModal = ref(false);

// 获取发送者名称
const getSenderName = () => {
  switch (props.message.type) {
    case 'system':
      return '系统通知';
    case 'apply':
      return `用户申请`;
    default:
      return `用户 ${props.message.sender_id}`;
  }
};

// 获取消息预览内容（简短版本）
const getMessagePreview = () => {
  const data = props.message.data;
  let content = '';
  
  if (typeof data === 'string') {
    content = data;
  } else if (typeof data === 'object' && data !== null) {
    if (data.data) {
      content = data.data;
    } else if (data.message) {
      content = data.message;
    } else {
      content = JSON.stringify(data);
    }
  } else {
    content = '暂无内容';
  }
  
  // 限制预览内容长度
  return content.length > 50 ? content.substring(0, 50) + '...' : content;
};

// 获取完整消息内容
const getMessageContent = () => {
  const data = props.message.data;
  if (typeof data === 'string') {
    return data;
  } else if (typeof data === 'object' && data !== null) {
    if (data.data) {
      return data.data;
    } else if (data.message) {
      return data.message;
    }
    return JSON.stringify(data);
  }
  return '暂无内容';
};

// 获取消息类型样式
const getTypeClass = () => {
  switch (props.message.type) {
    case 'system':
      return 'type-system';
    case 'apply':
      return 'type-apply';
    case 'comment':
      return 'type-comment';
    case 'like':
      return 'type-like';
    case 'follow':
      return 'type-follow';
    default:
      return '';
  }
};

// 获取消息类型标签
const getTypeLabel = () => {
  switch (props.message.type) {
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

// 获取类型标签颜色
const getTypeTagType = () => {
  switch (props.message.type) {
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

// 获取 action 标签
const getActionLabel = () => {
  switch (props.message.action) {
    case 'create':
      return '创建';
    case 'update':
      return '更新';
    case 'done':
      return '已完成';
    default:
      return '';
  }
};

// 获取 action 标签颜色
const getActionTagType = () => {
  switch (props.message.action) {
    case 'create':
      return 'info';
    case 'update':
      return 'warning';
    case 'done':
      return 'success';
    default:
      return 'default';
  }
};

// 检查按钮显示状态
const shouldShowCreateButton = computed(() => props.message.action === 'create');
const shouldShowUpdateButton = computed(() => props.message.action === 'update');
const shouldShowDoneButton = computed(() => props.message.action === 'done');

// 判断是否可以同意/拒绝
const canApprove = () => {
  return props.message.type === 'apply';
};

const canReject = () => {
  return props.message.type === 'apply';
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

// 处理卡片点击事件
const handleCardClick = () => {
  showDetailModal.value = true;
  if (!props.message.isRead) {
    emit('mark-read', props.message);
  }
};

// 事件处理
const handleMarkRead = () => {
  emit('mark-read', props.message);
};

const handleApprove = () => {
  emit('approve', props.message);
};

const handleReject = () => {
  emit('reject', props.message);
};

const handleReply = () => {
  emit('reply', props.message);
};

const handleDelete = () => {
  emit('delete', props.message);
  showDetailModal.value = false;
};

const handleCreate = () => {
  emit('create', props.message);
};

const handleUpdate = () => {
  emit('update', props.message);
};

const handleDone = () => {
  emit('done', props.message);
};
</script>
