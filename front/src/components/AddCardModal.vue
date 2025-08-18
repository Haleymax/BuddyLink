<template>
  <!-- 创建/编辑卡片模态框 -->
  <n-modal v-model:show="showModal" preset="dialog" :title="editingCard ? '编辑搭子卡片' : '发布搭子需求'"
    style="width: 90%; max-width: 800px;" class="add-card-modal">
    <template #header>
      <div>{{ editingCard ? '编辑搭子卡片' : '发布搭子需求' }}</div>
    </template>
    <n-form :model="cardData" :rules="formRules" ref="formRef" label-width="100px" size="medium">
      <n-form-item label="活动标题" path="title">
        <n-input v-model:value="cardData.title" placeholder="输入搭子活动标题" />
      </n-form-item>

      <n-form-item label="活动类型" path="type">
        <n-select v-model:value="cardData.type" :options="activityTypeOptions" placeholder="选择活动类型" />
      </n-form-item>

      <n-form-item label="活动描述" path="content">
        <n-input v-model:value="cardData.content" type="textarea" placeholder="详细描述活动内容、要求等..." :rows="4" />
      </n-form-item>

      <n-form-item label="活动时间" path="date">
        <n-date-picker v-model:value="dateTimestamp" type="datetime" placeholder="选择活动时间"
          :is-date-disabled="disablePastDates" :actions="['clear', 'now', 'confirm']"
          :time-picker-props="{
            actions: ['clear', 'now', 'confirm'],
            isHourDisabled: (hour: number) => hour < new Date().getHours() && isToday(dateTimestamp),
            isMinuteDisabled: (minute: number, hour: number) => hour === new Date().getHours() && minute < new Date().getMinutes() && isToday(dateTimestamp)
          }" clearable class="datetime-picker" />
      </n-form-item>

      <n-form-item label="活动地点" path="location">
        <n-space vertical size="medium">
          <!-- 位置信息显示 -->
          <div v-if="selectedLocation">
            <n-card size="small" class="location-selected-card">
              <template #header>
                <n-space align="center">
                  <n-icon color="#52c41a">
                    <svg viewBox="0 0 24 24">
                      <path fill="currentColor"
                        d="M12,2A7,7 0 0,1 19,9C19,14.25 12,22 12,22C12,22 5,14.25 5,9A7,7 0 0,1 12,2M12,4A5,5 0 0,0 7,9C7,10 7,12 12,18.71C17,12 17,10 17,9A5,5 0 0,0 12,4M12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5Z" />
                    </svg>
                  </n-icon>
                  <span class="location-selected-text">已选择位置</span>
                </n-space>
              </template>
              <n-descriptions :column="1" size="small">
                <n-descriptions-item label="地址">
                  {{ selectedLocation.address || '未知地址' }}
                </n-descriptions-item>
                <n-descriptions-item label="坐标">
                  {{ selectedLocation.longitude.toFixed(6) }}, {{ selectedLocation.latitude.toFixed(6) }}
                </n-descriptions-item>
              </n-descriptions>
              <template #action>
                <n-space>
                  <n-button size="small" @click="showLocationSelector = true" type="primary" ghost>
                    重新选择
                  </n-button>
                  <n-button size="small" @click="clearSelectedLocation" type="error" ghost>
                    清除位置
                  </n-button>
                </n-space>
              </template>
            </n-card>
          </div>

          <!-- 位置选择按钮 -->
          <div v-else>
            <n-button type="primary" @click="showLocationSelector = true" block size="large" dashed>
              <template #icon>
                <n-icon>
                  <svg viewBox="0 0 24 24">
                    <path fill="currentColor"
                      d="M12,2A7,7 0 0,1 19,9C19,14.25 12,22 12,22C12,22 5,14.25 5,9A7,7 0 0,1 12,2M12,4A5,5 0 0,0 7,9C7,10 7,12 12,18.71C17,12 17,10 17,9A5,5 0 0,0 12,4M12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5Z" />
                  </svg>
                </n-icon>
              </template>
              点击选择活动地点
            </n-button>
            <n-text depth="3" class="location-help-text">
              支持地图选点、GPS定位、地址搜索
            </n-text>
          </div>
        </n-space>
      </n-form-item>

      <n-form-item label="需要人数">
        <n-input-number v-model:value="cardData.people_required" min="1" max="50" placeholder="不限" />
      </n-form-item>

      <n-form-item label="性别要求">
        <n-select v-model:value="cardData.gender_required" :options="genderOptions" placeholder="性别要求" />
      </n-form-item>

      <n-form-item label="标签">
        <n-dynamic-tags v-model:value="tagList" :max="5" :input-props="{
          placeholder: '输入标签按回车添加'
        }" @update:value="handleTagsUpdate" />
        <span class="tags-help-text">最多添加5个标签，按回车确认</span>
      </n-form-item>

      <n-form-item label="隐私设置">
        <n-switch v-model:value="cardData.is_private" />
        <span class="privacy-help-text">{{ cardData.is_private ? '仅自己可见位置' : '公开位置信息' }}</span>
      </n-form-item>
    </n-form>
    <template #action>
      <n-space>
        <n-button @click="closeModal">取消</n-button>
        <n-dropdown :options="saveOptions" @select="handleSave">
          <n-button type="primary" :loading="submitting">
            {{ editingCard ? '保存修改' : '发布卡片' }}
          </n-button>
        </n-dropdown>
      </n-space>
    </template>
  </n-modal>

  <!-- 位置选择模态框 -->
  <n-modal v-model:show="showLocationSelector" preset="card" title="选择活动地点"
    style="width: 80vw; max-width: 1000px; height: 90vh;" :show-icon="false" :closable="true" :mask-closable="false" class="location-modal">
    <template #header>
      <n-space align="center" justify="space-between" class="modal-header-space">
        <n-space align="center">
          <n-icon color="#1890ff">
            <svg viewBox="0 0 24 24">
              <path fill="currentColor"
                d="M12,2A7,7 0 0,1 19,9C19,14.25 12,22 12,22C12,22 5,14.25 5,9A7,7 0 0,1 12,2M12,4A5,5 0 0,0 7,9C7,10 7,12 12,18.71C17,12 17,10 17,9A5,5 0 0,0 12,4M12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5Z" />
            </svg>
          </n-icon>
          <span class="modal-header-title">选择活动地点</span>
        </n-space>
        <n-tag type="info" size="small">
          <template #icon>
            <n-icon>
              <svg viewBox="0 0 24 24">
                <path fill="currentColor"
                  d="M11,9H13V7H11M12,20C7.59,20 4,16.41 4,12C4,7.59 7.59,4 12,4C16.41,4 20,7.59 20,12C20,16.41 16.41,20 12,20M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M11,17H13V11H11V17Z" />
              </svg>
            </n-icon>
          </template>
          点击地图中"确认使用此位置"按钮完成选择
        </n-tag>
      </n-space>
    </template>

    <div style="height: calc(90vh - 80px); width: 100%; padding: 16px;" class="location-selector-modal">
      <MapComponent @location-selected="handleLocationSelected" :initial-location="selectedLocation" :is-modal="true" />
    </div>
  </n-modal>
