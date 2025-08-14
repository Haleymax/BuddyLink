import request from '../utils/request';

// 社交卡片相关接口

// 获取社交卡片列表
export const getSocialCards = (params?: {
  page?: number;
  limit?: number;
  search?: string;
  privacy?: 'public' | 'private' | '';
  sort?: 'date' | 'likes' | 'comments';
  user_id?: number;
}) => {
  return request({
    url: '/api/v1/social-cards',
    method: 'get',
    params
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
export interface CreateSocialCardData {
  title: string;
  content: string;
  images?: string;
  location?: string;
  tags?: string;
  is_private?: boolean;
}

export const createSocialCard = (data: CreateSocialCardData) => {
  return request({
    url: '/api/v1/social-cards',
    method: 'post',
    data
  });
};

// 更新社交卡片
export const updateSocialCard = (id: number, data: Partial<CreateSocialCardData>) => {
  return request({
    url: `/api/v1/social-cards/${id}`,
    method: 'put',
    data
  });
};

// 删除社交卡片
export const deleteSocialCard = (id: number) => {
  return request({
    url: `/api/v1/social-cards/${id}`,
    method: 'delete'
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
