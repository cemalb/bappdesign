package Auth

import (
	"Cemal_Proj_2/Database"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var username = r.Form.Get("identity")
	var password = r.Form.Get("password")
	h := md5.New()
	io.WriteString(h, password)
	var query = "SELECT id FROM `users` WHERE `username` LIKE '"+username+"' AND `password` LIKE '"+hex.EncodeToString(h.Sum(nil))+"'";
	query2 := Database.Db.QueryRow(query)
	var id int
	query2.Scan(&id)
	if id != 0 {
		expiration := time.Now().Add(365 * 24 * time.Hour)
		cookie := http.Cookie{Name: "login", Value: "ok", Expires: expiration}
		http.SetCookie(w, &cookie)

		response, _ := json.Marshal(map[string]string{"message": "successfully logged in"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(response)
	} else {
		response, _ := json.Marshal(map[string]string{"wrong": "true"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(response)
	}
}
