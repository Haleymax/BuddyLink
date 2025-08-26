<template>
  <div 
    class="message-card"
    :class="{ 'unread': !message.is_read }"
    @click="handleClick"
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
      <div v-if="!message.is_read" class="unread-dot"></div>
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
      <n-icon v-if="!message.is_read" size="8" color="#ff4d4f">
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
  </div>
</template>

<script lang="ts" setup>
import type { BaseMessage } from '../model/message';

// 扩展 BaseMessage 接口以包含额外的显示字段
interface DisplayMessage extends BaseMessage {
  id: number;
  created_at: string;
  updated_at: string;
  is_read: boolean;
  status?: 'pending' | 'approved' | 'rejected';
}

interface Props {
  message: DisplayMessage;
}

interface Emits {
  (e: 'click', message: DisplayMessage): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

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

// 处理点击事件
const handleClick = () => {
  emit('click', props.message);
};
</script>

<style scoped>
.message-card {
  display: flex;
  align-items: center;
  gap: 12px;
  padding: 12px 16px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  border-bottom: 1px solid #f0f0f0;
}

.message-card:hover {
  background: #fafafa;
}

.message-card.unread {
  background: #f6ffed;
  border-left: 4px solid #52c41a;
}

.message-card.unread:hover {
  background: #f0f9ff;
}

.message-card:last-child {
  border-bottom: none;
}

.message-avatar {
  position: relative;
  flex-shrink: 0;
}

.unread-dot {
  position: absolute;
  top: -2px;
  right: -2px;
  width: 8px;
  height: 8px;
  border-radius: 50%;
  background: #52c41a;
  border: 2px solid white;
}

.message-content {
  flex: 1;
  min-width: 0;
  overflow: hidden;
}

.message-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 4px;
}

.sender-name {
  font-weight: 600;
  color: #262626;
  font-size: 14px;
}

.message-type {
  display: inline-block;
  padding: 2px 6px;
  border-radius: 10px;
  font-size: 11px;
  font-weight: 500;
}

.type-system {
  background: #e6f7ff;
  color: #1890ff;
}

.type-apply {
  background: #fff7e6;
  color: #fa8c16;
}

.type-comment {
  background: #f6ffed;
  color: #52c41a;
}

.type-like {
  background: #fff0f6;
  color: #eb2f96;
}

.type-follow {
  background: #f9f0ff;
  color: #722ed1;
}

.message-text {
  color: #595959;
  font-size: 13px;
  line-height: 1.4;
  margin-bottom: 4px;
  overflow: hidden;
  text-overflow: ellipsis;
  white-space: nowrap;
}

.message-time {
  color: #8c8c8c;
  font-size: 12px;
}

.message-indicator {
  display: flex;
  align-items: center;
  gap: 4px;
  flex-shrink: 0;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .message-card {
    padding: 10px 12px;
  }
  
  .message-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 2px;
  }
  
  .sender-name {
    font-size: 13px;
  }
  
  .message-text {
    font-size: 12px;
  }
}
</style>
