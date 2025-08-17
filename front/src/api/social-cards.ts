import type { SocialCard } from '../model/social-cards';
import request from '../utils/request';


// 获取社交卡片列表
export const getSocialCards = (params?: {
    page?: number;
    limit?: number;
    search?: string;
    privacy?: 'public' | 'private' | '';
    sort?: 'date' | 'likes' | 'comments';
    user_id?: number;
    token:string;
}) => {
    let url = '/api/v1/social_card/card/';
    if (params?.user_id) {
        url = url + "user/" +params?.user_id
    }
    return request({
        url: url,
        method: 'get',
        headers: {
            "Authorization": params?.token,
        },
    });
};


// 获取单个社交卡片详情
export const getSocialCard = (id: number) => {
    return request({
        url: `/api/v1/social-cards/${id}`,
        method: 'get'
    });
};

// 创建社交卡片
export const createSocialCard = (token:string, data: SocialCard) => {

    return request({
        url: '/api/v1/social_card/card',
        method: 'post',
        headers: {
            "Authorization": token,
        },
        data
    });
};

// 更新社交卡片
export const updateSocialCard = (token:string, id: number, data: Partial<SocialCard>) => {
    return request({
        url: `/api/v1/social_card/card/${id}`,
        method: 'put',
        headers: {
            "Authorization": token,
        },
        data
    });
};

// 删除社交卡片
export const deleteSocialCard = (token:string, id: number) => {
    return request({
        url: `/api/v1/social_card/card/${id}`,
        method: 'delete',
        headers: {
            "Authorization": token,
        }
    });
};

// 点赞/取消点赞
export const toggleLikeSocialCard = (id: number) => {
    return request({
        url: `/api/v1/social-cards/${id}/like`,
        method: 'post'
    });
};

// 获取卡片评论
export const getSocialCardComments = (cardId: number, params?: {
    page?: number;
    limit?: number;
}) => {
    return request({
        url: `/api/v1/social-cards/${cardId}/comments`,
        method: 'get',
        params
    });
};

// 添加评论
export interface CreateCommentData {
    content: string;
    parent_id?: number; // 父评论ID，用于回复
}

export const createSocialCardComment = (cardId: number, data: CreateCommentData) => {
    return request({
        url: `/api/v1/social-cards/${cardId}/comments`,
        method: 'post',
        data
    });
};

// 删除评论
export const deleteSocialCardComment = (cardId: number, commentId: number) => {
    return request({
        url: `/api/v1/social-cards/${cardId}/comments/${commentId}`,
        method: 'delete'
    });
};

// 上传图片
export const uploadCardImage = (file: File) => {
    const formData = new FormData();
    formData.append('image', file);

    return request({
        url: '/api/v1/social-cards/upload-image',
        method: 'post',
        data: formData,
        headers: {
            'Content-Type': 'multipart/form-data'
        }
    });
};

// 获取用户的社交卡片
export const getUserSocialCards = (userId: number, params?: {
    page?: number;
    limit?: number;
    include_private?: boolean;
}) => {
    return request({
        url: `/api/v1/users/${userId}/social-cards`,
        method: 'get',
        params
    });
};
