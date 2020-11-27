package controllers

import "github.com/gerokkos/clerk/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/populate", middlewares.SetMiddlewareJSON(s.Populate)).Methods("POST")
	s.Router.HandleFunc("/clerks", middlewares.SetMiddlewareJSON(s.Clerks)).Methods("GET")

}
