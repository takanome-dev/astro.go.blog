interface Validation {
  valid: boolean;
  message: string;
}

export const validateTwitterUsername = (username: string): Validation => {
  if (username.length > 15)
    return {
      valid: false,
      message: "Username must be less than 15 characters",
    };
  else if (
    ["admin", "twitter"].some((reserved) =>
      username.toLocaleLowerCase().includes(reserved)
    )
  )
    return {
      valid: false,
      message: "Username contains reserved word 'admin' or 'twitter'",
    };
  else if (!String(username).match(/^\w{0,15}$/))
    return {
      valid: false,
      message: "Username can only contain letters, numbers, and underscores",
    };
  else return { valid: true, message: "" };
};

export const validateGithubUsername = (username: string): Validation => {
  if (username.length > 39)
    return {
      valid: false,
      message: "Username must be less than 40 characters",
    };
  else if (!String(username).match(/^[a-z\d](?:[a-z\d]|-(?=[a-z\d])){0,38}$/i))
    return {
      valid: false,
      message:
        "Username can only contain alphanumeric characters or single hyphens, and cannot begin or end with a hyphen",
    };
  else return { valid: true, message: "" };
};

export const validateHttpsUrl = (url: string): Validation => {
  try {
    const urlObj = new URL(url);
    if (urlObj.protocol !== "https:") {
      return {
        valid: false,
        message: "URL must start with https://",
      };
    }
    return { valid: true, message: "" };
  } catch (_) {
    return {
      valid: false,
      message: "Invalid URL",
    };
  }
};
