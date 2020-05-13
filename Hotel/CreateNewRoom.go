package Hotel

import (
	"Cemal_Proj_2/Database"
	"encoding/json"
	"net/http"
)

func CreateNewRoom(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var roomType = r.Form.Get("Type")
	var subType = r.Form.Get("SubType")
	var isSmokingAvailable = r.Form.Get("IsSmokingAvailable")

	query2, err := Database.Db.Query("INSERT INTO `rooms` (`ID`, `Type`, `SubType`, `Status`, `IsSmokingAvailable`) VALUES (NULL, '"+roomType+"', '"+subType+"', 'Available', '"+isSmokingAvailable+"');")
	if err != nil {
		panic(err.Error())
	}
	defer query2.Close()

	response, _ := json.Marshal(map[string]string{"message": "new room added"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}