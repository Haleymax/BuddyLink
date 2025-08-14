<template>
  <div class="social-cards-container">
    <!-- 页面头部 -->
    <div class="page-header">
      <div class="header-title">
        <h2>我的搭子卡片</h2>
        <p>发布搭子需求，寻找志同道合的伙伴</p>
      </div>
      <div class="header-actions">
        <n-button type="primary" @click="showCreateModal = true" size="large">
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z"/>
              </svg>
            </n-icon>
          </template>
          发布搭子需求
        </n-button>
      </div>
    </div>

    <!-- 工具栏 -->
    <div class="toolbar">
      <div class="toolbar-left">
        <n-space>
          <n-input 
            v-model:value="searchKeyword" 
            placeholder="搜索我的卡片..."
            clearable
            style="width: 280px"
          >
            <template #prefix>
              <n-icon>
                <svg viewBox="0 0 24 24">
                  <path fill="currentColor" d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z"/>
                </svg>
              </n-icon>
            </template>
          </n-input>
          <n-select 
            v-model:value="filterStatus" 
            :options="statusOptions" 
            placeholder="状态筛选"
            style="width: 120px"
          />
          <n-select 
            v-model:value="sortBy" 
            :options="sortOptions" 
            placeholder="排序方式"
            style="width: 120px"
          />
        </n-space>
      </div>
      <div class="toolbar-right">
        <n-space>
          <n-button 
            v-if="selectedCards.length > 0" 
            type="error" 
            @click="handleBatchDelete"
            :loading="batchDeleting"
          >
            <template #icon>
              <n-icon>
                <svg viewBox="0 0 24 24">
                  <path fill="currentColor" d="M9,3V4H4V6H5V19A2,2 0 0,0 7,21H17A2,2 0 0,0 19,19V6H20V4H15V3H9M7,6H17V19H7V6M9,8V17H11V8H9M13,8V17H15V8H13Z"/>
                </svg>
              </n-icon>
            </template>
            删除选中 ({{ selectedCards.length }})
          </n-button>
          <span class="card-count">共 {{ filteredCards.length }} 张卡片</span>
        </n-space>
      </div>
    </div>

    <!-- 卡片网格 -->
    <div class="cards-grid">
      <n-card 
        v-for="card in filteredCards" 
        :key="card.id" 
        class="buddy-card"
        :class="{ 'selected': selectedCards.includes(card.id) }"
        hoverable
        @click="handleCardClick(card)"
      >
        <!-- 卡片头部 -->
        <template #header>
          <div class="card-header">
            <div class="card-header-left">
              <div class="card-selector" @click.stop="toggleCardSelection(card.id)">
                <n-checkbox 
                  :checked="selectedCards.includes(card.id)"
                  @update:checked="(checked: boolean) => toggleCardSelection(card.id, checked)"
                />
              </div>
              <div class="card-title-section">
                <h3 class="card-title">{{ card.title }}</h3>
                <n-tag type="info" size="small" class="activity-tag">
                  {{ card.activity_type }}
                </n-tag>
              </div>
            </div>
            <div class="card-header-right">
              <n-tag 
                :type="getStatusType(card.status)" 
                size="small"
                class="status-tag"
              >
                {{ getStatusText(card.status) }}
              </n-tag>
            </div>
          </div>
        </template>

        <!-- 卡片内容 -->
        <div class="card-content-area">
          <!-- 活动描述 -->
          <div class="card-description">
            <p>{{ truncateText(card.content, 100) }}</p>
          </div>

          <!-- 时间和位置信息 -->
          <div class="card-info-grid">
            <div class="info-item">
              <div class="info-label">
                <n-icon size="16" color="#667eea">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M16.2,16.2L11,13V7H12.5V12.2L17,14.9L16.2,16.2Z"/>
                  </svg>
                </n-icon>
                <span>活动时间</span>
              </div>
              <div class="info-value">{{ formatActivityDate(card.activity_date) }}</div>
            </div>

            <div class="info-item" v-if="card.location">
              <div class="info-label">
                <n-icon size="16" color="#52c41a">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5M12,2A7,7 0 0,0 5,9C5,14.25 12,22 12,22S19,14.25 19,9A7,7 0 0,0 12,2Z"/>
                  </svg>
                </n-icon>
                <span>活动地点</span>
              </div>
              <div class="info-value">
                {{ card.location }}
                <n-tag v-if="card.is_private" type="warning" size="tiny" class="privacy-tag">
                  私密
                </n-tag>
              </div>
            </div>

            <div class="info-item">
              <div class="info-label">
                <n-icon size="16" color="#fa8c16">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M16 4c0-1.11.89-2 2-2s2 .89 2 2-.89 2-2 2-2-.89-2-2zm4 18v-6h2.5l-2.54-7.63A2.99 2.99 0 0 0 16.96 6c-.8 0-1.54.37-2.01.97L12 10.5 8.05 6.97A2.99 2.99 0 0 0 5.04 6c-1.23 0-2.3.75-2.76 1.89L0 16h2.5v6h5v-6H9L7.5 8.5 12 13l4.5-4.5L15 16h1.5v6h5z"/>
                  </svg>
                </n-icon>
                <span>需要人数</span>
              </div>
              <div class="info-value">{{ card.required_people || '不限' }}</div>
            </div>

            <div class="info-item">
              <div class="info-label">
                <n-icon size="16" color="#eb2f96">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z"/>
                  </svg>
                </n-icon>
                <span>性别要求</span>
              </div>
              <div class="info-value">{{ card.gender_requirement || '不限' }}</div>
            </div>
          </div>

          <!-- 标签 -->
          <div class="card-tags" v-if="card.tags">
            <n-space size="small">
              <n-tag 
                v-for="tag in parseTags(card.tags)" 
                :key="tag" 
                size="small"
                type="info"
                round
              >
                #{{ tag }}
              </n-tag>
            </n-space>
          </div>
        </div>

        <!-- 卡片底部 -->
        <template #footer>
          <div class="card-footer">
            <div class="card-stats">
              <div class="stat-item">
                <n-icon size="16" color="#1890ff">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12,9A3,3 0 0,0 9,12A3,3 0 0,0 12,15A3,3 0 0,0 15,12A3,3 0 0,0 12,9M12,17A5,5 0 0,1 7,12A5,5 0 0,1 12,7A5,5 0 0,1 17,12A5,5 0 0,1 12,17M12,4.5C7,4.5 2.73,7.61 1,12C2.73,16.39 7,19.5 12,19.5C17,19.5 21.27,16.39 23,12C21.27,7.61 17,4.5 12,4.5Z"/>
                  </svg>
                </n-icon>
                <span>{{ card.views_count || 0 }} 浏览</span>
              </div>
              <div class="stat-item">
                <n-icon size="16" color="#f5222d">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M12,21.35L10.55,20.03C5.4,15.36 2,12.27 2,8.5 2,5.41 4.42,3 7.5,3C9.24,3 10.91,3.81 12,5.08C13.09,3.81 14.76,3 16.5,3C19.58,3 22,5.41 22,8.5C22,12.27 18.6,15.36 13.45,20.03L12,21.35Z"/>
                  </svg>
                </n-icon>
                <span>{{ card.interested_count || 0 }} 感兴趣</span>
              </div>
              <div class="stat-item">
                <n-icon size="16" color="#52c41a">
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor" d="M9,22A1,1 0 0,1 8,21V18H4A2,2 0 0,1 2,16V4C2,2.89 2.9,2 4,2H20A2,2 0 0,1 22,4V16A2,2 0 0,1 20,18H13.9L10.2,21.71C10,21.9 9.75,22 9.5,22V22H9Z"/>
                  </svg>
                </n-icon>
                <span>{{ card.applications_count || 0 }} 申请</span>
              </div>
            </div>
            <div class="card-actions">
              <n-button size="small" type="primary" ghost @click.stop="editCard(card)">
                编辑
              </n-button>
              <n-button size="small" type="error" ghost @click.stop="deleteCard(card)">
                删除
              </n-button>
            </div>
          </div>
        </template>
      </n-card>
    </div>

    <!-- 空状态 -->
    <div v-if="filteredCards.length === 0" class="empty-state">
      <n-empty description="还没有发布搭子卡片">
        <template #icon>
          <n-icon size="60" color="#ccc">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor" d="M16 4c0-1.11.89-2 2-2s2 .89 2 2-.89 2-2 2-2-.89-2-2zm4 18v-6h2.5l-2.54-7.63A2.99 2.99 0 0 0 16.96 6c-.8 0-1.54.37-2.01.97L12 10.5 8.05 6.97A2.99 2.99 0 0 0 5.04 6c-1.23 0-2.3.75-2.76 1.89L0 16h2.5v6h5v-6H9L7.5 8.5 12 13l4.5-4.5L15 16h1.5v6h5z"/>
            </svg>
          </n-icon>
        </template>
        <template #extra>
          <n-button type="primary" @click="showCreateModal = true">
            发布第一张搭子卡片
          </n-button>
        </template>
      </n-empty>
    </div>

    <!-- 创建/编辑卡片模态框 -->
    <n-modal v-model:show="showCreateModal" preset="dialog" :title="editingCard ? '编辑搭子卡片' : '发布搭子需求'">
      <template #header>
        <div>{{ editingCard ? '编辑搭子卡片' : '发布搭子需求' }}</div>
      </template>
      <n-form 
        :model="newCard" 
        :rules="formRules"
        ref="formRef"
        label-width="100px"
        size="medium"
      >
        <n-form-item label="活动标题" path="title">
          <n-input v-model:value="newCard.title" placeholder="输入搭子活动标题" />
        </n-form-item>
        
        <n-form-item label="活动类型" path="activity_type">
          <n-select 
            v-model:value="newCard.activity_type" 
            :options="activityTypeOptions"
            placeholder="选择活动类型"
          />
        </n-form-item>

        <n-form-item label="活动描述" path="content">
          <n-input 
            v-model:value="newCard.content" 
            type="textarea" 
            placeholder="详细描述活动内容、要求等..."
            :rows="4"
          />
        </n-form-item>

        <n-form-item label="活动时间" path="activity_date">
          <n-date-picker 
            v-model:value="newCard.activity_date" 
            type="datetime"
            placeholder="选择活动时间"
            style="width: 100%"
          />
        </n-form-item>

        <n-form-item label="活动地点" path="location">
          <n-input v-model:value="newCard.location" placeholder="活动地点" />
        </n-form-item>

        <n-form-item label="需要人数">
          <n-input-number v-model:value="newCard.required_people" min="1" max="50" placeholder="不限" />
        </n-form-item>

        <n-form-item label="性别要求">
          <n-select 
            v-model:value="newCard.gender_requirement" 
            :options="genderOptions"
            placeholder="性别要求"
          />
        </n-form-item>

        <n-form-item label="标签">
          <n-input v-model:value="newCard.tags" placeholder="用逗号分隔，如：运动,健身,羽毛球" />
        </n-form-item>

        <n-form-item label="隐私设置">
          <n-switch v-model:value="newCard.is_private" />
          <span style="margin-left: 8px;">{{ newCard.is_private ? '仅自己可见位置' : '公开位置信息' }}</span>
        </n-form-item>
      </n-form>
      <template #action>
        <n-space>
          <n-button @click="closeCreateModal">取消</n-button>
          <n-button type="primary" @click="saveCard" :loading="submitting">
            {{ editingCard ? '保存修改' : '发布卡片' }}
          </n-button>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, reactive } from 'vue'
