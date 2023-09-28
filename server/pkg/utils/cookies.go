package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/securecookie"
)

var secure = securecookie.New([]byte(os.Getenv("COOKIE_HASH_KEY")), []byte(os.Getenv("COOKIE_BLOCK_KEY")))
const cookieName = string("auth_token")

func EncodeCookie(token string, exp time.Time) (*http.Cookie, error) {
	value := map[string]string{
		"access_token": token,
	}
	
	encoded, err := secure.Encode(cookieName, value); 
	if err != nil { 
		return nil, err
	}

	cookie := &http.Cookie{
		Name:  cookieName,
		Value: encoded,
		Expires: exp,
		Path:  "/",
		// TODO: set secure true on prod
		Secure: false,
		HttpOnly: true,
		SameSite: http.SameSite(3),
	}

	return cookie, nil
}

func DecodeCookie(cookie *http.Cookie) error {
	value := make(map[string]string)
	return secure.Decode(cookieName, cookie.Value, &value);
}