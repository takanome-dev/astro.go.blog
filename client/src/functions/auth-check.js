// src/functions/auth-check.js
export default function authCheck(req, res) {
  const cookie = req.cookies["auth_token"]; // replace 'your-cookie-name' with your actual cookie name

  if (!cookie) {
    const params = new URLSearchParams({
      redirectedFrom: req.url,
    });
    res.redirect(302, `/login?${params.toString()}`);
  } else {
    // If the user is authenticated, continue to the requested page
    res.end();
  }
}
