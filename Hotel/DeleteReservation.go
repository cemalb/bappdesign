package Hotel

import (
	"Cemal_Proj_2/Database"
	"encoding/json"
	"net/http"
)

func DeleteReservation(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	var reservationId = r.Form.Get("ID")

	currentReservation := Database.Db.QueryRow("SELECT ID, RoomID FROM reservations WHERE `reservations`.`ID` = "+reservationId+";")
	var RoomID = ""
	var ID = ""
	currentReservation.Scan(&ID, &RoomID)
	pastRoomQuery, err := Database.Db.Query("UPDATE `rooms` SET `Status` = 'Available' WHERE `rooms`.`ID` = "+RoomID+";")
	if err != nil {
		panic(err.Error())
	}
	pastRoomQuery.Close();

	deleteQuery, err := Database.Db.Query("DELETE FROM `reservations` WHERE `reservations`.`ID` = "+reservationId+";")
	if err != nil {
		panic(err.Error())
	}
	deleteQuery.Close();


	response, _ := json.Marshal(map[string]string{"message": "reservation successfully deleted"})
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(response)
}