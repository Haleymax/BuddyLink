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
                                <path fill="currentColor" d="M19 13h-6v6h-2v-6H5v-2h6V5h2v6h6v2z" />
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
                    <n-input v-model:value="searchKeyword" placeholder="搜索我的卡片..." clearable class="search-input">
                        <template #prefix>
                            <n-icon>
                                <svg viewBox="0 0 24 24">
                                    <path fill="currentColor"
                                        d="M15.5 14h-.79l-.28-.27C15.41 12.59 16 11.11 16 9.5 16 5.91 13.09 3 9.5 3S3 5.91 3 9.5 5.91 16 9.5 16c1.61 0 3.09-.59 4.23-1.57l.27.28v.79l5 4.99L20.49 19l-4.99-5zm-6 0C7.01 14 5 11.99 5 9.5S7.01 5 9.5 5 14 7.01 14 9.5 11.99 14 9.5 14z" />
                                </svg>
                            </n-icon>
                        </template>
                    </n-input>
                    <n-select v-model:value="filterStatus" :options="statusOptions" placeholder="状态筛选"
                        class="filter-select" />
                    <n-select v-model:value="sortBy" :options="sortOptions" placeholder="排序方式" class="sort-select" />
                </n-space>
            </div>
            <div class="toolbar-right">
                <n-space>
                    <n-button v-if="selectedCards.length > 0" type="error" @click="handleBatchDelete"
                        :loading="batchDeleting">
                        <template #icon>
                            <n-icon>
                                <svg viewBox="0 0 24 24">
                                    <path fill="currentColor"
                                        d="M9,3V4H4V6H5V19A2,2 0 0,0 7,21H17A2,2 0 0,0 19,19V6H20V4H15V3H9M7,6H17V19H7V6M9,8V17H11V8H9M13,8V17H15V8H13Z" />
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
            <buddy-card v-for="card in paginatedCards" :key="card.id" :card="card"
                :selected="card.id ? selectedCards.includes(card.id) : false"
                @select="(checked) => card.id && toggleCardSelection(card.id, checked)" @click="handleCardClick(card)"
                @edit="editCard(card)" @delete="deleteCard(card)" />
        </div>

        <!-- 分页 -->
        <div class="pagination-container">
            <n-pagination v-model:page="currentPage" v-model:page-size="pageSize" :item-count="filteredCards.length"
                :page-sizes="[12, 24, 36, 48]" show-size-picker show-quick-jumper @update:page="handlePageChange"
                @update:page-size="handlePageSizeChange" />
        </div>

        <!-- 空状态 -->
        <div v-if="filteredCards.length === 0" class="empty-state">
            <n-empty description="还没有发布搭子卡片">
                <template #icon>
                    <n-icon size="60" color="#ccc">
                        <svg viewBox="0 0 24 24">
                            <path fill="currentColor"
                                d="M16 4c0-1.11.89-2 2-2s2 .89 2 2-.89 2-2 2-2-.89-2-2zm4 18v-6h2.5l-2.54-7.63A2.99 2.99 0 0 0 16.96 6c-.8 0-1.54.37-2.01.97L12 10.5 8.05 6.97A2.99 2.99 0 0 0 5.04 6c-1.23 0-2.3.75-2.76 1.89L0 16h2.5v6h5v-6H9L7.5 8.5 12 13l4.5-4.5L15 16h1.5v6h5z" />
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

        <!-- 创建/编辑卡片模态框组件 -->
        <AddCardModal 
            v-model:show="showCreateModal" 
            :editing-card="editingCard"
            @save="handleCardSave"
            @cancel="handleCardCancel"
        />
    </div>
</template>

<script lang="ts" setup>
import { ref, computed, onMounted } from 'vue'
import { useMessage } from 'naive-ui'
import { type SocialCard } from '../../model/social-cards'
import BuddyCard from '../../components/BuddyCard.vue'
import AddCardModal from '../../components/AddCardModal.vue'
import '../../styles/SocialCards.css'
import { createSocialCard, deleteSocialCard, getSocialCards, updateSocialCard } from '../../api/social-cards'
import type { ApiResponse } from '../../api/apiResponse'
import { useAuthStore } from '../../stores/auth.store'
import router from '../../router'

const message = useMessage()
const authStore = useAuthStore()

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

const cards = ref<SocialCard[]>([])

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

// 处理卡片保存
const handleCardSave = async (data: { cardData: SocialCard; saveType: string }) => {
    try {
        submitting.value = true
        
        if (editingCard.value) {
            // 编辑模式
            const index = cards.value.findIndex(c => c.id === editingCard.value!.id)
            if (index > -1) {
                Object.assign(cards.value[index], {
                    ...data.cardData,
                    status: data.saveType === 'draft' ? 'draft' : 'active'
                })

                const response = await updateSocialCard(authStore.token ?? '', cards.value[index].id!, cards.value[index]) as unknown as ApiResponse
                if (response.code !== 200) {
                    message.error(response.message || '更新卡片失败，请稍后再试')
                    return
                }
                message.success(data.saveType === 'draft' ? '已保存为草稿' : '卡片更新成功')
            }
        } else {
            // 创建模式
            message.info(String(authStore.$id))
            const newCardData: SocialCard = {
                ...data.cardData,
                user_id: Number(authStore.user?.id),
                images: '',
                people_count: 0,
                status: data.saveType === 'draft' ? 'draft' : 'active',
                value: null
            }

            const token = authStore.token
            if (!token) {
                message.error('请先登录')
                router.push('/login')
                return
            }

            const response = await createSocialCard(token, newCardData) as unknown as ApiResponse
            if (response.code !== 200) {
                message.error(response.message || '创建卡片失败，请稍后再试')
                return
            }
            message.success('卡片创建成功')
            cards.value.unshift(newCardData)
        }

        showCreateModal.value = false
        editingCard.value = null
    } catch (error) {
        message.error('操作失败，请重试')
    } finally {
        submitting.value = false
    }
}

// 处理卡片取消
const handleCardCancel = () => {
    showCreateModal.value = false
    editingCard.value = null
}

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
    showCreateModal.value = true
}

const deleteCard = async (card: SocialCard) => {
    const index = cards.value.findIndex(c => c.id === card.id)
    if (index > -1) {
        cards.value.splice(index, 1)

        const response = await deleteSocialCard(authStore.token ?? '', card.id as number) as unknown as ApiResponse
        if (response.code === 200) {
            message.success('卡片已删除')
        } else {
            message.error('删除失败，请重试')
        }
    }
}

onMounted(async () => {
    try {
        const response = await getSocialCards({
            user_id: authStore.user?.id,
            token: authStore.token ?? ''
        }) as unknown as ApiResponse

        if (response.code !== 200) {
            message.error(response.message || '获取卡片列表失败，请稍后再试')
            return
        }

        cards.value = response.data.map((card: any) => ({
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

        console.log('Cards loaded successfully:', cards.value)
    } catch (error) {
        console.error('Error fetching cards:', error)
        message.error('加载卡片列表时发生错误')
    }
});
</script>
