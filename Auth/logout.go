package Auth

import (
	"net/http"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	cookie := http.Cookie{Name: "login", Value: "not_ok", MaxAge: -1}
	http.SetCookie(w, &cookie)
}