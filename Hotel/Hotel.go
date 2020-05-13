package Hotel

import (
	"Cemal_Proj_2/Database"
	"github.com/go-chi/render"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)


type User struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type HotelRoom struct {
	ID     string `json:"id"`
	Type string  `json:"type"`
	SubType string  `json:"sub_type"`
	Status  string `json:"status"`
	IsSmokingAvailable  string `json:"is_smoking_available"`
}
var hotelRooms = []HotelRoom{}

func GetAvailableHotelRooms(w http.ResponseWriter, r *http.Request) {
	results, err := Database.Db.Query("SELECT ID, Type, Subtype, Status, IsSmokingAvailable  FROM rooms WHERE `Status` = 'Available'")
	if err != nil {
		panic(err.Error())
	}
	hotelRooms = []HotelRoom{}
	for results.Next() {
		var room HotelRoom
		err = results.Scan(&room.ID, &room.Type, &room.SubType, &room.Status, &room.IsSmokingAvailable)
		if err != nil {
			panic(err.Error())
		}
		hotelRooms = append(hotelRooms, HotelRoom{ID:room.ID, Type: room.Type, SubType: room.SubType, Status: room.Status, IsSmokingAvailable: room.IsSmokingAvailable});
	}
	defer results.Close()

	if err := render.RenderList(w, r, HotelRoomListResponse(hotelRooms)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

}

func GetAllHotelRooms(w http.ResponseWriter, r *http.Request) {
	results, err := Database.Db.Query("SELECT ID, Type, Subtype, Status, IsSmokingAvailable  FROM rooms")
	if err != nil {
		panic(err.Error())
	}
	hotelRooms = []HotelRoom{}
	for results.Next() {
		var room HotelRoom
		err = results.Scan(&room.ID, &room.Type, &room.SubType, &room.Status, &room.IsSmokingAvailable)
		if err != nil {
			panic(err.Error())
		}
		hotelRooms = append(hotelRooms, HotelRoom{ID:room.ID, Type: room.Type, SubType: room.SubType, Status: room.Status, IsSmokingAvailable: room.IsSmokingAvailable});
	}
	defer results.Close()

	if err := render.RenderList(w, r, HotelRoomListResponse(hotelRooms)); err != nil {
		render.Render(w, r, ErrRender(err))
		return
	}

}

func HotelRoomListResponse(hotelRoom []HotelRoom) []render.Renderer {
	list := []render.Renderer{}
	for _, hRoom := range hotelRoom {
		list = append(list, HotelRoomResponse(hRoom))
	}
	return list
}

func HotelRoomResponse(hotelRoom HotelRoom) *RoomResponse {
	resp := &RoomResponse{HotelRoom: hotelRoom}
	return resp
}

type RoomResponse struct {
	HotelRoom
	Elapsed int64 `json:"elapsed"`
}



func (rd *RoomResponse) Render(w http.ResponseWriter, r *http.Request) error {
	// Pre-processing before a response is marshalled and sent across the wire
	rd.Elapsed = 10
	return nil
}

func ErrRender(err error) render.Renderer {
	return &ErrResponse{
		Err:            err,
		HTTPStatusCode: 422,
		StatusText:     "Error rendering response.",
		ErrorText:      err.Error(),
	}
}




type ErrResponse struct {
	Err            error `json:"-"` // low-level runtime error
	HTTPStatusCode int   `json:"-"` // http response status code

	StatusText string `json:"status"`          // user-level status message
	AppCode    int64  `json:"code,omitempty"`  // application-specific error code
	ErrorText  string `json:"error,omitempty"` // application-level error message, for debugging
}

func (e *ErrResponse) Render(w http.ResponseWriter, r *http.Request) error {
	render.Status(r, e.HTTPStatusCode)
	return nil
}



var ErrNotFound = &ErrResponse{HTTPStatusCode: 404, StatusText: "Resource not found."}
