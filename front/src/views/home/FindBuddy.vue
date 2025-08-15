<template>
  <div class="find-buddy-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-title">
        <h2>寻找搭子</h2>
        <p>浏览日历，发现身边的有趣活动和志趣相投的伙伴</p>
      </div>
      <div class="header-actions">
        <n-space>
          <n-select 
            v-model:value="filterType" 
            :options="activityTypeOptions" 
            placeholder="活动类型"
            style="width: 160px"
            clearable
          />
          <n-select 
            v-model:value="filterGender" 
            :options="genderFilterOptions" 
            placeholder="性别要求"
            style="width: 120px"
            clearable
          />
          <n-button @click="refreshActivities" type="primary" ghost>
            <template #icon>
              <n-icon>
                <svg viewBox="0 0 24 24">
                  <path fill="currentColor" d="M17.65,6.35C16.2,4.9 14.21,4 12,4A8,8 0 0,0 4,12A8,8 0 0,0 12,20C15.73,20 18.84,17.45 19.73,14H17.65C16.83,16.33 14.61,18 12,18A6,6 0 0,1 6,12A6,6 0 0,1 12,6C13.66,6 15.14,6.69 16.22,7.78L13,11H20V4L17.65,6.35Z"/>
                </svg>
              </n-icon>
            </template>
            刷新
          </n-button>
        </n-space>
      </div>
    </div>

    <!-- 日历组件 -->
    <div class="calendar-section">
      <n-card title="活动日历" :bordered="false" class="calendar-card">
        <template #header-extra>
          <n-space>
            <n-tag type="info" size="small">
              <template #icon>
                <n-icon>
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M19,3H18V1H16V3H8V1H6V3H5A2,2 0 0,0 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V5A2,2 0 0,0 19,3M19,19H5V8H19V19Z"/>
                  </svg>
                </n-icon>
              </template>
              点击日期查看活动
            </n-tag>
            <n-button size="small" @click="goToToday" type="primary" ghost>
              今天
            </n-button>
          </n-space>
        </template>
        
        <n-calendar 
          v-model:value="selectedDate"
          @update:value="handleDateSelect"
          :is-date-disabled="isDateDisabled"
          class="activity-calendar"
        >
          <template #date="{ year, month, date }">
            <div class="calendar-date">
              <div class="date-number">{{ date }}</div>
              <div class="activity-dots" v-if="getActivitiesForDate(year, month, date).length > 0">
                <div 
                  v-for="activity in getActivitiesForDate(year, month, date).slice(0, 4)" 
                  :key="activity.id"
                  class="activity-dot"
                  :style="{ backgroundColor: getActivityColor(activity.type) }"
                  :title="`${activity.type}: ${activity.title}`"
                ></div>
                <div 
                  v-if="getActivitiesForDate(year, month, date).length > 4"
                  class="activity-dot more"
                  :title="`还有 ${getActivitiesForDate(year, month, date).length - 4} 个活动`"
                >
                  +{{ getActivitiesForDate(year, month, date).length - 4 }}
                </div>
              </div>
            </div>
          </template>
        </n-calendar>
      </n-card>
    </div>

    <!-- 选中日期的活动列表 -->
    <div class="activities-section">
      <n-card 
        :title="`${formatSelectedDate()} 的活动`" 
        :bordered="false" 
        class="activities-card"
      >
        <template #header-extra>
          <n-tag type="success" size="small" v-if="selectedDateActivities.length > 0">
            {{ selectedDateActivities.length }} 个活动
          </n-tag>
        </template>

        <div v-if="selectedDateActivities.length > 0" class="activities-grid">
          <div 
            v-for="activity in selectedDateActivities" 
            :key="activity.id"
            class="activity-card"
            @click="handleActivityClick(activity)"
          >
            <div class="activity-header">
              <div class="activity-type-badge">
                <n-tag :type="getActivityTagType(activity.type)" size="small">
                  {{ activity.type }}
                </n-tag>
              </div>
              <div class="activity-time">
                {{ formatActivityTime(activity.date) }}
              </div>
            </div>
            
            <div class="activity-content">
              <h3 class="activity-title">{{ activity.title }}</h3>
              <p class="activity-description">{{ activity.content }}</p>
            </div>

            <div class="activity-info">
              <div class="location-info" v-if="activity.location.address">
                <n-icon size="14" color="#666">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12,2A7,7 0 0,1 19,9C19,14.25 12,22 12,22C12,22 5,14.25 5,9A7,7 0 0,1 12,2M12,4A5,5 0 0,0 7,9C7,10 7,12 12,18.71C17,12 17,10 17,9A5,5 0 0,0 12,4M12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5Z"/>
                  </svg>
                </n-icon>
                <span>{{ activity.location.address }}</span>
              </div>
              
              <div class="participant-info">
                <n-icon size="14" color="#666">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M16,4C18.21,4 20,5.79 20,8C20,10.21 18.21,12 16,12C13.79,12 12,10.21 12,8C12,5.79 13.79,4 16,4M16,14C18.67,14 24,15.34 24,18V20H8V18C8,15.34 13.33,14 16,14M8.5,4C10.71,4 12.5,5.79 12.5,8C12.5,10.21 10.71,12 8.5,12C6.29,12 4.5,10.21 4.5,8C4.5,5.79 6.29,4 8.5,4M8.5,14C11.17,14 16.5,15.34 16.5,18V20H0V18C0,15.34 5.33,14 8.5,14Z"/>
                  </svg>
                </n-icon>
                <span>
                  {{ activity.people_count }}/{{ activity.people_required || '不限' }} 人
                  <span v-if="activity.gender_required !== 'any'">
                    · {{ activity.gender_required }}
                  </span>
                </span>
              </div>
            </div>

            <div class="activity-tags" v-if="activity.tags && activity.tags.length > 0">
              <n-tag 
                v-for="tag in activity.tags.slice(0, 3)" 
                :key="tag"
                size="small"
                type="info"
                class="activity-tag"
              >
                {{ tag }}
              </n-tag>
            </div>

            <div class="activity-footer">
              <n-button 
                type="primary" 
                size="small" 
                @click.stop="joinActivity(activity)"
                :disabled="activity.people_required && activity.people_count >= activity.people_required"
              >
                <template #icon>
                  <n-icon>
                    <svg viewBox="0 0 24 24">
                      <path fill="currentColor" d="M19,13H13V19H11V13H5V11H11V5H13V11H19V13Z"/>
                    </svg>
                  </n-icon>
                </template>
                {{ activity.people_required && activity.people_count >= activity.people_required ? '已满员' : '加入' }}
              </n-button>
              
              <n-button size="small" @click.stop="viewActivityDetail(activity)">
                查看详情
              </n-button>
            </div>
          </div>
        </div>

        <div v-else class="empty-activities">
          <n-empty description="该日期暂无活动">
            <template #icon>
              <n-icon size="60" color="#ccc">
                <svg viewBox="0 0 24 24">
                  <path fill="currentColor" d="M19,3H18V1H16V3H8V1H6V3H5A2,2 0 0,0 3,5V19A2,2 0 0,0 5,21H19A2,2 0 0,0 21,19V5A2,2 0 0,0 19,3M19,19H5V8H19V19M7,10H12V15H7"/>
                </svg>
              </n-icon>
            </template>
            <template #extra>
              <n-button type="primary" @click="goToCreateActivity">
                发布新活动
              </n-button>
            </template>
          </n-empty>
        </div>
      </n-card>
    </div>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue'
