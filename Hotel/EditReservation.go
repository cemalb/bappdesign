package Hotel

import (
	"Cemal_Proj_2/Database"
	"encoding/json"
	"net/http"
)

func EditReservation(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var reservationId = r.Form.Get("ID")
	var roomId = r.Form.Get("RoomID")
	var day = r.Form.Get("DayCount")
	var guest = r.Form.Get("GuestName")

	currentReservation := Database.Db.QueryRow("SELECT ID, RoomID FROM reservations WHERE `reservations`.`ID` = "+reservationId+";")
	var RoomID = ""
	var ID = ""
	currentReservation.Scan(&ID, &RoomID)
	if RoomID != roomId {
		pastRoomQuery, err := Database.Db.Query("UPDATE `rooms` SET `Status` = 'Available' WHERE `rooms`.`ID` = "+RoomID+";")
		if err != nil {
			panic(err.Error())
		}
		pastRoomQuery.Close();
	}

	query, err := Database.Db.Query("UPDATE `rooms` SET `Status` = 'Reserved' WHERE `rooms`.`ID` = "+roomId+";")
	if err != nil {
		panic(err.Error())
	}
	defer query.Close()
	query2, err := Database.Db.Query("UPDATE `reservations` SET `RoomID` = '"+roomId+"', `DayCount` = '"+day+"', `Guest` = '"+guest+"' WHERE `reservations`.`ID` = "+reservationId+";")
	defer query2.Close()

	response, _ := json.Marshal(map[string]string{"message": "reservation successfully updated"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}