import { useMessage, type FormInst, type FormRules } from 'naive-ui'
import '../../styles/SocialCards.css'

// 定义接口类型
interface User {
  id: number
  username: string
  avatar?: string
}

interface SocialCard {
  id: number
  user_id: number
  content: string
  title: string
  activity_type: string
  activity_date: number
  required_people?: number
  gender_requirement?: string
  status: 'active' | 'expired' | 'completed'
  images?: string
  views_count: number
  interested_count: number
  applications_count: number
  is_private: boolean
  location?: string
  date: string
  tags?: string
  user?: User
}

const message = useMessage()
const formRef = ref<FormInst | null>(null)

// 响应式数据
const searchKeyword = ref('')
const filterStatus = ref('')
const sortBy = ref('date')
const showCreateModal = ref(false)
const submitting = ref(false)
const batchDeleting = ref(false)
const selectedCards = ref<number[]>([])
const editingCard = ref<SocialCard | null>(null)

// 模拟数据
const cards = ref<SocialCard[]>([
  {
    id: 1,
    user_id: 1,
    title: '周末羽毛球搭子',
    content: '想找个羽毛球搭子，周末一起运动。我是新手，希望找个耐心的伙伴一起学习进步。',
    activity_type: '运动健身',
    activity_date: new Date('2024-08-17T10:00:00').getTime(),
    required_people: 1,
    gender_requirement: '不限',
    status: 'active',
    views_count: 45,
    interested_count: 8,
    applications_count: 3,
    is_private: false,
    location: '体育馆羽毛球场',
    date: '2024-08-14T09:30:00Z',
    tags: '运动,羽毛球,周末',
    user: {
      id: 1,
      username: '当前用户',
      avatar: ''
    }
  },
  {
    id: 2,
    user_id: 1,
    title: '电影院观影搭子',
    content: '新上映的电影想去看，但是一个人去有点无聊，找个伙伴一起去看电影聊天。',
    activity_type: '娱乐休闲',
    activity_date: new Date('2024-08-16T19:30:00').getTime(),
    required_people: 2,
    gender_requirement: '女生',
    status: 'active',
    views_count: 23,
    interested_count: 5,
    applications_count: 1,
    is_private: true,
    location: '万达影城',
    date: '2024-08-13T20:00:00Z',
    tags: '电影,娱乐,周六',
    user: {
      id: 1,
      username: '当前用户',
      avatar: ''
    }
  }
])

