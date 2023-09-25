export interface Post {
  id: string;
  title: string;
  image: string;
  body: string;
  is_draft: boolean;
  is_published: boolean;
  created_at: string;
  updated_at: string;
  // TODO: Add author
  // user_id: string;
}
