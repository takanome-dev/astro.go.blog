package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/gorilla/securecookie"
)

type Cookie struct {
	AccessToken string `json:"access_token"`
}

var secure = securecookie.New([]byte(os.Getenv("COOKIE_HASH_KEY")), []byte(os.Getenv("COOKIE_BLOCK_KEY")))
const CookieName = string("auth_token")

func EncodeCookie(token string, exp time.Time) (*http.Cookie, error) {
	value := map[string]string{
		CookieName: token,
	}
	
	encoded, err := secure.Encode(CookieName, value); 
	if err != nil { 
		return nil, err
	}

	cookie := &http.Cookie{
		Name:  CookieName,
		Value: encoded,
		Expires: exp,
		Path:  "/",
		// TODO: set secure true on prod
		Secure: false,
		HttpOnly: true,
		// SameSite: http.SameSite(3),
	}

	return cookie, nil
}

func DecodeCookie(cookie *http.Cookie) (Cookie, error) {
	value := make(map[string]string)
	
	err := secure.Decode(CookieName, cookie.Value, &value)
	if err != nil {
		return Cookie{}, err
	}

	token := Cookie{
		AccessToken: value[CookieName],
	}

	return token, nil
}