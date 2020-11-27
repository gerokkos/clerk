package controllers

import "github.com/gerokkos/clerk/middlewares"

func (s *Server) initializeRoutes() {

	s.Router.HandleFunc("/populate", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")

}
