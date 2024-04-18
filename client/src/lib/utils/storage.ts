import type { User } from "@/types/user";

export const getAuthToken = () => {
  const token = localStorage.getItem("token");
  return token ? token : null;
};

export const getCurrentUser = () => {
  const user = localStorage.getItem("user");
  return user ? (JSON.parse(user) as User) : null;
};
