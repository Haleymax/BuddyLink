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
              {{ card.type }}
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
          <div class="info-value">{{ formatActivityDate(card.date) }}</div>
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
          <div class="info-value">{{ card.people_required || '不限' }}</div>
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
          <div class="info-value">{{ card.gender_required || '不限' }}</div>
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
            <n-icon size="16" color="#52c41a">
              <svg viewBox="0 0 24 24">
                <path fill="currentColor" d="M16,13C15.71,13 15.38,13 15.03,13.05C16.19,13.89 17,15 17,16.5V19H23V16.5C23,14.17 18.33,13 16,13M8,13C5.67,13 1,14.17 1,16.5V19H15V16.5C15,14.17 10.33,13 8,13M8,11A3,3 0 0,0 11,8A3,3 0 0,0 8,5A3,3 0 0,0 5,8A3,3 0 0,0 8,11M16,11A3,3 0 0,0 19,8A3,3 0 0,0 16,5A3,3 0 0,0 13,8A3,3 0 0,0 16,11Z" />
              </svg>
            </n-icon>
            <span>{{ card.people_count }} / {{ card.people_required || '不限' }} 人</span>
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

const formatActivityDate = (dateStr: string): string => {
  const date = new Date(dateStr)
  if (isNaN(date.getTime())) {
    return '时间未设置'
  }
  return date.toLocaleString('zh-CN', {
    year: 'numeric',
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

const parseTags = (tags: string[] | undefined): string[] => {
  if (!tags) return []
  return tags
}

const getStatusType = (status: string) => {
  switch (status) {
    case 'active': return 'success'
    case 'expired': return 'warning'
    case 'completed': return 'info'
    case 'closed': return 'default'
    case 'draft': return 'default'
    case 'deleted': return 'error'
    default: return 'default'
  }
}

const getStatusText = (status: string) => {
  switch (status) {
    case 'active': return '进行中'
    case 'expired': return '已过期'
    case 'completed': return '已完成'
    case 'closed': return '已关闭'
    case 'draft': return '草稿'
    case 'deleted': return '已删除'
    default: return '未知'
  }
}
</script>