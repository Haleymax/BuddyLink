<template>
  <n-card 
    class="buddy-card"
    :class="{ 'selected': selected }"
    hoverable
    @click="handleClick"
  >
    <!-- 卡片头部 -->
    <template #header>
      <div class="card-header">
        <div class="card-header-left">
          <div class="card-selector" @click.stop="() => handleSelect()">
            <n-checkbox 
              :checked="selected"
              @update:checked="handleSelect"
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
                <path fill="currentColor" d="M12,2A10,10 0 0,0 2,12A10,10 0 0,0 12,22A10,10 0 0,0 22,12A10,10 0 0,0 12,2M16.2,16.2L11,13V7H12.5V12.2L17,14.9L16.2,16.2Z" />
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
                <path fill="currentColor" d="M12,11.5A2.5,2.5 0 0,1 9.5,9A2.5,2.5 0 0,1 12,6.5A2.5,2.5 0 0,1 14.5,9A2.5,2.5 0 0,1 12,11.5M12,2A7,7 0 0,0 5,9C5,14.25 12,22 12,22S19,14.25 19,9A7,7 0 0,0 12,2Z" />
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
                <path fill="currentColor" d="M16 4c0-1.11.89-2 2-2s2 .89 2 2-.89 2-2 2-2-.89-2-2zm4 18v-6h2.5l-2.54-7.63A2.99 2.99 0 0 0 16.96 6c-.8 0-1.54.37-2.01.97L12 10.5 8.05 6.97A2.99 2.99 0 0 0 5.04 6c-1.23 0-2.3.75-2.76 1.89L0 16h2.5v6h5v-6H9L7.5 8.5 12 13l4.5-4.5L15 16h1.5v6h5z" />
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
                <path fill="currentColor" d="M12 12c2.21 0 4-1.79 4-4s-1.79-4-4-4-4 1.79-4 4 1.79 4 4 4zm0 2c-2.67 0-8 1.34-8 4v2h16v-2c0-2.66-5.33-4-8-4z" />
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
                <path fill="currentColor" d="M12,9A3,3 0 0,0 9,12A3,3 0 0,0 12,15A3,3 0 0,0 15,12A3,3 0 0,0 12,9M12,17A5,5 0 0,1 7,12A5,5 0 0,1 12,7A5,5 0 0,1 17,12A5,5 0 0,1 12,17M12,4.5C7,4.5 2.73,7.61 1,12C2.73,16.39 7,19.5 12,19.5C17,19.5 21.27,16.39 23,12C21.27,7.61 17,4.5 12,4.5Z" />
              </svg>
            </n-icon>
            <span>{{ card.views_count || 0 }} 浏览</span>
          </div>
          <div class="stat-item">
            <n-icon size="16" color="#f5222d">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M12,21.35L10.55,20.03C5.4,15.36 2,12.27 2,8.5 2,5.41 4.42,3 7.5,3C9.24,3 10.91,3.81 12,5.08C13.09,3.81 14.76,3 16.5,3C19.58,3 22,5.41 22,8.5C22,12.27 18.6,15.36 13.45,20.03L12,21.35Z" />
              </svg>
            </n-icon>
            <span>{{ card.interested_count || 0 }} 感兴趣</span>
          </div>
          <div class="stat-item">
            <n-icon size="16" color="#52c41a">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M9,22A1,1 0 0,1 8,21V18H4A2,2 0 0,1 2,16V4C2,2.89 2.9,2 4,2H20A2,2 0 0,1 22,4V16A2,2 0 0,1 20,18H13.9L10.2,21.71C10,21.9 9.75,22 9.5,22V22H9Z" />
              </svg>
            </n-icon>
            <span>{{ card.applications_count || 0 }} 申请</span>
          </div>
        </div>
        <div class="card-actions">
          <n-button size="small" type="primary" ghost @click.stop="$emit('edit')">
            编辑
          </n-button>
          <n-button size="small" type="error" ghost @click.stop="$emit('delete')">
            删除
          </n-button>
        </div>
      </div>
    </template>
  </n-card>
</template>

<script setup lang="ts">
import { type SocialCard } from '../model/social-cards'
import '../styles/BuddyCard.css'

// 组件属性定义
interface Props {
  card: SocialCard
  selected?: boolean
}

// 组件事件定义
interface Emits {
  (e: 'select', value: boolean): void
  (e: 'edit'): void
  (e: 'delete'): void
  (e: 'click'): void
}

const props = withDefaults(defineProps<Props>(), {
  selected: false
})

const emit = defineEmits<Emits>()

// 方法
const handleClick = () => {
  emit('click')
}

const handleSelect = (checked?: boolean) => {
  emit('select', checked ?? !props.selected)
}

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
</script>