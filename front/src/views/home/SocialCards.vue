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
      <buddy-card
        v-for="card in paginatedCards"
        :key="card.id"
        :card="card"
        :selected="selectedCards.includes(card.id)"
        @select="(checked) => toggleCardSelection(card.id, checked)"
        @click="handleCardClick(card)"
        @edit="editCard(card)"
        @delete="deleteCard(card)"
      />
    </div>

    <!-- 分页 -->
    <div class="pagination-container">
      <n-pagination
        v-model:page="currentPage"
        v-model:page-size="pageSize"
        :item-count="filteredCards.length"
        :page-sizes="[12, 24, 36, 48]"
        show-size-picker
        show-quick-jumper
        @update:page="handlePageChange"
        @update:page-size="handlePageSizeChange"
      />
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
import { type SocialCard } from '../../model/social-cards'
import BuddyCard from '../../components/BuddyCard.vue'
import '../../styles/SocialCards.css'

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

// 分页相关
const currentPage = ref(1)
const pageSize = ref(12)

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
    result = result.filter((card: SocialCard) => 
      card.title.toLowerCase().includes(searchKeyword.value.toLowerCase()) ||
      card.content.toLowerCase().includes(searchKeyword.value.toLowerCase()) ||
      card.activity_type.toLowerCase().includes(searchKeyword.value.toLowerCase())
    )
  }

  // 状态过滤
  if (filterStatus.value) {
    result = result.filter((card: SocialCard) => card.status === filterStatus.value)
  }

  // 排序
  result.sort((a: SocialCard, b: SocialCard) => {
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

// 分页后的卡片列表
const paginatedCards = computed(() => {
  const start = (currentPage.value - 1) * pageSize.value
  const end = start + pageSize.value
  return filteredCards.value.slice(start, end)
})

// 分页相关方法
const handlePageChange = (page: number) => {
  currentPage.value = page
  window.scrollTo({ top: 0, behavior: 'smooth' })
}

const handlePageSizeChange = (size: number) => {
  pageSize.value = size
  // 当每页数量改变时，可能需要调整当前页码
  const maxPage = Math.ceil(filteredCards.value.length / size)
  if (currentPage.value > maxPage) {
    currentPage.value = maxPage
  }
}// 方法
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
