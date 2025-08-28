
export interface BaseMessage {
    id: number
    created_at: string
    updated_at: string
    sender_id: number
    receiver_id: number
    type: 'apply' | 'comment' | 'like' | 'follow' | 'system'
    status: string
    data: any
    action: 'create' | 'update' | 'done'
    isRead?: boolean
} 