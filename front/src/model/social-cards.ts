
export interface SocialCard {
  id: number;
  user_id: number;
  content: string;
  title: string;
  type: string;
  images: string;
  gender_required: 'male' | 'female' | 'any';
  people_required: number | null;
  people_count: number;
  is_private: boolean;
  location: string;
  status: 'draft' | 'active' | 'closed' | 'deleted' | 'completed';
  date: string;
  tags?: string;
}
                                               