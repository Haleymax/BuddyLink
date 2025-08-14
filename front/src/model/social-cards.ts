// 社交卡片类型定义

export interface User {
  id: number;
  username: string;
  avatar?: string;
}

export interface SocialCard {
  id: number;
  user_id: number;
  title: string;
  content: string;
  activity_type: string;
  activity_date: number;
  required_people?: number;
  gender_requirement?: string;
  status: 'active' | 'expired' | 'completed';
  views_count: number;
  interested_count: number;
  applications_count: number;
  is_private: boolean;
  location?: string;
  tags?: string;
  date: string;
  user?: User;
}
