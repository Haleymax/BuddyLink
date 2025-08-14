<template>
  <div class="home-container">
    <!-- 左侧菜单 -->
    <div class="sidebar" :class="{ collapsed: collapsed }">
      <div class="sidebar-header">
        <div class="logo">
          <n-icon size="32" color="#18a058">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M12 2C6.48 2 2 6.48 2 12s4.48 10 10 10 10-4.48 10-10S17.52 2 12 2zm-2 15l-5-5 1.41-1.41L10 14.17l7.59-7.59L19 8l-9 9z"/>
            </svg>
          </n-icon>
          <span class="logo-text">BuddyLink</span>
        </div>
        <div class="user-info">
          <n-avatar round size="small" :src="user_info.avatar || undefined">
            <template v-if="!user_info.avatar">
              <n-icon>
                <svg viewBox="0 0 24 24">
                  <path fill="currentColor" d="M12,4A4,4 0 0,1 16,8A4,4 0 0,1 12,12A4,4 0 0,1 8,8A4,4 0 0,1 12,4M12,14C16.42,14 20,15.79 20,18V20H4V18C4,15.79 7.58,14 12,14Z"/>
                </svg>
              </n-icon>
            </template>
          </n-avatar>
          <span class="username">{{ user_info.username ? `欢迎回来，${user_info.username}` : '欢迎回来' }}</span>
        </div>
      </div>

      <div class="menu-container">
        <n-menu
          v-model:value="activeKey"
          :collapsed="collapsed"
          :collapsed-width="64"
          :collapsed-icon-size="22"
          :options="menuOptions"
          @update:value="handleMenuSelect"
        />
      </div>

      <div class="sidebar-footer">
        <n-button
          quaternary
          circle
          @click="collapsed = !collapsed"
          class="collapse-btn"
        >
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24" style="transition: transform 0.3s ease;">
                <path fill="currentColor" :d="collapsed ? 'M4 6h16v2H4zm0 5h16v2H4zm0 5h16v2H4z' : 'M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z'" />
              </svg>
            </n-icon>
          </template>
        </n-button>
      </div>
    </div>

    <!-- 右侧内容区域 -->
    <div class="main-content">
      <!-- 顶部工具栏 -->
      <div class="content-header">
        <div class="breadcrumb">
          <n-breadcrumb>
            <n-breadcrumb-item>首页</n-breadcrumb-item>
            <n-breadcrumb-item>{{ getCurrentPageTitle() }}</n-breadcrumb-item>
          </n-breadcrumb>
        </div>
        <div class="header-actions">
          <n-button quaternary circle>
            <template #icon>
              <n-icon>
                <svg viewBox="0 0 24 24">
                  <path fill="currentColor" d="M12 22c1.1 0 2-.9 2-2h-4c0 1.1.89 2 2 2zm6-6v-5c0-3.07-1.64-5.64-4.5-6.32V4c0-.83-.67-1.5-1.5-1.5s-1.5.67-1.5 1.5v.68C7.63 5.36 6 7.92 6 11v5l-2 2v1h16v-1l-2-2z"/>
                </svg>
              </n-icon>
            </template>
          </n-button>
          <n-dropdown :options="userMenuOptions" @select="handleUserMenuSelect">
            <n-button quaternary circle>
              <template #icon>
                <n-icon>
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12 2C13.1 2 14 2.9 14 4C14 5.1 13.1 6 12 6C10.9 6 10 5.1 10 4C10 2.9 10.9 2 12 2ZM21 9V7L15 1H5C3.89 1 3 1.89 3 3V21C3 22.11 3.89 23 5 23H19C20.11 23 21 22.11 21 21V9M19 9H14V4H19V9Z"/>
                  </svg>
                </n-icon>
              </template>
            </n-button>
          </n-dropdown>
        </div>
      </div>

      <!-- 动态内容区域 -->
      <div class="content-body">
        <component :is="currentComponent" />
      </div>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, h, onMounted} from 'vue';
import { useMessage } from 'naive-ui';
import type { MenuOption } from 'naive-ui';
import type { User } from '../stores/auth.store';
import '../styles/Home.css';
import { useAuthStore } from '../stores/auth.store'

// 导入业务组件
import Dashboard from './home/Dashboard.vue';
import Profile from './home/Profile.vue';
import Settings from './home/Settings.vue';
import Messages from './home/Messages.vue';
import MapView from './home/MapView.vue';
import { useFetchUserInfo } from '../common/get_user_info';
import router from '../router';

const message = useMessage();
const collapsed = ref(false);
const activeKey = ref('dashboard');

const user_info = ref<User>({})
const authStore = useAuthStore()
const { fetchUserInfoOnce } = useFetchUserInfo()