import { useMessage } from 'naive-ui'
import type { SocialCard } from '../../model/social-cards'
import '../../styles/FindBuddy.css'

const message = useMessage()

// 响应式数据
const selectedDate = ref<number>(Date.now())
const filterType = ref<string>('')
const filterGender = ref<string>('')
const activities = ref<SocialCard[]>([])
const loading = ref(false)

// 活动类型选项
const activityTypeOptions = [
  { label: '全部类型', value: '' },
  { label: '运动健身', value: '运动健身' },
  { label: '娱乐休闲', value: '娱乐休闲' },
  { label: '学习交流', value: '学习交流' },
  { label: '旅行出游', value: '旅行出游' },
  { label: '美食探店', value: '美食探店' },
  { label: '文化艺术', value: '文化艺术' },
  { label: '其他', value: '其他' }
]

// 性别筛选选项
const genderFilterOptions = [
  { label: '不限性别', value: '' },
  { label: '男生', value: '男生' },
  { label: '女生', value: '女生' }
]

// 计算选中日期的活动
const selectedDateActivities = computed(() => {
  const selectedDateStr = new Date(selectedDate.value).toDateString()
  let filteredActivities = activities.value.filter(activity => {
    const activityDateStr = new Date(activity.date).toDateString()
    return activityDateStr === selectedDateStr && activity.status === 'active'
  })

  // 应用筛选条件
  if (filterType.value) {
    filteredActivities = filteredActivities.filter(activity => 
      activity.type === filterType.value
    )
  }

  if (filterGender.value) {
    filteredActivities = filteredActivities.filter(activity => 
      activity.gender_required === filterGender.value || activity.gender_required === 'any'
    )
  }

  return filteredActivities
})

