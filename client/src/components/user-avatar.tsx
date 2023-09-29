import type { AvatarProps } from "@radix-ui/react-avatar";

import { User } from "lucide-react";
import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";

interface UserAvatarProps extends AvatarProps {
  name: string;
  image?: string;
}

export function UserAvatar({ name, image, ...props }: UserAvatarProps) {
  return (
    <Avatar {...props}>
      {image ? (
        <AvatarImage alt="Picture" src={image} />
      ) : (
        <AvatarFallback>
          <span className="sr-only">{name}</span>
          <User className="h-4 w-4" />
        </AvatarFallback>
      )}
    </Avatar>
  );
}
