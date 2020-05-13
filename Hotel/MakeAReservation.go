package Hotel

import (
	"Cemal_Proj_2/Database"
	"encoding/json"
	"net/http"
	"time"
)

func MakeAReservation(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var roomId = r.Form.Get("RoomID")
	var day = r.Form.Get("DayCount")
	var guest = r.Form.Get("GuestName")

	query, err := Database.Db.Query("UPDATE `rooms` SET `Status` = 'Reserved' WHERE `rooms`.`ID` = "+roomId+";")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	currentTime := time.Now()
	var date = currentTime.Format("2006-01-02 15:04:05")
	query2, err := Database.Db.Query("INSERT INTO `reservations` (`ID`, `RoomID`, `DayCount`, `ReservedAt`, `Guest`) VALUES (NULL, '"+roomId+"', '"+day+"', '"+date+"', '"+guest+"');")
	defer query2.Close()

	response, _ := json.Marshal(map[string]string{"message": "reservation completed"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}