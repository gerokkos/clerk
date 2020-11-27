package controllers

import (
	"github.com/gerokkos/clerk/api/middleware"
)

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/populate", middleware.SetMiddlewareJSON(s.Populate)).Methods("POST")
	s.Router.HandleFunc("/clerks", middleware.SetMiddlewareJSON(s.Clerks)).Methods("GET")

}
