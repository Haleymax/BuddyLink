<template>
  <div class="dashboard">
    <div class="welcome-section">
      <n-card class="welcome-card" :bordered="false">
        <div class="welcome-content">
          <div class="welcome-text">
            <h2>欢迎回来！</h2>
            <p>今天是个美好的一天，让我们开始工作吧～</p>
          </div>
          <div class="welcome-image">
            <n-icon size="120" color="#18a058">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12 2l3.09 6.26L22 9.27l-5 4.87 1.18 6.88L12 17.77l-6.18 3.25L7 14.14 2 9.27l6.91-1.01L12 2z"/>
              </svg>
            </n-icon>
          </div>
        </div>
      </n-card>
    </div>

    <div class="stats-section">
      <n-grid :cols="4" :x-gap="16">
        <n-gi>
          <n-card class="stat-card" :bordered="false">
            <div class="stat-content">
              <div class="stat-icon">
                <n-icon size="40" color="#1890ff">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M16 6l2.29 2.29-4.88 4.88-4-4L2 16.59 3.41 18l6-6 4 4 6.3-6.29L22 12V6z"/>
                  </svg>
                </n-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">1,234</div>
                <div class="stat-label">总访问量</div>
              </div>
            </div>
          </n-card>
        </n-gi>
        <n-gi>
          <n-card class="stat-card" :bordered="false">
            <div class="stat-content">
              <div class="stat-icon">
                <n-icon size="40" color="#52c41a">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
                  </svg>
                </n-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">89</div>
                <div class="stat-label">已完成任务</div>
              </div>
            </div>
          </n-card>
        </n-gi>
        <n-gi>
          <n-card class="stat-card" :bordered="false">
            <div class="stat-content">
              <div class="stat-icon">
                <n-icon size="40" color="#faad14">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-1 17.93c-3.94-.49-7-3.85-7-7.93 0-.62.08-1.21.21-1.79L9 15v1c0 1.1.9 2 2 2v1.93zm6.9-2.54c-.26-.81-1-1.39-1.9-1.39h-1v-3c0-.55-.45-1-1-1H8v-2h2c.55 0 1-.45 1-1V7h2c1.1 0 2-.9 2-2v-.41c2.93 1.19 5 4.06 5 7.41 0 2.08-.8 3.97-2.1 5.39z"/>
                  </svg>
                </n-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">25</div>
                <div class="stat-label">进行中项目</div>
              </div>
            </div>
          </n-card>
        </n-gi>
        <n-gi>
          <n-card class="stat-card" :bordered="false">
            <div class="stat-content">
              <div class="stat-icon">
                <n-icon size="40" color="#f5222d">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M20 6h-2.18c.11-.31.18-.65.18-1a2.996 2.996 0 0 0-5.5-1.65l-.5.67-.5-.68C10.96 2.54 10.05 2 9 2 7.34 2 6 3.34 6 5c0 .35.07.69.18 1H4c-1.11 0-2 .89-2 2 0 .75.4 1.38 1 1.73V19c0 1.11.89 2 2 2h14c1.11 0 2-.89 2-2V9.73c.6-.35 1-.98 1-1.73 0-1.11-.89-2-2-2zM9 4c.55 0 1 .45 1 1s-.45 1-1 1-1-.45-1-1 .45-1 1-1zm6 0c.55 0 1 .45 1 1s-.45 1-1 1-1-.45-1-1 .45-1 1-1z"/>
                  </svg>
                </n-icon>
              </div>
              <div class="stat-info">
                <div class="stat-number">156</div>
                <div class="stat-label">待办事项</div>
              </div>
            </div>
          </n-card>
        </n-gi>
      </n-grid>
    </div>

    <div class="charts-section">
      <n-grid :cols="2" :x-gap="16">
        <n-gi>
          <n-card title="最近活动" :bordered="false">
            <div class="activity-list">
              <div class="activity-item" v-for="(item, index) in activities" :key="index">
                <div class="activity-icon">
                  <n-icon size="24" :color="item.color">
                    <component :is="item.icon" />
                  </n-icon>
                </div>
                <div class="activity-content">
                  <div class="activity-title">{{ item.title }}</div>
                  <div class="activity-time">{{ item.time }}</div>
                </div>
              </div>
            </div>
          </n-card>
        </n-gi>
        <n-gi>
          <n-card title="快速操作" :bordered="false">
            <div class="quick-actions">
              <n-button type="primary" size="large" block @click="handleQuickAction('create')">
                <template #icon>
                  <n-icon>
                    <svg viewBox="0 0 24 24">
                      <path fill="currentColor" d="M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z"/>
                    </svg>
                  </n-icon>
                </template>
                创建新项目
              </n-button>
              <n-button type="default" size="large" block @click="handleQuickAction('upload')">
                <template #icon>
                  <n-icon>
                    <svg viewBox="0 0 24 24">
                      <path fill="currentColor" d="M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z"/>
                    </svg>
                  </n-icon>
                </template>
                上传文档
              </n-button>
              <n-button type="default" size="large" block @click="handleQuickAction('invite')">
                <template #icon>
                  <n-icon>
                    <svg viewBox="0 0 24 24">
                      <path fill="currentColor" d="M16 4c0-1.11.89-2 2-2s2 .89 2 2-.89 2-2 2-2-.89-2-2zm4 18v-6h2.5l-2.54-7.63A2.99 2.99 0 0 0 16.96 6c-.8 0-1.54.37-2.01.97L12 10.5 8.05 6.97A2.99 2.99 0 0 0 5.04 6c-1.23 0-2.3.75-2.76 1.89L0 16h2.5v6h5v-6H9L7.5 8.5 12 13l4.5-4.5L15 16h1.5v6h5z"/>
                    </svg>
                  </n-icon>
                </template>
                邀请成员
              </n-button>
            </div>
          </n-card>
        </n-gi>
      </n-grid>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, h } from 'vue';
