package main

import (
	"Cemal_Proj_2/Auth"
	"Cemal_Proj_2/Database"
	"Cemal_Proj_2/Hotel"
	"Cemal_Proj_2/User"
	"database/sql"
	"flag"
	"fmt"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)



func main() {
	var err error
	dbSource := fmt.Sprintf("root:%s@tcp(%s:%s)/%s?charset=utf8",  Database.DbPass, Database.DbHost, Database.DbPort, Database.DbName)

	Database.Db, err = sql.Open("mysql", dbSource)

	catch(err)

	flag.Parse()

	r := chi.NewRouter()
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300, // Maximum value not ignored by any of major browsers
	}))
	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.Route("/rooms", func(r chi.Router) {
		r.Post("/all", Hotel.GetAllHotelRooms)
		r.Post("/available", Hotel.GetAvailableHotelRooms)
		r.Post("/create-new-room", Hotel.CreateNewRoom)
	})
	r.Route("/reservation", func(r chi.Router) {
		r.Post("/make", Hotel.MakeAReservation)
		r.Post("/list", Hotel.ListReservations)
		r.Post("/edit", Hotel.EditReservation)
		r.Post("/delete", Hotel.DeleteReservation)
	})

	r.Route("/user", func(r chi.Router) {
		r.Post("/add", User.Add)
	})
	r.Route("/auth", func(r chi.Router) {
		r.Post("/login", Auth.Login)
		r.Post("/logout", Auth.Logout)
	})


	http.ListenAndServe(":3333", r)
	defer Database.Db.Close()
}
func catch(err error) {
	if err != nil {
		panic(err)
	}
}

