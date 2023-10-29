export const formatDate = (date: string) => {
  return Intl.DateTimeFormat("en-US", {
    year: "numeric",
    month: "long",
    day: "numeric",
  }).format(new Date(date));
};

export const timeAgo = (date: Date) => {
  const now = new Date();
  const secondsAgo = Math.round((+now - +date) / 1000);
  const minutesAgo = Math.round(secondsAgo / 60);
  const hoursAgo = Math.round(minutesAgo / 60);
  const daysAgo = Math.round(hoursAgo / 24);

  const rtf = new Intl.RelativeTimeFormat("en", { numeric: "auto" });

  if (daysAgo > 0) {
    return rtf.format(-daysAgo, "day");
  } else if (hoursAgo > 0) {
    return rtf.format(-hoursAgo, "hour");
  } else if (minutesAgo > 0) {
    return rtf.format(-minutesAgo, "minute");
  } else {
    return rtf.format(-secondsAgo, "second");
  }
};