</template>

<script lang="ts" setup>
import { ref, reactive, watch, computed } from 'vue'
import { useMessage, type FormInst, type FormRules, type DropdownOption } from 'naive-ui'
import { type SocialCard } from '../model/social-cards'
import type { Location } from '../model/location'
import MapComponent from './MapComponent.vue'
import '../styles/AddCardModal.css'

// Props
interface Props {
  show: boolean
  editingCard?: SocialCard | null
}

const props = withDefaults(defineProps<Props>(), {
  editingCard: null
})

// Emits
interface Emits {
  (e: 'update:show', value: boolean): void
  (e: 'save', data: { cardData: SocialCard; saveType: string }): void
  (e: 'cancel'): void
}

const emit = defineEmits<Emits>()

// Reactive data
const message = useMessage()
const formRef = ref<FormInst | null>(null)
const submitting = ref(false)
const showLocationSelector = ref(false)
const selectedLocation = ref<Location | null>(null)
const tagList = ref<string[]>([])

// Computed
const showModal = computed({
  get: () => props.show,
  set: (value: boolean) => emit('update:show', value)
})

// Card data
const cardData = reactive<SocialCard>({
  user_id: 1,
  title: '',
  content: '',
  type: '',
  images: '',
  gender_required: 'any',
  people_required: null,
  people_count: 0,
  location: { longitude: 0, latitude: 0, address: '' },
  is_private: false,
  status: 'draft',
  date: new Date().toISOString(),
  tags: [],
  id: 0,
  value: null
})

