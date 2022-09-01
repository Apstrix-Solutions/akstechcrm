package controllers

import (
	"net/http"

	"github.com/praveencs87/akstechcrm/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Welcome To Akstech CRM API")

}
