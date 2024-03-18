package handler

import (
	"greatcomcatengineering.com/backend/utils"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {

}

func BaseApiHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {

	case "GET":
		utils.RespondWithJSON(w, http.StatusOK, "Great Comcat Engineering API", "Welcome to the Great Comcat Engineering API")

	default:
		utils.RespondWithError(w, http.StatusMethodNotAllowed, "Unsupported HTTP method")
	}
}