// 用于日期选择器的时间戳值
const dateTimestamp = ref<number | null>(Date.now())

// 监听时间戳变化，同步到 cardData
watch(dateTimestamp, (newTimestamp) => {
  if (newTimestamp) {
    cardData.date = new Date(newTimestamp).toISOString()
  }
})

// Utility functions
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

// Options
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
  { label: '不限', value: 'any' },
  { label: '男生', value: '男生' },
  { label: '女生', value: '女生' }
]

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

// Form validation rules
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
    {
      required: true,
      validator: (_rule: any, value: Location) => {
        if (!value || !value.address) {
          return new Error('请选择活动地点')
        }
        return true
      },
      trigger: 'change'
    }
  ]
}

// Methods
const handleTagsUpdate = (tags: string[]) => {
  tagList.value = tags
  cardData.tags = tags
}

const handleLocationSelected = (location: Location) => {
  selectedLocation.value = location
  cardData.location = location
  showLocationSelector.value = false
  message.success('位置选择成功！')
}

const clearSelectedLocation = () => {
  selectedLocation.value = null
  cardData.location = { longitude: 0, latitude: 0, address: '' }
  message.info('已清除位置信息')
}

const closeModal = () => {
  emit('update:show', false)
  emit('cancel')
  resetForm()
}

const resetForm = () => {
  selectedLocation.value = null
  dateTimestamp.value = Date.now()
  Object.assign(cardData, {
    title: '',
    content: '',
    type: '',
    images: '',
    location: { longitude: 0, latitude: 0, address: '' },
    people_required: null,
    gender_required: 'any',
    people_count: 0,
    tags: [],
    is_private: false,
    date: new Date().toISOString(),
    value: null
  })
  tagList.value = []
}

const handleSave = async (key: string) => {
  if (!formRef.value) return

  try {
    await formRef.value.validate()
    submitting.value = true

    // 确保使用最新的时间戳更新 cardData.date
    if (dateTimestamp.value) {
      cardData.date = new Date(dateTimestamp.value).toISOString()
    }

    // Emit save event with card data and save type
    emit('save', {
      cardData: { ...cardData },
      saveType: key
    })

  } catch (error) {
    message.error('请检查表单输入')
  } finally {
    submitting.value = false
  }
}

// Watch for editing card changes
watch(() => props.editingCard, (editingCard) => {
  if (editingCard) {
    Object.assign(cardData, {
      id: editingCard.id,
      title: editingCard.title,
      content: editingCard.content,
      type: editingCard.type,
      location: editingCard.location,
      people_required: editingCard.people_required,
      gender_required: editingCard.gender_required,
      tags: editingCard.tags || [],
      is_private: editingCard.is_private,
      date: editingCard.date
    })
    tagList.value = editingCard.tags || []
    selectedLocation.value = editingCard.location
    // 将 ISO 字符串转换为时间戳
    dateTimestamp.value = editingCard.date ? new Date(editingCard.date).getTime() : Date.now()
  } else {
    resetForm()
  }
}, { immediate: true })

// Watch for modal show/hide
watch(() => props.show, (show) => {
  if (!show) {
    submitting.value = false
  }
})
</script>
