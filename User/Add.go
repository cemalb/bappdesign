package User

import (
	"Cemal_Proj_2/Database"
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"io"
	"net/http"
)

func Add(w http.ResponseWriter, r *http.Request) {
	cookie, _ := r.Cookie("login");
	if cookie.Value != "ok" {
		response, _ := json.Marshal(map[string]string{"error": "cookie problem"})
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(200)
		w.Write(response)
		return
	}
	r.ParseForm()
	var username = r.Form.Get("username")
	var password = r.Form.Get("password")
	h := md5.New()
	io.WriteString(h, password)


	query2, err := Database.Db.Query("INSERT INTO `users` (`id`, `username`, `password`) VALUES (NULL, '"+username+"', '"+hex.EncodeToString(h.Sum(nil))+"');")
	if err != nil {
		panic(err.Error())
	}
	defer query2.Close()

	response, _ := json.Marshal(map[string]string{"message": "new room added"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}
