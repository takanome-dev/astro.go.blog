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
