<template>
  <div class="message-detail-card">
    <div class="detail-header">
      <div class="sender-info">
        <div class="avatar-wrapper">
          <n-avatar :size="72" class="sender-avatar">
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
          <div class="avatar-badge" :class="getTypeClass()">
            <n-icon size="14">
              <svg viewBox="0 0 24 24" v-if="message.type === 'system'">
                <path fill="currentColor" d="M12,15.5A3.5,3.5 0 0,1 8.5,12A3.5,3.5 0 0,1 12,8.5A3.5,3.5 0 0,1 15.5,12A3.5,3.5 0 0,1 12,15.5M19.43,12.97C19.47,12.65 19.5,12.33 19.5,12C19.5,11.67 19.47,11.34 19.43,11L21.54,9.37C21.73,9.22 21.78,8.95 21.66,8.73L19.66,5.27C19.54,5.05 19.27,4.96 19.05,5.05L16.56,6.05C16.04,5.66 15.5,5.32 14.87,5.07L14.5,2.42C14.46,2.18 14.25,2 14,2H10C9.75,2 9.54,2.18 9.5,2.42L9.13,5.07C8.5,5.32 7.96,5.66 7.44,6.05L4.95,5.05C4.73,4.96 4.46,5.05 4.34,5.27L2.34,8.73C2.22,8.95 2.27,9.22 2.46,9.37L4.57,11C4.53,11.34 4.5,11.67 4.5,12C4.5,12.33 4.53,12.65 4.57,12.97L2.46,14.63C2.27,14.78 2.22,15.05 2.34,15.27L4.34,18.73C4.46,18.95 4.73,19.03 4.95,18.95L7.44,17.94C7.96,18.34 8.5,18.68 9.13,18.93L9.5,21.58C9.54,21.82 9.75,22 10,22H14C14.25,22 14.46,21.82 14.5,21.58L14.87,18.93C15.5,18.68 16.04,18.34 16.56,17.94L19.05,18.95C19.27,19.03 19.54,18.95 19.66,18.73L21.66,15.27C21.78,15.05 21.73,14.78 21.54,14.63L19.43,12.97Z"/>
              </svg>
              <svg viewBox="0 0 24 24" v-else-if="message.type === 'apply'">
                <path fill="currentColor" d="M9,22A1,1 0 0,1 8,21V18H4A2,2 0 0,1 2,16V4C2,2.89 2.9,2 4,2H20A2,2 0 0,1 22,4V16A2,2 0 0,1 20,18H13.9L10.2,21.71C10,21.9 9.75,22 9.5,22V22H9Z"/>
              </svg>
              <svg viewBox="0 0 24 24" v-else>
                <path fill="currentColor" d="M12,21.35L10.55,20.03C5.4,15.36 2,12.27 2,8.5C2,5.41 4.42,3 7.5,3C9.24,3 10.91,3.81 12,5.08C13.09,3.81 14.76,3 16.5,3C19.58,3 22,5.41 22,8.5C22,12.27 18.6,15.36 13.45,20.03L12,21.35Z"/>
              </svg>
            </n-icon>
          </div>
        </div>
        <div class="sender-details">
          <h3 class="sender-name">{{ getSenderName() }}</h3>
          <p class="sender-subtitle">{{ getSenderSubtitle() }}</p>
        </div>
      </div>
      <div class="message-meta">
        <div class="message-type" :class="getTypeClass()">
          {{ getTypeLabel() }}
        </div>
        <div class="action-type" :class="getActionClass()">
          {{ getActionLabel() }}
        </div>
        <div class="message-time">{{ formatTime(message.created_at) }}</div>
        <div v-if="!message.is_read" class="unread-badge">未读</div>
      </div>
    </div>

    <div class="detail-content">
      <h4 class="content-title">消息内容</h4>
      <div class="message-content" v-html="getFormattedContent()"></div>
      
      <!-- 如果是申请类型消息，显示额外信息 -->
      <div v-if="message.type === 'apply' && getApplicationDetails()" class="application-details">
        <h4 class="content-title">申请详情</h4>
        <div class="application-info">
          <div v-for="(value, key) in getApplicationDetails()" :key="key" class="detail-item">
            <span class="detail-label">{{ getDetailLabel(key) }}：</span>
            <span class="detail-value">{{ value }}</span>
          </div>
        </div>
      </div>

      <!-- 系统消息的额外信息 -->
      <div v-if="message.type === 'system' && getSystemDetails()" class="system-details">
        <h4 class="content-title">系统信息</h4>
        <div class="system-info">
          <div v-for="(value, key) in getSystemDetails()" :key="key" class="detail-item">
            <span class="detail-label">{{ getDetailLabel(key) }}：</span>
            <span class="detail-value">{{ value }}</span>
          </div>
        </div>
      </div>
    </div>

    <div class="detail-actions">
      <n-space>
        <n-button v-if="!message.is_read" type="primary" @click="handleMarkAsRead">
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12,2A7,7 0 0,1 19,9C19,11.38 17.81,13.47 16,14.74V17A1,1 0 0,1 15,18H9A1,1 0 0,1 8,17V14.74C6.19,13.47 5,11.38 5,9A7,7 0 0,1 12,2M9,21H15A1,1 0 0,1 15,22H9A1,1 0 0,1 9,21Z"/>
              </svg>
            </n-icon>
          </template>
          标记为已读
        </n-button>
        
        <!-- 基于 action 的按钮 -->
        <n-button v-if="message.action === 'create'" type="info" @click="handleCreate">
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M19,13H13V19H11V13H5V11H11V5H13V11H19V13Z"/>
              </svg>
            </n-icon>
          </template>
          创建操作
        </n-button>
        
        <n-button v-if="message.action === 'update'" type="warning" @click="handleUpdate">
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z"/>
              </svg>
            </n-icon>
          </template>
          更新操作
        </n-button>
        
        <!-- 基于消息类型的按钮 -->
        <n-button v-if="message.type === 'apply' && canRespond()" type="success" @click="handleApprove">
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M9,20.42L2.79,14.21L5.62,11.38L9,14.77L18.88,4.88L21.71,7.71L9,20.42Z"/>
              </svg>
            </n-icon>
          </template>
          同意申请
        </n-button>
        
        <n-button v-if="message.type === 'apply' && canRespond()" type="error" @click="handleReject">
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
</template>

