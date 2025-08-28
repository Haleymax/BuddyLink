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
              @mark-read="handleMarkRead"
              @approve="handleApprove"
              @reject="handleReject"
              @reply="handleReply"
              @delete="handleDelete"
              @create="handleCreate"
              @update="handleUpdate"
              @done="handleDone"
            />
          </div>
        </n-tab-pane>

        <n-tab-pane name="apply" tab="申请消息">
          <div class="message-list">
            <MessageCard 
              v-for="message in applyMessages" 
              :key="message.id"
              :message="message"
              @mark-read="handleMarkRead"
              @approve="handleApprove"
              @reject="handleReject"
              @reply="handleReply"
              @delete="handleDelete"
              @create="handleCreate"
              @update="handleUpdate"
              @done="handleDone"
            />
          </div>
        </n-tab-pane>

        <n-tab-pane name="system" tab="系统消息">
          <div class="message-list">
            <MessageCard 
              v-for="message in systemMessages" 
              :key="message.id"
              :message="message"
              @mark-read="handleMarkRead"
              @approve="handleApprove"
              @reject="handleReject"
              @reply="handleReply"
              @delete="handleDelete"
              @create="handleCreate"
              @update="handleUpdate"
              @done="handleDone"
            />
          </div>
        </n-tab-pane>
      </n-tabs>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue';
import { useMessage } from 'naive-ui';
import { useAuthStore } from '../../stores/auth.store';
import type { BaseMessage } from '../../model/message';
import type { ApiResponse } from '../../api/apiResponse';
import { deleteMessageById, getMineMessage } from '../../api/message';
import MessageCard from '../../components/MessageCard.vue';
import '../../styles/Messages.css';

const message = useMessage();
const authStore = useAuthStore();
const activeTab = ref('all');

// 响应式数据
const messages = ref<BaseMessage[]>([]);
const loading = ref(false);

// 计算各类消息
const allMessages = computed(() => messages.value);
const applyMessages = computed(() => messages.value.filter(m => m.type === 'apply'));
const systemMessages = computed(() => messages.value.filter(m => m.type === 'system'));

// 事件处理函数 - 接收来自 MessageCard 的事件
const handleMarkRead = (msg: BaseMessage) => {
  const messageIndex = messages.value.findIndex(m => m.id === msg.id);
  if (messageIndex > -1) {
    messages.value[messageIndex] = { ...msg, isRead: true };
  }
  message.success('已标记为已读');
};

const handleApprove = (msg: BaseMessage) => {
  console.log('Approving message:', msg.id);
  message.success('已同意申请');
  // TODO: 调用API处理申请
};

const handleReject = (msg: BaseMessage) => {
  console.log('Rejecting message:', msg.id);
  message.warning('已拒绝申请');
  // TODO: 调用API处理申请
};

const handleReply = (msg: BaseMessage) => {
  console.log('Replying to message:', msg.id);
  message.info('回复功能');
  // TODO: 实现回复功能
};

const handleDelete = async (msg: BaseMessage) => {
  const index = messages.value.findIndex(m => m.id === msg.id);
  if (index > -1) {
    messages.value.splice(index, 1);
    const token = authStore.token ?? '';
    const response = await deleteMessageById(token, msg.id) as unknown as ApiResponse;

    if (response.code === 200) {
      message.success('消息已删除');
    } else {
      message.error('删除消息失败');
    }
  }
};

const handleCreate = (msg: BaseMessage) => {
  console.log('Creating operation for message:', msg.id);
  message.info('创建操作');
  // TODO: 实现创建操作
};

const handleUpdate = (msg: BaseMessage) => {
  console.log('Updating operation for message:', msg.id);
  message.info('更新操作');
  // TODO: 实现更新操作
};

const handleDone = (msg: BaseMessage) => {
  console.log('Completing operation for message:', msg.id);
  message.success('操作已完成');
  // TODO: 实现完成操作
};

// 标记全部已读
const markAllRead = () => {
  messages.value.forEach(msg => {
    msg.isRead = true;
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
          isRead: msg.is_read || false
        }));
      } else if (response.data && typeof response.data === 'object') {
        messages.value = [{
          ...response.data,
          isRead: response.data.is_read || false
        }];
      } else {
        messages.value = [];
      }
    } else {
      message.error(response.message || '获取消息失败');
      // 使用模拟数据
      messages.value = [
        {
          id: 1,
          sender_id: 2,
          receiver_id: 1,
          type: 'apply',
          data: { message: '申请加入您的活动' },
          action: 'create',
          isRead: false,
          created_at: new Date().toISOString(),
          updated_at: new Date().toISOString(),
          status: 'pending'
        },
        {
          id: 2,
          sender_id: 0,
          receiver_id: 1,
          type: 'system',
          data: { message: '系统维护通知' },
          action: 'done',
          isRead: true,
          created_at: new Date(Date.now() - 3600000).toISOString(),
          updated_at: new Date(Date.now() - 3600000).toISOString(),
          status: 'active'
        }
      ];
    }
  } catch (error) {
    console.error('加载消息失败:', error);
    message.error('加载消息失败，使用模拟数据');
    // 使用模拟数据
    messages.value = [
      {
        id: 1,
        sender_id: 2,
        receiver_id: 1,
        type: 'apply',
        data: { message: '申请加入您的活动' },
        action: 'create',
        isRead: false,
        created_at: new Date().toISOString(),
        updated_at: new Date().toISOString(),
        status: 'pending'
      },
      {
        id: 2,
        sender_id: 0,
        receiver_id: 1,
        type: 'system',
        data: { message: '系统维护通知' },
        action: 'done',
        isRead: true,
        created_at: new Date(Date.now() - 3600000).toISOString(),
        updated_at: new Date(Date.now() - 3600000).toISOString(),
        status: 'active'
      }
    ];
  } finally {
    loading.value = false;
  }
};

onMounted(() => {
  loadMessages();
});
</script>
