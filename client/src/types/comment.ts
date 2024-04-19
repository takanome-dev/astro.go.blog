export interface Comment {
  id: string;
  body: string;
  user_id: string;
  post_id: string;
  created_at: string;
  updated_at: string;
  deleted_at?: string;
  edited_at?: string;
}
