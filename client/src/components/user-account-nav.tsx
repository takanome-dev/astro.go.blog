import {
  DropdownMenu,
  DropdownMenuContent,
  DropdownMenuItem,
  DropdownMenuPortal,
  DropdownMenuSeparator,
  DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import { UserAvatar } from "@/components/user-avatar";
import { getCurrentUser } from "../lib/utils/storage";

export default function UserAccountNav() {
  const user = getCurrentUser();

  return (
    <DropdownMenu>
      <DropdownMenuTrigger className="focus:ring-0">
        <UserAvatar name={user?.username ?? ""} className="h-8 w-8" />
      </DropdownMenuTrigger>
      <DropdownMenuPortal>
        <DropdownMenuContent
          align="end"
          className="border-slate-400 bg-background"
        >
          <div className="flex items-center justify-start gap-2 p-2">
            <UserAvatar name={user?.username ?? ""} className="h-10 w-10" />
            <div className="flex flex-col space-y-1 leading-none">
              <p className="font-medium">{user?.username}</p>
              <p className="w-[200px] truncate text-sm text-muted-foreground">
                {user?.email}
              </p>
            </div>
          </div>
          <DropdownMenuSeparator />
          <DropdownMenuItem asChild>
            <a href="/dashboard">Dashboard</a>
          </DropdownMenuItem>
          <DropdownMenuSeparator />
          <DropdownMenuItem
            className="cursor-pointer"
            onSelect={(event) => {
              event.preventDefault();
              localStorage.removeItem("user");
              localStorage.removeItem("token");
              window.location.href = "/";
            }}
          >
            Sign out
          </DropdownMenuItem>
        </DropdownMenuContent>
      </DropdownMenuPortal>
    </DropdownMenu>
  );
}
