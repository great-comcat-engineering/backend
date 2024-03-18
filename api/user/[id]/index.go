package handler

import (
	"greatcomcatengineering.com/backend/services/user"
	"greatcomcatengineering.com/backend/utils"
	"net/http"
)

func GetUserByEmail(w http.ResponseWriter, r *http.Request) {

	switch r.Method {

	case "POST":
		user.HandleGetUserByEmail(w, r)

	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Unsupported HTTP method")
	}
}
