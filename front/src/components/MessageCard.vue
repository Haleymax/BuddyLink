<template>
  <div 
    class="message-card"
    :class="{ 'unread': !message.is_read }"
    @click="handleClick"
  >
    <div class="message-avatar">
      <n-avatar :size="48">
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
        <div class="message-time">{{ formatTime(message.created_at) }}</div>
      </div>
      <div class="message-text">{{ getMessageContent() }}</div>
      <div class="message-type" :class="getTypeClass()">
        {{ getTypeLabel() }}
      </div>
    </div>
    
    <div class="message-actions" @click.stop>
      <n-dropdown :options="menuOptions" @select="handleAction">
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
</template>

<script lang="ts" setup>
import { computed } from 'vue';
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
  (e: 'action', action: string, message: DisplayMessage): void;
}

const props = defineProps<Props>();
const emit = defineEmits<Emits>();

// 获取发送者名称
const getSenderName = () => {
  switch (props.message.type) {
    case 'system':
      return '系统通知';
    case 'apply':
      return `用户 ${props.message.sender_id}`;
    default:
      return `用户 ${props.message.sender_id}`;
  }
};

// 获取消息内容
const getMessageContent = () => {
  const data = props.message.data;
  if (typeof data === 'string') {
    return data;
  } else if (typeof data === 'object' && data !== null) {
    if (data.data) {
      return data.data;
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

// 计算菜单选项
const menuOptions = computed(() => {
  const options = [];
  
  if (!props.message.is_read) {
    options.push({
      label: '标记为已读',
      key: 'mark-read'
    });
  }
  
  options.push({
    label: '查看详情',
    key: 'view-detail'
  });
  
  options.push({
    label: '删除消息',
    key: 'delete'
  });
  
  return options;
});

// 处理点击事件
const handleClick = () => {
  emit('click', props.message);
};

// 处理操作事件
const handleAction = (action: string) => {
  emit('action', action, props.message);
};
</script>

<style scoped>
.message-card {
  display: flex;
  align-items: flex-start;
  gap: 16px;
  padding: 16px;
  border-radius: 8px;
  cursor: pointer;
  transition: all 0.3s ease;
  border-bottom: 1px solid #f0f0f0;
}

.message-card:hover {
  background: #fafafa;
  transform: translateY(-1px);
  box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
}

.message-card.unread {
  background: #e6f4ff;
  border-left: 4px solid #1890ff;
}

.message-card.unread:hover {
  background: #d6f2ff;
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
  width: 12px;
  height: 12px;
  border-radius: 50%;
  background: #ff4d4f;
  border: 2px solid white;
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0% {
    box-shadow: 0 0 0 0 rgba(255, 77, 79, 0.7);
  }
  70% {
    box-shadow: 0 0 0 10px rgba(255, 77, 79, 0);
  }
  100% {
    box-shadow: 0 0 0 0 rgba(255, 77, 79, 0);
  }
}

.message-content {
  flex: 1;
  min-width: 0;
}

.message-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 8px;
}

.sender-name {
  font-weight: 600;
  color: #262626;
  font-size: 16px;
}

.message-time {
  color: #8c8c8c;
  font-size: 14px;
}

.message-text {
  color: #595959;
  font-size: 14px;
  line-height: 1.5;
  margin-bottom: 8px;
  display: -webkit-box;
  -webkit-line-clamp: 2;
  line-clamp: 2;
  -webkit-box-orient: vertical;
  overflow: hidden;
}

.message-type {
  display: inline-block;
  padding: 2px 8px;
  border-radius: 12px;
  font-size: 12px;
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

.message-actions {
  flex-shrink: 0;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .message-card {
    padding: 12px;
  }
  
  .message-header {
    flex-direction: column;
    align-items: flex-start;
    gap: 4px;
  }
}
</style>