import { useMessage } from 'naive-ui';
import '../../styles/Dashboard.css';

const message = useMessage();

// 活动数据
const activities = ref([
  {
    title: '完成了项目 BuddyLink 的初始化',
    time: '2 小时前',
    color: '#18a058',
    icon: () => h('svg', { viewBox: '0 0 24 24' }, [
      h('path', { 
        fill: 'currentColor', 
        d: 'M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z' 
      })
    ])
  },
  {
    title: '上传了新的设计文档',
    time: '5 小时前',
    color: '#1890ff',
    icon: () => h('svg', { viewBox: '0 0 24 24' }, [
      h('path', { 
        fill: 'currentColor', 
        d: 'M14,2H6A2,2 0 0,0 4,4V20A2,2 0 0,0 6,22H18A2,2 0 0,0 20,20V8L14,2M18,20H6V4H13V9H18V20Z' 
      })
    ])
  },
  {
    title: '邀请了 3 位新成员加入团队',
    time: '1 天前',
    color: '#722ed1',
    icon: () => h('svg', { viewBox: '0 0 24 24' }, [
      h('path', { 
        fill: 'currentColor', 
        d: 'M16 4c0-1.11.89-2 2-2s2 .89 2 2-.89 2-2 2-2-.89-2-2zm4 18v-6h2.5l-2.54-7.63A2.99 2.99 0 0 0 16.96 6c-.8 0-1.54.37-2.01.97L12 10.5 8.05 6.97A2.99 2.99 0 0 0 5.04 6c-1.23 0-2.3.75-2.76 1.89L0 16h2.5v6h5v-6H9L7.5 8.5 12 13l4.5-4.5L15 16h1.5v6h5z' 
      })
    ])
  },
  {
    title: '创建了新的会议安排',
    time: '2 天前',
    color: '#fa8c16',
    icon: () => h('svg', { viewBox: '0 0 24 24' }, [
      h('path', { 
        fill: 'currentColor', 
        d: 'M19 3h-1V1h-2v2H8V1H6v2H5c-1.11 0-1.99.9-1.99 2L3 19c0 1.1.89 2 2 2h14c1.1 0 2-.9 2-2V5c0-1.1-.9-2-2-2zm0 16H5V8h14v11zM7 10h5v5H7z' 
      })
    ])
  }
]);

// 快速操作处理
const handleQuickAction = (action: string) => {
  switch (action) {
    case 'create':
      message.success('开始创建新项目');
      break;
    case 'upload':
      message.info('准备上传文档');
      break;
    case 'invite':
      message.info('打开邀请面板');
      break;
  }
};
</script>
