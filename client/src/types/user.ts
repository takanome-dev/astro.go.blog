import type { Comment } from "@/types/comment";
import type { Post } from "@/types/post";

export interface User {
  id: string;
  username: string;
  email: string;
  created_at: string;
  updated_at: string;
  name?: string;
  bio?: string;
  image?: string;
  location?: string;
  website_url?: string;
  github_username?: string;
  twitter_username?: string;
  title?: string;
}

export interface CurrentUserKPIs {
  user: User;
  last_three_posts: {
    id: string;
    title: string;
    image: string;
    body: string;
    is_draft: boolean;
    is_published: boolean;
    created_at: string;
    updated_at: string;
    user_id?: string;
  }[];
  last_three_comments: Comment[];
}