<script lang="ts" setup>
import type { BaseMessage } from '../model/message';

// 扩展 BaseMessage 接口
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
  (e: 'mark-read', message: DisplayMessage): void;
  (e: 'approve', message: DisplayMessage): void;
  (e: 'reject', message: DisplayMessage): void;
  (e: 'reply', message: DisplayMessage): void;
  (e: 'delete', message: DisplayMessage): void;
  (e: 'create', message: DisplayMessage): void;
  (e: 'update', message: DisplayMessage): void;
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

// 获取发送者副标题
const getSenderSubtitle = () => {
  switch (props.message.type) {
    case 'system':
      return '系统自动发送';
    case 'apply':
      return '搭子活动申请';
    case 'comment':
      return '评论通知';
    case 'like':
      return '点赞通知';
    case 'follow':
      return '关注通知';
    default:
      return '';
  }
};

// 获取格式化的消息内容
const getFormattedContent = () => {
  const data = props.message.data;
  if (typeof data === 'string') {
    return data.replace(/\n/g, '<br>');
  } else if (typeof data === 'object' && data !== null) {
    if (data.message || data.data) {
      const content = data.message || data.data;
      return typeof content === 'string' ? content.replace(/\n/g, '<br>') : JSON.stringify(content);
    }
  }
  return '暂无内容';
};