// 新建卡片表单
const newCard = reactive({
  title: '',
  content: '',
  activity_type: '',
  activity_date: null as number | null,
  location: '',
  required_people: null as number | null,
  gender_requirement: '',
  tags: '',
  is_private: false
})

// 表单验证规则
const formRules: FormRules = {
  title: [
    { required: true, message: '请输入活动标题', trigger: 'blur' }
  ],
  content: [
    { required: true, message: '请输入活动描述', trigger: 'blur' }
  ],
  activity_type: [
    { required: true, message: '请选择活动类型', trigger: 'change' }
  ],
  activity_date: [
    { required: true, message: '请选择活动时间', trigger: 'change' }
  ],
  location: [
    { required: true, message: '请输入活动地点', trigger: 'blur' }
  ]
}

// 选项数据
const statusOptions = [
  { label: '全部', value: '' },
  { label: '进行中', value: 'active' },
  { label: '已过期', value: 'expired' },
  { label: '已完成', value: 'completed' }
]

const sortOptions = [
  { label: '最新发布', value: 'date' },
  { label: '活动时间', value: 'activity_date' },
  { label: '最多关注', value: 'interested' },
  { label: '最多申请', value: 'applications' }
]

const activityTypeOptions = [
  { label: '运动健身', value: '运动健身' },
  { label: '娱乐休闲', value: '娱乐休闲' },
  { label: '学习交流', value: '学习交流' },
  { label: '旅行出游', value: '旅行出游' },
  { label: '美食探店', value: '美食探店' },
  { label: '文化艺术', value: '文化艺术' },
  { label: '其他', value: '其他' }
]

