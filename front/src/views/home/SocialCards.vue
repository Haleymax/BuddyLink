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
            v-model:value="newCard.type" 
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
            v-model:value="newCard.date"
            type="datetime"
            placeholder="选择活动时间"
            :is-date-disabled="disablePastDates"
            :actions="['clear', 'now', 'confirm']"
            @update:value="(value: number | null) => newCard.date = value ? new Date(value).toISOString() : ''"
            :time-picker-props="{
              actions: ['clear', 'now', 'confirm'],
              isHourDisabled: (hour: number) => hour < new Date().getHours() && isToday(newCard.date ? Date.parse(newCard.date) : null),
              isMinuteDisabled: (minute: number, hour: number) => hour === new Date().getHours() && minute < new Date().getMinutes() && isToday(newCard.activity_date)
            }"
            clearable
            style="width: 100%"
          />
        </n-form-item>

        <n-form-item label="活动地点" path="location">
          <n-input v-model:value="newCard.location" placeholder="活动地点" />
        </n-form-item>

        <n-form-item label="需要人数">
          <n-input-number v-model:value="newCard.people_required" min="1" max="50" placeholder="不限" />
        </n-form-item>

        <n-form-item label="性别要求">
          <n-select 
            v-model:value="newCard.gender_required" 
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
          <n-dropdown :options="saveOptions" @select="handleSave">
            <n-button type="primary" :loading="submitting">
              {{ editingCard ? '保存修改' : '发布卡片' }}
            </n-button>
          </n-dropdown>
        </n-space>
      </template>
    </n-modal>
  </div>
</template>

<script lang="ts" setup>
import { ref, computed, reactive } from 'vue'
import { useMessage, type FormInst, type FormRules, type DropdownOption } from 'naive-ui'
import { type SocialCard } from '../../model/social-cards'
import BuddyCard from '../../components/BuddyCard.vue'
import '../../styles/SocialCards.css'

const disablePastDates = (timestamp: number) => {
  const today = new Date()
  today.setHours(0, 0, 0, 0)
  return timestamp < today.getTime()
}

const isToday = (timestamp: number | null): boolean => {
  if (!timestamp) return false
  const date = new Date(timestamp)
  const today = new Date()
  return date.getDate() === today.getDate() &&
         date.getMonth() === today.getMonth() &&
         date.getFullYear() === today.getFullYear()
}

const message = useMessage()
const formRef = ref<FormInst | null>(null)

const searchKeyword = ref('')
const filterStatus = ref('')
const sortBy = ref('date')
const showCreateModal = ref(false)
const submitting = ref(false)
const batchDeleting = ref(false)
const selectedCards = ref<number[]>([])
const editingCard = ref<SocialCard | null>(null)

const currentPage = ref(1)
const pageSize = ref(12)

const cards = ref<SocialCard[]>([
  {
    id: 1,
    user_id: 1,
    title: '周末羽毛球搭子',
    content: '想找个羽毛球搭子，周末一起运动。我是新手，希望找个耐心的伙伴一起学习进步。',
    type: '运动健身',
    images: '',
    gender_required: 'any',
    people_required: 1,
    people_count: 0,
    status: 'active',
    is_private: false,
    location: '体育馆羽毛球场',
    date: '2024-08-14T09:30:00Z',
    tags: '运动,羽毛球,周末'
  },
  {
    id: 2,
    user_id: 1,
    title: '电影院观影搭子',
    content: '新上映的电影想去看，但是一个人去有点无聊，找个伙伴一起去看电影聊天。',
    type: '娱乐休闲',
    images: '',
    gender_required: 'any',
    people_required: 2,
    people_count: 0,
    status: 'active',
    is_private: true,
    location: '万达影城',
    date: '2024-08-13T20:00:00Z',
    tags: '电影,娱乐,周六'
  }
])


const newCard = reactive<SocialCard>({
    user_id: 1,
    title: '',
    content: '',
    type: '',
    images: '',
    gender_required: 'any',
    people_required: null,
    people_count: 0,
    location: '',
    is_private: false,
    status: 'draft',
    date: new Date().toISOString(),
    tags: '',
    id: 0
})

// 表单验证规则
const formRules: FormRules = {
  title: [
    { required: true, message: '请输入活动标题', trigger: 'blur' }
  ],
  content: [
    { required: true, message: '请输入活动描述', trigger: 'blur' }
  ],
  type: [
    { required: true, message: '请选择活动类型', trigger: 'change' }
  ],
  location: [
    { required: true, message: '请输入活动地点', trigger: 'blur' }
  ]
}

// 保存选项
const saveOptions: DropdownOption[] = [
  {
    label: '立即发布',
    key: 'publish'
  },
  {
    label: '保存为草稿',
    key: 'draft'
  }
]

// 选项数据
const statusOptions = [
  { label: '全部', value: '' },
  { label: '进行中', value: 'active' },
  { label: '已过期', value: 'expired' },
  { label: '已完成', value: 'completed' },
  { label: '草稿', value: 'draft' }
]

const sortOptions = [
  { label: '最新发布', value: 'date' }
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
      card.type.toLowerCase().includes(searchKeyword.value.toLowerCase())
    )
  }

  // 状态过滤
  if (filterStatus.value) {
    result = result.filter((card: SocialCard) => card.status === filterStatus.value)
  }

  // 排序
  result.sort((a: SocialCard, b: SocialCard) => {
    switch (sortBy.value) {
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
    if (card.id) toggleCardSelection(card.id)
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
    
    cards.value = cards.value.filter(card => !card.id || !selectedCards.value.includes(card.id))
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
    type: card.type,
    location: card.location,
    people_required: card.people_required,
    gender_required: card.gender_required,
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
    type: '',
    images: '',
    location: '',
    people_required: null,
    gender_required: 'any',
    people_count: 0,
    tags: '',
    is_private: false,
    date: new Date().toISOString()
  })
}

const handleSave = async (key: string) => {
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
            type: newCard.type,
            location: newCard.location,
            people_required: newCard.people_required,
            gender_required: newCard.gender_required,
            tags: newCard.tags,
            is_private: newCard.is_private,
            status: key === 'draft' ? 'draft' : 'active',
            date: new Date().toISOString()
          })
          message.success(key === 'draft' ? '已保存为草稿' : '卡片更新成功')
        }
      } else {
        // 创建模式
        const newCardData: SocialCard = {
          user_id: 1,
          title: newCard.title,
          content: newCard.content,
          type: newCard.type,
          images: '',
          location: newCard.location,
          people_required: newCard.people_required,
          people_count: 0,
          gender_required: newCard.gender_required,
          tags: newCard.tags,
          is_private: newCard.is_private,
          status: key === 'draft' ? 'draft' : 'active',
          date: new Date().toISOString()
        }

        console.log(newCardData)

        cards.value.unshift(newCardData)
        message.success(key === 'draft' ? '已保存为草稿' : '搭子卡片发布成功')
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