// 获取申请详情
const getApplicationDetails = () => {
  if (props.message.type !== 'apply') return null;
  
  const data = props.message.data;
  if (typeof data === 'object' && data !== null) {
    const details: Record<string, any> = {};
    
    if (data.activity_title) details.activity_title = data.activity_title;
    if (data.activity_time) details.activity_time = data.activity_time;
    if (data.activity_location) details.activity_location = data.activity_location;
    if (data.applicant_name) details.applicant_name = data.applicant_name;
    if (data.applicant_phone) details.applicant_phone = data.applicant_phone;
    if (data.reason) details.reason = data.reason;
    
    return Object.keys(details).length > 0 ? details : null;
  }
  return null;
};

// 获取系统详情
const getSystemDetails = () => {
  if (props.message.type !== 'system') return null;
  
  const data = props.message.data;
  if (typeof data === 'object' && data !== null) {
    const details: Record<string, any> = {};
    
    if (data.version) details.version = data.version;
    if (data.update_time) details.update_time = data.update_time;
    if (data.affected_features) details.affected_features = data.affected_features;
    if (data.maintenance_time) details.maintenance_time = data.maintenance_time;
    
    return Object.keys(details).length > 0 ? details : null;
  }
  return null;
};

// 获取详情标签
const getDetailLabel = (key: string) => {
  const labelMap: Record<string, string> = {
    activity_title: '活动标题',
    activity_time: '活动时间',
    activity_location: '活动地点',
    applicant_name: '申请人',
    applicant_phone: '联系方式',
    reason: '申请理由',
    version: '版本',
    update_time: '更新时间',
    affected_features: '影响功能',
    maintenance_time: '维护时间'
  };
  return labelMap[key] || key;
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
      return '系统通知';
    case 'apply':
      return '申请消息';
    case 'comment':
      return '评论通知';
    case 'like':
      return '点赞通知';
    case 'follow':
      return '关注通知';
    default:
      return '';
  }
};

// 获取 action 样式
const getActionClass = () => {
  switch (props.message.action) {
    case 'create':
      return 'action-create';
    case 'update':
      return 'action-update';
    default:
      return '';
  }
};

// 获取 action 标签
const getActionLabel = () => {
  switch (props.message.action) {
    case 'create':
      return '创建';
    case 'update':
      return '更新';
    default:
      return '';
  }
};

// 判断是否可以响应申请
const canRespond = () => {
  if (props.message.type !== 'apply') return false;
  const data = props.message.data;
  return typeof data === 'object' && data !== null && (!data.status || data.status === 'pending');
};

// 格式化时间
const formatTime = (timeStr: string) => {
  const time = new Date(timeStr);
  return time.toLocaleString('zh-CN', {
    year: 'numeric',
    month: '2-digit',
    day: '2-digit',
    hour: '2-digit',
    minute: '2-digit',
    second: '2-digit'
  });
};

// 事件处理
const handleMarkAsRead = () => {
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
};

const handleCreate = () => {
  emit('create', props.message);
};

const handleUpdate = () => {
  emit('update', props.message);
};
</script>