const genderOptions = [
  { label: '不限', value: '' },
  { label: '男生', value: '男生' },
  { label: '女生', value: '女生' }
]

// 计算属性
const filteredCards = computed(() => {
  let result = cards.value

  // 搜索过滤
  if (searchKeyword.value) {
    result = result.filter(card => 
      card.title.toLowerCase().includes(searchKeyword.value.toLowerCase()) ||
      card.content.toLowerCase().includes(searchKeyword.value.toLowerCase()) ||
      card.activity_type.toLowerCase().includes(searchKeyword.value.toLowerCase())
    )
  }

  // 状态过滤
  if (filterStatus.value) {
    result = result.filter(card => card.status === filterStatus.value)
  }

  // 排序
  result.sort((a, b) => {
    switch (sortBy.value) {
      case 'activity_date':
        return (a.activity_date || 0) - (b.activity_date || 0)
      case 'interested':
        return b.interested_count - a.interested_count
      case 'applications':
        return b.applications_count - a.applications_count
      case 'date':
      default:
        return new Date(b.date).getTime() - new Date(a.date).getTime()
    }
  })

  return result
})

// 方法

const formatActivityDate = (timestamp: number): string => {
  const date = new Date(timestamp)
  return date.toLocaleString('zh-CN', {
    month: 'short',
    day: 'numeric',
    hour: '2-digit',
    minute: '2-digit'
  })
}

const truncateText = (text: string, maxLength: number): string => {
  if (text.length <= maxLength) return text
  return text.slice(0, maxLength) + '...'
}

const parseTags = (tags: string): string[] => {
  if (!tags) return []
  return tags.split(',').map(tag => tag.trim()).filter(tag => tag)
}

const getStatusType = (status: string) => {
  switch (status) {
    case 'active': return 'success'
    case 'expired': return 'warning'
    case 'completed': return 'info'
    default: return 'default'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'active': return '进行中'
    case 'expired': return '已过期'
    case 'completed': return '已完成'
    default: return '未知'
  }
}

