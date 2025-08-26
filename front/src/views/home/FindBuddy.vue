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
            :class="{ 'own-activity': isOwnActivity(activity) }"
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
                :disabled="activity.people_required && activity.people_count >= activity.people_required || isOwnActivity(activity) || isActivityApplied(activity)"
                :class="{ 
                  'own-activity-btn': isOwnActivity(activity),
                  'applied-activity-btn': isActivityApplied(activity)
                }"
              >
                <template #icon>
                  <n-icon>
                    <svg viewBox="0 0 24 24" v-if="!isActivityApplied(activity)">
                      <path fill="currentColor" d="M19,13H13V19H11V13H5V11H11V5H13V11H19V13Z"/>
                    </svg>
                    <svg viewBox="0 0 24 24" v-else>
                      <path fill="currentColor" d="M21,7L9,19L3.5,13.5L4.91,12.09L9,16.17L19.59,5.59L21,7Z"/>
                    </svg>
                  </n-icon>
                </template>
                {{ 
                  isOwnActivity(activity) ? '我的活动' :
                  isActivityApplied(activity) ? '已申请' :
                  (activity.people_required && activity.people_count >= activity.people_required ? '已满员' : '加入')
                }}
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

    <!-- 活动详情模态框 -->
    <n-modal 
      v-model:show="showDetailModal" 
      preset="card" 
      title="活动详情"
      style="width: 400px; max-width: 30vw;"
      :mask-closable="true"
    >
      <BuddyCard 
        v-if="selectedActivity"
        :card="selectedActivity"
        :selected="false"
        :hideActions="true"
        @edit="handleEditActivity"
        @delete="handleDeleteActivity"
        @click="() => {}"
        @select="() => {}"
      />
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue'
import { useMessage } from 'naive-ui'
import type { SocialCard } from '../../model/social-cards'
import '../../styles/FindBuddy.css'
import { getSocialCards } from '../../api/social-cards'
import type { ApiResponse } from '../../api/apiResponse'
import { useAuthStore } from '../../stores/auth.store'
import BuddyCard from '../../components/BuddyCard.vue'
import type { BaseMessage } from '../../model/message'
import { addMessage, getAllMessageByUserId, getMineMessage } from '../../api/message'

const message = useMessage()
const authStore = useAuthStore()


const selectedDate = ref<number>(Date.now())
const filterType = ref<string>('')
const filterGender = ref<string>('')
const activities = ref<SocialCard[]>([])
const loading = ref(false)
const showDetailModal = ref(false)
const selectedActivity = ref<SocialCard | null>(null)

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


const Message = ref<BaseMessage | null>({
  sender_id: authStore.user?.id ?? 0,
  receiver_id: 0,
  type: 'apply',
  data: {
    card_id: 0
  }
})

const appliedMessages = ref<BaseMessage[]>([])


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

// 判断是否为自己发布的活动
const isOwnActivity = (activity: SocialCard) => {
  return authStore.user?.id === activity.user_id
}

// 判断是否已经申请过该活动
const isActivityApplied = (activity: SocialCard) => {
  if (!appliedMessages.value || appliedMessages.value.length === 0) return false
  
  // 查找是否有对该活动发布者发送的申请消息
  return appliedMessages.value.some(msg => 
    msg.sender_id === authStore.user?.id && 
    msg.receiver_id === activity.user_id && 
    msg.type === 'apply'
  )
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
  
  if (isActivityApplied(activity)) {
    message.info('您已经申请过这个活动了')
    return
  }
  
  Message.value = {
    sender_id: authStore.user?.id ?? 0,
    receiver_id: activity.user_id,
    type: 'apply',
    data: {
      "data": "Apply to join activity"
    },
  }

  try {
    const response = await addMessage(authStore.token ?? '', Message.value) as unknown as ApiResponse

    if (response.code !== 200) {
      message.error(`发送消息失败：${response.message}`)
      return
    } 
    
    message.success(`发送消息成功：${response.message}`)

    await getMineMessage(authStore.token ?? '', authStore.user?.id ?? 0, 'apply', 'active')

  } catch (error) {
    console.error('申请活动失败:', error)
    message.error('申请活动失败')
  }
}


const FindAllAppliedMessages = async () => {
  try {
    const response = await getMineMessage(authStore.token ?? '', authStore.user?.id ?? 0, 'apply', 'active') as unknown as ApiResponse

    console.log('API 响应:', response) // 调试信息
    
    if (response.code === 200) {
      if (Array.isArray(response.data)) {
        appliedMessages.value = response.data as BaseMessage[]
      } else {
        console.warn('API 返回的数据不是数组:', response.data)
        appliedMessages.value = []
      }
    } else {
      message.error(response.message || '获取申请消息失败')
      appliedMessages.value = []
    }
  } catch (error) {
    console.error('获取申请消息失败:', error)
    message.error('获取申请消息失败')
    appliedMessages.value = []
  }
  
  return appliedMessages.value
}

const viewActivityDetail = (activity: SocialCard) => {
  selectedActivity.value = activity
  showDetailModal.value = true
}

const handleEditActivity = () => {
  message.info('编辑功能暂未实现')
  showDetailModal.value = false
}

const handleDeleteActivity = () => {
  message.info('删除功能暂未实现')
  showDetailModal.value = false
}

const goToCreateActivity = () => {
  message.info('即将跳转到创建活动页面')
}

// 加载活动数据
const loadActivities = async () => {
  try {
    loading.value = true
    const response = await getSocialCards({token: authStore.token ?? ''}) as unknown as ApiResponse
    if (response.code !== 200) {
      message.error(response.message || '获取卡片失败')
      return
    }

    activities.value = response.data.map((card: any) => ({
            id: card.id,
            user_id: card.user_id,
            title: card.title,
            content: card.content,
            type: card.type,
            images: card.images || '',
            gender_required: card.gender_required,
            people_required: card.people_required,
            people_count: card.people_count,
            location: card.location,
            is_private: card.is_private,
            status: card.status,
            date: card.date,
            tags: card.tags
        }))
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
  FindAllAppliedMessages()
})
</script>
