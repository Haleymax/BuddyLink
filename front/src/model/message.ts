
export interface BaseMessage {
    sender_id: number
    receiver_id: number
    type: 'apply' | 'comment' | 'like' | 'follow' | 'system'
    data: any
}