<style scoped>
.message-detail-card {
  background: linear-gradient(135deg, #ffffff 0%, #fafbfc 100%);
  border-radius: 20px;
  padding: 32px;
  box-shadow: 
    0 20px 40px rgba(0, 0, 0, 0.08),
    0 8px 24px rgba(0, 0, 0, 0.04),
    0 2px 8px rgba(0, 0, 0, 0.02);
  max-width: 800px;
  margin: 0 auto;
  border: 1px solid rgba(255, 255, 255, 0.8);
  position: relative;
  overflow: hidden;
}

.message-detail-card::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  height: 4px;
  background: linear-gradient(90deg, #667eea 0%, #764ba2 100%);
}

.detail-header {
  display: flex;
  justify-content: space-between;
  align-items: flex-start;
  margin-bottom: 28px;
  padding-bottom: 20px;
  background: rgba(255, 255, 255, 0.6);
  border-radius: 16px;
  padding: 20px;
  border: 1px solid rgba(240, 240, 240, 0.5);
  position: relative;
}

.sender-info {
  display: flex;
  gap: 20px;
  align-items: center;
}

.avatar-wrapper {
  position: relative;
}

.sender-avatar {
  box-shadow: 
    0 8px 16px rgba(0, 0, 0, 0.1),
    0 0 0 4px rgba(255, 255, 255, 0.8),
    0 0 0 6px rgba(102, 126, 234, 0.2);
  transition: all 0.3s ease;
}

.sender-avatar:hover {
  transform: scale(1.05);
  box-shadow: 
    0 12px 24px rgba(0, 0, 0, 0.15),
    0 0 0 4px rgba(255, 255, 255, 0.9),
    0 0 0 8px rgba(102, 126, 234, 0.3);
}

.avatar-badge {
  position: absolute;
  bottom: -2px;
  right: -2px;
  width: 24px;
  height: 24px;
  border-radius: 50%;
  display: flex;
  align-items: center;
  justify-content: center;
  border: 2px solid white;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
}

.sender-details h3 {
  margin: 0 0 6px 0;
  font-size: 22px;
  font-weight: 700;
  color: #1a1a1a;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.sender-details p {
  margin: 0;
  color: #6b7280;
  font-size: 14px;
  font-weight: 500;
}

.message-meta {
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 10px;
}

.message-type {
  padding: 6px 16px;
  border-radius: 20px;
  font-size: 12px;
  font-weight: 600;
  white-space: nowrap;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.action-type {
  padding: 4px 12px;
  border-radius: 16px;
  font-size: 11px;
  font-weight: 500;
  white-space: nowrap;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  box-shadow: 0 1px 4px rgba(0, 0, 0, 0.1);
}

.avatar-badge.type-system {
  background: linear-gradient(135deg, #1890ff 0%, #096dd9 100%);
  color: white;
}

.avatar-badge.type-apply {
  background: linear-gradient(135deg, #fa8c16 0%, #d48806 100%);
  color: white;
}

.avatar-badge.type-comment {
  background: linear-gradient(135deg, #52c41a 0%, #389e0d 100%);
  color: white;
}

.avatar-badge.type-like {
  background: linear-gradient(135deg, #eb2f96 0%, #c41d7f 100%);
  color: white;
}

.avatar-badge.type-follow {
  background: linear-gradient(135deg, #722ed1 0%, #531dab 100%);
  color: white;
}

.type-system {
  background: linear-gradient(135deg, #e6f7ff 0%, #bae7ff 100%);
  color: #1890ff;
  border: 1px solid rgba(24, 144, 255, 0.2);
}

.type-apply {
  background: linear-gradient(135deg, #fff7e6 0%, #ffd591 100%);
  color: #fa8c16;
  border: 1px solid rgba(250, 140, 22, 0.2);
}

.type-comment {
  background: linear-gradient(135deg, #f6ffed 0%, #d9f7be 100%);
  color: #52c41a;
  border: 1px solid rgba(82, 196, 26, 0.2);
}

.type-like {
  background: linear-gradient(135deg, #fff0f6 0%, #ffadd2 100%);
  color: #eb2f96;
  border: 1px solid rgba(235, 47, 150, 0.2);
}

.type-follow {
  background: linear-gradient(135deg, #f9f0ff 0%, #d3adf7 100%);
  color: #722ed1;
  border: 1px solid rgba(114, 46, 209, 0.2);
}

.action-create {
  background: linear-gradient(135deg, #e6f7ff 0%, #bae7ff 100%);
  color: #1890ff;
  border: 1px solid rgba(24, 144, 255, 0.2);
}

.action-update {
  background: linear-gradient(135deg, #fff7e6 0%, #ffd591 100%);
  color: #fa8c16;
  border: 1px solid rgba(250, 140, 22, 0.2);
}

.message-time {
  color: #9ca3af;
  font-size: 13px;
  font-weight: 500;
  display: flex;
  align-items: center;
  gap: 6px;
}

.message-time::before {
  content: '🕒';
  font-size: 12px;
}

.unread-badge {
  background: linear-gradient(135deg, #ff4d4f 0%, #ff7875 100%);
  color: white;
  padding: 4px 12px;
  border-radius: 12px;
  font-size: 11px;
  font-weight: 600;
  text-transform: uppercase;
  letter-spacing: 0.5px;
  box-shadow: 0 2px 8px rgba(255, 77, 79, 0.3);
  animation: pulse 2s infinite;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.8; }
}

.detail-content {
  margin-bottom: 28px;
}

.content-title {
  margin: 0 0 16px 0;
  font-size: 18px;
  font-weight: 700;
  color: #1f2937;
  display: flex;
  align-items: center;
  gap: 8px;
}

.content-title::before {
  content: '💬';
  font-size: 16px;
}

.message-content {
  background: linear-gradient(135deg, #f8fafc 0%, #f1f5f9 100%);
  padding: 24px;
  border-radius: 16px;
  line-height: 1.7;
  color: #374151;
  margin-bottom: 24px;
  border: 1px solid rgba(240, 240, 240, 0.8);
  position: relative;
  font-size: 15px;
}

.message-content::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 4px;
  height: 100%;
  background: linear-gradient(180deg, #667eea 0%, #764ba2 100%);
  border-radius: 2px;
}

.application-details,
.system-details {
  margin-top: 24px;
}

.application-details .content-title::before {
  content: '📋';
}

.system-details .content-title::before {
  content: '⚙️';
}

.application-info {
  background: linear-gradient(135deg, #f0f9ff 0%, #e0f2fe 100%);
  padding: 24px;
  border-radius: 16px;
  border: 1px solid rgba(14, 165, 233, 0.1);
  position: relative;
}

.system-info {
  background: linear-gradient(135deg, #f6ffed 0%, #f0f9ff 100%);
  padding: 24px;
  border-radius: 16px;
  border: 1px solid rgba(82, 196, 26, 0.1);
  position: relative;
}

.application-info::before,
.system-info::before {
  content: '';
  position: absolute;
  top: 0;
  left: 0;
  width: 4px;
  height: 100%;
  border-radius: 2px;
}

.application-info::before {
  background: linear-gradient(180deg, #0ea5e9 0%, #0284c7 100%);
}

.system-info::before {
  background: linear-gradient(180deg, #52c41a 0%, #389e0d 100%);
}

.detail-item {
  display: flex;
  margin-bottom: 12px;
  padding: 8px 0;
  border-bottom: 1px solid rgba(240, 240, 240, 0.5);
}

.detail-item:last-child {
  margin-bottom: 0;
  border-bottom: none;
}

.detail-label {
  font-weight: 600;
  color: #374151;
  min-width: 100px;
  flex-shrink: 0;
  font-size: 14px;
}

.detail-value {
  color: #6b7280;
  word-break: break-all;
  font-size: 14px;
  line-height: 1.5;
}

.detail-actions {
  padding-top: 24px;
  border-top: 2px solid rgba(240, 240, 240, 0.5);
  background: rgba(255, 255, 255, 0.4);
  margin: 0 -32px -32px -32px;
  padding: 24px 32px;
  border-radius: 0 0 20px 20px;
}

/* 按钮样式增强 */
.detail-actions :deep(.n-button) {
  border-radius: 12px;
  font-weight: 600;
  transition: all 0.3s ease;
}

.detail-actions :deep(.n-button:hover) {
  transform: translateY(-2px);
  box-shadow: 0 8px 16px rgba(0, 0, 0, 0.15);
}

.detail-actions :deep(.n-button--info-type) {
  background: linear-gradient(135deg, #1890ff 0%, #096dd9 100%);
  border: none;
}

.detail-actions :deep(.n-button--warning-type) {
  background: linear-gradient(135deg, #fa8c16 0%, #d48806 100%);
  border: none;
}

.detail-actions :deep(.n-button--primary-type) {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  border: none;
}

.detail-actions :deep(.n-button--success-type) {
  background: linear-gradient(135deg, #52c41a 0%, #389e0d 100%);
  border: none;
}

.detail-actions :deep(.n-button--error-type) {
  background: linear-gradient(135deg, #ff4d4f 0%, #ff7875 100%);
  border: none;
}

/* 响应式设计 */
@media (max-width: 768px) {
  .message-detail-card {
    padding: 20px;
    margin: 0 12px;
    border-radius: 16px;
  }
  
  .detail-header {
    flex-direction: column;
    gap: 16px;
    padding: 16px;
  }
  
  .sender-info {
    justify-content: center;
    text-align: center;
  }
  
  .sender-details h3 {
    font-size: 20px;
  }
  
  .message-meta {
    align-items: center;
    flex-direction: row;
    justify-content: center;
    flex-wrap: wrap;
    gap: 8px;
  }
  
  .message-type {
    font-size: 11px;
    padding: 4px 12px;
  }
  
  .action-type {
    font-size: 10px;
    padding: 3px 8px;
  }
  
  .detail-item {
    flex-direction: column;
    gap: 4px;
    padding: 12px 0;
  }
  
  .detail-label {
    min-width: auto;
    font-size: 13px;
  }
  
  .detail-value {
    font-size: 13px;
  }
  
  .detail-actions {
    margin: 0 -20px -20px -20px;
    padding: 20px;
  }
  
  .content-title {
    font-size: 16px;
  }
  
  .message-content {
    padding: 20px;
    font-size: 14px;
  }
}

/* 动画效果 */
@keyframes fadeInUp {
  from {
    opacity: 0;
    transform: translateY(20px);
  }
  to {
    opacity: 1;
    transform: translateY(0);
  }
}

.message-detail-card {
  animation: fadeInUp 0.6s ease-out;
}

.detail-header {
  animation: fadeInUp 0.8s ease-out;
}

.detail-content {
  animation: fadeInUp 1s ease-out;
}

.detail-actions {
  animation: fadeInUp 1.2s ease-out;
}
@media (max-width: 768px) {
  .message-detail-card {
    margin: 8px;
    border-radius: 12px;
  }
  
  .message-header {
    padding: 16px;
  }
  
  .avatar-wrapper {
    width: 48px;
    height: 48px;
  }
  
  .avatar {
    width: 48px;
    height: 48px;
  }
  
  .avatar-badge {
    width: 16px;
    height: 16px;
    font-size: 10px;
  }
  
  .sender-info h3 {
    font-size: 16px;
  }
  
  .sender-info p {
    font-size: 12px;
  }
  
  .time {
    font-size: 12px;
  }
  
  .message-content {
    padding: 16px;
  }
  
  .content-text {
    font-size: 14px;
    line-height: 1.5;
  }
  
  .detail-section {
    padding: 12px 16px;
  }
  
  .detail-section h4 {
    font-size: 14px;
  }
  
  .detail-section p {
    font-size: 13px;
  }
  
  .message-actions {
    padding: 16px;
    gap: 8px;
  }
  
  .action-btn {
    padding: 8px 16px;
    font-size: 13px;
    flex: 1;
  }
}

@media (max-width: 480px) {
  .message-detail-card {
    margin: 4px;
    border-radius: 8px;
  }
  
  .message-header {
    padding: 12px;
  }
  
  .avatar-wrapper {
    width: 40px;
    height: 40px;
  }
  
  .avatar {
    width: 40px;
    height: 40px;
  }
  
  .sender-info h3 {
    font-size: 14px;
  }
  
  .sender-info p {
    font-size: 11px;
  }
  
  .message-content {
    padding: 12px;
  }
  
  .detail-section {
    padding: 8px 12px;
  }
  
  .message-actions {
    padding: 12px;
    flex-direction: column;
  }
  
  .action-btn {
    width: 100%;
    justify-content: center;
  }
}
</style>