// 菜单选项
const menuOptions: MenuOption[] = [
  {
    label: '仪表盘',
    key: 'dashboard',
    icon: () => h('div', { class: 'menu-icon' }, [
      h('svg', { viewBox: '0 0 24 24' }, [
        h('path', { 
          fill: 'currentColor', 
          d: 'M3 13h8V3H3v10zm0 8h8v-6H3v6zm10 0h8V11h-8v10zm0-18v6h8V3h-8z' 
        })
      ])
    ])
  },
  {
    label: '消息中心',
    key: 'messages',
    icon: () => h('div', { class: 'menu-icon' }, [
      h('svg', { viewBox: '0 0 24 24' }, [
        h('path', { 
          fill: 'currentColor', 
          d: 'M20 2H4c-1.1 0-1.99.9-1.99 2L2 22l4-4h14c1.1 0 2-.9 2-2V4c0-1.1-.9-2-2-2zm-2 12H6v-2h12v2zm0-3H6V9h12v2zm0-3H6V6h12v2z' 
        })
      ])
    ])
  },
  {
    label: '地图定位',
    key: 'map',
    icon: () => h('div', { class: 'menu-icon' }, [
      h('svg', { viewBox: '0 0 24 24' }, [
        h('path', { 
          fill: 'currentColor', 
          d: 'M12,2C15.31,2 18,4.66 18,7.95C18,12.41 12,19 12,19S6,12.41 6,7.95C6,4.66 8.69,2 12,2M12,6A2,2 0 0,0 10,8A2,2 0 0,0 12,10A2,2 0 0,0 14,8A2,2 0 0,0 12,6Z' 
        })
      ])
    ])
  },
  {
    label: '个人资料',
    key: 'profile',
    icon: () => h('div', { class: 'menu-icon' }, [
      h('svg', { viewBox: '0 0 24 24' }, [
        h('path', { 
          fill: 'currentColor', 
          d: 'M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z' 
        })
      ])
    ])
  },
  {
    label: '设置',
    key: 'settings',
    icon: () => h('div', { class: 'menu-icon' }, [
      h('svg', { viewBox: '0 0 24 24' }, [
        h('path', { 
          fill: 'currentColor', 
          d: 'M19.14,12.94c0.04-0.3,0.06-0.61,0.06-0.94c0-0.32-0.02-0.64-0.07-0.94l2.03-1.58c0.18-0.14,0.23-0.41,0.12-0.61 l-1.92-3.32c-0.12-0.22-0.37-0.29-0.59-0.22l-2.39,0.96c-0.5-0.38-1.03-0.7-1.62-0.94L14.4,2.81c-0.04-0.24-0.24-0.41-0.48-0.41 h-3.84c-0.24,0-0.43,0.17-0.47,0.41L9.25,5.35C8.66,5.59,8.12,5.92,7.63,6.29L5.24,5.33c-0.22-0.08-0.47,0-0.59,0.22L2.74,8.87 C2.62,9.08,2.66,9.34,2.86,9.48l2.03,1.58C4.84,11.36,4.82,11.69,4.82,12s0.02,0.64,0.07,0.94l-2.03,1.58 c-0.18,0.14-0.23,0.41-0.12,0.61l1.92,3.32c0.12,0.22,0.37,0.29,0.59,0.22l2.39-0.96c0.5,0.38,1.03,0.7,1.62,0.94l0.36,2.54 c0.05,0.24,0.24,0.41,0.48,0.41h3.84c0.24,0,0.44-0.17,0.47-0.41l0.36-2.54c0.59-0.24,1.13-0.56,1.62-0.94l2.39,0.96 c0.22,0.08,0.47,0,0.59-0.22l1.92-3.32c0.12-0.22,0.07-0.47-0.12-0.61L19.14,12.94z M12,15.6c-1.98,0-3.6-1.62-3.6-3.6 s1.62-3.6,3.6-3.6s3.6,1.62,3.6,3.6S13.98,15.6,12,15.6z' 
        })
      ])
    ])
  }
];

// 用户菜单选项
const userMenuOptions = [
  {
    label: '个人中心',
    key: 'profile'
  },
  {
    label: '账户设置',
    key: 'account-settings'
  },
  {
    type: 'divider'
  },
  {
    label: '退出登录',
    key: 'logout'
  }
];

// 当前显示的组件
const currentComponent = computed(() => {
  switch (activeKey.value) {
    case 'dashboard':
      return Dashboard;
    case 'messages':
      return Messages;
    case 'map':
      return MapView;
    case 'profile':
      return Profile;
    case 'settings':
      return Settings;
    default:
      return Dashboard;
  }
});

// 获取当前页面标题
const getCurrentPageTitle = () => {
  const pageMap: Record<string, string> = {
    dashboard: '仪表盘',
    messages: '消息中心',
    map: '地图定位',
    profile: '个人资料',
    settings: '设置'
  };
  return pageMap[activeKey.value] || '仪表盘';
};

// 菜单选择处理
const handleMenuSelect = (key: string) => {
  activeKey.value = key;
};

// 用户菜单选择处理
const handleUserMenuSelect = (key: string) => {
  switch (key) {
    case 'profile':
      activeKey.value = 'profile';
      break;
    case 'account-settings':
      activeKey.value = 'settings';
      break;
    case 'logout':
      message.success('退出登录成功');
      // 这里添加退出登录的逻辑
      break;
  }
};

onMounted(async () => {
  console.log('Home component mounted, checking authentication...');
  console.log('Current token:', authStore.token);
  
  if (!authStore.token) {
    console.log('No token found, redirecting to login');
    message.error('未登录，请先登录');
    router.push('/login');
    return;
  }

  const res = await fetchUserInfoOnce({error: message.error});
  console.log('fetchUserInfoOnce result:', res);
  console.log('User info:', res.data);
  if (res.ok && res.data) {
    user_info.value = res.data;
    console.log('User info loaded successfully:', res.data);
  } else {
    console.log('Failed to get user info, clearing auth and redirecting');
    message.error('获取用户信息失败，请重新登录');
    authStore.clearAuth();
    router.push('/login');
  }
})

</script>