// 获取指定日期的活动
const getActivitiesForDate = (year: number, month: number, date: number) => {
  const targetDate = new Date(year, month - 1, date).toDateString()
  return activities.value.filter(activity => {
    const activityDate = new Date(activity.date).toDateString()
    return activityDate === targetDate && activity.status === 'active'
  })
}

// 获取活动类型对应的颜色
const getActivityColor = (type: string) => {
  const colorMap: Record<string, string> = {
    '运动健身': '#52c41a',      // 绿色
    '娱乐休闲': '#1890ff',      // 蓝色
    '学习交流': '#722ed1',      // 紫色
    '旅行出游': '#fa8c16',      // 橙色
    '美食探店': '#eb2f96',      // 粉红色
    '文化艺术': '#13c2c2',      // 青色
    '其他': '#8c8c8c'           // 灰色
  }
  return colorMap[type] || '#8c8c8c'
}

// 获取活动标签类型
const getActivityTagType = (type: string) => {
  const typeMap: Record<string, string> = {
    '运动健身': 'success',
    '娱乐休闲': 'info',
    '学习交流': 'warning',
    '旅行出游': 'error',
    '美食探店': 'success',
    '文化艺术': 'info',
    '其他': 'default'
  }
  return typeMap[type] || 'default'
}

// 禁用过去的日期
const isDateDisabled = (timestamp: number) => {
  const today = new Date()
  today.setHours(0, 0, 0, 0)
  return timestamp < today.getTime()
}

// 格式化选中日期
const formatSelectedDate = () => {
  const date = new Date(selectedDate.value)
  return date.toLocaleDateString('zh-CN', {
    year: 'numeric',
    month: 'long',
    day: 'numeric',
    weekday: 'long'
  })
}

// 格式化活动时间
const formatActivityTime = (dateStr: string) => {
  const date = new Date(dateStr)
  return date.toLocaleTimeString('zh-CN', {
    hour: '2-digit',
    minute: '2-digit'
  })
}

// 事件处理函数
const handleDateSelect = (value: number) => {
  selectedDate.value = value
}

const goToToday = () => {
  selectedDate.value = Date.now()
}

const refreshActivities = async () => {
  await loadActivities()
  message.success('活动列表已刷新')
}

const handleActivityClick = (activity: SocialCard) => {
  console.log('点击活动:', activity)
  // 这里可以显示活动详情模态框
}

const joinActivity = async (activity: SocialCard) => {
  console.log('加入活动:', activity)
  message.success(`已申请加入活动：${activity.title}`)
  // 这里应该调用加入活动的API
}

const viewActivityDetail = (activity: SocialCard) => {
  console.log('查看活动详情:', activity)
  // 这里可以跳转到活动详情页面或显示详情模态框
}

const goToCreateActivity = () => {
  // 这里可以跳转到创建活动页面
  message.info('即将跳转到创建活动页面')
}

