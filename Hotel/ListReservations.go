package Hotel

import (
	"Cemal_Proj_2/Database"
	"github.com/go-chi/render"
	"net/http"
)

type Reservation struct {
	ID     string `json:"id"`
	RoomID string  `json:"room_id"`
	DayCount string  `json:"day_count"`
	ReservedAt  string `json:"reserved_at"`
	GuestName  string `json:"guest_name"`
}
var reservations = []Reservation{}

func ListReservations(w http.ResponseWriter, r *http.Request) {
	results, err := Database.Db.Query("SELECT ID, RoomID, DayCount, ReservedAt, Guest  FROM reservations")
	if err != nil {
		panic(err.Error())
	}
	reservations = []Reservation{}
	for results.Next() {
		var reserv Reservation
		err = results.Scan(&reserv.ID, &reserv.RoomID, &reserv.DayCount, &reserv.ReservedAt, &reserv.GuestName)
		if err != nil {
			panic(err.Error())
		}
		reservations = append(reservations, Reservation{ID:reserv.ID, RoomID: reserv.RoomID, DayCount: reserv.DayCount, ReservedAt: reserv.ReservedAt, GuestName: reserv.GuestName});
	}
	defer results.Close()

	if err := render.RenderList(w, r, ReservationsListResponse(reservations)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}
}

func ReservationsListResponse(reservations []Reservation) []render.Renderer {
	list := []render.Renderer{}
	for _, hRoom := range reservations {
		list = append(list, ReservationsResponse(hRoom))
	}
	return list
}

func ReservationsResponse(reservations Reservation) *ReservationResponse {
	resp := &ReservationResponse{Reservation: reservations}
	return resp
}

type ReservationResponse struct {
	Reservation
	Elapsed int64 `json:"elapsed"`
}



func (rd *ReservationResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	rd.Elapsed = 10
	return nil
}