const toggleCardSelection = (cardId: number, checked?: boolean) => {
  if (checked === undefined) {
    // 点击卡片区域时切换选择
    const index = selectedCards.value.indexOf(cardId)
    if (index > -1) {
      selectedCards.value.splice(index, 1)
    } else {
      selectedCards.value.push(cardId)
    }
  } else {
    // checkbox 状态变化
    if (checked) {
      if (!selectedCards.value.includes(cardId)) {
        selectedCards.value.push(cardId)
      }
    } else {
      const index = selectedCards.value.indexOf(cardId)
      if (index > -1) {
        selectedCards.value.splice(index, 1)
      }
    }
  }
}

const handleCardClick = (card: SocialCard) => {
  // 点击卡片查看详情或切换选择
  if (selectedCards.value.length > 0) {
    toggleCardSelection(card.id)
  } else {
    message.info(`查看卡片详情：${card.title}`)
  }
}

const handleBatchDelete = async () => {
  if (selectedCards.value.length === 0) return
  
  try {
    batchDeleting.value = true
    // 模拟删除
    await new Promise(resolve => setTimeout(resolve, 1000))
    
    cards.value = cards.value.filter(card => !selectedCards.value.includes(card.id))
    message.success(`成功删除 ${selectedCards.value.length} 张卡片`)
    selectedCards.value = []
  } catch (error) {
    message.error('删除失败，请重试')
  } finally {
    batchDeleting.value = false
  }
}

const editCard = (card: SocialCard) => {
  editingCard.value = card
  Object.assign(newCard, {
    title: card.title,
    content: card.content,
    activity_type: card.activity_type,
    activity_date: card.activity_date,
    location: card.location,
    required_people: card.required_people,
    gender_requirement: card.gender_requirement || '',
    tags: card.tags || '',
    is_private: card.is_private
  })
  showCreateModal.value = true
}

const toggleCardStatus = (card: SocialCard) => {
  const newStatus = card.status === 'active' ? 'completed' : 'active'
  card.status = newStatus
  message.success(`卡片状态已更新为：${getStatusText(newStatus)}`)
}

const deleteCard = (card: SocialCard) => {
  const index = cards.value.findIndex(c => c.id === card.id)
  if (index > -1) {
    cards.value.splice(index, 1)
    message.success('卡片已删除')
  }
}

const closeCreateModal = () => {
  showCreateModal.value = false
  editingCard.value = null
  // 重置表单
  Object.assign(newCard, {
    title: '',
    content: '',
    activity_type: '',
    activity_date: null,
    location: '',
    required_people: null,
    gender_requirement: '',
    tags: '',
    is_private: false
  })
}

const saveCard = async () => {
  if (!formRef.value) return
  
  try {
    await formRef.value.validate()
    submitting.value = true
    
    // 模拟API调用
    setTimeout(() => {
      if (editingCard.value) {
        // 编辑模式
        const index = cards.value.findIndex(c => c.id === editingCard.value!.id)
        if (index > -1) {
          Object.assign(cards.value[index], {
            title: newCard.title,
            content: newCard.content,
            activity_type: newCard.activity_type,
            activity_date: newCard.activity_date,
            location: newCard.location,
            required_people: newCard.required_people,
            gender_requirement: newCard.gender_requirement,
            tags: newCard.tags,
            is_private: newCard.is_private
          })
          message.success('卡片更新成功')
        }
      } else {
        // 创建模式
        const newCardData: SocialCard = {
          id: cards.value.length + 1,
          user_id: 1,
          title: newCard.title,
          content: newCard.content,
          activity_type: newCard.activity_type,
          activity_date: newCard.activity_date || Date.now(),
          location: newCard.location,
          required_people: newCard.required_people || undefined,
          gender_requirement: newCard.gender_requirement || undefined,
          tags: newCard.tags,
          is_private: newCard.is_private,
          status: 'active',
          views_count: 0,
          interested_count: 0,
          applications_count: 0,
          date: new Date().toISOString(),
          user: {
            id: 1,
            username: '当前用户',
            avatar: ''
          }
        }
        
        cards.value.unshift(newCardData)
        message.success('搭子卡片发布成功')
      }
      
      closeCreateModal()
      submitting.value = false
    }, 1000)
    
  } catch (error) {
    message.error('请检查表单输入')
    submitting.value = false
  }
}
</script>