// 示例活动数据
const mockActivities: SocialCard[] = [
  {
    id: 1,
    title: '周末篮球约战',
    content: '寻找篮球爱好者一起打球，技术不限，重在参与和锻炼身体！',
    type: '运动健身',
    date: new Date(Date.now() + 24 * 60 * 60 * 1000).toISOString(), // 明天
    location: {
      latitude: 39.9042,
      longitude: 116.4074,
      address: '朝阳公园篮球场'
    },
    people_count: 3,
    people_required: 6,
    gender_required: 'any',
    status: 'active',
    is_private: false,
    tags: ['篮球', '运动', '朋友'],
    user_id: 1,
    images: '',
    value: null
  },
  {
    id: 2,
    title: '咖啡馆读书分享',
    content: '一起在安静的咖啡馆读书，分享读书心得，欢迎爱书人士加入！',
    type: '学习交流',
    date: new Date(Date.now() + 2 * 24 * 60 * 60 * 1000).toISOString(), // 后天
    location: {
      latitude: 39.9042,
      longitude: 116.4074,
      address: '三里屯星巴克'
    },
    people_count: 2,
    people_required: 4,
    gender_required: 'any',
    status: 'active',
    is_private: false,
    tags: ['读书', '咖啡', '分享'],
    user_id: 2,
    images: '',
    value: null
  },
  {
    id: 3,
    title: '周末爬香山',
    content: '天气不错，约几个人一起爬香山看红叶，拍照打卡，享受自然风光！',
    type: '旅行出游',
    date: new Date(Date.now() + 3 * 24 * 60 * 60 * 1000).toISOString(), // 3天后
    location: {
      latitude: 39.9959,
      longitude: 116.1873,
      address: '香山公园'
    },
    people_count: 1,
    people_required: 5,
    gender_required: 'any',
    status: 'active',
    is_private: false,
    tags: ['爬山', '摄影', '自然'],
    user_id: 3,
    images: '',
    value: null
  },
  {
    id: 4,
    title: '密室逃脱挑战',
    content: '新开的密室逃脱主题很有趣，需要团队配合，寻找喜欢解谜的小伙伴！',
    type: '娱乐休闲',
    date: new Date(Date.now() + 24 * 60 * 60 * 1000).toISOString(), // 明天
    location: {
      latitude: 39.9042,
      longitude: 116.4074,
      address: '王府井密室逃脱馆'
    },
    people_count: 2,
    people_required: 4,
    gender_required: 'any',
    status: 'active',
    is_private: false,
    tags: ['密室逃脱', '解谜', '团队'],
    user_id: 4,
    images: '',
    value: null
  },
  {
    id: 5,
    title: '日式料理制作课',
    content: '学习制作正宗日式料理，从寿司到拉面，一起动手制作美食！',
    type: '美食探店',
    date: new Date(Date.now() + 4 * 24 * 60 * 60 * 1000).toISOString(), // 4天后
    location: {
      latitude: 39.9042,
      longitude: 116.4074,
      address: '日料烹饪教室'
    },
    people_count: 3,
    people_required: 6,
    gender_required: 'any',
    status: 'active',
    is_private: false,
    tags: ['日料', '烹饪', '美食'],
    user_id: 5,
    images: '',
    value: null
  },
  {
    id: 6,
    title: '摄影作品展览',
    content: '参观当代摄影艺术展，一起欣赏和讨论摄影作品，提升艺术鉴赏力！',
    type: '文化艺术',
    date: new Date(Date.now() + 5 * 24 * 60 * 60 * 1000).toISOString(), // 5天后
    location: {
      latitude: 39.9042,
      longitude: 116.4074,
      address: '中国美术馆'
    },
    people_count: 1,
    people_required: 3,
    gender_required: 'any',
    status: 'active',
    is_private: false,
    tags: ['摄影', '艺术', '展览'],
    user_id: 6,
    images: '',
    value: null
  },
  {
    id: 7,
    title: '夜跑健身团',
    content: '每周三次夜跑，强身健体，结识跑友，一起坚持运动，保持健康生活！',
    type: '运动健身',
    date: new Date(Date.now() + 6 * 24 * 60 * 60 * 1000).toISOString(), // 6天后
    location: {
      latitude: 39.9042,
      longitude: 116.4074,
      address: '奥林匹克森林公园'
    },
    people_count: 4,
    people_required: 8,
    gender_required: 'any',
    status: 'active',
    is_private: false,
    tags: ['跑步', '健身', '夜跑'],
    user_id: 7,
    images: '',
    value: null
  },
  {
    id: 8,
    title: '桌游聚会',
    content: '周末桌游聚会，有各种经典桌游，适合新手和老手，快来享受策略游戏的乐趣！',
    type: '娱乐休闲',
    date: new Date(Date.now() + 2 * 24 * 60 * 60 * 1000).toISOString(), // 后天
    location: {
      latitude: 39.9042,
      longitude: 116.4074,
      address: '桌游咖啡厅'
    },
    people_count: 3,
    people_required: 6,
    gender_required: 'any',
    status: 'active',
    is_private: false,
    tags: ['桌游', '策略', '聚会'],
    user_id: 8,
    images: '',
    value: null
  }
]

// 加载活动数据
const loadActivities = async () => {
  try {
    loading.value = true
    // 使用示例数据
    await new Promise(resolve => setTimeout(resolve, 500)) // 模拟网络延迟
    activities.value = mockActivities
    message.success('活动数据加载完成')
  } catch (error) {
    console.error('加载活动失败:', error)
    message.error('加载活动失败')
  } finally {
    loading.value = false
  }
}

onMounted(() => {
  loadActivities()
})
</script>
