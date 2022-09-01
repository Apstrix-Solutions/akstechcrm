package controllers

import "github.com/praveencs87/akstechcrm/api/middlewares"

func (s *Server) initializeRoutes() {

	// Home Route
	s.Router.HandleFunc("/", middlewares.SetMiddlewareJSON(s.Home)).Methods("GET")

	// Login Route
	s.Router.HandleFunc("/login", middlewares.SetMiddlewareJSON(s.Login)).Methods("POST")

	//Users routes
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.CreateUser)).Methods("POST")
	s.Router.HandleFunc("/users", middlewares.SetMiddlewareJSON(s.GetUsers)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(s.GetUser)).Methods("GET")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateUser))).Methods("PUT")
	s.Router.HandleFunc("/users/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteUser)).Methods("DELETE")

	//Lead routes
	s.Router.HandleFunc("/leads", middlewares.SetMiddlewareJSON(s.CreateLead)).Methods("POST")
	s.Router.HandleFunc("/leads", middlewares.SetMiddlewareJSON(s.GetLeads)).Methods("GET")
	s.Router.HandleFunc("/leads/{id}", middlewares.SetMiddlewareJSON(s.GetLead)).Methods("GET")
	s.Router.HandleFunc("/leads/{id}", middlewares.SetMiddlewareJSON(middlewares.SetMiddlewareAuthentication(s.UpdateLead))).Methods("PUT")
	s.Router.HandleFunc("/leads/{id}", middlewares.SetMiddlewareAuthentication(s.DeleteLead)).Methods("DELETE")
}
